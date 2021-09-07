package config

type AlterManager struct {
	Url string `yaml:"url"`
	Name string `yaml:"name"`
}

type Configuration struct {
	Server struct{
		Port string `yaml:"port"`
	} `yaml:"server"`
	AlertManagers [] AlterManager `yaml:"alert_managers"`
}
