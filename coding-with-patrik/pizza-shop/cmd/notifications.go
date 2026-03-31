package main

import "sync"

type NotificationManager struct {
	/*
	   "order1234": {
	     0xc0000a4000: true,
	     0xc0000a4060: true,
	   },
	 */
	clients map[string]map[chan string]bool
	mtx sync.RWMutex
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		clients: make(map[string]map[chan string]bool),
	}
}

func (this *NotificationManager) AddClient(key string, client chan string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()

	if this.clients[key] == nil {
		this.clients[key] = make(map[chan string]bool)
	}
	this.clients[key][client] = true
}

func (this *NotificationManager) RemoveClient(key string, client chan string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()

	var clients = this.clients[key]
	if clients != nil {
		delete(clients, client)
		if len(clients) == 0 {
			delete(this.clients, key)
		}
	}
	close(client)
}

func (this *NotificationManager) Notify(key string, message string) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()

	for client := range this.clients[key] {
		select {
		case client <- message:
		default:
		}
	}
}
