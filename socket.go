package socket

import (
	"io"
	"net"
	"sync"
	"time"
)

// Socket ...
type Socket struct {
	net.Conn
	isUse bool
	mu    sync.RWMutex
}

// Release ...
func (s *Socket) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isUse = false

}

// IsUse ...
func (s *Socket) IsUse() (b bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b = s.isUse
	return
}

// Ping to detect whether the socket is closed.
func (s *Socket) Ping() (b bool, err error) {
	one := []byte{}
	s.SetReadDeadline(time.Now())
	if _, err := s.Read(one); err == io.EOF {
		s.Close()
		return false, err
	}
	var zero time.Time
	s.SetReadDeadline(zero)
	return true, err
}
