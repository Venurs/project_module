package conf

import (
	"os"
	"reflect"
	"sync"

	"github.com/crgimenes/goconfig"
)

const (
	LocalMode       string = "local"
	HuoHuoMode      string = "huohuo"
	ProdMode        string = "prod"
	DockerModel     string = "docker"
	ProdDockerModel string = "prod_docker"
)

var modeName = LocalMode

type configMain struct {
	HuoHuo                configHuoHuo
	Prod                  configProd
	Docker                configDocker
	GOModel               string `cfgDefault:"local"`
	Domain                string
	DatabaseSource        string `cfgDefault:"sg:sg123456@tcp(127.0.0.1:3306)/sellergrowth"`
	DatabaseSourceReplica string `cfgDefault:""`
	RedisAddr             string `cfgDefault:"127.0.0.1:6379"`
	RedisPassword         string `cfgDefault:""`
	InfoLogFile           string
	ErrorLogFile          string
	ActionLogFile         string
	PolyvLiveAppId        string `cfgDefault:"emwn1txkpa"`
	PolyvLiveAppSecret    string `cfgDefault:"9254d15500bb4fb18b3ba2e9e337e4eb"`
	PolyvLiveUserId       string `cfgDefault:"3297899785"`
	UnicornServiceRpcUrl  string `cfgDefault:"http://127.0.0.1:8001/rpc/"`
	RABBITMQHOST          string `cfgDefault:"amqp://guest:guest@120.24.68.34:5672/"`
}

var (
	Settings configMain
	once     sync.Once
)

func init() {
	once.Do(func() {
		Settings = New()
	})
}

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func New() configMain {

	config := configMain{}
	err := goconfig.Parse(&config)
	if err != nil {
		panic("config initialization failed.")
	}
	projectPath := GetProjectPath()
	config.InfoLogFile = projectPath + "/logs/info.log"
	config.ErrorLogFile = projectPath + "/logs/error.log"
	config.ActionLogFile = projectPath + "/logs/action.log"
	local_setting := getLocalSetting()
	if local_setting == "huohuo" {
		sv := reflect.ValueOf(config.HuoHuo)
		resetConfig(&configHuoHuo{}, sv, &config)
	} else if local_setting == "prod" {
		sv := reflect.ValueOf(config.Prod)
		resetConfig(&configProd{}, sv, &config)
	} else if local_setting == "docker" {
		sv := reflect.ValueOf(config.Docker)
		resetConfig(&configDocker{}, sv, &config)
	} else if local_setting == "prod_docker" {
		sv := reflect.ValueOf(config.Docker)
		resetConfig(&configDocker{}, sv, &config)
	}

	return config
}

func resetConfig(configObject interface{}, sv reflect.Value, config interface{}) {
	st := reflect.TypeOf(configObject)
	v_elem := reflect.ValueOf(config).Elem()
	refField := st.Elem()
	for i := 0; i < refField.NumField(); i++ {
		name := refField.Field(i).Name
		if v_elem.FieldByName(name).CanSet() {
			v_elem.FieldByName(name).Set(sv.FieldByName(name))
		}
	}
}

func getLocalSetting() string {
	localSetting := os.Getenv("GO_LOCAL_SETTING")
	if len(localSetting) == 0 {
		setMode(LocalMode)
	} else {
		setMode(localSetting)
	}
	return modeName
}

func setMode(value string) {
	switch value {
	case LocalMode:
		modeName = LocalMode
	case HuoHuoMode:
		modeName = HuoHuoMode
	case ProdMode:
		modeName = ProdMode
	case DockerModel:
		modeName = DockerModel
	case ProdDockerModel:
		modeName = ProdDockerModel
	default:
		panic("local settings unknown: " + value)
	}
}
