package _chan

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestPush 通过 chan 控制发送速率
func TestPushSpeedControl(t *testing.T) {
	pushChan := make(chan int, 3)
	pushClose := make(chan int, 1)
	controlChan := make(chan time.Duration, 1)

	pushTicker := time.NewTicker(time.Millisecond * 1000) // 发送速率
	println("当前速率", time.Millisecond*1000)
	closeAfter := time.After(time.Second * 10) // 到时间关闭
	changeAfter := time.After(time.Second * 5) // 到时间修改速率
	changeSpeed := time.Millisecond * 500      // 之后修改的速率

	for {
		select {
		case <-pushTicker.C: // 发送数据
			a := rand.Int()
			println("发送数据：", a)
			pushChan <- a
		case data := <-pushChan: // 接收数据
			println("接收数据：", data)
		case <-controlChan: // 接收速率变更
			pushTicker.Reset(changeSpeed)
			println("接收 速率控制：", changeSpeed)
		case <-changeAfter: // 5秒后开始变更发送速率
			controlChan <- changeSpeed
			println("发送 速率控制：", changeSpeed)
		case <-closeAfter: // 10秒后结束发送，并回收所有chan
			pushClose <- 1
			println("发送 push 关闭信号")
		case <-pushClose: // 关闭信号
			pushTicker.Stop()
			println("time.ticker 关闭")
			close(controlChan)
			println("controlChan close")
			close(pushChan)
			println("pushChan close")
			close(pushClose)
			goto END
		}
	}

END:
	assert.True(t, true)
	return
}
