// +build !linux android

package device

import (
	"github.com/cyberscaling/wireguard/conn"
	"github.com/cyberscaling/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
