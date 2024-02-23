package socketmanager

import (
	"context"
	"fmt"
	"sync"
)

type SocketManagerID string

const (
	SOCKETMANAGER SocketManagerID = "socketmanager"
)

type SocketManager interface {
	IsActive(id string) bool
	UpdateDue(id string) bool
	SetUpdateDue(id string, due bool)
	Remove(id string)
	GetArbs(id string) map[string]interface{}
	GetArb(id string, key string) interface{}
	SetArb(id string, key string, arb interface{}) error
}

type SocketConnection struct {
	SocketID  string
	UpdateDue bool
	Arbs      map[string]interface{}
}

type ArbResult struct {
	Value interface{}
	Err   error
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

// Add the socket manager to context
func (sm *SimpleSocketManager) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, SOCKETMANAGER, sm)
}

// Get socket manager from context
func GetSocketManagerFromContext(ctx context.Context) (*SimpleSocketManager, error) {
	sm := ctx.Value(SOCKETMANAGER)
	switch t := sm.(type) {
	case *SimpleSocketManager:
		return t, nil
	}
	return nil, fmt.Errorf("socket manager is not in context")
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

// returns the map of interfaces attached to a socket
func (sm *SimpleSocketManager) GetArbs(id string) (map[string]interface{}, error) {

	s, active := sm.GetSocket(id)
	if !active {
		return nil, fmt.Errorf("socket %v is not active", id)
	}
	return s.Arbs, nil
}

// returns an interface attached to a socket
func (sm *SimpleSocketManager) GetArb(id string, key string) ArbResult {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	socket, ok := sm.activeSockets[id]
	if !ok {
		return ArbResult{Err: fmt.Errorf("socket %v is not active", id)}
	}

	arb, ok := socket.Arbs[key]
	if !ok {
		return ArbResult{Err: fmt.Errorf("%v not found for socket %v", key, id)}
	}

	return ArbResult{Value: arb}
}

// sets arbitrary data for a socket
func (sm *SimpleSocketManager) SetArb(id string, key string, arb interface{}) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	socket, ok := sm.activeSockets[id]
	if !ok {
		return fmt.Errorf("socket %v is not active", id)
	}

	// If nil then instantiate it
	if socket.Arbs == nil {
		socket.Arbs = make(map[string]interface{})
	}

	socket.Arbs[key] = arb
	sm.activeSockets[id] = socket
	return nil
}
