package client

import "github.com/Yiivgeny/incotex-mercury-client/protocol"

func (c *Client) Request(address protocol.Address, request protocol.PDU, target Unmarshaller) error {
	err, pdu := c.request(address, request, target != nil)
	if err != nil {
		return err
	}
	if target == nil {
		return nil
	}
	return target.Unmarshall(pdu)
}

type Unmarshaller interface {
	Unmarshall(pdu protocol.PDU) error
}
