package runtime_util

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
}

/**
 * 定时任务
 */
func TestTimer(t *testing.T) {
	ticker := time.Tick(time.Second)
	for {
		select {
		case <-ticker:
			fmt.Println("定时任务1s")
		}
	}
}

/**
 * 生产者消费者测试
 */
func TestWork(t *testing.T) {
	taskCh := make(chan int, 100)
	Work(taskCh)
	for i := 0; i < 10; i++ {
		taskCh <- i
		time.Sleep(1 * time.Second)
	}
	select {
	case <-time.After(time.Hour):
	}
}

func Work(taskCh <-chan int) {
	//启用5个消费者
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("消费者{%d},正在取任务{%d}\n", id, task)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

/**
 * 控制并发数
 */
func TestConcurrencyNumControl(t *testing.T) {
	//控制并发数目为3
	var token = make(chan int, 3)
	for i := 0; i < 100; i++ {
		go func() {
			token <- 1
			fmt.Println("正在工作")
			time.Sleep(time.Second)
			<-token
		}()
	}
	time.Sleep(time.Hour)
}
