package admin

import (
	"net"
)

func IsAddressValid(address string) error {
	_, _, err := net.ParseCIDR(address)
	if err != nil {
		return err
	}

	return nil
}
