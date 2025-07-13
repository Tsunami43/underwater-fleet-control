package usecase

import (
	"errors"
	"fmt"

	"github.com/Tsunami43/underwater-fleet-control/internal/domain"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/communication"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/logger"
)

type PacketService struct {
	Modem  communication.Modem
	Logger *logger.Logger
}

// Создаём сервис
func NewPacketService(modem communication.Modem, logger *logger.Logger) *PacketService {
	return &PacketService{Modem: modem, Logger: logger}
}

// Отправить пакет, проверить его валидность
func (ps *PacketService) SendPacket(packet *domain.Packet) error {
	if !packet.IsValid() {
		return errors.New("packet is invalid")
	}
	ps.Logger.Log("Sending packet: " + packet.String())
	return ps.Modem.Send(packet)
}

// Обработка входящего пакета

func (ps *PacketService) HandleIncoming(packet *domain.Packet) {
	if !packet.IsValid() {
		ps.Logger.Log("Received invalid packet: " + packet.String())
		fmt.Println("Received INVALID:", packet.String())
		return
	}
	ps.Logger.Log("Received valid packet: " + packet.String())
	fmt.Println("Received:", packet.String())
}
