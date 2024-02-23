package socketmanager_test

import (
	"context"
	"testing"

	"github.com/mperkins808/socketmanager/go/pkg/socketmanager"
)

func TestIsActive(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()

	sm.Add("foo", "bar")

	exists := sm.IsActive("foo")

	if !exists {
		t.Error("socket foo. Want true. Was false")
	}

	exists = sm.IsActive("fizz")
	if exists {
		t.Error("socket fizz. Want false. Was true")
	}
}

func TestRemove(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	sm.Add("foo", "bar")

	sm.Remove("foo")

	exists := sm.IsActive("foo")

	if exists {
		t.Error("socket foo. Want false. Was true")
	}

}

func TestSetUpdateDue(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	sm.Add("foo", "bar")

	sm.SetUpdateDue("foo", true)

	active, ok := sm.GetSocket("foo")
	if !ok {
		t.Error("socket foo. Want found. Was not found")
	}

	if !active.UpdateDue {
		t.Error("socket foo active. Want true. Was false")
	}

	sm.SetUpdateDue("foo", false)
	active, _ = sm.GetSocket("foo")

	if active.UpdateDue {
		t.Error("socket foo active. Want false. Was true")
	}
}

func TestGetSocket(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	sm.Add("foo", "bar")

	s, ok := sm.GetSocket("foo")
	if !ok {
		t.Error("socket foo GetSocket. Want true. Was false")
	}

	if s.SocketID != "bar" {
		t.Error("socketid id. want bar. was " + s.SocketID)
	}

	s, ok = sm.GetSocket("fizz")
	if ok {
		t.Error("socket fiz GetSocket. Want false. Was true")
	}

}

func TestGetActiveSockets(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	sm.Add("foo", "bar")
	sm.Add("fizz", "bar")

	sockets := sm.GetActiveSockets()
	if len(sockets) != 2 {
		t.Errorf("socketmanager active sockets length. Want 2. Was %v", len(sockets))
	}

	if sockets["foo"].SocketID != "bar" {
		t.Errorf("socket foo. socket id want bar. Was %s", sockets["foo"].SocketID)
	}
}

func TestSocketContext(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	ctx := sm.WithContext(context.Background())
	sm.Add("foo", "bar")
	smC, err := socketmanager.GetSocketManagerFromContext(ctx)
	if err != nil {
		t.Errorf("%v", err)
	}

	s, ok := smC.GetSocket("foo")
	if !ok {
		t.Error("getting socket manager from context. socket foo was not found")
	}

	if s.SocketID != "bar" {
		t.Errorf("socket id. want bar. was %v", s.SocketID)
	}

	ctx = context.Background()
	_, err = socketmanager.GetSocketManagerFromContext(ctx)
	if err == nil || err.Error() != "socket manager is not in context" {
		t.Errorf("error was %v. want socket manager is not in context", err.Error())
	}
}

func TestSocketArbs(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	sm.Add(id, "bar")

	key := "fizz"
	data := "string string"

	err := sm.SetArb(id, key, data)
	if err != nil {
		t.Fatalf("error was %v", err)
	}

	arb, err := sm.GetArb(id, key).String()
	if err != nil {
		t.Fatalf("error was %v", err)
	}

	if arb != data {
		t.Errorf("arbitrary data was not the same as when it was set. input %v. output %v", data, arb)

	}

	_, err = sm.GetArb(id, "somebad key").String()
	if err == nil {
		t.Errorf("getting an arb for a invalid key did not error")
	}

	_, err = sm.GetArbs("matty p")
	if err == nil {
		t.Errorf("getting an arb for a invalid id did not error")
	}

}

func TestMultipleSocketsArbs(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	id2 := "bud"
	sm.Add(id, "bar")
	sm.Add(id2, "biz")

	key := "fizz"
	data := "string string"
	data2 := "lalalalala"
	err := sm.SetArb(id, key, data)
	if err != nil {
		t.Fatalf("error was %v", err)
	}

	err = sm.SetArb(id2, key, data2)
	if err != nil {
		t.Fatalf("error was %v", err)
	}

	arb, err := sm.GetArb(id, key).String()
	arb2, err2 := sm.GetArb(id2, key).String()

	if err != nil {
		t.Fatalf("error was %v", err)
	}

	if err2 != nil {
		t.Fatalf("error was %v", err2)
	}

	if arb != data {
		t.Errorf("arbitrary data was not the same as when it was set. input %v. output %v", data, arb)
	}
	if arb2 != data2 {
		t.Errorf("arbitrary data was not the same as when it was set. input %v. output %v", data, arb2)
	}
}
