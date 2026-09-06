# Конспект 13. select, sync.Mutex, WaitGroup, Once

## 1. select — "switch" для каналов

`select` позволяет горутине одновременно ждать на нескольких каналах и обработать тот, который "готов" первым.

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "из первого канала" }()
go func() { ch2 <- "из второго канала" }()

select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
}
// выполнится тот case, чей канал "выстрелит" первым; если оба готовы одновременно — выбор случайный
```

### select с default — неблокирующая операция с каналом

```go
select {
case msg := <-ch:
    fmt.Println("получено:", msg)
default:
    fmt.Println("данных пока нет, не блокируемся")
}
```

### Таймаут через select + time.After — очень частый практический паттерн

```go
select {
case result := <-resultChan:
    fmt.Println("успели:", result)
case <-time.After(2 * time.Second):
    fmt.Println("таймаут — операция заняла слишком много времени")
}
```

### Отмена работы через context.Context (стандартный подход в реальных проектах)

```go
func worker(ctx context.Context, results chan<- int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("получен сигнал отмены:", ctx.Err())
            return
        case results <- rand.Intn(100):
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    results := make(chan int)
    go worker(ctx, results)

    for {
        select {
        case r := <-results:
            fmt.Println("результат:", r)
        case <-ctx.Done():
            fmt.Println("main: завершаем работу")
            return
        }
    }
}
```

`context.Context` — стандартный способ прокидывать сигнал отмены и дедлайны через цепочку вызовов функций/горутин. Практически обязателен в любом реальном Go-сервисе (HTTP-хендлеры, работа с БД принимают `ctx` первым аргументом по конвенции).

## 2. sync.Mutex — защита разделяемых данных

Когда несколько горутин читают и пишут одну и ту же переменную (не через канал), нужен мьютекс, чтобы избежать гонки данных.

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock() // defer гарантирует разблокировку даже при панике внутри
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    wg.Wait()
    fmt.Println(counter.Value()) // всегда ровно 1000, благодаря мьютексу
}
```

**Без мьютекса** тот же код с `count++` из 1000 горутин почти гарантированно даст число меньше 1000 — операция `count++` не атомарна (это чтение + инкремент + запись, между этими шагами может вклиниться другая горутина).

### sync.RWMutex — отдельные блокировки для чтения и записи

Если читателей значительно больше, чем писателей — `RWMutex` позволяет множеству горутин читать одновременно, блокируя только на время записи:

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()         // несколько горутин могут одновременно держать RLock
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()          // Lock() эксклюзивен — блокирует и читателей, и писателей
    defer c.mu.Unlock()
    c.data[key] = value
}
```

## 3. sync.WaitGroup — дождаться завершения группы горутин

```go
var wg sync.WaitGroup

for i := 1; i <= 5; i++ {
    wg.Add(1) // увеличиваем счётчик ДО запуска горутины
    go func(id int) {
        defer wg.Done() // уменьшаем счётчик по завершении — ОБЯЗАТЕЛЬНО через defer
        fmt.Println("горутина", id, "завершена")
    }(i) // передаём i как аргумент, а не захватываем через замыкание — важно для Go < 1.22!
}

wg.Wait() // блокируется, пока счётчик не станет 0
fmt.Println("все горутины завершены")
```

**Частые ошибки с WaitGroup:**
- Вызывать `wg.Add(1)` **внутри** горутины, а не перед её запуском — гонка между `Wait()` и `Add()`.
- Забыть `wg.Done()` в одной из веток кода (например, при раннем `return`) — тогда `Wait()` зависнет навсегда (deadlock).

## 4. sync.Once — выполнить код ровно один раз

```go
var once sync.Once
var config *Config

func GetConfig() *Config {
    once.Do(func() {
        fmt.Println("загружаем конфиг...")
        config = loadConfigFromDisk()
    })
    return config
}
```

Сколько бы горутин ни вызвали `GetConfig()` одновременно, функция внутри `once.Do` выполнится ровно один раз — классический потокобезопасный способ реализовать паттерн singleton/ленивую инициализацию.

## 5. sync.Map — потокобезопасная мапа "из коробки"

Используется точечно — для сценариев с редкой записью и частым чтением разными горутинами (обычный `map` + `Mutex` в большинстве случаев быстрее и понятнее):

```go
var m sync.Map

m.Store("key", "value")
value, ok := m.Load("key")
m.Delete("key")

m.Range(func(key, value interface{}) bool {
    fmt.Println(key, value)
    return true // false — прервать итерацию
})
```

---

## Проверь себя

1. Что делает `select` и как выбирается case, если готовы сразу несколько каналов?
2. Как реализовать таймаут ожидания результата с помощью `select`?
3. Почему `count++` без мьютекса из нескольких горутин даёт непредсказуемый результат?
4. В чём разница между `Lock()`/`Unlock()` и `RLock()`/`RUnlock()` у `sync.RWMutex`?
5. Что произойдёт, если забыть вызвать `wg.Done()` в одной из горутин?

**Далее:** Конспект 14 — паттерны конкурентности, race conditions, deadlocks.
