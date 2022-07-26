package client

import (
	"net"
	"time"
)

type TransportTCP struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTransportTCP(cfg *Config) (*TransportTCP, error) {
	addr, err := net.ResolveTCPAddr("tcp", cfg.Host)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil && conn != nil {
			_ = conn.Close()
		}
	}()
	err = conn.SetNoDelay(true)
	if err != nil {
		return nil, err
	}
	return &TransportTCP{
		conn:    conn,
		timeout: cfg.ResponseTimeout,
	}, nil
}

func (t *TransportTCP) Read(p []byte) (n int, err error) {
	err = t.conn.SetReadDeadline(time.Now().Add(t.timeout))
	return t.conn.Read(p)
}

func (t *TransportTCP) Write(p []byte) (n int, err error) {
	return t.conn.Write(p)
}

func (t *TransportTCP) Close() error {
	return t.conn.Close()
}
