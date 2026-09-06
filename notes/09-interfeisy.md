# Конспект 9. Интерфейсы

## 1. Что такое интерфейс в Go

Интерфейс — это набор сигнатур методов. Любой тип, у которого есть **все** эти методы, автоматически считается реализующим интерфейс — **без явного объявления** (никакого `implements`, как в Java/C#).

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Rectangle и Circle НИГДЕ явно не сказано "implements Shape" —
// они автоматически подходят под интерфейс, потому что имеют нужные методы.

func printInfo(s Shape) {
    fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimeter())
}

printInfo(Rectangle{Width: 4, Height: 5})
printInfo(Circle{Radius: 3})
```

Это называется **структурная типизация** или **duck typing** ("если ходит как утка и крякает как утка — это утка"): важно не заявленное родство типов, а фактическое наличие нужного поведения.

## 2. Пустой интерфейс interface{} и any

```go
var x interface{} // может содержать значение ЛЮБОГО типа
x = 42
x = "строка"
x = Rectangle{}

// С Go 1.18 есть алиас — читается лучше:
var y any = 42
```

Пустой интерфейс — способ написать функцию, принимающую значение любого типа (пример — `fmt.Println(a ...interface{})`), но использовать его нужно осторожно: теряется проверка типов на этапе компиляции, приходится делать type assertion в рантайме.

## 3. Type assertion — "достать" конкретный тип из интерфейса

```go
var s Shape = Circle{Radius: 5}

c, ok := s.(Circle) // "безопасная" форма — ok == false, если тип не совпал (а не паника)
if ok {
    fmt.Println("Это круг, радиус:", c.Radius)
}

c2 := s.(Circle) // "небезопасная" форма — PANIC, если тип не Circle!
```

**Правило:** используй форму с `ok`, если не уверен на 100% в типе значения. Форму без `ok` — только когда логика гарантирует конкретный тип (и паника в случае ошибки — это на самом деле сигнал о баге, который лучше не скрывать).

## 4. type switch — обработка нескольких возможных типов

```go
func describe(i interface{}) string {
    switch v := i.(type) {
    case int:
        return fmt.Sprintf("целое число: %d", v)
    case string:
        return fmt.Sprintf("строка длиной %d", len(v))
    case Shape:
        return fmt.Sprintf("фигура с площадью %.2f", v.Area())
    case nil:
        return "значение отсутствует (nil)"
    default:
        return fmt.Sprintf("неизвестный тип: %T", v)
    }
}
```

## 5. Маленькие интерфейсы — идиома Go

**Важный принцип культуры Go**, который часто спрашивают на собеседовании: "принимай интерфейсы, возвращай структуры" (accept interfaces, return structs), и интерфейсы должны быть максимально маленькими.

Классический пример из стандартной библиотеки:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Композиция маленьких интерфейсов в больший:
type ReadWriter interface {
    Reader
    Writer
}
```

`io.Reader` и `io.Writer` — вероятно, самые важные интерфейсы стандартной библиотеки. Файлы, сетевые соединения, буферы в памяти, HTTP-тела запросов — всё это реализует `Reader`/`Writer`, что позволяет писать код, не зависящий от конкретного источника данных.

## 6. Интерфейс как контракт для тестирования (Dependency Injection)

Одна из главных практических причин использовать интерфейсы — возможность подменить реализацию в тестах:

```go
type UserRepository interface {
    GetByID(id int) (*User, error)
}

type PostgresRepository struct { /* реальное подключение к БД */ }
func (r *PostgresRepository) GetByID(id int) (*User, error) { /* SQL-запрос */ return nil, nil }

type MockRepository struct{} // используется только в тестах
func (r *MockRepository) GetByID(id int) (*User, error) {
    return &User{Name: "Тестовый пользователь"}, nil
}

type UserService struct {
    repo UserRepository // зависит от ИНТЕРФЕЙСА, а не от конкретной БД
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}
```

В продакшене передаём `&PostgresRepository{}`, в тестах — `&MockRepository{}`. `UserService` вообще не знает, откуда на самом деле берутся данные.

## 7. nil-интерфейс — известная ловушка

**Очень частый вопрос на собеседованиях**, вызывает удивление у новичков:

```go
type MyError struct{}
func (e *MyError) Error() string { return "моя ошибка" }

func doSomething() error {
    var err *MyError = nil
    // ... код, который не устанавливает err ...
    return err // ВОЗВРАЩАЕТ НЕ nil!
}

result := doSomething()
fmt.Println(result == nil) // false !!
```

**Почему так происходит:** интерфейс в Go внутри себя хранит ДВЕ вещи — тип значения и само значение. `err` — это `*MyError(nil)`: тип известен (`*MyError`), значение — nil. Когда это присваивается интерфейсу `error`, интерфейс становится "(*MyError, nil)" — а это **не то же самое**, что интерфейс без типа и значения (`(nil, nil)`), который и является настоящим `nil`.

**Правильный способ** — явно возвращать `nil` литералом интерфейса, а не через типизированную nil-переменную:

```go
func doSomething() error {
    var myErr *MyError
    if somethingWentWrong {
        myErr = &MyError{}
    }
    if myErr != nil {
        return myErr
    }
    return nil // явный, "чистый" nil
}
```

---

## Проверь себя

1. Как тип "узнаёт", что реализует интерфейс — через явное объявление или автоматически?
2. Чем безопасная форма type assertion (`v, ok := x.(T)`) отличается от небезопасной (`v := x.(T)`)?
3. Почему в Go принято делать интерфейсы маленькими (1-2 метода)?
4. Зачем в тестах передают mock-реализацию интерфейса вместо реальной?
5. Почему `err != nil` может быть true, даже если в переменную явно писали `nil`?

**Далее:** Конспект 10 — обработка ошибок, panic/recover, errors.Is/As.
