package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	//"time"
	"strings"
	"syscall"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "ps2:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			// log.Println("read:", err) // too noisy for closed connection
			break
		}
		log.Printf("recv: %s", message)
		for i := 0; i < 10; i++ {

			err = c.WriteMessage(mt, []byte(strings.Repeat("0", 10000)))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
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

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
