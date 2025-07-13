# 🚢 Underwater Fleet Control

### Система управления групповыми задачами подводных аппаратов с обработкой и логированием связи.

### 📂 Структура проекта (основное)
```plaintext
cmd/app/main.go                  // Тестовый запуск приложения
internal/domain/                 // Бизнес-структуры (Robot, Packet)
internal/usecase/                // Логика обработки пакетов
internal/infra/communication/    // Реализация модема (Mock / Production)
internal/infra/logger/           // Логирование событий
internal/delivery/               // Слушатель и обработчик пакетов
```

🚀 Быстрый запуск для тестов (эмуляция модема)

```bash
go run cmd/app/main.go
```
Ожидаемый результат:
Отправка и прием пакетов

🛠 Сборка бинарника (для деплоя на реальном оборудовании)
```bash
go build -o fleet-control ./cmd/app
./fleet-control
```
-----
### 📡 Использование с реальным гидроакустическим модемом

Что необходимо:
* Драйвер или SDK твоего реального модема.
* Подключение будет зависеть от модели устройства. Обычно используется: TCP / Serial / USB интерфейс. Библиотека от производителя.

Реализация интерфейса Modem вместо MockModem:

В internal/infra/communication нужно создать новый файл:

```go
type RealModem struct {...}
func (r *RealModem) Send(pkt *domain.Packet) error { ... }
func (r *RealModem) Receive() <-chan *domain.Packet { ... }
```

```go
modem := communication.NewRealModem() // вместо Mock
```

#### 📄 Логирование
Все события записываются в файл:

```
fleet.log
```
