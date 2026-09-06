# Конспект 15. Дженерики (Generics, Go 1.18+)

## 1. Зачем нужны дженерики

До версии 1.18 (2022 год) в Go не было дженериков — для работы с разными типами приходилось либо копировать код под каждый тип, либо использовать `interface{}` с потерей типовой безопасности и type assertion в рантайме. Дженерики решают это на уровне компиляции.

**Проблема без дженериков:**

```go
func SumInts(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func SumFloats(nums []float64) float64 { // дублирование кода под каждый тип
    total := 0.0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## 2. Синтаксис — параметры типа

```go
func Sum[T int | float64](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

fmt.Println(Sum([]int{1, 2, 3}))          // 6, T выведен как int
fmt.Println(Sum([]float64{1.1, 2.2}))      // 3.3, T выведен как float64
```

`[T int | float64]` — объявление параметра типа `T` с **ограничением** (constraint): `T` может быть либо `int`, либо `float64`. `int | float64` — это перечисление допустимых типов через `|`.

## 3. Встроенные constraints и пакет constraints

```go
import "golang.org/x/exp/constraints" // до включения в стандартную библиотеку

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

fmt.Println(Max(3, 5))          // 5
fmt.Println(Max(3.14, 2.71))     // 3.14
fmt.Println(Max("apple", "banana")) // "banana" — строки тоже Ordered (сравнение лексикографическое)
```

`constraints.Ordered` — интерфейс-ограничение, объединяющий все типы, которые можно сравнивать через `<`, `>` (числа, строки).

Собственные пользовательские ограничения объявляются как обычные интерфейсы:

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
```

### comparable — встроенное ограничение

```go
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item { // оператор == доступен благодаря ограничению comparable
            return true
        }
    }
    return false
}

fmt.Println(Contains([]int{1, 2, 3}, 2))          // true
fmt.Println(Contains([]string{"a", "b"}, "c"))     // false
```

`any` (алиас `interface{}`) — тоже валидное ограничение, означающее "любой тип без требований к операциям над ним".

## 4. Обобщённые структуры

Дженерики применимы не только к функциям, но и к типам:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    last := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return last, true
}

intStack := &Stack[int]{}
intStack.Push(1)
intStack.Push(2)
val, ok := intStack.Pop() // val == 2, ok == true

stringStack := &Stack[string]{}
stringStack.Push("hello")
```

`var zero T` — идиоматичный способ получить "нулевое значение" параметризованного типа `T`, каким бы он ни был.

## 5. Обобщённые методы стандартной библиотеки (пакет slices и maps, Go 1.21+)

С версией 1.21 в стандартную библиотеку добавили дженерик-пакеты для типовых операций:

```go
import "slices"

s := []int{3, 1, 2}
slices.Sort(s)                       // [1 2 3]
fmt.Println(slices.Contains(s, 2))    // true
fmt.Println(slices.Index(s, 2))       // 1
reversed := slices.Clone(s)
slices.Reverse(reversed)              // [3 2 1]

import "maps"

m := map[string]int{"a": 1, "b": 2}
keys := maps.Keys(m)   // итератор по ключам (Go 1.23+) или срез в более ранних версиях пакета
```

## 6. Когда НЕ стоит использовать дженерики

Философия Go — не злоупотреблять абстракциями. Дженерики стоит применять, когда:
- Есть реальное дублирование кода под разные типы (контейнеры, математические операции).
- Нужна типобезопасная универсальная структура данных (стек, очередь, дерево).

Дженериками **не стоит** заменять простые интерфейсы там, где хватило бы одного метода — это часто только усложняет читаемость без выигрыша.

---

## Проверь себя

1. Какую проблему решают дженерики, которую нельзя было решить через `interface{}` без потерь?
2. Что означает конструкция `[T int | float64]` в сигнатуре функции?
3. Чем ограничение `comparable` отличается от `any`?
4. Как получить "нулевое значение" параметризованного типа `T` внутри обобщённой функции?
5. Приведи пример, когда использование дженериков оправдано, и пример, когда это избыточно.

**Далее:** Конспект 16 — стандартная библиотека: strings, strconv, time, os, io, bufio.
