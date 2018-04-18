package main

// Функция для пула потоков.
type ThreadTaskFunc func()

// Функция для пула потоков.
type ThreadTask struct {
	f ThreadTaskFunc
}

// Создает функцию для пула потоков.
func NewThreadTask(f ThreadTaskFunc) *ThreadTask {
	return &ThreadTask{f}
}

// Запускает функцию для пула потоков на выполнение.
func (task ThreadTask) Run() {
	task.f()
}