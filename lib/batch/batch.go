package batch

import (
	"context"
	"golang.org/x/sync/errgroup"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	users := make(chan user, n)
	errG, _ := errgroup.WithContext(context.Background())
	errG.SetLimit(int(pool))
	for i := 0; i < int(n); i++ {
		z := i
		errG.Go(func() error {
			users <- getOne(int64(z))
			return nil
		})
	}
	errG.Wait()

	var r []user
	for i := 0; i < int(n); i++ {
		x := <-users
		r = append(r, x)
	}

	return r
}
