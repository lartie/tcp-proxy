package main

import (
	"flag"
	"github.com/lartie/tcp-proxy/proxy"
	"github.com/lartie/tcp-proxy/utils"
	"sync"
)

var (
	env = flag.String("env", "dev", ".env.{environment} file")
)

func init() {
	flag.Parse()

	utils.LoadEnv(*env)
}

func main() {
	ipt := loadIPTable()
	listenPorts := loadListenPorts()

	wg := sync.WaitGroup{}
	wg.Add(1)

	p := proxy.NewDefaultTCPProxy(&wg)

	for _, port := range listenPorts {
		go p.ListenAndProxy(port, ipt)
	}

	wg.Wait()
}

func loadIPTable() proxy.IPTable {
	ipt := make(proxy.IPTable)
	utils.FillStructFromConfig("IP_TABLE_CONFIG", &ipt)

	return ipt
}

func loadListenPorts() proxy.Ports {
	var ports proxy.Ports
	utils.FillStructFromConfig("LISTEN_PORTS_CONFIG", &ports)

	return ports
}
