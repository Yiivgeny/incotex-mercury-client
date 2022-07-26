package client

import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

type Client struct {
	transport   io.ReadWriter
	sendTimer   *time.Timer
	sendTimeout time.Duration
	mutex       *sync.Mutex
}

func NewClient(cfg *Config, transport io.ReadWriter) Client {
	return Client{
		transport:   transport,
		sendTimer:   time.NewTimer(time.Nanosecond),
		sendTimeout: cfg.FrameTimeout,
		mutex:       &sync.Mutex{},
	}
}

func (c *Client) send(adu protocol.ADU) error {
	<-c.sendTimer.C
	n, err := c.transport.Write(adu)
	if err != nil {
		return err
	}
	if n != len(adu) {
		return errors.New("unknown send error")
	}
	c.sendTimer.Reset(c.sendTimeout)

	return nil
}

func (c *Client) recv() (protocol.ADU, error) {
	result := make(protocol.ADU, 1+255+2)
	n, err := c.transport.Read(result)
	if err != nil {
		return []byte{}, err
	}
	return result[:n], nil
}

func (c *Client) request(address protocol.Address, req protocol.PDU, withResponse bool) (error, protocol.PDU) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.send(protocol.NewADU(address, req)); err != nil {
		return err, nil
	}
	if !withResponse {
		return nil, nil
	}

	rAdu, err := c.recv()
	if err != nil {
		return err, nil
	}

	pdu := rAdu.PDU()
	if !rAdu.Verify() {
		err = errors.New("invalid checksum")
	} else if rAdu.Address() != address {
		err = errors.New("addresses mismatch")
	}

	return err, pdu
}
