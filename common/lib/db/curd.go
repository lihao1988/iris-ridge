package db

import (
	"errors"

	"ridge/common/model"

	"gorm.io/gorm"
)

// First finds the first record ordered by primary key
func First(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Where(m).Scopes(options...).First(m).Error
}

// FindOne finds the first record
func FindOne(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Limit(1).Where(m).Scopes(options...).Find(m).Error
}

// Find finds the records
func Find(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Where(m).Scopes(options...).Find(m).Error
}

// IsExists check if the record exists
func IsExists(m model.Impl, options ...func(*gorm.DB) *gorm.DB) bool {
	if err := m.GetDB().Scopes(options...).Where(m).First(m).Error; err != nil {
		return false
	}

	return m.GetID() > 0
}

// IsGormNotFoundErr record not found error
func IsGormNotFoundErr(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// Create creates the record by model_instance
func Create(m model.Impl) error {
	return m.GetDB().Create(m).Error
}

// Update updates the record by model_instance
func Update(m model.Impl) error {
	return m.GetDB().Save(m).Error
}

// UpdateMap updates the record by model attribute and options. data must be a map.
func UpdateMap(m model.Impl, data map[string]interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Model(m).Where(m).Scopes(options...).Updates(data).Error
}

// UpdateMapScope updates the record by options. data must be a map.
func UpdateMapScope(m model.Impl, data map[string]interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Model(m).Scopes(options...).Updates(data).Error
}

// UpdateMapByID updates the record by 'id'. data must be a map.
func UpdateMapByID(m model.Impl, id uint, data map[string]interface{}) error {
	return m.GetDB().Model(m).Where("`id` = ?", id).Updates(data).Error
}

// UpdateColumn updates the column of record.
func UpdateColumn(m model.Impl, key string, value interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Model(m).Scopes(options...).UpdateColumn(key, value).Error
}

// UpdateColumns updates the columns of record.
func UpdateColumns(m model.Impl, values interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Model(m).Scopes(options...).UpdateColumns(values).Error
}

// Delete Deletes performs a soft delete (update 'delete_at' field)
func Delete(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Scopes(options...).Delete(m).Error
}

// DeleteForHard Deletes performs a hard delete
func DeleteForHard(m model.Impl, options ...func(*gorm.DB) *gorm.DB) error {
	return m.GetDB().Unscoped().Scopes(options...).Delete(m).Error
}

// DeleteByID Deletes a record by primary key
func DeleteByID(m model.Impl, id uint, unscoped bool) error {
	// unscoped 'true', performs a hard delete
	if unscoped {
		return DeleteForHard(m, func(tx *gorm.DB) *gorm.DB {
			return tx.Where("id = ?", id)
		})
	}

	// unscoped 'false', performs a soft delete
	return Delete(m, func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	})
}

// List get the list of data
func List(b model.Impl, res interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	err := b.GetDB().Where(b).Scopes(options...).Find(res).Error
	if err != nil && !IsGormNotFoundErr(err) {
		return err
	}

	return nil
}

// Scan get the list of data (table_name)
func Scan(m model.Impl, res interface{}, options ...func(*gorm.DB) *gorm.DB) error {
	// only table_name
	err := m.GetDB().Table(m.TableName()).Where(m).Scopes(options...).Scan(res).Error
	if err != nil && !IsGormNotFoundErr(err) {
		return err
	}

	return nil
}

// Count get the count of list
func Count(m model.Impl, options ...func(*gorm.DB) *gorm.DB) int64 {
	var count int64
	err := m.GetDB().Model(m).Where(m).Scopes(options...).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

// PageList get the page of records
func PageList(m model.Impl, res interface{}, page, pageSize int,
	options ...func(*gorm.DB) *gorm.DB) (int64, error) {
	// get count
	count := Count(m, options...)

	// get data
	options = append(options, WithPage(page, pageSize))
	err := List(m, res, options...)
	return count, err
}
