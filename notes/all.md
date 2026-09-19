# Go — основы, функции и структуры

## 1. Структура программы Go

Минимальная программа:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

### `package main`

Указывает, что файл относится к пакету `main`.

Именно пакет `main` используется для запускаемой программы.

### `import`

Подключает пакеты.

Например:

```go
import "fmt"
```

### `func main()`

`main()` — точка входа в программу.

Выполнение программы начинается с неё.

---

# 2. Переменные

Объявление:

```go
var age int
```

Можно сразу задать значение:

```go
var age int = 25
```

Короткое объявление:

```go
age := 25
```

`:=` одновременно создаёт переменную, определяет её тип и присваивает значение.

`:=` используется внутри функций.

---

# 3. Базовые типы

Основные типы:

```go
int
string
bool
```

Примеры:

```go
var age int = 25
var name string = "Ivan"
var ok bool = true
```

`bool` может иметь два значения:

```text
true
false
```

---

# 4. Нулевые значения

Если переменная объявлена без значения, Go автоматически задаёт zero value.

```go
var n int
```

Получит:

```text
0
```

```go
var name string
```

Получит:

```text
""
```

```go
var ok bool
```

Получит:

```text
false
```

---

# 5. Арифметические операции

Основные операции:

```go
+
-
*
/
%
```

`%` возвращает остаток от деления.

```go
10 % 2 // 0
11 % 2 // 1
```

Проверка чётности:

```go
n%2 == 0
```

Проверка нечётности:

```go
n%2 != 0
```

---

# 6. Вывод

Используется пакет `fmt`.

```go
fmt.Println("Привет")
```

Несколько значений:

```go
fmt.Println("Возраст:", age)
```

`Println` добавляет перевод строки.

Для приглашения ко вводу удобно:

```go
fmt.Print("Введите число: ")
```

---

# 7. Ввод через `fmt.Scan`

```go
var age int

fmt.Scan(&age)
```

Можно считывать несколько значений:

```go
var name string
var age int

fmt.Scan(&name, &age)
```

`Scan` записывает введённые данные в переданные переменные.

Поэтому передаётся адрес:

```go
fmt.Scan(&age)
```

`&age` — адрес переменной `age`.

Практическое правило:

```text
fmt.Scan(&variable)
```

---

# 8. Ошибки `fmt.Scan`

`fmt.Scan` может вернуть количество успешно считанных значений и ошибку:

```go
n, err := fmt.Scan(&a, &b)
```

Если количество считанных значений не нужно:

```go
_, err := fmt.Scan(&a, &b)
```

Проверка:

```go
if err != nil {
	fmt.Println("Ошибка ввода")
	return
}
```

`nil` означает отсутствие ошибки.

`_` используется, когда возвращаемое значение намеренно не нужно.

---

# 9. `return`

`return` прекращает выполнение текущей функции.

Пример:

```go
if err != nil {
	fmt.Println("Ошибка")
	return
}
```

После `return` функция больше не выполняется.

`return` также может вернуть значение:

```go
return value
```

Например:

```go
func square(n int) int {
	return n * n
}
```

Здесь функция одновременно завершается и возвращает результат.

### Главное различие

```text
return
→ выйти из функции

return value
→ выйти из функции и вернуть значение
```

---

# 10. Условия

```go
if age >= 18 {
	fmt.Println("Совершеннолетний")
}
```

`if` выполняет код только при истинном условии.

### Операторы сравнения

```go
==
!=
>
<
>=
<=
```

`=` — присваивание.

```go
age = 25
```

`==` — сравнение.

```go
age == 25
```

### `if / else`

```go
if age >= 18 {
	fmt.Println("Совершеннолетний")
} else {
	fmt.Println("Несовершеннолетний")
}
```

### `else if`

Используется для нескольких вариантов:

```go
if age < 13 {
	fmt.Println("Ребёнок")
} else if age < 18 {
	fmt.Println("Подросток")
} else {
	fmt.Println("Взрослый")
}
```

Условия проверяются сверху вниз.

---

# 11. Логические операторы

### `&&` — И

Обе части должны быть истинными:

```go
age >= 18 && age < 65
```

### `||` — ИЛИ

Достаточно одной истинной части.

### `!` — отрицание

```go
!true // false
```

---

# 12. `switch`

`switch` позволяет выбрать один из нескольких вариантов.

```go
switch day {
case 1:
	fmt.Println("Понедельник")
case 2:
	fmt.Println("Вторник")
case 3:
	fmt.Println("Среда")
default:
	fmt.Println("Неизвестный день")
}
```

Несколько значений в одном `case`:

```go
switch day {
case 1, 2, 3, 4, 5:
	fmt.Println("Рабочий день")
case 6, 7:
	fmt.Println("Выходной")
default:
	fmt.Println("Ошибка")
}
```

`default` используется, если ни один `case` не подошёл.

---

# 13. Цикл `for`

Классический вариант:

```go
for i := 1; i <= 10; i++ {
	fmt.Println(i)
}
```

Состав:

```text
начало
условие
изменение
```

### `i++`

Увеличивает значение на 1:

```go
i++
```

эквивалентно:

```go
i = i + 1
```

или:

```go
i += 1
```

### Бесконечный цикл

```go
for {
	...
}
```

### `break`

Останавливает текущий цикл:

```go
for {
	fmt.Scan(&n)

	if n == 0 {
		break
	}
}
```

`break` выходит из цикла.

`return` выходит из функции.

---

# 14. Счётчик

Счётчик используется для подсчёта количества элементов.

```go
count := 0

count++
```

Например:

```go
if n%2 == 0 {
	count++
}
```

Здесь считается количество чётных чисел.

---

# 15. Накопитель

Накопитель используется для накопления результата.

```go
sum := 0
```

Затем:

```go
sum += n
```

что эквивалентно:

```go
sum = sum + n
```

Разница:

```text
count++
→ количество

sum += n
→ сумма значений
```

---

# 16. Диапазоны

Диапазон от `a` до `b` включительно:

```go
for i := a; i <= b; i++ {
	...
}
```

Если `a > b`, значения можно поменять местами:

```go
a, b = b, a
```

После этого:

```go
for i := a; i <= b; i++ {
	...
}
```

будет проходить диапазон от меньшего к большему.

---

# 17. Область видимости

Переменная доступна только в своей области видимости.

Например:

```go
if true {
	x := 10
	fmt.Println(x)
}
```

После `if` переменная `x` недоступна.

Переменная цикла:

```go
for i := 0; i < 10; i++ {
	...
}
```

существует в области этого `for`.

---

# 18. Функции

Функция — именованный блок кода, который выполняет отдельную задачу.

Пример:

```go
func add(a, b int) int {
	return a + b
}
```

Вызов:

```go
result := add(5, 7)
```

Получим:

```text
result = 12
```

### Параметры

В:

```go
func add(a, b int) int
```

`a` и `b` — параметры.

### Аргументы

В:

```go
add(5, 7)
```

`5` и `7` — аргументы.

---

# 19. Несколько параметров одного типа

Можно писать:

```go
func multiply(a, b int) int {
	return a * b
}
```

вместо:

```go
func multiply(a int, b int) int {
	return a * b
}
```

Это одинаковые варианты записи.

---

# 20. Несколько возвращаемых значений

Go позволяет функции возвращать сразу несколько значений.

```go
func calculate(a, b int) (int, int) {
	return a + b, a * b
}
```

Получение результатов:

```go
sum, product := calculate(6, 4)
```

Получим:

```text
sum = 10
product = 24
```

Порядок важен.

```go
return a + b, a * b
```

означает:

```text
первое значение → sum
второе значение → product
```

Если переменные уже существуют:

```go
sum, product = calculate(a, b)
```

используется `=`, а не `:=`.

---

# 21. Результат + `error`

Очень распространённый паттерн Go:

```go
result, err := someFunction()
```

Функция возвращает:

```text
результат
+
ошибку
```

Пример:

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}

	return a / b, nil
}
```

При успешном выполнении:

```go
result, err := divide(10, 2)
```

получим:

```text
result = 5
err = nil
```

`nil` означает, что ошибки нет.

---

# 22. Проверка `error`

Стандартный шаблон:

```go
result, err := divide(a, b)

if err != nil {
	fmt.Println("Ошибка:", err)
	return
}

fmt.Println("Результат:", result)
```

Логика:

```text
получить result и err
↓
err != nil?
├── да → обработать ошибку → return
└── нет → использовать result
```

После ошибки `return` прекращает выполнение функции.

Важно:

> Обнаружить ошибку недостаточно. Нужно понять, можно ли после неё продолжать выполнение.

Если продолжение небезопасно, используется `return`.

---

# 23. Создание ошибки

Для создания ошибки можно использовать:

```go
errors.New("сообщение")
```

Например:

```go
return 0, errors.New("деление на ноль")
```

Также можно использовать:

```go
fmt.Errorf("сообщение")
```

---

# 24. Разделение ответственности

Хорошая функция выполняет одну понятную задачу.

Например:

```text
main
 ↓
получить ввод
 ↓
вызвать функцию
 ↓
получить результат / error
 ↓
обработать ошибку
 ↓
вывести результат
```

А функция:

```text
получить данные
 ↓
выполнить конкретную задачу
 ↓
return
```

Например, `divide` должна заниматься делением, а `main` — взаимодействием с пользователем.

---

# 25. Именованные возвращаемые значения

Go позволяет назвать возвращаемые значения:

```go
func calculate(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b

	return
}
```

`sum` и `diff` становятся переменными внутри функции.

Пустой:

```go
return
```

возвращает их текущие значения.

Можно также вернуть значения явно:

```go
func minMax(a, b int) (min int, max int) {
	if a > b {
		return b, a
	}

	return a, b
}
```

Именованные результаты не требуют обязательного использования пустого `return`.

---

# 26. `struct`

`struct` объединяет несколько связанных данных в один тип.

```go
type Student struct {
	Name  string
	Age   int
	Grade int
}
```

`Student` — новый пользовательский тип.

Поля:

```text
Name
Age
Grade
```

имеют свои типы.

---

# 27. Создание `struct`

Можно создать значение структуры:

```go
student := Student{
	Name:  "Ivan",
	Age:   29,
	Grade: 5,
}
```

Доступ к полям:

```go
student.Name
student.Age
student.Grade
```

Например:

```go
fmt.Println(student.Name)
fmt.Println(student.Age)
fmt.Println(student.Grade)
```

Результат:

```text
Ivan
29
5
```

---

# 28. Создание структуры через функцию

Структуру можно создавать внутри отдельной функции:

```go
func createStudent(name string, age, grade int) Student {
	return Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
```

В `main`:

```go
student := createStudent(name, age, grade)
```

Схема:

```text
ввод
 ↓
createStudent(...)
 ↓
Student
 ↓
student.Name
student.Age
student.Grade
```

---

# 29. Передача `struct` в функцию

Структуру можно передавать в функцию как обычное значение:

```go
func isAdult(student Student) bool {
	return student.Age >= 18
}
```

Вызов:

```go
adult := isAdult(student)
```

Результат:

```text
true
```

или:

```text
false
```

Для поведения, логически принадлежащего определённому типу, в Go используются методы.

---

# 30. Обмен значений

В Go можно поменять значения двух переменных местами:

```go
a, b = b, a
```

Например:

```text
a = 10
b = 4
```

после:

```go
a, b = b, a
```

получим:

```text
a = 4
b = 10
```

Это удобно для нормализации диапазона.

---

# 31. Типичные схемы

### Простая функция

```go
func add(a, b int) int {
	return a + b
}
```

```go
result := add(a, b)
```

### Несколько результатов

```go
func calculate(a, b int) (int, int) {
	return a + b, a - b
}
```

```go
sum, diff := calculate(a, b)
```

### Результат + ошибка

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}

	return a / b, nil
}
```

```go
result, err := divide(a, b)

if err != nil {
	fmt.Println(err)
	return
}
```

### `struct`

```go
type Student struct {
	Name  string
	Age   int
	Grade int
}
```

```go
student := createStudent(name, age, grade)
```

---

# 32. Что нужно уверенно понимать

```text
переменные
↓
условия
↓
циклы
↓
счётчики и накопители
↓
функции
↓
параметры
↓
return
↓
несколько возвращаемых значений
↓
result + error
↓
именованные результаты
↓
struct
```

Главный принцип:

> Сначала понять, что должна делать программа, затем разделить задачу на понятные действия и только после этого писать код.

---

# 33. Методы

Метод — функция, которая связана с определённым типом.

Например:

```go
type Student struct {
	Name  string
	Age   int
	Grade int
}
```

Метод:

```go
func (s Student) IsAdult() bool {
	return s.Age >= 18
}
```

Теперь метод вызывается через объект:

```go
student.IsAdult()
```

а не как обычная функция:

```go
isAdult(student)
```

### Receiver

Часть:

```go
(s Student)
```

называется **receiver**.

`s` — имя receiver.

`Student` — тип, к которому относится метод.

---

# 34. Метод без параметров

Пример:

```go
func (s Student) IsAdult() bool {
	return s.Age >= 18
}
```

Вызов:

```go
adult := student.IsAdult()
```

Метод использует данные самого `Student`.

Внутри метода:

```go
s.Age
```

означает возраст текущего студента.

---

# 35. Метод с параметром

Метод может принимать дополнительные параметры.

```go
func (s Student) IsGradeAtLeast(minGrade int) bool {
	return s.Grade >= minGrade
}
```

Вызов:

```go
gradeOK := student.IsGradeAtLeast(4)
```

Здесь:

```text
s
→ текущий Student

minGrade
→ обычный параметр метода
```

---

# 36. Value receiver

Метод:

```go
func (s Student) IsAdult() bool
```

использует **value receiver**.

На базовом уровне полезно думать так:

```text
Student
→ метод получает значение структуры
→ удобно, когда мы только читаем данные
```

Например:

```go
func (s Student) IsAdult() bool {
	return s.Age >= 18
}
```

Такой метод не должен изменять исходный `Student`.

---

# 37. Pointer receiver

Если метод должен изменять исходную структуру, используется pointer receiver:

```go
func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}
```

Вызов:

```go
student.SetGrade(5)
```

После вызова:

```text
student.Grade = 5
```

Практическое правило:

```text
только читаем Student
→ Student

изменяем Student
→ *Student
```

---

# 38. Поле и метод — разные вещи

Поле структуры:

```go
student.Grade
```

Это значение.

Метод:

```go
student.SetGrade(5)
```

Это действие.

Поэтому:

```go
oldGrade := student.Grade
```

получает значение поля.

А:

```go
student.SetGrade(5)
```

изменяет структуру.

---

# 39. Метод, который возвращает значение

Пример:

```go
func (s Student) IsExcellent() bool {
	return s.Grade >= 5
}
```

Можно:

```go
excellent := student.IsExcellent()
```

потому что метод возвращает `bool`.

Получится:

```text
excellent = true
```

или:

```text
excellent = false
```

---

# 40. Метод, который ничего не возвращает

Пример:

```go
func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}
```

Правильно:

```go
student.SetGrade(5)
```

Неправильно:

```go
newGrade := student.SetGrade(5)
```

Почему?

Потому что `SetGrade` ничего не возвращает.

Получить новое значение можно отдельно:

```go
student.SetGrade(5)

newGrade := student.Grade
```

Главная идея:

```text
метод может изменить объект
≠
метод обязательно что-то возвращает
```

---

# 41. Методы, изменяющие состояние

Например:

```go
func (s *Student) SetAge(age int) {
	s.Age = age
}
```

Использование:

```go
oldAge := student.Age
student.SetAge(30)
```

Теперь:

```text
oldAge = 24
student.Age = 30
```

Состояние структуры изменилось.

---

# 42. Порядок выполнения и состояние объекта

Программа выполняется последовательно.

Например:

```go
adult := student.IsAdult()

student.SetAge(30)
```

`adult` содержит результат проверки **в момент вызова** `IsAdult()`.

Он не пересчитывается автоматически после изменения `student`.

Если нужно проверить новое состояние:

```go
student.SetAge(30)
adult = student.IsAdult()
```

То же правило действует для других вычисленных значений.

---

# 43. Метод, возвращающий строку

Метод может формировать строку:

```go
func (s Student) Info() string {
	return fmt.Sprintf(
		"Имя: %s, возраст: %d, оценка: %d",
		s.Name,
		s.Age,
		s.Grade,
	)
}
```

Вызов:

```go
info := student.Info()
```

или:

```go
fmt.Println(student.Info())
```

`Info()` только читает данные, поэтому использует:

```go
func (s Student) Info() string
```

### `fmt.Sprintf`

`fmt.Sprintf` похож на `fmt.Printf`, но не выводит строку на экран.

Он **возвращает готовую строку**:

```go
text := fmt.Sprintf("Возраст: %d", age)
```

---

# 44. Метод с несколькими условиями

Метод может объединять уже известные проверки.

Например:

```go
func (s Student) CanGraduate() bool {
	return s.Age >= 18 && s.Grade >= 4
}
```

Оба условия должны быть истинными:

```text
Age >= 18
И
Grade >= 4
```

Только тогда:

```text
CanGraduate() → true
```

---

# 45. Метод с `error`

Метод может не только изменять структуру, но и возвращать ошибку.

Например:

```go
func (s *Student) SetGrade(grade int) error {
	if grade < 1 || grade > 5 {
		return errors.New("Неверная оценка")
	}

	s.Grade = grade
	return nil
}
```

Логика:

```text
grade < 1 или grade > 5
↓
error
↓
Grade НЕ изменяется

1 <= grade <= 5
↓
Grade изменяется
↓
nil
```

---

# 46. Проверка `error` в методе, изменяющем объект

Вызов:

```go
err := student.SetGrade(5)

if err != nil {
	fmt.Println("Ошибка:", err)
	return
}
```

Если значение недопустимо:

```go
err != nil
```

и изменение структуры не происходит.

Если всё правильно:

```text
err == nil
```

и структура изменяется.

---

# 47. Сначала проверка, потом изменение

Очень важен порядок действий.

Правильно:

```go
func (s *Student) SetGrade(grade int) error {
	if grade < 1 || grade > 5 {
		return errors.New("Неверная оценка")
	}

	s.Grade = grade
	return nil
}
```

Неправильно:

```go
func (s *Student) SetGrade(grade int) error {
	s.Grade = grade

	if grade < 1 || grade > 5 {
		return errors.New("Неверная оценка")
	}

	return nil
}
```

Во втором варианте неправильное значение уже записано в структуру до проверки.

Общий принцип:

```text
проверка
↓
ошибка?
├── да → return error
└── нет → изменить состояние
```

---

# 48. `struct` + несколько методов

Один тип может иметь несколько методов.

Например:

```go
type Student struct {
	Name  string
	Age   int
	Grade int
}
```

Методы:

```go
func (s Student) IsAdult() bool {
	return s.Age >= 18
}
```

```go
func (s Student) IsGradeAtLeast(minGrade int) bool {
	return s.Grade >= minGrade
}
```

```go
func (s Student) IsExcellent() bool {
	return s.Grade >= 5
}
```

```go
func (s Student) CanGraduate() bool {
	return s.Age >= 18 && s.Grade >= 4
}
```

Изменяющие методы:

```go
func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}
```

```go
func (s *Student) SetAge(age int) {
	s.Age = age
}
```

Можно представить тип так:

```text
Student
├── IsAdult()             → читает
├── IsGradeAtLeast(...)   → читает
├── IsExcellent()         → читает
├── CanGraduate()         → читает
├── Info()                → читает
├── SetGrade(...)         → изменяет
└── SetAge(...)            → изменяет
```

---

# 49. Разница между value receiver и pointer receiver

### Value receiver

```go
func (s Student) IsAdult() bool
```

Используется, когда метод читает данные.

Упрощённо:

```text
Student
→ получаем значение
→ читаем данные
```

### Pointer receiver

```go
func (s *Student) SetGrade(grade int)
```

Используется, когда метод должен изменить исходный объект.

Упрощённо:

```text
*Student
→ работаем с исходным объектом
→ меняем данные
```

Практическое правило для текущего уровня:

```text
читаем → Student

изменяем → *Student
```

---

# 50. Важное различие: действие и результат

Метод может:

### Только вернуть результат

```go
adult := student.IsAdult()
```

### Только изменить объект

```go
student.SetGrade(5)
```

### Изменить объект и вернуть ошибку

```go
err := student.SetGrade(10)
```

Поэтому всегда нужно смотреть на объявление метода:

```go
func (...) ...
```

Часть после `)` показывает, что метод возвращает.

Например:

```go
func (s Student) IsAdult() bool
```

возвращает `bool`.

А:

```go
func (s *Student) SetGrade(grade int)
```

ничего не возвращает.

А:

```go
func (s *Student) SetGrade(grade int) error
```

возвращает `error`.

---

# 51. Типичные ошибки при работе с методами

### Вызывать поле как функцию

Неправильно:

```go
student.Grade()
```

Правильно:

```go
student.Grade
```

### Использовать `SetGrade` как возвращающую функцию

Неправильно:

```go
newGrade := student.SetGrade(5)
```

если `SetGrade` ничего не возвращает.

Правильно:

```go
student.SetGrade(5)
newGrade := student.Grade
```

### Изменять структуру через value receiver

Неправильно для метода, который должен менять исходный объект:

```go
func (s Student) SetGrade(grade int) {
	s.Grade = grade
}
```

Для изменения исходного `Student` используется:

```go
func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}
```

### Игнорировать `error`

Нежелательно:

```go
student.SetGrade(10)
```

если функция возвращает `error`.

Правильнее:

```go
err := student.SetGrade(10)

if err != nil {
	fmt.Println("Ошибка:", err)
	return
}
```

---

# 52. Общая схема работы с `struct` и методами

```text
получить данные
↓
создать struct
↓
вызвать методы чтения
↓
при необходимости изменить struct
↓
обработать error
↓
получить новое состояние
↓
вывести результат
```

Например:

```text
name, age, grade
↓
Student
↓
IsAdult()
↓
IsGradeAtLeast()
↓
SetGrade()
↓
student.Grade
```

---

# 53. Контрольная схема текущего уровня

### Поле

```go
student.Grade
```

→ значение внутри структуры.

### Метод чтения

```go
student.IsAdult()
```

→ возвращает результат.

### Метод с параметром

```go
student.IsGradeAtLeast(4)
```

→ получает дополнительное значение и возвращает результат.

### Метод изменения

```go
student.SetGrade(5)
```

→ изменяет существующий `Student`.

### Метод изменения + ошибка

```go
err := student.SetGrade(10)
```

→ пытается изменить `Student` и сообщает, удалось ли это.

---

# 54. Что нужно уверенно понимать после этого этапа

```text
переменные
типы
zero values
fmt.Scan
fmt.Println
&
_
if / else / else if
switch
for
break
++
+=
%
счётчики
накопители
диапазоны
обмен значений
область видимости

functions
parameters
arguments
return
несколько возвращаемых значений
result + error
named returns

struct
поля
создание структуры
передача структуры
методы
receiver
value receiver
pointer receiver
методы с параметрами
методы с return
методы без return
изменение состояния
error в методах
порядок изменения состояния
```

---

# 55. Главный принцип текущего этапа

При разборе программы всегда задавай себе четыре вопроса:

```text
1. Что хранится в переменной?

2. Что возвращает функция или метод?

3. Изменяет ли метод исходный объект?

4. В каком состоянии находится объект в этот момент?
```

Большая часть ошибок, которые мы находили в практике, была связана именно с:

```text
неправильным порядком действий
↓
неправильным использованием результата
↓
изменением состояния
↓
игнорированием error
↓
путаницей между полем и методом
```
## Работа над ошибками после контрольной

Результат контрольной: **60 / 100**.

Основные темы для закрепления:

- `&` и работа с `fmt.Scan`
- `error`, `nil`, `return`
- `%`
- `break` vs `return`
- счётчик и накопитель
- параметры и аргументы
- value receiver и pointer receiver
- методы и обычные функции
- проверка данных перед изменением состояния
- порядок вычисления и сохранения результатов

Следующий этап — углублять работу со структурами и переходить к следующим фундаментальным структурам данных Go.
