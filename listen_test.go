package main

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	// nil은 object c에서 파생된 단어로 객체 참조 전용 null 값으로 Go에서는 null 대신 nil(닐) 사용
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())
}
