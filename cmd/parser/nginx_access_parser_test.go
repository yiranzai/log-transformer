package parser

import (
	"testing"

	"github.com/nxadm/tail"
	"github.com/stretchr/testify/assert"
	"github.com/yiranzai/log-transformer/conf"
)

func TestParseVarNames(t *testing.T) {
	names := parseVarNames("$remote_addr - $remote_user [$time_local] \"$request\"  $status $body_bytes_sent" +
		" \"$http_referer\" \"$http_user_agent\" \"$http_x_forwarded_for\"")
	assert.True(t, len(names) > 0)
}

func TestGetVarRegex(t *testing.T) {
	regex := getVarRegex(
		"$remote_addr - $remote_user [$time_local] \"$request\" $status $body_bytes_sent \"$http_referer\" \"$http_user_agent\" \"$http_x_forwarded_for\"",
	)
	text := parseText(
		regex,
		"127.0.0.1 - - [06/Aug/2021:17:05:32 +0800] \"GET /debug/default/toolbar?tag=610cfb5b6dd87 HTTP/1.1\" 404 171 \"https://wechatapp.futunn.com/news/flash/8499372\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36\" \"-\"",
	)
	assert.True(t, len(text) > 0)
}

func TestParse(t *testing.T) {
	cfg := conf.Init("../conf/config.yaml")
	ng := NewParser(cfg.Parser)
	parse, _ := ng.Parse(tail.NewLine("127.0.0.1 - - [06/Aug/2021:17:05:32 +0800] \"GET /debug/default/toolbar?tag"+
		"=610cfb5b6dd87 HTTP/1.1\" 404 171 \"https://wechatapp.futunn.com/news/flash/8499372\" \"Mozilla/5.0 ("+
		"Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537."+
		"36\" \"-\"", 1))
	assert.True(t, len(parse) > 0)
}

func Test_parseFile(t *testing.T) {
	cfg := conf.Init("../conf/config.yaml")
	ng := NewParser(cfg.Parser)
	file := "/data/var/nginx/log/abc.access.log"
	s, s2 := ng.ParseFile(file)
	assert.Equal(t, s, "abc")
	assert.Equal(t, s2, "abc.access.log")

	file = "/data/var/nginx/log/abc___access.log"
	s, s2 = ng.ParseFile(file)
	assert.Equal(t, s, "abc")
	assert.Equal(t, s2, "abc___access.log")
}

func Test_parseDirOrFile(t *testing.T) {
	cfg := conf.Init("../conf/config.yaml")
	ng := NewParser(cfg.Parser)
	file := "/Users/yiranzai/work/logs/nginx"
	result := ng.ParseDirOrFile(file)

	assert.True(t, len(result) > 0)
}
