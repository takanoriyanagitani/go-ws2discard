package ws2discard

import (
	"iter"

	ws "golang.org/x/net/websocket"
)

type WsConn struct{ *ws.Conn }

type Buffer struct{ raw []byte }

func (c WsConn) Close() error { return c.Conn.Close() }

func (c WsConn) RecvBytes(dst *Buffer) error {
	return ws.Message.Receive(c.Conn, &dst.raw)
}

func (c WsConn) ToBytesIter() iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		var buf Buffer
		for {
			err := c.RecvBytes(&buf)
			if !yield(buf.raw, err) {
				return
			}
		}
	}
}
