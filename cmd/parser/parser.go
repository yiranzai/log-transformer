package parser

import (
	"fmt"

	"github.com/golang-module/carbon"
	"github.com/nxadm/tail"
)

const TypeNginx = "nginx"
const TypeNginxAccess = "nginx_access"

const NginxTimeLocal = "02/Jan/2006:15:04:05 -0700"

//IParser Parser接口
type IParser interface {
	//Parse 解析一行内容
	Parse(line *tail.Line) (map[string]interface{}, *carbon.Carbon)
	//ParseDirOrFile 解析一个目录
	ParseDirOrFile(path string) map[string][]string
	//ParseFile 解析单个文件
	ParseFile(file string) (string, string)
}

//NewParser 实例化
func NewParser(cfg IParserConfig) IParser {
	var parser IParser
	switch cfg.GetType() {
	case TypeNginx:
		parser = NewNginxAccessParser(cfg)
	case TypeNginxAccess:
		parser = NewNginxAccessParser(cfg)
	default:
		panic(fmt.Sprintf("parser type %s not support", cfg.GetType()))
	}
	return parser
}

//IParserConfig Parser所需配置文件接口
type IParserConfig interface {
	GetType() string
	GetTimeField() string
	GetFormat() string
}
