package network

import (
	"net"
	"syscall"
	"testing"
	"time"
)

// 1
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		// 2
		Control: func(_, addr string, c syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection time out",
				Name:        addr,
				Server:      "!27.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TestDialTimeout(t *testing.T) {
	// 3
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		c.Close()
		t.Fatal("connection did not time out")
	}
	// 4
	nErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err)
	}
	// 5
	if !nErr.Timeout() {
		t.Fatal("error is not a timeout")
	}
}
