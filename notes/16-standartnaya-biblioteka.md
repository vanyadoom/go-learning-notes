# Конспект 16. Стандартная библиотека: strings, strconv, time, os, io, bufio

## 1. strings — работа со строками

```go
import "strings"

strings.Contains("golang", "lang")       // true
strings.HasPrefix("golang", "go")         // true
strings.HasSuffix("golang", "lang")       // true
strings.Index("golang", "lang")           // 2 (индекс первого вхождения, -1 если не найдено)
strings.ToUpper("golang")                 // "GOLANG"
strings.ToLower("GOLANG")                 // "golang"
strings.TrimSpace("  golang  ")           // "golang"
strings.Trim("--golang--", "-")           // "golang"
strings.Split("a,b,c", ",")               // []string{"a", "b", "c"}
strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"
strings.Replace("aaa", "a", "b", 2)        // "bba" — заменить не более 2 вхождений
strings.ReplaceAll("aaa", "a", "b")        // "bbb"
strings.Repeat("ab", 3)                    // "ababab"
strings.Fields("  a  b   c ")              // []string{"a", "b", "c"} — разбить по пробелам, без пустых
```

### strings.Builder — эффективная конкатенация строк в цикле

Обычная конкатенация строк через `+=` в цикле создаёт новую строку на каждой итерации (строки в Go неизменяемы!) — это дорого для больших объёмов. `strings.Builder` избегает лишних аллокаций:

```go
var sb strings.Builder
for i := 0; i < 1000; i++ {
    sb.WriteString("x")
}
result := sb.String()
```

## 2. strconv — преобразования строка ↔ число (повтор из конспекта 2, с деталями)

```go
n, err := strconv.Atoi("42")               // string -> int
s := strconv.Itoa(42)                       // int -> string
f, err := strconv.ParseFloat("3.14", 64)    // string -> float64
i64, err := strconv.ParseInt("42", 10, 64)  // string -> int64, base 10, размер 64 бита
b, err := strconv.ParseBool("true")         // string -> bool
s2 := strconv.FormatFloat(3.14159, 'f', 2, 64) // float64 -> "3.14"
```

## 3. time — работа со временем

```go
import "time"

now := time.Now()
fmt.Println(now.Year(), now.Month(), now.Day())
fmt.Println(now.Format("2006-01-02 15:04:05")) // ВАЖНО: формат задаётся через "опорную дату"!

duration := 5 * time.Second
time.Sleep(duration)

start := time.Now()
// ... какая-то работа ...
elapsed := time.Since(start)
fmt.Println("заняло:", elapsed)

deadline := time.Now().Add(10 * time.Minute)

t, err := time.Parse("2006-01-02", "2024-03-15")
```

**Особенность форматирования дат в Go (частый источник путаницы):** вместо `YYYY-MM-DD` или `%Y-%m-%d`, как в других языках, Go использует **опорную дату** `Mon Jan 2 15:04:05 MST 2006` (запоминается как последовательность 1,2,3,4,5,6,7 — месяц/день/час/минута/секунда/год/часовой пояс). Чтобы отформатировать дату, пишешь строку-образец в этом формате:

```go
fmt.Println(now.Format("02.01.2006")) // европейский формат даты: "15.03.2024"
fmt.Println(now.Format(time.RFC3339)) // готовая константа: "2024-03-15T10:30:00Z"
```

## 4. os — работа с окружением, файлами, аргументами

```go
import "os"

args := os.Args // []string, args[0] — путь к бинарнику, args[1:] — аргументы командной строки

value := os.Getenv("PATH")
os.Setenv("MY_VAR", "value")

file, err := os.Open("data.txt")        // открыть для чтения
defer file.Close()

file2, err := os.Create("output.txt")   // создать/перезаписать для записи
defer file2.Close()

data, err := os.ReadFile("data.txt")    // прочитать файл целиком в память
err = os.WriteFile("output.txt", []byte("данные"), 0644)

os.Exit(1) // немедленный выход с кодом возврата (defer НЕ выполнятся!)
```

**Важный нюанс:** `os.Exit` завершает программу немедленно, минуя все отложенные `defer` — в отличие от обычного `return` из `main()`.

## 5. io и bufio — потоковая обработка данных

`io.Reader` и `io.Writer` (см. конспект 9) — универсальные интерфейсы для последовательного чтения/записи байт, без знания конкретного источника (файл, сеть, память).

```go
import (
    "bufio"
    "os"
)

file, _ := os.Open("data.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text() // построчное чтение — самый частый паттерн
    fmt.Println(line)
}
if err := scanner.Err(); err != nil {
    fmt.Println("ошибка чтения:", err)
}

// bufio.Reader — более гибкое буферизованное чтение
reader := bufio.NewReader(os.Stdin)
line, err := reader.ReadString('\n') // читать до символа-разделителя

// bufio.Writer — буферизованная запись (важно для производительности!)
writer := bufio.NewWriter(os.Stdout)
writer.WriteString("привет\n")
writer.Flush() // ОБЯЗАТЕЛЬНО — иначе данные могут остаться в буфере и не попасть в вывод!
```

**Почему буферизация важна:** каждый прямой вызов `os.File.Write` — это системный вызов (syscall), который относительно дорог. `bufio` накапливает данные в памяти и сбрасывает их пачками, что на порядки быстрее при большом количестве мелких записей.

## 6. Пример: комбинирование всего вместе — чтение CSV-подобного файла

```go
func processFile(path string) ([]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("не удалось открыть файл: %w", err)
    }
    defer file.Close()

    var numbers []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        n, err := strconv.Atoi(line)
        if err != nil {
            return nil, fmt.Errorf("некорректная строка %q: %w", line, err)
        }
        numbers = append(numbers, n)
    }
    return numbers, scanner.Err()
}
```

---

## Проверь себя

1. Почему `strings.Builder` эффективнее, чем конкатенация строк через `+=` в цикле?
2. Что означает "опорная дата" `Mon Jan 2 15:04:05 MST 2006` в форматировании времени Go?
3. Чем `os.Exit(1)` отличается от обычного `return` из `main()` в контексте `defer`?
4. Зачем нужен `writer.Flush()` после записи через `bufio.Writer`?
5. Почему прямые системные вызовы записи в файл дороже, чем буферизованная запись?

**Далее:** Конспект 17 — JSON и базовая работа с net/http.
