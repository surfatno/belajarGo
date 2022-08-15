package Section4

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait()
	fmt.Println("Selesai")
}

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter", counter)
}

func TestPool(t *testing.T) {
	var pool sync.Pool
	//pool:= sync.Pool{}
	pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value) // untuk menyimpan data ke map
}
func TestSyncMap(t *testing.T) {
	var data sync.Map
	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(&data, i, group)
	}
	group.Wait()

	data.Range(func(key, value interface{}) bool { //digunakan untuk melakukan iterasi seluruh data di map
		fmt.Println(key, ":", value)
		return true
	})
	//delete(key) untuk menghapus data di map menggunakan key
	//load(key) untuk mengambil data dari map menggunakan key
}

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()
	group.Wait()
}
