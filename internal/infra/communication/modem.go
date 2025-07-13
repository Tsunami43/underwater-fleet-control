package communication

import (
	"errors"
	"github.com/Tsunami43/underwater-fleet-control/internal/domain"
	"sync"
	"time"
)

// Интерфейс для работы с модемом
type Modem interface {
	Send(packet *domain.Packet) error
	Receive() (*domain.Packet, error)
}

// Мок-реализация гидроакустического модема для теста
type MockModem struct {
	inChan  chan *domain.Packet
	outChan chan *domain.Packet
	closed  bool
	mu      sync.Mutex
}

func NewMockModem() *MockModem {
	return &MockModem{
		inChan:  make(chan *domain.Packet, 10),
		outChan: make(chan *domain.Packet, 10),
	}
}

// Отправка пакета (кладём в канал)
func (m *MockModem) Send(packet *domain.Packet) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		return errors.New("modem closed")
	}
	m.outChan <- packet
	return nil
}

// Получение пакета (читаем из канала с тайм-аутом)
func (m *MockModem) Receive() (*domain.Packet, error) {
	select {
	case pkt := <-m.inChan:
		return pkt, nil
	case <-time.After(3 * time.Second):
		return nil, errors.New("timeout waiting for packet")
	}
}

// Для теста: внутренняя функция эмулирует приём пакетов из outChan обратно в inChan
func (m *MockModem) StartEcho() {
	go func() {
		for pkt := range m.outChan {
			time.Sleep(500 * time.Millisecond) // эмуляция задержки
			m.inChan <- pkt                    // возвращаем тот же пакет обратно
		}
	}()
}

// Закрыть модем
func (m *MockModem) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.closed {
		close(m.outChan)
		close(m.inChan)
		m.closed = true
	}
}
