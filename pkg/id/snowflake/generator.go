package snowflake

import (
	"math/rand"
	"sync"
	"time"

	"github.com/ingot-cloud/ingot-go/internal/app/support/errors"
)

const (
	// 2017-01-18 23:46:01, 起始时间戳，用于用当前时间戳减去这个时间戳，算出偏移量。41位时间位可以使用69年。
	// 起始的时间戳，可以修改为服务第一次启动的时间。
	// 若服务已经开始使用，起始时间戳就不应该改变。
	startTimestamp = int64(1484754361114)
	// workerId占用10位
	workIDBits = uint64(10)
	// 序列占用12位
	sequenceBits = uint64(12)
	// 最大workId为1023，此处 ^ 为取反
	maxWordID          = ^(int64(-1) << workIDBits)
	sequenceMask       = ^(int64(-1) << sequenceBits)
	workIDLeftShift    = int64(sequenceBits)
	timestampLeftShift = int64(sequenceBits) + int64(workIDLeftShift)
)

// Generator snowflake 实现
type Generator struct {
	// 工作ID 0-1023
	WorkID        int64
	lastTimestamp int64
	sequence      int64
	lock          sync.Mutex
}

// NextID 获取ID
func (g *Generator) NextID() (int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	timestamp := g.timeGen()

	if timestamp < g.lastTimestamp {
		offset := g.lastTimestamp - timestamp
		if offset > 5 {
			return 0, errors.ErrIDClockBack
		}
		// 若在允许时间回拨的毫秒量范围内，则允许等待2倍的偏移量后重新获取
		var sleepMillisecond int64 = offset << 1
		time.Sleep(time.Millisecond * time.Duration(sleepMillisecond))

		timestamp = g.timeGen()
		if timestamp < g.lastTimestamp {
			return 0, errors.ErrIDClockBack
		}
	}

	// 相同时间窗口
	if timestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & sequenceMask
		// seq 为0的时候表示是下一毫秒时间开始对seq做随机
		if g.sequence == 0 {
			g.sequence = int64(rand.Intn(100))
			timestamp = g.tilNextMillis(g.lastTimestamp)
		}
	} else {
		g.sequence = int64(rand.Intn(100))
	}

	g.lastTimestamp = timestamp

	return (timestamp-startTimestamp)<<timestampLeftShift | (g.WorkID << workIDLeftShift) | g.sequence, nil
}

// 当前时间毫秒
func (g *Generator) timeGen() int64 {
	return time.Now().UnixNano() / 1e6
}

func (g *Generator) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := g.timeGen()
	for {
		if timestamp > lastTimestamp {
			break
		}
		timestamp = g.timeGen()
	}
	return timestamp
}
