package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/golang-module/carbon"
	"github.com/nxadm/tail"
)

//NginxAccessParser NginxAccessLog解析器
type NginxAccessParser struct {
	VarNames   []string
	TimeFiled  string
	parseRegex *regexp.Regexp
}

//NewNginxAccessParser Nginx文件处理的构造函数
func NewNginxAccessParser(cfg IParserConfig) IParser {
	ng := &NginxAccessParser{TimeFiled: cfg.GetTimeField(), parseRegex: getVarRegex(cfg.GetFormat())}
	names := parseVarNames(cfg.GetFormat())
	ng.VarNames = make([]string, 0, len(names))
	for _, name := range names {
		ng.VarNames = append(ng.VarNames, name[1])
	}
	return ng
}

//Parse 解析一行文件
func (n *NginxAccessParser) Parse(line *tail.Line) (map[string]interface{}, *carbon.Carbon) {
	var time carbon.Carbon
	result := make(map[string]interface{})
	matched := parseText(n.parseRegex, line.Text)
	if len(matched) == 0 {
		return result, &time
	}
	vars := matched[0][1:]
	for i, name := range n.VarNames {
		result[name] = vars[i]
	}

	time = carbon.Time2Carbon(line.Time)
	if result[n.TimeFiled] != nil {
		str := result[n.TimeFiled].(string)
		if len(str) != 0 {
			// 先判断是否nginx通用格式
			time = carbon.ParseByLayout(str, NginxTimeLocal)
			if time.Error != nil {
				// 再判断是否通用可兼容的格式
				time = carbon.Parse(str)
				if time.Error != nil {
					time = carbon.Time2Carbon(line.Time)
				}
			}
		}
	}
	return result, &time
}

//ParseDirOrFile 解析目录或文件中的文件(无递归)
func (n *NginxAccessParser) ParseDirOrFile(path string) map[string][]string {
	var result map[string][]string
	result = make(map[string][]string)
	fi, err := os.Stat(path)
	if err != nil {
		return result
	}
	mode := fi.Mode()
	if mode.IsDir() {
		files, _ := ioutil.ReadDir(path)
		for _, fileInfo := range files {
			if fileInfo.Mode().IsRegular() {
				svrName, file := n.ParseFile(path + fileInfo.Name())
				if len(svrName) != 0 {
					if _, ok := result[svrName]; !ok {
						result[svrName] = make([]string, 0, 0)
					}
					result[svrName] = append(result[svrName], file)
				}
			}
		}
	}
	if mode.IsRegular() {
		svrName, file := n.ParseFile(path)
		if len(svrName) != 0 {
			if _, ok := result[svrName]; !ok {
				result[svrName] = make([]string, 0, 0)
			}
			result[svrName] = append(result[svrName], file)
		}
	}
	return result
}

//ParseFile 识别access文件名
func (n *NginxAccessParser) ParseFile(file string) (string, string) {
	base := filepath.Base(file)
	result := regexp.MustCompile(`(.*[a-zA-Z])[^a-zA-Z]+access[^a-zA-Z]+log$`).FindAllStringSubmatch(base, -1)
	if result == nil {
		result = regexp.MustCompile(`access[^a-zA-Z]+(.*[a-zA-Z])[^a-zA-Z]+log$`).FindAllStringSubmatch(base, -1)
		if result == nil {
			return "", file
		}
	}

	return result[0][1], file
}

//parseText 匹配一段文本
func parseText(regex *regexp.Regexp, str string) [][]string {
	return regex.FindAllStringSubmatch(str, -1)
}

//parseVarNames 解析nginx配置文件格式
func parseVarNames(format string) [][]string {
	reg := getFormatRegex()
	return parseText(reg, format)
}

//getFormatRegex 获取格式化nginx配置文件的正则
func getFormatRegex() *regexp.Regexp {
	return regexp.MustCompile(`\$(\w+)\b`)
}

//getVarRegex 获取nginx日志解析的正则
func getVarRegex(str string) *regexp.Regexp {
	str = regexp.QuoteMeta(str)
	result := regexp.MustCompile(`\\\$(\w+)\b`).ReplaceAllString(str, "(.*)")
	return regexp.MustCompile(result)
}
