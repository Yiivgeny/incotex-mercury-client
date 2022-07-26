package client

import (
	"errors"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

type Auth struct {
	AccessLevel protocol.AccessLevel
	Password    []byte
}

func (c *Client) TestCommunication(address protocol.Address) error {
	request := protocol.PDU{
		byte(protocol.MethodTestCommunication),
	}
	err, pdu := c.request(address, request, true)
	if err != nil {
		return err
	}
	if len(pdu) != 1 || (pdu[0] != 0x00 && pdu[0] != 0x80) {
		return errors.New("unexpected response")
	}
	return nil
}

func (c *Client) OpenChannel(address protocol.Address, auth Auth) error {
	request := protocol.PDU{
		byte(protocol.MethodOpenChannel),
		byte(auth.AccessLevel),
	}
	request = append(request, auth.Password[:6]...)
	err, pdu := c.request(address, request, true)
	if err != nil {
		return err
	}
	if len(pdu) != 1 || pdu[0] != 0x00 {
		return errors.New("unexpected response")
	}
	return nil
}

func (c *Client) CloseChannel(address protocol.Address) error {
	request := protocol.PDU{
		byte(protocol.MethodCloseChannel),
	}
	err, pdu := c.request(address, request, true)
	if err != nil {
		return err
	}
	if len(pdu) != 1 || pdu[0] != 0x00 {
		return errors.New("unexpected response")
	}
	return nil
}
