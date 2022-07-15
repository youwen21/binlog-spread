package store

import (
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"binlog_spread/conf"
	"fmt"
	"sync"
	"time"
)

var (
	// 逻辑中使用的某个变量
	eventStreamData           = make([]models.ApiBinlog2, 0, 200)
	eventStreamLastUpdateTime time.Time
	// 与变量对应的使用互斥锁
	dataGuard sync.Mutex

	defaultMaxRows          = 200
	durationThreshold int64 = 300
)

func init() {
	if conf.Config.ModelStreamFlushRows != 0 {
		defaultMaxRows = conf.Config.ModelStreamFlushRows
	}
	if conf.Config.ModelStreamFlushDuration != 0 {
		durationThreshold = conf.Config.ModelStreamFlushDuration
	}
}

func streamNeedUpdate() bool {
	if len(eventStreamData) >= defaultMaxRows {
		return true
	}

	if eventStreamLastUpdateTime.IsZero() {
		return false
	}

	if time.Now().Unix()-eventStreamLastUpdateTime.Unix() > durationThreshold {
		return true
	}

	return false
}

func StreamAddRows(ss []models.ApiBinlog2) {
	streamAddToMem(ss)

	if !streamNeedUpdate() {
		return
	}

	ok, _ := streamStore()
	if ok {
		eventStreamLastUpdateTime = time.Now()
	}
}

func streamAddToMem(ss []models.ApiBinlog2) {
	// 锁定
	dataGuard.Lock()
	// 在函数退出时解除锁定
	defer dataGuard.Unlock()

	eventStreamData = append(eventStreamData, ss...)
}

func streamStore() (bool, error) {
	dataGuard.Lock()
	defer dataGuard.Unlock()

	session := comps.GetSession()
	session.Table("api_binlog").Create(eventStreamData)

	err := session.Error
	fmt.Println(err)

	//eventStreamData = make([]models.DddEventStream, 200)
	eventStreamData = eventStreamData[0:0]

	return true, nil
}
