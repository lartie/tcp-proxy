package proxy

import (
	"context"
	"fmt"
	"github.com/lartie/tcp-proxy/utils"
	"net"
	"strconv"
	"sync"
	"time"
)

// TCPProxy config
type TCPProxy struct {
	wg *sync.WaitGroup
	lc net.ListenConfig
	ctx context.Context
	network string
	dialTimeout time.Duration
}

// NewDefaultTCPProxy creates an TCPProxy with default parameters.
// Network: tcp
// Context: Background
// Timeout: 10 seconds
func NewDefaultTCPProxy(wg *sync.WaitGroup) *TCPProxy {
	t := new(TCPProxy)

	t.ctx = context.Background()
	t.network = "tcp"
	t.wg = wg
	t.lc = net.ListenConfig{}
	t.dialTimeout = 10 * time.Second

	return t
}

// ListenAndProxy port and proxy connections
func (t TCPProxy) ListenAndProxy(port Port, table IPTable) {
	utils.WriteInfoLog(fmt.Sprintf("Listening tcp on :%d", port))

	l, err := t.lc.Listen(t.ctx, t.network, fmt.Sprintf(":%d", port))
	utils.CheckErrWithWaitGroup(err, t.wg)

	defer l.Close()

	for {
		conn, err := l.Accept()
		utils.CheckErrWithWaitGroup(err, t.wg)

		rh, _, err := net.SplitHostPort(conn.RemoteAddr().String())
		utils.CheckErr(err)

		target, ok := table[rh]

		if ok {
			utils.WriteInfoLog(fmt.Sprintf("Proxy addr has found %s -> %s", rh, target))

			t.dial(conn, target, port)
		} else {
			utils.WriteCritLog(fmt.Sprintf("Unexpected remote addr %s", rh))
			err = conn.Close()

			if err != nil {
				utils.WriteCritLog(fmt.Sprintf("Got error while conn is closing: %s", err.Error()))
			}
		}
	}
}

// dial connects to the address and proxy packets
// When packets exchange will be terminated, it will be closed.
func (t TCPProxy) dial(local net.Conn, target string, port Port) {
	target = net.JoinHostPort(target, strconv.Itoa(int(port)))

	remote, err := net.DialTimeout(t.network, target, t.dialTimeout)

	if err != nil {
		utils.WriteCritLog(err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go proxy(local, remote, &wg)
	go proxy(remote, local, &wg)

	wg.Wait()
}

func proxy(from, to net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	_, _, err := net.SplitHostPort(from.LocalAddr().String())
	utils.CheckErr(err)

	b := make([]byte, 10240)
	for {
		n, err := from.Read(b)
		if err != nil {
			break
		}

		if n > 0 {
			to.Write(b[:n])
		}
	}

	from.Close()
	to.Close()
}
