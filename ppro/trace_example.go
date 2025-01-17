package main

import (
    "flag"
    "log"
    "net/http"
    _ "net/http/pprof"
    "sync"
    "time"
)


func main() {
    flag.Parse()

    //这里实现了远程获取pprof数据的接口
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go work(&wg)
    }

    wg.Wait()
    // Wait to see the global run queue deplete.
    time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
    time.Sleep(time.Second)

    var counter int
    for i := 0; i < 1e10; i++ {
        time.Sleep(time.Millisecond * 100)
        counter++
    }
    wg.Done()
}