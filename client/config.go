package client

import (
	"time"

	"github.com/Yiivgeny/incotex-mercury-client/protocol"
)

type Config struct {
	Host            string
	Baud            uint
	ResponseTimeout time.Duration
	FrameTimeout    time.Duration
}

func NewConfig(baud uint, m int) Config {
	c := Config{
		Baud: baud,
	}
	c.FrameTimeout = protocol.FrameTimeout(baud, false) * time.Duration(m)
	c.ResponseTimeout = protocol.ResponseTimeout(baud) * time.Duration(m)

	return c
}
