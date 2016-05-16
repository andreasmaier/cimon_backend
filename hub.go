package main

import "fmt"

type hub struct {
    connections map[*connection]bool

    broadcast chan []byte

    register chan *connection

    unregister chan *connection
}

var h = hub{
    broadcast: make(chan []byte),
    register: make(chan *connection),
    unregister: make(chan *connection),
    connections: make(map[*connection]bool),
}

func (h *hub) run() {
    for {
        select {
        case c := <- h.register:
            fmt.Println("Register received")
            h.connections[c] = true
        case c := <- h.unregister:
            fmt.Println("Unregister received")
            if _, ok := h.connections[c]; ok {
                delete(h.connections, c)
                close(c.send)
            }
        case m := <- h.broadcast:
            fmt.Println("Broadcast received")

            for c := range h.connections {
                select {
                case c.send <- m:
                default:
                    close(c.send)
                    delete(h.connections, c)
                }
            }
        }
    }
}
