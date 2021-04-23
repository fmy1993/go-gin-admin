package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// ini包是定义好了一个File结构体，处理一个或者多个ini配置文件
	/*
		File represents a combination of one or more INI files in memory.
				type File struct {
			    options     LoadOptions
			    dataSources []dataSource

			    // Should make things safe, but sometimes doesn't matter.
			    BlockMode bool
			    lock      sync.RWMutex // 加了读写锁

			    // To keep data in order.
			    sectionList []string
			    // To keep track of the index of a section with same name.
			    // This meta list is only used with non-unique section names are allowed.
			    sectionIndexes []int

			    // Actual data is stored here.
			    sections map[string][]*Section
				// 实际上数据以map[string][]的形式存储
			    NameMapper
			    ValueMapper
			}
	*/
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

// 处理app.ini文件的内容，根据ini文件得到一个定义好的配置文件类(有点spring beanfactory的意思)，方便别的地方可以直接读取这个类的值
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini") // 这里Cfg已经有了所有的信息
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	//配置文件主要分成三段，根据这三段做一个各写一个函数，方便检查和修改
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	//											这个似乎是一个正则表达式
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
