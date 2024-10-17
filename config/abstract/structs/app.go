package structs

// RegConfigMap the registered instance map of Configure abs
var RegConfigMap = map[string]interface{}{}

// AttrMap different with file name(case.title)
// if the attribute of "CGlobal" different with file name(case.title),
// you need set the map of "attrMap"
var AttrMap = map[string]string{
	// typedef.AppName: "application", // if file name "application.yml"
}

// init App
func init() {
	RegConfigMap["app"] = &App{} // app.yml
}

// App 应用配置
type App struct {
	Env             string `yaml:"env" json:"env"`
	Host            string `yaml:"host" json:"host"`
	Port            string `yaml:"port" json:"port"`
	WithRoute       bool   `yaml:"with_route" json:"with_route"`
	WithTLS         bool   `yaml:"with_TLS" json:"with_TLS"`
	CrtFile         string `yaml:"crt_file" json:"crt_file"`
	KeyFile         string `yaml:"key_file" json:"key_file"`
	ViewDir         string `yaml:"view_dir" json:"view_dir"`
	PublicDir       string `yaml:"public_dir" json:"public_dir"`
	WithSwagger     bool   `yaml:"with_swagger" json:"with_swagger"`
	WithPProf       bool   `yaml:"with_pprof" json:"with_pprof"`
	DefaultTimezone string `yaml:"default_timezone" json:"default_timezone"`
}
