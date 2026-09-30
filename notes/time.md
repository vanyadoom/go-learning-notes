# Пакет `time` в Go

Пакет `time` используется для работы с датой, временем и временными интервалами.

Подключение:

```go
import "time"
```

## 1. `time.Now()`

Функция:

```go
time.Now()
```

возвращает текущие дату и время.

Пример:

```go
now := time.Now()

fmt.Println(now)
```

Переменная `now` имеет тип:

```go
time.Time
```

Коротко:

```text
time.Time → конкретный момент времени
```

---

## 2. Получение частей даты

Из значения `time.Time` можно получать отдельные части даты.

```go
now := time.Now()

fmt.Println(now.Year())
fmt.Println(now.Month())
fmt.Println(now.Day())
```

Основные методы:

```go
now.Year()
now.Month()
now.Day()
```

Они позволяют получить:

```text
Year  → год
Month → месяц
Day   → день месяца
```

---

## 3. Получение времени

Из `time.Time` также можно получить часы, минуты и секунды.

```go
now := time.Now()

fmt.Println(now.Hour())
fmt.Println(now.Minute())
fmt.Println(now.Second())
```

Основные методы:

```text
Hour()   → часы
Minute() → минуты
Second() → секунды
```

---

## 4. `time.Duration`

`time.Duration` используется для представления временного интервала.

Главное различие:

```text
time.Time     → конкретный момент времени
time.Duration → промежуток времени
```

Например:

```go
duration := 5 * time.Second
```

Здесь `duration` представляет пять секунд.

---

## 5. Единицы времени

В пакете `time` есть готовые значения:

```go
time.Second
time.Minute
time.Hour
```

Примеры:

```go
5 * time.Second
```

```go
10 * time.Minute
```

```go
2 * time.Hour
```

Коротко:

```text
time.Second → секунда
time.Minute → минута
time.Hour   → час
```

---

## 6. `time.Sleep`

`time.Sleep` приостанавливает выполнение текущей goroutine на указанное время.

Пример:

```go
fmt.Println("Начало")

time.Sleep(2 * time.Second)

fmt.Println("Конец")
```

Между первым и вторым выводом пройдёт примерно две секунды.

Главная схема:

```text
Sleep → подождать указанный промежуток времени
```

---

## 7. `Add`

Метод `Add` позволяет прибавить временной интервал к `time.Time`.

Пример:

```go
now := time.Now()

future := now.Add(2 * time.Hour)
```

`future` будет представлять момент на два часа позже `now`.

Можно двигаться и назад:

```go
past := now.Add(-2 * time.Hour)
```

Здесь получится момент на два часа раньше.

---

## 8. `time.Time` и `time.Duration`

Важно не путать:

```text
time.Time
→ момент времени
→ когда?
```

```text
time.Duration
→ временной интервал
→ сколько времени?
```

Пример:

```go
now := time.Now()
```

`now` — `time.Time`.

А:

```go
delay := 5 * time.Second
```

`delay` — `time.Duration`.

---

## Шпаргалка

```text
time.Now()
→ получить текущие дату и время
```

```text
time.Time
→ конкретный момент времени
```

```text
time.Duration
→ временной промежуток
```

```text
time.Second
time.Minute
time.Hour
→ единицы времени
```

```text
now.Year()
now.Month()
now.Day()
→ части даты
```

```text
now.Hour()
now.Minute()
now.Second()
→ части времени
```

```text
now.Add(2 * time.Hour)
→ получить момент на 2 часа позже
```

```text
now.Add(-2 * time.Hour)
→ получить момент на 2 часа раньше
```

```text
time.Sleep(2 * time.Second)
→ приостановить выполнение примерно на 2 секунды
```

## Главное запомнить

```text
time.Time     → когда?
time.Duration → сколько времени?
```

И базовая цепочка:

```text
time.Now()
→ time.Time
→ можем получить год, месяц, день, часы, минуты и секунды
```