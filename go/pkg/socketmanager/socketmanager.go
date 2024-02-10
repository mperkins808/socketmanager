package socketmanager

import (
	"sync"
)

type SocketManager interface {
	IsActive(id string) bool
	UpdateDue(id string) bool
	SetUpdateDue(id string, due bool)
	Remove(id string)
}

type SocketConnection struct {
	SocketID  string
	UpdateDue bool
}

type SimpleSocketManager struct {
	activeSockets map[string]SocketConnection
	mu            sync.Mutex
}

// creates an instance of SimpleSocketManager
func NewSimpleSocketManager() *SimpleSocketManager {
	return &SimpleSocketManager{
		activeSockets: make(map[string]SocketConnection),
	}
}

// returns based on if the socket id is found in socket manager
func (sm *SimpleSocketManager) IsActive(id string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	_, ok := sm.activeSockets[id]
	return ok
}

// add a socket to socket manager
func (sm *SimpleSocketManager) Add(id, socketID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.activeSockets[id] = SocketConnection{
		SocketID:  socketID,
		UpdateDue: false,
	}
}

// remove a socket from socket manager
func (sm *SimpleSocketManager) Remove(id string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.activeSockets, id)
}

// return all active sockets
func (sm *SimpleSocketManager) GetActiveSockets() map[string]SocketConnection {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	result := make(map[string]SocketConnection, len(sm.activeSockets))
	for k, v := range sm.activeSockets {
		result[k] = SocketConnection{
			SocketID:  v.SocketID,
			UpdateDue: v.UpdateDue,
		}
	}
	return result
}

// return a socket
func (sm *SimpleSocketManager) GetSocket(id string) (SocketConnection, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	var socket SocketConnection
	s, ok := sm.activeSockets[id]
	if !ok {
		return socket, false
	}
	return s, true
}

// sets a socket due state
func (sm *SimpleSocketManager) SetUpdateDue(id string, due bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	s, ok := sm.activeSockets[id]
	if ok {
		update := SocketConnection{
			SocketID:  s.SocketID,
			UpdateDue: due,
		}
		sm.activeSockets[id] = update
	}

}

// checks if a socket is due for update
func (sm *SimpleSocketManager) UpdateDue(id string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	s, ok := sm.activeSockets[id]
	if !ok {
		return false
	}
	return s.UpdateDue
}
