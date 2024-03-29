package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "ps2:8080", "http service address")

func main() {
	if runtime.GOOS == "linux" {
		// Increase resources limitations
		var rLimit syscall.Rlimit
		if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}
		rLimit.Cur = rLimit.Max
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			panic(err)
		}
	}
	var wg sync.WaitGroup
	var v, _ = strconv.Atoi(os.Args[1])
	for k := 0; k < v; k++ {
		wg.Add(1)
		go goZilla(&wg, k)
	}
	t1 := time.Now()
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t1.String())
	fmt.Println(t2.String())

}

func goZilla(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %v: Started\n", id)
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	//log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	received := 0
	enough, _ := strconv.Atoi(os.Args[2])

	go func() {
		defer close(done)
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			received++
			if received%50 == 0 {
				//log.Printf("recv: %s", message)
				log.Printf("thread:" + strconv.Itoa(id) + ",received:" + strconv.Itoa(received))
			}
			if received >= enough {
				fmt.Printf("Worker %v: Finished\n", id)
				return
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			fmt.Printf("Worker %v: Finished\n", id)
			return
		}
	}

}
