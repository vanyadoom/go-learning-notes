# Go Learning Notes

Мой учебный репозиторий по Go.

Здесь я последовательно изучаю Go с нуля, сохраняю конспекты, практические задания, разборы ошибок и небольшие проекты.

Основная цель — не просто изучить синтаксис, а постепенно дойти до уровня **Junior Go Backend Developer / Go Internship** и уметь самостоятельно писать, читать, отлаживать и постепенно проектировать Go-код.

---

## Current status

**Текущий этап:** Standard Library / укрепление Go Fundamentals  
**Текущий урок:** `38`  
**Текущая тема:** пакет `time`  
**Текущая практика:** `practice/38-time`  
**Статус урока 38:** теория изучается, практика впереди

Недавно закрыты:

```text
32 — strings / bytes / rune / UTF-8
33 — range по string
34 — поиск символов
35 — замена символов
36 — package strings
37 — package strconv
38 — package time          ← сейчас здесь
```

Текущая приблизительная статистика:

```text
Освоение уже изученного материала    ~85%

Go Fundamentals                      ~60%
Путь до Junior Go Backend            ~15%
Полный Go / Backend roadmap           ~7%
```

Проценты являются учебной метрикой. Они показывают не количество просмотренных тем, а степень фактически изученного и закреплённого материала.

---

## Текущая позиция

```text
Go Fundamentals
        ↓
Strings / UTF-8
        ↓
Standard Library
        ↓
strings        ✅
strconv        ✅
time           🔄
os             ⏳
io             ⏳
bufio          ⏳
        ↓
JSON / HTTP
        ↓
Testing
        ↓
Concurrency
        ↓
SQL / PostgreSQL
        ↓
Backend Projects
        ↓
Junior Go Backend Developer
```

Сейчас основной язык уже перестаёт быть просто набором отдельных конструкций. Следующая большая цель — перейти от упражнений к использованию стандартной библиотеки, а затем к настоящим backend-задачам.

---

## Как проходит обучение

Обычная структура урока:

```text
1. Цель урока
2. Связь с уже изученным
3. Короткое повторение
4. Новая концепция
5. Prediction
6. Мини-практика
7. Bug Hunt
8. Самостоятельная задача
9. Code Review
10. Закрепление слабых мест
11. Мини-контрольная
12. Конспект
13. Git commit
14. Progress Report
```

Главный принцип:

```text
увидел
→ понял
→ применил
→ ошибся
→ разобрал ошибку
→ повторил
→ начал использовать автоматически
```

Готовые решения заранее не используются, если задачу можно решить самостоятельно.

---

# Уже изучено

## Go basics

- `package main`
- `import`
- `func main`
- переменные
- `var`
- `:=`
- `int`
- `string`
- `bool`
- zero values
- арифметические операторы
- `%`
- `fmt.Println`
- `fmt.Scan`
- `fmt.Scanln`
- `error`
- `nil`
- `_`
- `return`

---

## Control flow

- `if`
- `else`
- `else if`
- операторы сравнения
- `&&`
- `||`
- `!`
- `switch`

---

## Loops

- `for`
- бесконечный `for`
- `break`
- счётчики
- накопители
- `range`
- граничные условия
- обмен значений

```go
a, b = b, a
```

---

## Functions

- объявление функций
- параметры
- аргументы
- возвращаемые значения
- множественный возврат
- вызов функций
- область видимости
- ранний `return`

---

## Arrays and slices

- arrays
- indexes
- `len`
- slices
- `cap`
- `append`
- backing array
- последний индекс
- границы slice

---

## Maps

- создание `map`
- `make`
- чтение и запись
- `value, ok`
- zero value отсутствующего ключа
- `range`
- особенности порядка обхода

---

## Structs

- создание `struct`
- поля
- изменение полей
- копирование struct
- отличие struct от map

---

## Methods

- methods
- receiver
- value receiver
- pointer receiver

---

## Pointers

- `&`
- `*`
- адрес переменной
- разыменование
- изменение значения через pointer

---

# Strings / UTF-8

Пройден отдельный блок по строкам.

Изучено:

- `string`
- `byte`
- `rune`
- UTF-8
- `len(string)`
- `[]rune`
- индексирование строк
- `range` по строке
- поиск символов
- замена символов
- построение новой строки

Ключевая модель:

```text
len(string)
→ общее количество байтов

len([]rune(string))
→ количество Unicode-символов

i в range
→ байтовый индекс начала rune

r в range
→ текущий rune

fmt.Println(r)
→ числовое Unicode-значение

fmt.Println(string(r))
→ сам символ
```

После контрольной тема была отдельно доведена до автоматизма.

Результаты закрепления:

```text
len(string)          9 / 10
i в range           10 / 10
смешанный раунд      8 / 10
финальный стресс     5 / 5
```

---

# Standard Library

## `strings` ✅

Изучено:

```go
strings.Contains
strings.Count
strings.ReplaceAll
strings.ToLower
strings.ToUpper
```

Также изучено комбинирование функций:

```text
ToLower + Contains
→ поиск без учёта регистра

ToLower + Count
→ подсчёт без учёта регистра
```

Мини-контрольная:

```text
5 / 5
```

Практика:

```text
practice/36-strings-package
```

Конспект:

```text
notes/strings-package.md
```

---

## `strconv` ✅

Изучено:

```go
strconv.Atoi
strconv.Itoa
```

Главная модель:

```text
Atoi
string → int

Itoa
int → string
```

Также закреплено:

```text
err == nil
→ ошибки нет

err != nil
→ ошибка есть

return
→ завершить текущую функцию
```

Пример:

```go
num, err := strconv.Atoi(text)

if err != nil {
    return
}

result := num + 5
textResult := strconv.Itoa(result)
```

Мини-контрольная:

```text
4 / 5
```

Практика:

```text
practice/37-strconv
```

Конспект:

```text
notes/strconv.md
```

---

## `time` 🔄

Текущая тема.

Уже изучено:

```go
time.Now()
time.Time
time.Duration

time.Second
time.Minute
time.Hour

time.Date()
time.Add()
time.Sub()
time.Sleep()
```

Главная модель:

```text
time.Time
→ конкретный момент времени
→ когда?

time.Duration
→ временной промежуток
→ сколько времени?
```

Примеры:

```go
now := time.Now()
```

```go
future := now.Add(24 * time.Hour)
```

```go
duration := end.Sub(start)
```

Следующий этап — самостоятельная практика.

Практика:

```text
practice/38-time
```

Конспект:

```text
notes/time.md
```

---

# Practice

Практические задания находятся в:

```text
practice/
```

Последний учебный блок:

```text
practice/32-strings-check
practice/33-string-range
practice/34-string-search
practice/35-string-replace
practice/36-strings-package
practice/37-strconv
practice/38-time
```

Каждая практика должна закреплять одну конкретную идею.

---

# Notes

Конспекты находятся в:

```text
notes/
```

Важные текущие конспекты:

```text
notes/strings-bytes-runes.md
notes/byte-rune-UTF-8-len-range
notes/strings-package.md
notes/strconv.md
notes/time.md
```

Также сохраняются более крупные теоретические материалы и накопительные заметки.

---

# Git workflow

Git является отдельной частью обучения.

Основной цикл:

```text
working tree
↓
git add
↓
staging area
↓
git commit
↓
local repository
↓
git push
↓
GitHub
```

Регулярно используются:

```bash
git status
git diff
git add
git diff --cached
git commit
git log --oneline
git push
```

Цель — довести работу с Git до автоматизма вместе с изучением Go.

---

# Progress tracking

После полноценных уроков отслеживаются:

```text
Текущая тема
Прогресс темы
Go Fundamentals
Junior Go Backend roadmap
Полный roadmap
Результаты практики
Ошибки
Сильные стороны
Что нужно повторить
Следующая тема
```

Дополнительно оцениваются:

```text
Knowledge
→ понимаю концепцию

Code
→ могу написать самостоятельно

Debug
→ могу найти и исправить ошибку

Retention
→ могу вспомнить и применить позже
```

---

# Контрольные проверки

Тестирование используется не ради оценки, а для поиска слабых мест.

Пройден большой тест:

```text
50 вопросов
41 правильный ответ
82%
```

Отдельная контрольная по блоку строк показала слабое понимание `byte / rune / len / range`, после чего тема была полностью повторена и закреплена отдельными тренировочными раундами.

Это отражает главный принцип обучения:

```text
ошибка
→ определить причину
→ вернуться к теме
→ повторить
→ проверить снова
→ двигаться дальше
```

---

# Roadmap

Обучение строится по персональному плану и дополнительно сверяется с:

```text
Develp10/golangroadmap2026
```

Основной маршрут:

```text
Go Fundamentals
↓
Standard Library
↓
Errors / Interfaces
↓
Testing
↓
Concurrency
↓
JSON
↓
HTTP / REST API
↓
SQL / PostgreSQL
↓
Docker / Linux
↓
Backend Projects
↓
Junior Go Backend Developer
```

Более сложные темы не добавляются только ради скорости прохождения roadmap. Сначала должна быть подготовлена необходимая база.

---

# Learning philosophy

Главная цель репозитория — сохранить реальный процесс обучения.

```text
задача
→ попытка
→ ошибка
→ понимание причины
→ исправление
→ повторение
→ самостоятельное применение
```

Цель постепенно перейти от:

```text
"я понимаю этот код"
```

к:

```text
"я могу написать этот код самостоятельно"
```

а затем:

```text
"я могу самостоятельно спроектировать и реализовать приложение"
```

---

# Next milestone

Текущая ближайшая цель:

```text
Закрыть урок 38 — time
↓
os
↓
io
↓
bufio
↓
JSON / HTTP
```

После завершения базовой стандартной библиотеки начнётся переход к первым полноценным backend-задачам.

---

**Status:** Learning in progress  
**Current lesson:** 38 — `time`