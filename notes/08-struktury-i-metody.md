# Конспект 8. Структуры и методы

## 1. Объявление и создание структур

```go
type User struct {
    Name    string
    Age     int
    Email   string
}

// Способы создания:
u1 := User{Name: "Alex", Age: 25, Email: "a@mail.com"} // именованные поля — рекомендуется
u2 := User{"Alex", 25, "a@mail.com"}                    // позиционно — хрупко, не рекомендуется
u3 := User{}                                             // все поля — нулевые значения
u4 := &User{Name: "Alex"}                                // указатель на структуру, остальные поля — нули
```

Именованная инициализация (`u1`) — идиоматичный стандарт в Go: позволяет менять порядок полей структуры без риска сломать код, а также опускать не важные поля.

## 2. Вложенные структуры

```go
type Address struct {
    City, Street string
}

type User struct {
    Name    string
    Address Address // вложенная структура
}

u := User{
    Name: "Alex",
    Address: Address{
        City:   "Moscow",
        Street: "Lenina",
    },
}
fmt.Println(u.Address.City)
```

## 3. Встраивание (embedding) — замена наследованию

В Go нет наследования классов. Вместо этого — **встраивание** (composition): структура может содержать анонимное поле другого типа, и тогда его поля/методы становятся "видны" напрямую у внешней структуры.

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " издаёт звук"
}

type Dog struct {
    Animal // встраивание — без имени поля, только тип
    Breed  string
}

d := Dog{
    Animal: Animal{Name: "Рекс"},
    Breed:  "Овчарка",
}

fmt.Println(d.Name)   // "Рекс" — поле Animal.Name доступно напрямую через d!
fmt.Println(d.Speak()) // метод Animal.Speak() тоже "поднялся" на уровень Dog
```

**Важно понимать:** это не наследование в ООП-смысле — это синтаксический сахар, который автоматически "прокидывает" (promotes) доступ к полям/методам вложенного типа. `Dog` не является подтипом `Animal` — это разные типы, между ними нет полиморфизма, только удобный доступ к полям.

Если у `Dog` есть свой метод с тем же именем, что и у `Animal` — метод `Dog` перекрывает (shadowing) метод `Animal`:

```go
func (d Dog) Speak() string {
    return d.Name + " лает"
}
fmt.Println(d.Speak())        // "Рекс лает" — метод Dog
fmt.Println(d.Animal.Speak()) // "Рекс издаёт звук" — явный доступ к методу Animal
```

## 4. Методы: value receiver vs pointer receiver

Метод в Go — это функция с "приёмником" (receiver), объявленным перед именем функции:

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver — получает КОПИЮ структуры
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver — получает указатель на оригинал, может изменять поля
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area()) // 50

rect.Scale(2) // Go автоматически берёт &rect — писать rect.Scale можно и для value-переменной
fmt.Println(rect.Width) // 20
```

**Правило выбора receiver (частый вопрос на собеседовании):**
- Используй **pointer receiver**, если метод должен изменять состояние структуры, или если структура большая (копирование дорого), или если структура содержит поля, которые не должны копироваться (например, `sync.Mutex`).
- Используй **value receiver**, если структура маленькая, неизменяемая по своей природе (например, `time.Time`, координаты точки), и метод только читает данные.
- **Если хотя бы один метод типа использует pointer receiver — принято делать все методы этого типа pointer receiver'ами** для консистентности.

## 5. Сравнение структур

Структуры сравнимы оператором `==`, если все их поля сравнимы (не содержат слайсов, мап, функций):

```go
type Point struct{ X, Y int }

p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2) // true — сравнение "по значению" всех полей
```

## 6. Теги структур (struct tags) — забегая вперёд к JSON (конспект 17)

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"` // omitempty — не включать поле в JSON, если оно пустое
    age   int    `json:"-"`                // неэкспортируемое поле — JSON-пакет его не видит вообще
}
```

Теги — это просто строки метаданных, прикреплённые к полю, которые читаются через reflection (пакетами вроде `encoding/json`, `encoding/xml`, ORM-библиотеками).

## 7. Конструкторы (идиома, а не языковая конструкция)

В Go нет ключевого слова `constructor` — принято писать функцию `NewXxx`, которая возвращает готовую структуру (обычно указатель), инкапсулируя валидацию/значения по умолчанию:

```go
type Config struct {
    Timeout int
    Retries int
}

func NewConfig(timeout int) *Config {
    if timeout <= 0 {
        timeout = 30 // значение по умолчанию
    }
    return &Config{
        Timeout: timeout,
        Retries: 3,
    }
}

cfg := NewConfig(0) // Timeout будет 30, не 0
```

---

## Проверь себя

1. Чем встраивание (embedding) в Go отличается от классического наследования в ООП?
2. Что произойдёт, если у Dog и встроенного Animal есть метод с одинаковым именем?
3. По какому правилу выбирают между value receiver и pointer receiver?
4. Почему структуру со слайсом внутри нельзя сравнить оператором `==`?
5. Что делает тег `json:"-"` у поля структуры?

**Далее:** Конспект 9 — интерфейсы, duck typing, type assertion.
