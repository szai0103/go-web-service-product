package product

import (
	"fmt"
	"github.com/schwarz/inventoryservice/product/data"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func ProductSocket(ws *websocket.Conn) {
	// we can verify that the origin is an allowed origin
	fmt.Printf("origin: %s\n", ws.Config().Origin)

	done := make(chan struct{})
	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("recieved message %s\n", msg.Data)
		}
		close(done)
	}(ws)

loop:
	for {
		select {
		case <-done:
			fmt.Println("connection was closed, lets break out of here")
			break loop
		default:
			fmt.Println("sending top 3 product list to the client")
			products, err := data.GetTopThreeProducts()
			if err != nil {
				log.Fatal(err)
			}
			if err := websocket.JSON.Send(ws, products); err != nil {
				log.Println(err)
				break
			}
			// pause for 10 seconds before sending again
			time.Sleep(10 * time.Second)
		}
	}
	fmt.Println("closing the connection")
	defer ws.Close()
}
