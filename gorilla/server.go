package main

import (
        "flag"
        "log"
        "net/http"
        //"time"
        "github.com/gorilla/websocket"
        "syscall"
        "strings"
)

var addr = flag.String("addr", "ps2:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
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
                //log.Printf("recv: %s", message)
                for i:=0;i<10;i++ {

                        err = c.WriteMessage(mt, message)
                        if err != nil {
                                log.Println("write:", err)
                                err = c.WriteMessage(mt, []byte(strings.Repeat("0",1000)))

                                break
                        }
                }
        }
}
func main() {
        // Increase resources limitations
        var rLimit syscall.Rlimit
        if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
                panic(err)
        }
        rLimit.Cur = rLimit.Max
        if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
                panic(err)
        }
        flag.Parse()
        log.SetFlags(0)
        http.HandleFunc("/echo", echo)
        log.Fatal(http.ListenAndServe(*addr, nil))
}
