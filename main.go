package main

import (
	"net/http"

	client "_/C_/Users/ThinkPad/Desktop/Go_React.js"

	"github.com/buynaa0314/Go_React.js/pkg/websocket"
)

func serverWS(pool *websocket.Pool, w http.ResponseWriter, r *httpRequest) {
	fmt.Println("websocket endpoint reached")

	conn, err :=websocket.Upgrade(w,r)

	if err!=nil{
		fmt.Fprint(w, "%+V\n", err)

	}

	client :=&websocket.client{
		Conn:conn,
		Pool:pool,
	}
	pool.Register <- client
	client.Read()
}
func setupRoutes(){
	pool:= websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func( w http.ResponseWriter, r *http.Request){
		 serveWS{pool, w, r}

	})
	func main(){
		fmt.Println("Buynaa's full stack chat project")
		setupRoutes()
		http.listenAndServe(" :9000", nil)
	}
