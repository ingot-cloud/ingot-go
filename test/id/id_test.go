package id

import (
	"sync"
	"testing"

	"github.com/ingot-cloud/ingot-go/pkg/id/snowflake"
)

func TestID(t *testing.T) {
	var wg sync.WaitGroup
	generator := &snowflake.Generator{
		WorkID: 1,
	}

	count := 1_000
	ch := make(chan int64, count)
	wg.Add(count)
	defer close(ch)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			id, _ := generator.NextID()
			ch <- id
		}()
	}
	wg.Wait()

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 判断是否有重复ID
		_, ok := m[id]
		if ok {
			t.Logf("发现重复id=%d", id)
			return
		}

		m[id] = i
	}

	t.Logf("所有ID生成完成，共计%d个", len(m))
}
