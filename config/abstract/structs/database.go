package structs

// init Database
func init() {
	RegConfigMap["database"] = &Database{} // database.yml
}

// Database 数据库连接配置
type Database struct {
	Default     string `yaml:"default" json:"default"`
	AutoMigrate bool   `yaml:"auto_migrate" json:"auto_migrate"`
	Connections `yaml:"connections" json:"connections"`
}

// Connections database connections
type Connections struct {
	Mysql *MysqlConn `yaml:"mysql" json:"mysql"`
}

// MysqlConn mysql connection
type MysqlConn struct {
	Host         string            `yaml:"host" json:"host"`
	Port         string            `yaml:"port" json:"port"`
	Database     string            `yaml:"database" json:"database"`
	Username     string            `yaml:"username" json:"username"`
	Password     string            `yaml:"password" json:"password"`
	Charset      string            `yaml:"charset" json:"charset"`
	MaxIdleCount int               `yaml:"maxIdle_count" json:"maxIdle_count"`
	MaxOpenCount int               `yaml:"max_open_count" json:"max_open_count"`
	Prefix       string            `yaml:"prefix" json:"prefix"`
	Params       map[string]string `yaml:"params" json:"params"`
}
