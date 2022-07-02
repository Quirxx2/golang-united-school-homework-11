package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(u chan user, task <-chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		id, ok := <-task
		if !ok {
			return
		}
		time.Sleep(time.Millisecond * 100)
		u <- user{ID: id}
	}
}

func getBatch(n int64, pool int64) (res []user) {

	users := make(chan user, n)
	var wg sync.WaitGroup
	wg.Add(int(pool))

	tasks := make(chan int64, pool)

	var i int64
	for i = 1; i <= pool; i++ {
		go getOne(users, tasks, &wg)
	}

	var j int64
	for j = 0; j < n; j++ {
		tasks <- j
	}

	close(tasks)

	wg.Wait()

	var y int64
	var r []user
	for y = 1; y <= n; y++ {
		x := <-users
		r = append(r, x)
	}

	close(users)

	return r
}
