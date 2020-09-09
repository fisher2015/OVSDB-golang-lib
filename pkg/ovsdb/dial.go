// +build linux darwin aix

package ovsdb

import "net"

func DialNet(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}
