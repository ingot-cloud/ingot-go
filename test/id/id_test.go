package id

import (
	"sync"
	"testing"

	"github.com/ingot-cloud/ingot-go/pkg/component/id/snowflake"
)

func TestID(t *testing.T) {
	var wg sync.WaitGroup
	generator := &snowflake.Generator{
		WorkID: 1,
	}

	count := 1_000_000
	ch := make(chan int64, count)
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup, out chan<- int64) {
			id, _ := generator.NextID()
			out <- id
			wg.Done()
		}(&wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	m := make(map[int64]int)
	for id := range ch {
		// 判断是否有重复ID
		_, ok := m[id]
		if ok {
			t.Logf("发现重复id=%d", id)
			return
		}

		m[id] = 1
	}

	t.Logf("所有ID生成完成，共计%d个", len(m))
}

func Test2(t *testing.T) {

}
