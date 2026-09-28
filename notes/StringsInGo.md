# Пакет `strings` в Go

Пакет `strings` содержит готовые функции для работы со строками.

Подключение:

```go
import "strings"
```

## 1. `strings.Contains`

Проверяет, содержится ли одна строка внутри другой.

```go
strings.Contains("Golang", "Go")
```

Результат:

```text
true
```

Возвращаемый тип:

```go
bool
```

То есть результат — `true` или `false`.

Важно: функция чувствительна к регистру.

```go
strings.Contains("Golang", "go")
```

вернёт:

```text
false
```

Коротко:

```text
Contains → есть ли подстрока
```

---

## 2. `strings.Count`

Считает количество непересекающихся вхождений подстроки.

```go
strings.Count("banana", "a")
```

Результат:

```text
3
```

Пример:

```go
strings.Count("banana", "na")
```

Результат:

```text
2
```

Функция возвращает число:

```text
Count → сколько раз встретилась подстрока
```

Пример:

```go
strings.Count("aaaa", "aa")
```

Результат:

```text
2
```

Совпадения считаются без пересечений:

```text
aa | aa
```

---

## 3. `strings.ReplaceAll`

Заменяет все найденные вхождения одной строки на другую.

```go
strings.ReplaceAll("banana", "a", "*")
```

Результат:

```text
b*n*n*
```

Другой пример:

```go
strings.ReplaceAll("hello hello", "hello", "Go")
```

Результат:

```text
Go Go
```

Коротко:

```text
ReplaceAll → заменить все вхождения
```

`ReplaceAll` возвращает новую строку.

---

## 4. `strings.ToLower`

Переводит строку в нижний регистр.

```go
strings.ToLower("Go LANG")
```

Результат:

```text
go lang
```

Коротко:

```text
ToLower → нижний регистр
```

---

## 5. `strings.ToUpper`

Переводит строку в верхний регистр.

```go
strings.ToUpper("Go lang")
```

Результат:

```text
GO LANG
```

Коротко:

```text
ToUpper → верхний регистр
```

---

## 6. Регистр имеет значение

Большинство операций со строками чувствительны к регистру.

Например:

```go
strings.Contains("Hello", "he")
```

вернёт:

```text
false
```

Потому что:

```text
"H" != "h"
```

То же самое относится к `Count` и `ReplaceAll`.

---

## 7. Поиск без учёта регистра

Если нужно искать слово независимо от регистра, можно сначала привести обе строки к одному регистру.

```go
text = strings.ToLower(text)
word = strings.ToLower(word)

fmt.Println(strings.Contains(text, word))
```

Например:

```text
text = "I Love Go"
word = "go"
```

после `ToLower`:

```text
text = "i love go"
word = "go"
```

После этого:

```go
strings.Contains(text, word)
```

вернёт:

```text
true
```

---

## 8. Подсчёт без учёта регистра

```go
text = strings.ToLower(text)
word = strings.ToLower(word)

fmt.Println(strings.Count(text, word))
```

Например:

```text
text = "Go go GO"
word = "go"
```

После нормализации:

```text
go go go
```

Результат:

```text
3
```

---

## 9. `Contains` + `Count`

Можно сначала проверить наличие слова:

```go
if strings.Contains(text, word) {
    fmt.Println(strings.Count(text, word))
} else {
    fmt.Println(0)
}
```

Логика:

```text
Contains → есть ли слово
Count    → сколько раз оно встретилось
```

---

## 10. Замена без учёта регистра

Один простой вариант:

```go
text = strings.ToLower(text)
old = strings.ToLower(old)

result := strings.ReplaceAll(text, old, new)
```

Важно:

```text
text → нормализуем
old  → нормализуем
new  → обычно не трогаем
```

`new` — это итоговое значение, которое мы хотим вставить.

Например:

```go
new := "JAVA"
```

Если сделать:

```go
new = strings.ToLower(new)
```

то мы потеряем нужный регистр и получим `"java"`.

---

## Главная шпаргалка

```text
strings.Contains   → есть ли подстрока → bool

strings.Count      → сколько раз встретилась → int

strings.ReplaceAll → заменить все вхождения → string

strings.ToLower    → нижний регистр → string

strings.ToUpper    → верхний регистр → string
```

Комбинации:

```text
ToLower + Contains
→ поиск без учёта регистра

ToLower + Count
→ подсчёт без учёта регистра

ToLower + ReplaceAll
→ простой вариант замены без учёта регистра
```

## Главное помнить

Функции пакета `strings` обычно не изменяют исходную строку напрямую.

Например:

```go
strings.ToLower(text)
```

само по себе не изменит `text`.

Нужно сохранить результат:

```go
text = strings.ToLower(text)
```

То же самое относится к другим функциям, возвращающим новую строку.