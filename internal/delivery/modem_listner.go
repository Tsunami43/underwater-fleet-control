package delivery

import (
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/communication"
	"github.com/Tsunami43/underwater-fleet-control/internal/usecase"
	"log"
)

func ListenAndHandle(modem communication.Modem, service *usecase.PacketService) {
	for {
		packet, err := modem.Receive()
		if err != nil {
			log.Println("Error receiving packet:", err)
			continue
		}
		service.HandleIncoming(packet)
	}
}

