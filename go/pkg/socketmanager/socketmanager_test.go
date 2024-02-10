package socketmanager_test

import (
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
