package socket_test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/mushroomsir/socket"
	"github.com/stretchr/testify/assert"
)

var (
	InitialCap = 5
	MaximumCap = 30
	network    = "tcp"
	address    = "127.0.0.1:7777"
	factory    = func() (net.Conn, error) { return net.Dial(network, address) }
)

func init() {
	// used for factory function
	go simpleTCPServer()
	time.Sleep(time.Millisecond * 300) // wait until tcp server has been settled
}
func TestSocket(t *testing.T) {
	t.Run("Socket use default options that should be", func(t *testing.T) {
		assert := assert.New(t)
		pool, err := socket.NewPool(1, 20, socket.NewFactory(address))
		if !assert.Nil(err) {
			return
		}
		assert.Equal(20, pool.Cap())
		assert.Equal(1, pool.Len())
		socket, err := pool.Get()
		if !assert.Nil(err) {
			return
		}
		b, err := socket.Ping()
		if !assert.Nil(err) {
			return
		}
		assert.Equal(true, b)
	})
}
func simpleTCPServer() {
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		buffer := make([]byte, 256)
		conn.Read(buffer)
	}
}
