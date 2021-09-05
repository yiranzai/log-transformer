package transformer

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-module/carbon"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
	"github.com/yiranzai/log-transformer/utils"
)

const TransformerLevelError = "error"
const TransformerLevelWarn = "warn"
const TransformerLevelInfo = "info"
const TransformerLevelDebug = "debug"

const TransformerTypeAccess = "access"       // 访问日志
const TransformerTypeWater = "water"         // 业务日志
const TransformerTypeFramework = "framework" // 框架日志

const TransformerTypeCommon = "common"

//ITransformer 转化工具
type ITransformer interface {
	//Transform 转化内容输出
	Transform(fields map[string]interface{}, time *carbon.Carbon)
	//Cleanup 清理文件监听
	Cleanup()
}

//NewTransformer 实例化转化工具
func NewTransformer(cfg ITransformerConfig, t *tail.Tail, svrName string) ITransformer {
	var transformer ITransformer
	switch cfg.GetType() {
	case TransformerTypeCommon:
		transformer = NewCommonTransformer(cfg, t, svrName)
	default:
		panic(fmt.Sprintf("transformer type %s not support", cfg.GetType()))
	}
	return transformer
}

//ITransformerConfig 转化配置
type ITransformerConfig interface {
	GetType() string
	GetMaxFiles() uint
	GetWritePath() string
	GetExtension() string
}

//CommonTransformer 通用转化工具(转成了json)
type CommonTransformer struct {
	cfg     ITransformerConfig
	Tail    *tail.Tail
	svrName string
	Logger  *logrus.Logger
}

//NewCommonTransformer 通用转化工具构造函数
func NewCommonTransformer(cfg ITransformerConfig, t *tail.Tail, svrName string) *CommonTransformer {
	transformer := &CommonTransformer{cfg: cfg, Tail: t, svrName: svrName}
	path := transformer.TransformerPath(svrName)
	//@see https://github.com/lestrrat-go/strftime
	rotateLogs, err := rotatelogs.New(
		path+cfg.GetExtension()+".%F",
		rotatelogs.WithLinkName(path+cfg.GetExtension()+""),
		rotatelogs.WithRotationCount(cfg.GetMaxFiles()),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Printf("failed to create rotatelogs: %s, svrName: %s", err, svrName)
		return nil
	}
	logger := logrus.New()
	logger.SetOutput(rotateLogs)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	transformer.Logger = logger
	return transformer
}

//Transform 转化内容输出
func (t CommonTransformer) Transform(fields map[string]interface{}, time *carbon.Carbon) {
	ip, _ := utils.LocalIp()
	fields["_service_name_"] = t.svrName
	fields["_level_"] = TransformerLevelInfo
	fields["_type_"] = TransformerTypeAccess
	fields["_local_ip_"] = ip

	fields["_datetime_"] = time.ToIso8601String()
	t.Logger.WithFields(fields).Info("")
}

//Cleanup 清理文件监听
func (t CommonTransformer) Cleanup() {
	t.Tail.Cleanup()
}

//TransformerPath 获取Transformer日志的路径
func (t *CommonTransformer) TransformerPath(s string) string {
	return t.cfg.GetWritePath() + s + utils.DirSplitChar + s + ".nginx." +
		"access."
}
