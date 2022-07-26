package protocol

import (
	"encoding/binary"
	"time"
)

func UnpackInteger(value []byte) uint32 {
	mask := byte(0xFF)
	for _, v := range value {
		mask = mask & v
	}
	if mask == 0xFF {
		return 0
	}

	r := append(make([]byte, 4-len(value), 4), value...)
	r[0], r[1], r[2], r[3] = r[1], r[0], r[3], r[2]
	return binary.BigEndian.Uint32(r)
}

func UnpackNormalShort(value []byte) uint16 {
	return binary.BigEndian.Uint16(value)
}

func Bdc2Int(value byte) int {
	result := int(value) >> 0 & 0b00001111 * 1
	result += int(value) >> 4 & 0b00001111 * 10

	return result
}

func SignedFloatDecode(v []byte) float32 {
	activeSign := 1
	reactiveSign := 1
	signByte := 0

	if len(v) == 4 {
		signByte = 1
	}

	if v[signByte]&0b10000000 > 0 {
		activeSign = -1
	}
	if v[signByte]&0b01000000 > 0 {
		reactiveSign = -1
	}
	v[signByte] = 0

	return float32(UnpackInteger(v)) / float32(100*activeSign*reactiveSign)
}

func FrameTimeout(baud uint, long bool) time.Duration {
	if baud >= 38400 && !long {
		return 2 * time.Millisecond
	} else if baud >= 19200 && !long {
		return 3 * time.Millisecond
	} else if baud >= 9600 && !long {
		return 5 * time.Millisecond
	} else if baud >= 4800 && !long {
		return 10 * time.Millisecond
	} else if baud >= 2400 && !long {
		return 20 * time.Millisecond
	} else if baud >= 2400 && long {
		return 25 * time.Millisecond
	} else if baud >= 1200 {
		return 40 * time.Millisecond
	} else if baud >= 600 {
		return 80 * time.Millisecond
	} else if baud >= 300 {
		return 160 * time.Millisecond
	} else {
		return 320 * time.Millisecond
	}
}

func ResponseTimeout(baud uint) time.Duration {
	if baud >= 9600 {
		return 150 * time.Millisecond
	} else if baud >= 4800 {
		return 180 * time.Millisecond
	} else if baud >= 2400 {
		return 210 * time.Millisecond
	} else if baud >= 1200 {
		return 400 * time.Millisecond
	} else if baud >= 600 {
		return 800 * time.Millisecond
	} else if baud >= 300 {
		return 1600 * time.Millisecond
	} else {
		return 3200 * time.Millisecond
	}
}