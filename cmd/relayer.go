package main

import "log"

type MessageRelayer struct {
	outgoing chan []byte
	incoming chan []byte
	done     chan struct{}
}

func (m MessageRelayer) relay() {
	for {
		select {
		case data, ok := <-m.incoming:
			if !ok {
				log.Print("Incoming channel closed")
				m.done <- struct{}{}
			} else {
				log.Printf("Received msg on incoming channel:\n %s...\n", data[:200])
				log.Print("Forwarding msg to outgoing channel")
				m.outgoing <- data

			}
		}

	}
}
