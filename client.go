package main

import(
  "github.com/gorilla/websocket"
)

// 各ユーザを示す
type client struct{
  socket *websocket.Conn // WebSocket
  send chan []byte // メッセージが送られるチャネル
  room *room // 参加しているチャットルーム
}
