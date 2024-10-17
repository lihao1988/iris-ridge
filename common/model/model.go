package model

import (
	"time"

	"gorm.io/gorm"
)

// Model the base model
type Model struct {
	// gorm_db instance
	db *gorm.DB

	// the table's base field
	ID        uint       `gorm:"primary_key" json:"id"`                                                                                        // 主键ID
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;not null;comment:'创建时间'" json:"createdAt" example:"2021-01-07T14:15:09+08:00"` // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime;not null;comment:'修改时间'" json:"updatedAt" example:"2021-01-07T14:15:09+08:00"` // 修改时间
	DeletedAt *time.Time `sql:"index" json:"-"`                                                                                                // 软删除
}

// TableName get the table name
func (m *Model) TableName() string {
	return ""
}

// TableComment table comment
func (m *Model) TableComment() string {
	return ""
}

// GetDB get db
func (m *Model) GetDB() *gorm.DB {
	return m.db
}

// SetDB set db
func (m *Model) SetDB(db *gorm.DB) {
	m.db = db
}

// GetID get record id
func (m *Model) GetID() uint {
	return m.ID
}

// SetID set record id
func (m *Model) SetID(id uint) {
	m.ID = id
}

// IsExists whether the record exist
func (m *Model) IsExists() bool {
	return m.ID > 0
}

// ResetCreatedAt reset create_at
func (m *Model) ResetCreatedAt() *Model {
	m.CreatedAt = time.Now()
	return m
}

// ResetUpdatedAt reset update_at
func (m *Model) ResetUpdatedAt() *Model {
	m.UpdatedAt = time.Now()
	return m
}

// Reset reset fields
func (m *Model) Reset() *Model {
	m.ResetUpdatedAt()
	m.ResetCreatedAt()
	m.SetID(0)
	return m
}
