package proxy

import (
	"encoding/json"
	"fmt"
	"net"
)

type Port uint16
type Ports []Port

// Association table for input and output ips.
type IPTable map[string]string

// UnmarshalJSON is checking ip addresses
func (i *IPTable) UnmarshalJSON(b []byte) error {
	tempStruct := make(map[string]string)
	err := json.Unmarshal(b, &tempStruct)

	err = checkIPTable(tempStruct)

	if err != nil {
		return err
	}

	for k, v := range tempStruct {
		(*i)[k] = v
	}

	return nil
}

func checkIPTable(ipt map[string]string) error {
	for ip1, ip2 := range ipt {
		if net.ParseIP(ip1) == nil {
			return fmt.Errorf("invalid ip %s", ip1)
		}

		if net.ParseIP(ip2) == nil {
			return fmt.Errorf("invalid ip %s", ip2)
		}
	}

	return nil
}
