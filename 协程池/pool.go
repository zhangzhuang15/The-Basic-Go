package pool

import (
	"runtime"
	"sync/atomic"
)

type Worker struct {
	dead   bool
	worker func(task interface{})
}

type Pool struct {
	// 任务队列
	tasks chan interface{}
	// 禁止加入新任务
	seal bool
	// 工作者列表
	workers []*Worker
	// 工作者最大数目
	maxWorkers int32
	// 任务队列中元素的最大数目
	maxTasks int32
	// 工作者数量
	workerSize int32
	// 任务数量
	taskSize int32
}

func New(maxWorkers, maxTasks int32) *Pool {
	return &Pool{
		taskSize:   0,
		workerSize: 0,
		seal:       false,
		maxWorkers: maxWorkers,
		maxTasks:   maxTasks,
		tasks:      make(chan interface{}, maxTasks),
		workers:    make([]*Worker, maxWorkers),
	}
}

func (pool *Pool) AddTask(task interface{}) bool {
	if pool.taskSize >= pool.maxTasks || pool.seal {
		return false
	}
	atomic.AddInt32(&pool.taskSize, 1)
	pool.tasks <- task
	return true
}

func (pool *Pool) Run(workerFunc func(interface{})) {
	num := int32(pool.taskSize / 2)
	if num > pool.maxWorkers {
		num = pool.maxWorkers
	}

	for i := int32(0); i < num; i++ {
		worker := &Worker{
			dead:   false,
			worker: workerFunc,
		}
		pool.workers[i] = worker
		go func() {
			for !worker.dead {
				task := <-pool.tasks
				if worker.dead {
					break
				}
				atomic.AddInt32(&pool.taskSize, -1)
				worker.worker(task)
			}
			atomic.AddInt32(&pool.workerSize, -1)
		}()
	}

	pool.workerSize = num
}

// Exit 进入加入新的任务，以阻塞的方式直到所有任务执行完毕
func (pool *Pool) Exit() {
	pool.seal = true
	for atomic.LoadInt32(&pool.taskSize) != 0 {
		runtime.Gosched()
	}
	workerSize := pool.workerSize
	for i := int32(0); i < workerSize; i++ {
		pool.workers[i].dead = true
	}
	close(pool.tasks)
	for atomic.LoadInt32(&pool.workerSize) != 0 {
	}
	pool.seal = false
}
