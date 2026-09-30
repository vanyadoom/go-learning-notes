# Пакет `strconv` в Go

Пакет `strconv` используется для преобразования строк в числа и чисел в строки.

Подключение:

```go
import "strconv"
```

## 1. `strconv.Atoi`

`Atoi` преобразует `string` в `int`.

```go
num, err := strconv.Atoi("42")
```

Результат:

```text
num = 42
err = nil
```

Главная схема:

```text
"42" → 42
string → int
```

Важно: `Atoi` возвращает два значения:

```go
num, err := strconv.Atoi(text)
```

Где:

```text
num → результат преобразования
err → ошибка
```

---

## 2. Ошибка при `Atoi`

Если строку нельзя преобразовать в число:

```go
num, err := strconv.Atoi("abc")
```

то:

```text
err != nil
```

То есть:

```text
err == nil → ошибки нет
err != nil → ошибка есть
```

Пример проверки:

```go
num, err := strconv.Atoi(text)

if err != nil {
    fmt.Println("Ошибка преобразования")
    return
}

fmt.Println(num)
```

---

## 3. Ранний `return`

`return` завершает выполнение текущей функции.

Например:

```go
if err != nil {
    fmt.Println("Ошибка")
    return
}

fmt.Println("Продолжаем")
```

Если ошибка есть:

```text
err != nil
→ заходим в if
→ return
→ функция заканчивается
→ код ниже не выполняется
```

Если ошибки нет:

```text
err == nil
→ if пропускается
→ программа идёт дальше
```

---

## 4. `strconv.Itoa`

`Itoa` делает обратное преобразование:

```text
int → string
```

Пример:

```go
text := strconv.Itoa(25)
```

Результат:

```text
"25"
```

То есть:

```text
25 → "25"
int → string
```

Важно: `Itoa` возвращает только одно значение.

```go
text := strconv.Itoa(num)
```

Здесь `err` не нужен.

---

## 5. Главное различие

```text
Atoi → string → int
Itoa → int → string
```

Примеры:

```go
num, err := strconv.Atoi("15")
```

```text
"15" → 15
```

И:

```go
text := strconv.Itoa(15)
```

```text
15 → "15"
```

---

## 6. Использование `Atoi` и `Itoa` вместе

Пример задачи:

```text
"15"
↓ Atoi
15
↓ +5
20
↓ Itoa
"20"
```

Код:

```go
text := "15"

num, err := strconv.Atoi(text)

if err != nil {
    fmt.Println("Ошибка")
    return
}

result := num + 5

textResult := strconv.Itoa(result)

fmt.Println(textResult)
```

---

## 7. Почему ошибку нужно проверять сразу

Нежелательно делать так:

```go
num, err := strconv.Atoi(text)

result := num + 5

if err != nil {
    fmt.Println(err)
}
```

Потому что `num` уже используется до проверки ошибки.

Лучше:

```go
num, err := strconv.Atoi(text)

if err != nil {
    fmt.Println(err)
    return
}

result := num + 5
```

Сначала проверяем ошибку, потом используем результат.

---

## Шпаргалка

```text
strconv.Atoi("42")
→ string в int
→ возвращает num и err
```

```text
strconv.Itoa(42)
→ int в string
→ возвращает только string
```

```text
err == nil
→ ошибки нет
```

```text
err != nil
→ ошибка есть
```

```text
return
→ закончить выполнение текущей функции
```

Главная формула урока:

```text
Atoi → string → int
Itoa → int → string
```