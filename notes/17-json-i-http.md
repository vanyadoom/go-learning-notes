# Конспект 17. JSON и базовая работа с net/http

## 1. Кодирование в JSON (Marshal)

```go
import "encoding/json"

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"` // не включать в вывод, если пусто
}

u := User{Name: "Alex", Age: 25}
data, err := json.Marshal(u)
fmt.Println(string(data)) // {"name":"Alex","age":25}

// с отступами, для читаемости (логи, отладка):
prettyData, err := json.MarshalIndent(u, "", "  ")
```

**Правило:** `encoding/json` работает только с **экспортируемыми** полями структуры (с большой буквы) — неэкспортируемые поля тег `json` вообще не спасёт, они игнорируются.

## 2. Декодирование из JSON (Unmarshal)

```go
jsonStr := `{"name":"Alex","age":25}`

var u User
err := json.Unmarshal([]byte(jsonStr), &u) // ОБЯЗАТЕЛЬНО указатель — иначе некуда записать результат
if err != nil {
    fmt.Println("ошибка парсинга:", err)
}
fmt.Println(u.Name, u.Age)
```

### Декодирование в неизвестную структуру — map[string]interface{}

Когда заранее не известна структура JSON (или она произвольная):

```go
var data map[string]interface{}
json.Unmarshal([]byte(`{"name":"Alex","age":25,"active":true}`), &data)

for key, value := range data {
    fmt.Printf("%s: %v (%T)\n", key, value, value)
}
// name: Alex (string)
// age: 25 (float64)      <- ВАЖНО: JSON-числа всегда декодируются как float64!
// active: true (bool)
```

**Частый вопрос на собеседовании:** почему `age` стало `float64`, а не `int`? — Формат JSON не различает целые и дробные числа как отдельные типы, поэтому пакет `encoding/json` по умолчанию декодирует любое число в `float64` при работе с `interface{}`. Если тип заранее известен (структура `User` с полем `Age int`), декодирование пройдёт корректно в `int`.

## 3. Потоковая работа с JSON: Decoder/Encoder (эффективнее для сети/файлов)

```go
resp, _ := http.Get("https://api.example.com/user")
defer resp.Body.Close()

var u User
decoder := json.NewDecoder(resp.Body)
err := decoder.Decode(&u) // читает и парсит прямо из потока, без промежуточного []byte
```

## 4. net/http — минимальный HTTP-сервер

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Привет, мир!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    user := User{Name: "Alex", Age: 25}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user) // сразу пишет JSON в ответ, без промежуточного Marshal
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/user", userHandler)

    log.Println("сервер запущен на :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

- `http.ResponseWriter` — интерфейс для записи ответа (реализует `io.Writer`).
- `*http.Request` — объект запроса: метод, заголовки, тело, параметры URL.
- `http.HandleFunc` — регистрирует функцию-обработчик на маршрут (роутинг здесь примитивный; в реальных проектах используют роутеры вроде `gorilla/mux`, `chi`, или встроенный `http.ServeMux` с шаблонами путей, начиная с Go 1.22).

## 5. Чтение параметров запроса и тела

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // query-параметры: /search?q=golang
    query := r.URL.Query().Get("q")

    // метод запроса
    if r.Method != http.MethodPost {
        http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    // тело запроса (например, JSON от клиента)
    var input User
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "некорректный JSON", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "получено: %+v, query: %s", input, query)
}
```

## 6. HTTP-клиент — исходящие запросы

```go
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close() // ОБЯЗАТЕЛЬНО — иначе утечка соединений

body, err := io.ReadAll(resp.Body)
fmt.Println(string(body))
fmt.Println("статус:", resp.StatusCode)

// POST-запрос с JSON-телом
payload, _ := json.Marshal(User{Name: "Alex", Age: 25})
resp2, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(payload),
)
defer resp2.Body.Close()

// Клиент с таймаутом (правильная практика — не использовать http.DefaultClient без таймаута в проде!)
client := &http.Client{Timeout: 5 * time.Second}
resp3, err := client.Get("https://api.example.com/data")
```

**Важное правило продакшн-кода:** всегда закрывай `resp.Body` (обычно через `defer`) — иначе соединения не освобождаются и в какой-то момент упрёшься в лимит открытых файловых дескрипторов.

---

## Проверь себя

1. Почему неэкспортируемые поля структуры не попадают в JSON, даже если у них есть тег `json`?
2. Почему число из JSON, декодированное в `interface{}`, оказывается типом `float64`?
3. В чём разница между `json.Marshal`/`Unmarshal` и `json.NewEncoder`/`NewDecoder`?
4. Зачем HTTP-клиенту явно задавать `Timeout`, если можно использовать `http.Get` напрямую?
5. Что произойдёт, если забыть закрыть `resp.Body` после HTTP-запроса?

**Далее:** Конспект 18 — тестирование: testing, table-driven tests, бенчмарки.
