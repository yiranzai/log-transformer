package conf

import (
	"github.com/spf13/viper"
	"github.com/yiranzai/log-transformer/utils"
)

//Config 配置文件
type Config struct {
	Monitor     *Monitor
	Parser      *Parser
	Transformer *Transformer
}

//Monitor 监控的配置
type Monitor struct {
	Path string
}

//Parser Nginx配置
type Parser struct {
	Type      string
	Format    string
	TimeField string
	Vars      []string
}

func (p *Parser) GetType() string {
	return p.Type
}

func (p *Parser) GetTimeField() string {
	return p.TimeField
}

func (p *Parser) GetFormat() string {
	return p.Format
}

//Transformer Transformer处理
type Transformer struct {
	Type      string
	Path      string
	MaxFiles  uint
	Extension string
}

func (f *Transformer) GetType() string {
	return f.Type
}

func (f *Transformer) GetMaxFiles() uint {
	return f.MaxFiles
}

func (f *Transformer) GetWritePath() string {
	return f.Path
}

func (f *Transformer) GetExtension() string {
	return f.Extension
}

//Init Transformer初始化
func (f *Transformer) Init() {
	err := utils.CreateDirNX(f.Path)
	if err != nil {
		panic(err)
	}
}

//Init init config
func Init(path string) *Config {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Loading config file failed, config file not found!")
		} else {
			panic("Loading config file failed, " + err.Error())
		}
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic("Unmarshal config file failed, " + err.Error())
	}
	cfg.Transformer.Init()
	return &cfg
}
