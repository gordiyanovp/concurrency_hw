package batch

import (
	"sync"
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
	ch := make(chan struct{}, pool)
	var wg sync.WaitGroup
	// var mu sync.Mutex
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			ch <- struct{}{}
			u := getOne(i)
			// mu.Lock()
			res = append(res, u)
			// mu.Unlock()
			<-ch
		}(int64(i))
	}
	wg.Wait()
	return res
}
