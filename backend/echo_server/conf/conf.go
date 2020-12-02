package conf

type Service struct {
	Port string `yaml:"port"`
}

type Database struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Type   string `yaml:"type"`
	Db     string `yaml:"db"`
	Conn   struct {
		Maxidle int `yaml:"maxidle"`
		Maxopen int `yaml:"maxopen"`
	}
}

type Conf struct {
	Name   string `yaml:"name"`
	Enable bool   `yaml:"enable"`
	Srv    Service
	Db     Database
}

type Protocol struct {
	TCPConf   Conf `yaml:"tcp"`
	HTTPConf  Conf `yaml:"http"`
	HTTPSConf Conf `yaml:"https"`
	GRPCConf  Conf `yaml:"grpc"`
}
