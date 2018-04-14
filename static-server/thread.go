package main

type ThreadTaskFunc func()

type ThreadTask struct {
	f ThreadTaskFunc
}

func NewThreadTask(f ThreadTaskFunc) *ThreadTask {
	return &ThreadTask{f}
}

func (task ThreadTask) Run() {
	task.f()
}