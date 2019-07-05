package proxy

import (
	"context"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewDefaultTCPProxy(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
		want *TCPProxy
	}{
		{
			"expected tcp proxy values",
			args{
				&sync.WaitGroup{},
			},
			&TCPProxy{
				&sync.WaitGroup{},
				net.ListenConfig{},
				context.Background(),
				"tcp",
				10 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultTCPProxy(tt.args.wg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultTCPProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTCPProxy_ListenAndProxy(t *testing.T) {
	type fields struct {
		wg          *sync.WaitGroup
		lc          net.ListenConfig
		ctx         context.Context
		network     string
		dialTimeout time.Duration
	}
	type args struct {
		port  Port
		table IPTable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := TCPProxy{
				wg:          tt.fields.wg,
				lc:          tt.fields.lc,
				ctx:         tt.fields.ctx,
				network:     tt.fields.network,
				dialTimeout: tt.fields.dialTimeout,
			}
			tp.ListenAndProxy(tt.args.port, tt.args.table)
		})
	}
}

func Test_proxy(t *testing.T) {
	type args struct {
		from net.Conn
		to   net.Conn
		wg   *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxy(tt.args.from, tt.args.to, tt.args.wg)
		})
	}
}
