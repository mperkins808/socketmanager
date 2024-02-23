package socketmanager_test

import (
	"testing"

	"github.com/mperkins808/socketmanager/go/pkg/socketmanager"
)

func TestGetIntArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Int()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetInt32Arb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	var d int32 = 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Int32()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetInt64Arb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	var d int64 = 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Int64()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetfloat32Arb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	var d float32 = 1.001
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Float32()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetfloat64Arb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := 1.001001
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Float64()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetIntArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]int, 1)
	d[0] = 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).IntArray()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetInt32ArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]int32, 1)
	d[0] = 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Int32Array()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetInt64ArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]int64, 1)
	d[0] = 1
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Int64Array()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetFloat32ArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]float32, 1)
	d[0] = 1.001
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Float32Array()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetFloat64ArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]float64, 1)
	d[0] = 1.00122
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Float64Array()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetBoolArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := false

	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).Bool()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}

func TestGetBoolArrArb(t *testing.T) {
	sm := socketmanager.NewSimpleSocketManager()
	id := "foo"
	key := "bar"
	sm.Add(id, "fizz")

	d := make([]bool, 1)
	d[0] = true
	err := sm.SetArb(id, key, d)
	if err != nil {
		t.Fatalf("error was %v", err)
	}
	_, err = sm.GetArb(id, key).BoolArray()
	if err != nil {
		t.Fatalf("error was %v", err)
	}
}
