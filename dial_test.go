package network

import (
	"io"
	"net"
	"testing"
)

// 3-4 예제
func TestDial(t *testing.T) {
	// 포트 미 지정시 랜덤 포트로 생성
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	// 1: 핸들러
	// 클라이언트 수신 연결
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for {
			// 2: 연결해서는 실패가 없이 로그만 출력
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			// 3:
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					// 4: 패킷을 받고 로그, 해당 패킷이 EOF이면 종료
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	// 5,6,7
	// 리스너로 연결
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	// 8
	conn.Close()
	<-done
	listener.Close()
	<-done
}
