package domain

import (
	"errors"
	"fmt"
	"hash/crc32"
)

type Packet struct {
	ID      string
	Payload []byte
	FromID  string
	ToID    string
	CRC     uint32
}

// Вычислить CRC для данных
func ComputeCRC(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// Проверить валидность пакета (целостность)
func (p *Packet) IsValid() bool {
	return p.CRC == ComputeCRC(p.Payload)
}

// Создать новый пакет с вычисленным CRC
func NewPacket(id, from, to string, payload []byte) *Packet {
	crc := ComputeCRC(payload)
	return &Packet{
		ID:      id,
		FromID:  from,
		ToID:    to,
		Payload: payload,
		CRC:     crc,
	}
}

func (p *Packet) String() string {
	return fmt.Sprintf("Packet[ID=%s, From=%s, To=%s, PayloadLen=%d, CRC=%d]",
		p.ID, p.FromID, p.ToID, len(p.Payload), p.CRC)
}

// Ошибка, если пакет не валиден
var ErrInvalidPacket = errors.New("invalid packet data")
