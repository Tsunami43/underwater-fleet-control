package main

import (
	// Импорт внутренних пакетов проекта
	"github.com/Tsunami43/underwater-fleet-control/internal/delivery"
	"github.com/Tsunami43/underwater-fleet-control/internal/domain"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/communication"
	"github.com/Tsunami43/underwater-fleet-control/internal/infra/logger"
	"github.com/Tsunami43/underwater-fleet-control/internal/usecase"
	"log"
	"time"
)

func main() {
	// Инициализация логгера. Все логи будут записываться в файл "fleet.log"
	logg, err := logger.NewLogger("fleet.log")
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}
	defer logg.Close() // Гарантированное закрытие файла лога в конце программы

	// Создаем мок-модем для имитации связи (реальный гидроакустический модем будет здесь в живой системе)
	modem := communication.NewMockModem()
	modem.StartEcho() // Стартуем эмуляцию получения ответов от "роботов"

	// Создаем сервис для отправки пакетов и обработки ответов
	packetService := usecase.NewPacketService(modem, logg)

	// Отдельной горутиной запускаем прослушивание входящих сообщений
	go delivery.ListenAndHandle(modem, packetService)

	// Отправляем тестовые пакеты (Ping) роботу с ID "robot1"
	for i := 1; i <= 3; i++ {
		pkt := domain.NewPacket(
			"pkt"+string(i+48),           // Генерация ID пакета: pkt1, pkt2, pkt3
			"server",                     // Отправитель — сервер
			"robot1",                     // Получатель — робот1
			[]byte("Ping "+string(i+48)), // Небольшой полезный груз (payload)
		)
		_ = packetService.SendPacket(pkt) // Отправляем пакет через сервис
		time.Sleep(1 * time.Second)       // Делаем небольшую задержку между отправками для читаемости вывода
	}

	// Даём системе немного времени на приём и обработку всех эхо-ответов
	time.Sleep(5 * time.Second)

	log.Println("Done.") // Конец тестового запуска
}
