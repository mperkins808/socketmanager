# socketmanager

Manage your socket connections across functions through context.

## Seting up the socket manager

```go
package main

import (
	"context"

	"github.com/mperkins808/socketmanager/go/pkg/socketmanager"
)

func main() {
	// initialising the socket manager
	sm := socketmanager.NewSimpleSocketManager()

	// OR

	// initialising context with socket manager attached
	ctx := socketmanager.NewSimpleSocketManager().WithContext(context.Background())

}

```

## Using within a socket handler function

Say you had a websocket that is connecting to `ws://localhost:3000/ws` and you also have a background task that needs to determine the number of active sockets connected to that endpoint.

```go
func main() {
	...

	// initialising the socket manager
	sm := socketmanager.NewSimpleSocketManager()

	// Adding socket manager to context
	ctx := sm.WithContext(context.Background())

	// Some background task that needs to know the active sockets
	go func(sm *socketmanager.SimpleSocketManager) {
		for {
			active := sm.GetActiveSockets()
			fmt.Println(active)
			time.Sleep(time.Second)
		}
	}(sm)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error(err)
			return
		}

		sID := uuid.New().String()
		// add the socket id to socket manager
		sm.Add(sID, sID)

		// remove the socket id if socket is closed
		defer sm.Remove(sID)
		defer conn.Close()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}

			time.Sleep(time.Second)
		}
		})

}
```

## Using context you dont need to parse socket manager as an argument. Perfect for Http handlers

```go
// your handler
func RequestGetData(w http.ResponseWriter, r *http.Request) {
	sm, err := socketmanager.GetSocketManagerFromContext(r.Context())
	if err != nil {
		// handle error
	}

}

// using the handler
func main() {
	// creates socket manager and adds to context
	ctx := socketmanager.NewSimpleSocketManager().WithContext(context.Background())

	mux.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		// socket manager can now be used within the function
		RequestGetData(w, r.WithContext(ctx))
	})

	// Some background task that also needs socket manager
	go func (ctx context.Context){
		for {
			smC, err := socketmanager.GetSocketManagerFromContext(ctx)
			if err != nil {
				// handle error
			}
			time.Sleep(time.Second)
		}
	}(ctx)

}
```
