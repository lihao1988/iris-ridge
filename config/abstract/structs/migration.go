package structs

// init Migration
func init() {
	RegConfigMap["migration"] = &Migration{}
}

// Migration 数据迁移相关配置
type Migration struct {
	MigrationDir string `yaml:"migration_dir" json:"migration_dir"`
	AutoMigrate  bool   `yaml:"auto_migrate" json:"auto_migrate"`
	AllowMissing bool   `yaml:"allow_missing" json:"allow_missing"`
}
