#!/bin/bash

export ENV_ADDR_PORT=127.0.0.1:17148

./cmd/ws2discard/ws2discard &
pid=$!
ps -o ppid -o pid -o args -p $pid
echo

node --eval '
  const pws = new Promise(res => {
		const ws = new WebSocket("ws://127.0.0.1:17148/discard");
		console.info("trying to connect...");
		ws.addEventListener("open", _ => res(ws), true);
	});

  pws.then(ws => {
		ws.send("helo");
		ws.close();
		return "ok"
	})
  .then(console.info)
	.catch(console.error)
	.finally(_ => console.info("done"))
  ;
'

sleep 1
kill $pid
