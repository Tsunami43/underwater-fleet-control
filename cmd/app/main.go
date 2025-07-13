package main

import (
	"github.com/Tsunami43/underwater-fleet-control/internal/delivery"
	"github.com/Tsunami43/underwater-fleet-control/internal/domain"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/communication"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/logger"
	"github.com/Tsunami43/underwater-fleet-control/internal/usecase"
	"log"
	"time"
)

func main() {
	logg, err := logger.NewLogger("fleet.log")
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}
	defer logg.Close()

	modem := communication.NewMockModem()
	modem.StartEcho()

	packetService := usecase.NewPacketService(modem, logg)

	go delivery.ListenAndHandle(modem, packetService)

	for i := 1; i <= 3; i++ {
		pkt := domain.NewPacket(
			"pkt"+string(i+48),
			"server",
			"robot1",
			[]byte("Ping "+string(i+48)),
		)
		_ = packetService.SendPacket(pkt)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Second)
	log.Println("Done.")
}
