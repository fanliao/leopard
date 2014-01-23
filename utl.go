package orm

import (
    "reflect"
)

//Future代表一个异步任务
type Future struct {
	chDone    chan int
	chTrigger chan int
	chFail    chan int
}

//Get函数将一直阻塞直到任务完成
func (this Future) Get() int {
	return <-this.chTrigger
}

func (this Future) Reslove(v int) {
	this.chDone <- v
}

func (this Future) start() {
	i := <-this.chDone
	callback()
	this.chTrigger <- i
	fmt.Println("is received")
}

func newFuture() *Future {
	f := &Future{make(chan int, 1), make(chan int, 1), make(chan int)}
	return f
}
func callback() {
	fmt.Println("callback")
}

func task() *Future {
	f := newFuture()
	go func() {
		time.Sleep(1 * time.Second)
		f.Reslove(10)
		fmt.Println("send done")
	}()
	go func() {
		f.start()
	}()
	fmt.Println("end start")
	return f
}
func testChan() {
	f := task()

	fmt.Println("begin receive")
	time.Sleep(2 * time.Second)
	fmt.Println("receive", f.Get())
}

//// BKDR Hash Function
//unsigned int BKDRHash(char *str)
//{
//    unsigned int seed = 131; // 31 131 1313 13131 131313 etc..
//    unsigned int hash = 0;

//    while (*str)
//    {
//        hash = hash * seed + (*str++);
//    }

//    return (hash & 0x7FFFFFFF);
//}

func BKDRHash(s string) uint {
	var seed uint32 = 131
	var hash uint32 = 0

	cs := []int(s)
	for _, c = range cs {
		hash = hash*seed + c
	}

	return (hash & 0x7FFFFFFF)
}

//如果对象是一个pointer,则返回被指向对象的类型，如果是slice，返回slice元素的类型，否则出错
func GetStructType(obj interface{}) reflect.Type, error {
    return nil, nil
}

