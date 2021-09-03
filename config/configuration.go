package config

type AlterManager struct {
	Url string `yaml:"url"`
	Name string `yaml:"name"`
}

type Configuration struct {
	AlertManagers [] AlterManager `yaml:"alert_managers"`
}
