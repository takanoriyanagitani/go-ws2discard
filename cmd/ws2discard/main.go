package main

import (
	"errors"
	"io"
	"iter"
	"log"
	"net/http"
	"os"

	wd "github.com/takanoriyanagitani/go-ws2discard"
	ws "golang.org/x/net/websocket"
)

func DiscardHandler(wc *ws.Conn) {
	log.Printf("handle start.\n")

	wcon := wd.WsConn{Conn: wc}
	defer wcon.Close()

	var msgs iter.Seq2[[]byte, error] = wcon.ToBytesIter()

	var cnt uint64 = 0
	var siz uint64 = 0

	for msg, err := range msgs {
		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Printf("%v\n", err)
			return
		}

		cnt += 1
		siz += uint64(len(msg))
	}

	log.Printf("msg cnt: %v\n", cnt)
	log.Printf("msgsize: %v\n", siz)
}

func main() {
	server := ws.Server{
		Config:    ws.Config{},
		Handshake: func(_ *ws.Config, _ *http.Request) error { return nil },
		Handler:   ws.Handler(DiscardHandler),
	}

	var addrPort string = os.Getenv("ENV_ADDR_PORT")
	log.Printf("ENV_ADDR_PORT: %s\n", addrPort)

	http.Handle("/discard", server)
	err := http.ListenAndServe(addrPort, nil)
	if nil != err {
		log.Printf("%v\n", err)
	}
}
