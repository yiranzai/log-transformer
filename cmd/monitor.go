package cmd

import (
	"io"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/nxadm/tail"
	"github.com/yiranzai/log-transformer/cmd/parser"
	"github.com/yiranzai/log-transformer/cmd/transformer"
	"github.com/yiranzai/log-transformer/conf"
	"github.com/yiranzai/log-transformer/utils"
)

//IMonitor 监控
type IMonitor interface {
	Do()
	WaitAndMonitor()
}

//Monitor 监控的实例
type Monitor struct {
	// 配置文件
	cfg *conf.Config
	// 解析器
	parser parser.IParser
	// 识别的服务
	svrs map[string]transformer.ITransformer
	wg   sync.WaitGroup
}

//Run 开始执行监控日志
func Run() {
	cfg := conf.Init(cfgPath)
	M := newMonitor(cfg)
	M.Do()
	M.WaitAndMonitor()
}

//newMonitor monitor构造函数
func newMonitor(cfg *conf.Config) IMonitor {
	m := &Monitor{cfg: cfg, svrs: make(map[string]transformer.ITransformer)}
	m.parser = parser.NewParser(cfg.Parser)
	return m
}

//Do 执行监控
func (m *Monitor) Do() {
	m.add(m.cfg.Monitor.Path, false)
}

//add 新增一个监控
func (m *Monitor) add(path string, isNew bool) {
	files := m.parseDirOrFile(path)
	for s, fileNames := range files {
		log.Println(s, fileNames)
		err := utils.CreateDirNX(m.cfg.Transformer.Path + s)
		if err != nil {
			panic(err)
		}
		m.monitorServer(s, fileNames, isNew)
	}
}

//add 移除监控
func (m *Monitor) remove(path string) {
	_, file := m.parser.ParseFile(path)
	m.Remove(file)
	m.wg.Done()
}

func (m *Monitor) Transform(svrName string, line *tail.Line) {
	t, ok := m.svrs[svrName]
	if !ok {
		return
	}
	parse, carbon := m.parser.Parse(line)
	t.Transform(parse, carbon)
}

//monitorExistFile 监控一个服务的日志
func (m *Monitor) monitorExistFile(svrName string, l string) {
	t := m.AddSvr(svrName, l, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Whence: io.SeekEnd}})
	m.wg.Add(1)
	go func() {
		log.Printf("[%s]监听开始...\n", l)
		for line := range t.Lines {
			log.Printf("[%s]-%s 新的内容 %s\n", svrName, l, line.Text)
			m.Transform(svrName, line)
		}
	}()
}

//monitorNewFile 监控一个服务的日志
func (m *Monitor) monitorNewFile(svrName string, l string) {
	t := m.AddSvr(svrName, l, tail.Config{Follow: true, ReOpen: true})
	m.wg.Add(1)
	go func() {
		log.Printf("[%s]监听开始...\n", l)
		for line := range t.Lines {
			log.Printf("[%s]-%s 新的内容 %s\n", svrName, l, line.Text)
			m.Transform(svrName, line)
		}
	}()
}

//Cancel 关闭所有的文件监听
func (m *Monitor) Cancel() {
	for _, svr := range m.svrs {
		m.wg.Done()
		svr.Cleanup()
	}
}

//AddSvr 新增服务的文件监听
func (m *Monitor) AddSvr(svrName, key string, config tail.Config) *tail.Tail {
	t, _ := tail.TailFile(key, config)
	svr := transformer.NewTransformer(m.cfg.Transformer, t, svrName)
	m.svrs[svrName] = svr
	return t
}

//Remove 移除服务的文件监听
func (m *Monitor) Remove(svrName string) {
	if svr, ok := m.svrs[svrName]; ok {
		svr.Cleanup()
	}
	delete(m.svrs, svrName)
}

//monitorServer 监控一个服务的日志
func (m *Monitor) monitorServer(svrName string, logs []string, isNew bool) {
	for _, l := range logs {
		if isNew {
			m.monitorNewFile(svrName, l)
			continue
		}
		m.monitorExistFile(svrName, l)
	}
}

//parseDirOrFile 解析目录或文件
func (m *Monitor) parseDirOrFile(path string) map[string][]string {
	return m.parser.ParseDirOrFile(path)
}

//WaitAndMonitor 等待并监控目录变化
func (m *Monitor) WaitAndMonitor() {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func(watch *fsnotify.Watcher) {
		err := watch.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(watch)
	//添加要监控的对象，文件或文件夹
	err = watch.Add(m.cfg.Monitor.Path)
	if err != nil {
		log.Fatal(err)
	}
	m.wg.Add(1)
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件 : ", ev.Name)
						m.add(ev.Name, true)
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("删除文件 : ", ev.Name)
						m.remove(ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.Println("重命名文件 : ", ev.Name)
						m.remove(ev.Name)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.Println("修改权限 : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					m.wg.Done()
					m.Cancel()
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
	m.wg.Wait()
}
