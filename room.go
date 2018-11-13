package main

type room struct{
  forward chan []byte // 他のクライアントに転送するメッセージを保持する
}
