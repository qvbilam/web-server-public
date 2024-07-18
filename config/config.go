package config

type ServerConfig struct {
	Home               string
	Name               string             `mapstructure:"name" json:"name"`
	Host               string             `mapstructure:"host" json:"host"`
	Port               int64              `mapstructure:"port" json:"port"`
	Tags               []string           `mapstructure:"tags" json:"tags"`
	OssConfig          OssConfig          `mapstructure:"oss" json:"oss"`
	UserServerConfig   UserServerConfig   `mapstructure:"user-server" json:"user-server"`
	PublicServerConfig PublicServerConfig `mapstructure:"public-server" json:"public-server"`
}

type OssConfig struct {
	Key         string `mapstructure:"key" json:"key"`
	Secret      string `mapstructure:"secret" json:"secret"`
	Host        string `mapstructure:"host" json:"host"`
	CallbackUrl string `mapstructure:"callback_url" json:"callback_url"`
	UploadDir   string `mapstructure:"upload_dir" json:"upload_dir"`
	ExpireTime  int64  `mapstructure:"expire_time" json:"expire_time"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type PublicServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
