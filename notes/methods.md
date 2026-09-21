# Go — методы (Methods)

## 1. Что такое метод

**Метод** — это функция, привязанная к определённому типу.

Например, есть структура:

```go
type Student struct {
    Name  string
    Age   int
    Grade int
}
```

Обычная функция:

```go
func printStudent(student Student) {
    fmt.Println(student.Name)
}
```

Вызов:

```go
printStudent(student)
```

Метод:

```go
func (student Student) Print() {
    fmt.Println(student.Name)
}
```

Вызов:

```go
student.Print()
```

Главное отличие:

```text
обычная функция:
printStudent(student)

метод:
student.Print()
```

---

# 2. Receiver

В методе появляется специальная часть:

```go
(student Student)
```

Например:

```go
func (student Student) Print() {
    fmt.Println(student.Name)
}
```

Эта часть называется **receiver** — получатель метода.

Она связывает метод с типом `Student`.

Структура:

```go
func (receiver Type) MethodName() {
    ...
}
```

Например:

```go
func (s Student) Print() {
    fmt.Println(s.Name)
}
```

Здесь:

* `s` — имя receiver;
* `Student` — тип receiver;
* `Print` — имя метода.

---

# 3. Имя receiver

Receiver можно назвать по-разному:

```go
func (student Student) Print()
```

или:

```go
func (s Student) Print()
```

Оба варианта корректны.

Если receiver называется `s`, внутри метода используется:

```go
s.Name
s.Age
s.Grade
```

Обычно в Go для receiver используют короткое имя:

```go
func (s Student) Print()
```

---

# 4. Метод без возвращаемого значения

Метод может просто выполнять действие:

```go
func (s Student) Print() {
    fmt.Println(s.Name, s.Age, s.Grade)
}
```

Вызов:

```go
student.Print()
```

---

# 5. Метод с возвращаемым значением

Метод может возвращать результат:

```go
func (s Student) GetGrade() int {
    return s.Grade
}
```

Использование:

```go
grade := student.GetGrade()
```

Метод может возвращать любой подходящий тип:

```go
func (s Student) IsAdult() bool {
    return s.Age >= 18
}
```

Использование:

```go
if student.IsAdult() {
    fmt.Println("Совершеннолетний")
}
```

---

# 6. Метод с параметрами

Метод может принимать дополнительные параметры:

```go
func (s Student) CheckGrade(minGrade int) bool {
    return s.Grade >= minGrade
}
```

Вызов:

```go
student.CheckGrade(4)
```

Здесь:

```text
student
   ↓
receiver

CheckGrade(4)
          ↑
       параметр
```

---

# 7. Value receiver

Метод может иметь receiver типа:

```go
Student
```

Например:

```go
func (s Student) SetGrade(grade int) {
    s.Grade = grade
}
```

Здесь `s` — **копия** исходной структуры.

Поэтому:

```go
student.SetGrade(4)
```

не изменит исходный `student`.

Условно:

```text
исходный Student
      │
      └── копия → s
```

Изменяется копия.

---

# 8. Pointer receiver

Если метод должен изменять исходную структуру, используется pointer receiver:

```go
func (s *Student) SetGrade(grade int) {
    s.Grade = grade
}
```

Теперь:

```go
student.SetGrade(4)
```

изменяет исходный `student`.

После вызова:

```go
fmt.Println(student.Grade)
```

будет:

```text
4
```

---

# 9. Value receiver vs pointer receiver

### Value receiver

```go
func (s Student) Method()
```

Получает копию.

```text
Student
  ↓
копия
  ↓
s
```

Изменения `s` не изменяют исходную структуру.

### Pointer receiver

```go
func (s *Student) Method()
```

Работает с исходным значением через указатель.

```text
Student
  ↓
указатель
  ↓
исходный объект
```

Изменения могут сохраняться в исходной структуре.

---

# 10. Вызов метода с pointer receiver

Если есть:

```go
func (s *Student) SetGrade(grade int) {
    s.Grade = grade
}
```

можно написать:

```go
student.SetGrade(4)
```

Не обязательно вручную писать:

```go
(&student).SetGrade(4)
```

В обычном случае Go автоматически берёт адрес переменной при вызове такого метода.

---

# 11. Методы могут только читать данные

Pointer receiver нужен не всегда.

Например:

```go
func (s Student) IsAdult() bool {
    return s.Age >= 18
}
```

Метод ничего не изменяет.

Он только читает данные и возвращает результат.

Поэтому здесь достаточно value receiver.

---

# 12. Методы могут изменять данные

Если метод должен изменить структуру:

```go
func (s *Student) SetGrade(grade int) {
    s.Grade = grade
}
```

Используется pointer receiver.

Например:

```go
student.SetGrade(4)
```

---

# 13. Методы и struct

Часто `struct` и методы используются вместе.

Например:

```go
type Student struct {
    Name  string
    Age   int
    Grade int
}

func (s Student) Print() {
    fmt.Println(s.Name, s.Age, s.Grade)
}

func (s Student) IsAdult() bool {
    return s.Age >= 18
}

func (s *Student) SetGrade(grade int) {
    s.Grade = grade
}
```

Получается:

```text
Student
│
├── данные
│   ├── Name
│   ├── Age
│   └── Grade
│
└── поведение
    ├── Print()
    ├── IsAdult()
    └── SetGrade()
```

---

# 14. Метод нельзя объявить внутри struct

Нельзя:

```go
type Student struct {
    Name string

    func Print() {
        ...
    }
}
```

Внутри `struct` находятся поля.

Методы объявляются отдельно:

```go
type Student struct {
    Name string
}

func (s Student) Print() {
    fmt.Println(s.Name)
}
```

---

# 15. У одного типа может быть много методов

Например:

```go
type Student struct {
    Name  string
    Age   int
    Grade int
}
```

У него могут быть:

```go
func (s Student) Print()
```

```go
func (s Student) IsAdult() bool
```

```go
func (s Student) GetGrade() int
```

```go
func (s *Student) SetGrade(grade int)
```

Все они являются методами типа `Student`.

---

# 16. Методы и обычные функции

Функция:

```go
func printStudent(s Student) {
    fmt.Println(s.Name)
}
```

Вызов:

```go
printStudent(student)
```

Метод:

```go
func (s Student) Print() {
    fmt.Println(s.Name)
}
```

Вызов:

```go
student.Print()
```

Метод можно воспринимать как функцию, которая **принадлежит определённому типу**.

---

# 17. Методы не являются классами

В Go нет классов в классическом смысле Java/C++.

Вместо этого используются:

```text
struct + methods
```

Например:

```go
type User struct {
    Name  string
    Email string
}

func (u User) Print() {
    fmt.Println(u.Name, u.Email)
}
```

То есть данные находятся в `struct`, а поведение реализуется методами.

---

# 18. Методы у собственных типов

Методы можно объявлять не только для структур.

Можно определить собственный тип:

```go
type Number int
```

И метод:

```go
func (n Number) Double() Number {
    return n * 2
}
```

Использование:

```go
n := Number(5)

fmt.Println(n.Double())
```

Результат:

```text
10
```

Нельзя просто добавить метод непосредственно стандартному `int`.

---

# 19. Экспортируемые методы

Имя метода с заглавной буквы:

```go
func (s Student) Print()
```

экспортируется из пакета.

Имя с маленькой буквы:

```go
func (s Student) print()
```

не экспортируется.

Это связано с правилами видимости Go.

---

# 20. Главное правило для запоминания

```text
Student
```

в receiver:

```go
func (s Student) Method()
```

→ работаем с копией.

```text
*Student
```

в receiver:

```go
func (s *Student) Method()
```

→ работаем с исходной структурой через указатель.

---

# 21. Мини-шпаргалка

```go
func (s Student) Print()
```

Метод типа `Student`.

```go
student.Print()
```

Вызов метода.

```go
func (s Student) GetGrade() int
```

Метод возвращает `int`.

```go
func (s Student) CheckGrade(min int) bool
```

Метод принимает параметр и возвращает `bool`.

```go
func (s *Student) SetGrade(grade int)
```

Pointer receiver; метод может изменить исходную структуру.

---

# 22. Что нужно знать на текущем этапе

Перед дальнейшим изучением важно уверенно понимать:

1. Что такое метод.
2. Чем метод отличается от обычной функции.
3. Что такое receiver.
4. Как вызывается метод.
5. Что такое value receiver.
6. Что такое pointer receiver.
7. Почему `Student` получает копию.
8. Почему `*Student` позволяет менять исходную структуру.
9. Что методы могут иметь параметры и возвращаемые значения.
10. Что `struct + methods` — один из основных способов организации типов в Go.
