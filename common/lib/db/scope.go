package db

import (
	"fmt"
	"gorm.io/gorm"
)

// BuildQuoteField set table_name and field
func BuildQuoteField(field string, tableName ...string) string {
	field = fmt.Sprintf("`%s`", field)
	if len(tableName) > 0 {
		field = fmt.Sprintf("`%s`.%s", tableName[0], field)
	}

	return field
}

// ScopeOpts scope func list
type ScopeOpts struct {
	opts []func(*gorm.DB) *gorm.DB // fun list
}

// NewScopeOpts new instance
func NewScopeOpts() *ScopeOpts {
	return &ScopeOpts{}
}

// Add append the func of query condition
func (s *ScopeOpts) Add(opt ...func(*gorm.DB) *gorm.DB) {
	s.opts = append(s.opts, opt...)
}

// Export export all func
func (s *ScopeOpts) Export() []func(*gorm.DB) *gorm.DB {
	return s.opts
}

// Copy clone scopeOpts
func (s *ScopeOpts) Copy() *ScopeOpts {
	sp := &ScopeOpts{}
	sp.opts = s.opts
	return sp
}

// WithSelect set select fields
func WithSelect(query interface{}, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	}
}

// WithWhere with where_options
func WithWhere(query interface{}, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

// WithOr with or_options
func WithOr(query interface{}, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Or(query, args...)
	}
}

// WithEq set field equal to value
func WithEq(field string, value interface{}, tableName ...string) func(*gorm.DB) *gorm.DB {
	field = BuildQuoteField(field, tableName...)
	return WithWhere(fmt.Sprintf("%s = ?", field), value)
}

// WithIn set field in values
func WithIn(field string, value interface{}, tableName ...string) func(*gorm.DB) *gorm.DB {
	field = BuildQuoteField(field, tableName...)
	return WithWhere(fmt.Sprintf("%s IN (?)", field), value)
}

// WithLike set field equal like value
func WithLike(field string, value interface{}, tableName ...string) func(*gorm.DB) *gorm.DB {
	field = BuildQuoteField(field, tableName...)
	return WithWhere(fmt.Sprintf("%s LIKE ?", field), value)
}

// WithJoins set select joins
func WithJoins(query string, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins(query, args...)
	}
}

// WithOrder set order
func WithOrder(field, sort string, tableName ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		field = BuildQuoteField(field, tableName...)
		return db.Order(fmt.Sprintf("%s %s", field, sort))
	}
}

// WithGroup support select group
func WithGroup(query string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Group(query)
	}
}

// WithHaving support group having
func WithHaving(query string, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Having(query, args...)
	}
}

// WithLimit with limit
func WithLimit(offset, limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

// WithLimitOne with limit one
func WithLimitOne() func(*gorm.DB) *gorm.DB {
	return WithLimit(0, 1)
}

// WithPage with limit for page
func WithPage(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize < 0 {
			pageSize = 20
		}
		if page < 0 {
			page = 1
		}

		start := (page - 1) * pageSize
		return db.Offset(start).Limit(pageSize)
	}
}

// WithUnscoped ignore scoped
func WithUnscoped() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

// WithLockForUpdate set select lock for update
func WithLockForUpdate() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Set("gorm:query_option", "FOR UPDATE")
	}
}

// WithPreload preload associations with given conditions
func WithPreload(column string, conditions ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(column, conditions...)
	}
}
