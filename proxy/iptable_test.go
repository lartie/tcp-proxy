package proxy

import (
	"encoding/json"
	"testing"
)

func TestIPTable_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		i       *IPTable
		args    args
		wantErr bool
	}{
		{
			"parse valid json to ip table",
			&IPTable{},
			args{
				[]byte(`{"127.0.0.1":"192.168.0.1","255.255.255.255":"0.0.0.0"}`),
			},
			false,
		},
		{
			"parse invalid json",
			&IPTable{},
			args{
				[]byte(`{"127.0.0.1":"192.168.0.1","255.255.255.255":"0.0.0}`),
			},
			true,
		},
		{
			"parse invalid ips",
			&IPTable{},
			args{
				[]byte(`{"275.0.0.1":"192.168.0.1"}`),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.Unmarshal(tt.args.b, tt.i); (err != nil) != tt.wantErr {
				t.Errorf("IPTable.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkIPTable(t *testing.T) {
	type args struct {
		ipt map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"ipv4 is valid",
			args{
				map[string]string{
					"127.0.0.1": "192.168.0.1",
					"255.255.255.255": "0.0.0.0",
					"0.0.0.0": "255.255.255.255",
				},
			},
			false,
		},
		{
			"ipv6 is valid",
			args{
				map[string]string{
					"0:0:0:0:0:ffff:7f00:1": "0:0:0:0:0:ffff:c0a8:1",
					"0:0:0:0:0:ffff:ffff:ffff": "0:0:0:0:0:ffff:0:0",
					"0:0:0:0:0:ffff:0:0": "0:0:0:0:0:ffff:ffff:ffff",
				},
			},
			false,
		},
		{
			"left ip is invalid",
			args{
				map[string]string{
					"hello": "0:0:0:0:0:ffff:c0a8:1",
				},
			},
			true,
		},
		{
			"right ip is invalid",
			args{
				map[string]string{
					"255.255.255.255": "127.0.0.256",
				},
			},
			true,
		},
		{
			"both ips are invalid",
			args{
				map[string]string{
					"0.255.255.1024": "127.0.0.256",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkIPTable(tt.args.ipt); (err != nil) != tt.wantErr {
				t.Errorf("checkIPTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}