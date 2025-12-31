package conf

type AppConfig struct {
	Server ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	OSS    OSSConfig    `mapstructure:"oss"`
	Alipay AlipayConfig `mapstructure:"alipay"`
	WxPay  WxpayConfig  `mapstructure:"wxpay"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type MySQLConfig struct {
	DSN string `mapstructure:"dsn" yaml:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

type OSSConfig struct {
	Endpoint    string `mapstructure:"endpoint" yaml:"endpoint"`
	AccessKey   string `mapstructure:"access_key" yaml:"access_key"`
	SecretKey   string `mapstructure:"secret_key" yaml:"secret_key"`
	VideoBucket string `mapstructure:"video_bucket" yaml:"video_bucket"`
	ImgBucket   string `mapstructure:"img_bucket" yaml:"img_bucket"`
}

type AlipayConfig struct {
	AppID      string `mapstructure:"app_id"`
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
}

type WxpayConfig struct {
	AppID      string `mapstructure:"app_id"`
	MchID      string `mapstructure:"mch_id"`
	PrivateKey string `mapstructure:"private_key"`
}
