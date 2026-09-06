# Конспект 18. Тестирование: testing, table-driven tests, бенчмарки

## 1. Базовый unit-тест

Тесты в Go — часть стандартной библиотеки, отдельный фреймворк не нужен. Файл с тестами называется `имя_test.go` и лежит рядом с тестируемым кодом, в том же пакете.

```go
// math.go
package mathutils

func Add(a, b int) int {
    return a + b
}
```

```go
// math_test.go
package mathutils

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; ожидалось %d", result, expected)
    }
}
```

**Правила именования (проверяются инструментом `go test`):**
- Файл: `*_test.go`.
- Функция теста: `func TestXxx(t *testing.T)` — обязательно с префиксом `Test` и заглавной буквой дальше.
- `t.Errorf` — фиксирует ошибку, но **не останавливает** выполнение остальной функции теста (в отличие от `t.Fatalf`, которая останавливает немедленно).

```bash
go test ./...           # запустить все тесты в проекте
go test -v ./...        # подробный вывод (какие тесты прошли/не прошли)
go test -run TestAdd     # запустить только тесты, чьё имя матчится с regex
go test -cover ./...     # процент покрытия кода тестами
```

## 2. Table-driven tests — идиоматичный стиль тестирования в Go

Вместо множества похожих отдельных функций-тестов — один тест с таблицей входов/ожидаемых результатов. Это стандарт де-факто в реальных Go-проектах:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"положительные числа", 2, 3, 5},
        {"отрицательные числа", -2, -3, -5},
        {"с нулём", 0, 5, 5},
        {"оба нуля", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) { // t.Run — подтест, виден в выводе отдельно
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; ожидалось %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

`t.Run` позволяет запустить конкретный подтест отдельно: `go test -run TestAdd/с_нулём`, а в выводе видно, какой именно случай из таблицы упал.

## 3. Тестирование ошибок

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name        string
        a, b        float64
        expectError bool
    }{
        {"нормальное деление", 10, 2, false},
        {"деление на ноль", 10, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := Divide(tt.a, tt.b)
            if (err != nil) != tt.expectError {
                t.Errorf("Divide(%v, %v): err = %v, ожидалось наличие ошибки: %v",
                    tt.a, tt.b, err, tt.expectError)
            }
        })
    }
}
```

## 4. testify — популярная сторонняя библиотека для тестов

`github.com/stretchr/testify` — не часть стандартной библиотеки, но фактический стандарт в индустрии, часто ожидается знание на собеседованиях:

```go
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result)             // мягкая проверка — тест продолжится дальше
    assert.NoError(t, err)                  // проверка на отсутствие ошибки
    assert.True(t, result > 0)
}

import "github.com/stretchr/testify/require"
// require.Equal(...) — то же самое, но останавливает тест немедленно при провале (как t.Fatalf)
```

## 5. Моки через интерфейсы (напоминание из конспекта 9)

```go
type Notifier interface {
    Send(message string) error
}

type MockNotifier struct {
    SentMessages []string
    ShouldFail   bool
}

func (m *MockNotifier) Send(message string) error {
    if m.ShouldFail {
        return errors.New("имитация ошибки отправки")
    }
    m.SentMessages = append(m.SentMessages, message)
    return nil
}

func TestOrderService_NotifiesOnSuccess(t *testing.T) {
    mock := &MockNotifier{}
    service := NewOrderService(mock)

    service.CompleteOrder(123)

    if len(mock.SentMessages) != 1 {
        t.Errorf("ожидалось одно уведомление, получено %d", len(mock.SentMessages))
    }
}
```

## 6. Бенчмарки — измерение производительности

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { // b.N подбирается фреймворком автоматически для стабильного измерения
        Add(2, 3)
    }
}
```

```bash
go test -bench=. ./...              # запустить все бенчмарки
go test -bench=. -benchmem ./...    # + показать количество аллокаций памяти
```

Вывод содержит наносекунды на операцию (`ns/op`) и, с `-benchmem`, байты и количество аллокаций на операцию (`B/op`, `allocs/op`) — важно при оптимизации горячих участков кода.

## 7. Test coverage (покрытие тестами)

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out   # открывает наглядный HTML-отчёт — какие строки не покрыты тестами
```

---

## Проверь себя

1. По какому правилу называется файл с тестами и функция-тест, чтобы `go test` их нашёл?
2. Чем `t.Errorf` отличается от `t.Fatalf`?
3. Зачем нужен table-driven test вместо нескольких отдельных функций-тестов?
4. Как имитировать зависимость (например, отправку email) в тесте без реального похода в сеть?
5. Что показывает флаг `-benchmem` при запуске бенчмарка?

**Далее:** Конспект 19 — идиомы Go, gofmt, линтеры, code review checklist.
