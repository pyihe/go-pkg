package monitor

import (
	"os"

	"github.com/robfig/cron/v3"
)

type myFile struct {
	c        *cron.Cron
	id       cron.EntryID
	spec     string //时间描述
	filePath string //单个配置文件的目录
	updateOn int64  //最后一次更新时间
	handle   func() //文件更新时需要直行的handler
}

func (m *myFile) init() {
	timer := cron.New(cron.WithSeconds())
	id, err := timer.AddJob(m.spec, m)
	if err != nil {
		panic(err)
	}
	m.id = id
	m.c = timer
	timer.Start()
}

func (m *myFile) Run() {
	file, err := os.Open(m.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	modifyTime := fileInfo.ModTime().Unix()
	if modifyTime > m.updateOn {
		//更新
		m.handle()
		m.updateOn = modifyTime
	}
}
