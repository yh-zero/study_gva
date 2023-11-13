package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	AddTaskByFunc(taskNane, spec string, task func(), option ...cron.Option) (cron.EntryID, error)
}

// timer 定时任务管理
type timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

// 新建定时器任务
func NewTimerTask() Timer {
	return &timer{taskList: make(map[string]*cron.Cron)}
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer) AddTaskByFunc(taskNane, spec string, task func(), option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskNane]; !ok {
		t.taskList[taskNane] = cron.New(option...)
	}
	id, err := t.taskList[taskNane].AddFunc(spec, task)
	t.taskList[taskNane].Start()
	return id, err
}
