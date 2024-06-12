# Устные вопросы

## 1. Какой самый эффективный способ конкатенации строк?
`string.Builder`\
Builder используется для эффективного формирования строки с помощью методов Write. Он минимизирует копирование памяти. Нулевое значение готово к использованию.
```go
var sb strings.Builder

for i := 0; i < 1000; i++ {
    sb.WriteString("a")
}

fmt.Println(sb.String())
```
## 2. Что такое интерфейсы, как они применяются в Go?
Интерфейс — это набор методов, которые могут быть реализованы типом. Иными словами, интерфейс — описание того, что может сделать тип. 

1. Основное назначение интерфейсов — реализация полиморфизма: с одной стороны, тип может реализовывать несколько интерфейсов в разных контекстах применения, с другой стороны, у нас есть возможность написания алгоритмов, работающих с разными типами данных.
2. C помощью интерфейсов можно написать код, абстрагированный от внешних модулей: при изменениях в них ничего не нужно переделывать в своём коде, и наоборот. 
3. Интерфейсы добавляют гибкости и снижают связность кода.

## 3. Чем отличаются RWMutex от Mutex?
RWMutex (Read-Write Mutex) - это взаимоисключающая блокировка читателя и писателя. Блокировка может принадлежать произвольному числу читателей или одному писателю.
Другими словами, читателям не нужно ждать друг друга. Они должны ждать только писателей, которые держат блокировку.
RWMutex предпочтительнее для данных, которые в основном читаются, а ресурс, который экономится по сравнению с Mutex, - это время.

## 4. Чем отличаются буферизированные и не буферизированные каналы?
По умолчанию каналы не буферизованы, это означает, что они будут принимать отправления (chan <-), только если есть соответствующий прием (<- chan), готовый принять отправленное значение. Буферизованные каналы принимают ограниченное количество значений без соответствующего приемника для этих значений. Буферизированные каналы сохраняют получаемые данные в виде кольцевой очереди. Операция отправки блокируется только тогда, когда буфер заполнен, а операция приёма блокируется только тогда, когда буфер пуст.

## 5. Какой размер у структуры struct{}{}? 
0 - размер зависит от типов, из которых он состоит, и порядка расположения полей в структуре. Компилятор оптимизирует пустые структуры, присваивая им размер 0 байт.

## 6. Есть ли в Go перегрузка методов или операторов?
Нет. 

### [Почему Go не поддерживает перегрузку методов и операторов?](https://go.dev/doc/faq#overloading)
Отправка метода упрощается, если не требуется также выполнять сопоставление по типам. Опыт работы с другими языками показал нам, что иногда полезно иметь различные методы с одним и тем же именем, но с разными подписями, но на практике это также может быть запутанным и хрупким. Сопоставление только по названию и требование единообразия типов было основным упрощающим решением в системе типов Go.

Что касается перегрузки операторов, то это кажется скорее удобством, чем абсолютным требованием. Опять же, без него все проще.

## 7. В какой последовательности будут выведены элементы map[int]int?
При итерации по мапе с помощью цикла range порядок итераций не задается и не гарантируется, что он будет одинаковым от одной итерации к другой. С момента выхода Go 1.0 среда выполнения рандомизирует порядок итераций мап. Программисты стали полагаться на стабильный порядок итераций ранних версий Go, который варьировался в разных реализациях, что приводило к ошибкам совместимости. Если вам нужен стабильный порядок итераций, вы должны поддерживать отдельную структуру данных, определяющую этот порядок.

## 8. В чем разница make и new?
New не инициализирует память, а только обнуляет ее. Он возвращает указатель на только что выделенное нулевое значение.
Make создает только слайсы, мапы и каналы и возвращает их инициализированными.

## 9. Сколько существует способов задать переменную типа slice или map?
```go
var m map[int]int // 1 способ
var m = map[int]int{} // 2 способ
m := make(map[int]int, 0) // 3 способ
```

## 10. Что выведет данная программа и почему?
```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```

Программа выведет: 
```
1
1
```
Аргументы функций передаются по значению. Когда указатель p передается в функцию update, создается копия этого указателя. Изменение копии указателя внутри функции update не влияет на оригинальный указатель p в функции main.

## 11. Что выведет данная программа и почему?
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```

```
fatal error: all goroutines are asleep - deadlock!
```
WaitGroup передается по значению в анонимную функцию. В результате чего, каждая горутина получает копию WaitGroup, а вызов wg.Done() внутри горутины уменьшает счётчик у этой копии, а не у оригинального WaitGroup. wg.Wait блокирует завершение программы, но, так как счетчик не станет равным нулю, произойдет deadlock.

## 12. Что выведет данная программа и почему?
```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Программа выведет: 
```
0
```
Внутри блока if создаётся новая локальная переменная n, которая "затеняет" переменную n, объявленную на уровне функции main. Изменения, произведённые с локальной переменной n внутри блока if, не влияют на переменную n снаружи блока.

## 13. Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
Программа выведет:
```
[100 2 3 4 5]
```
v[0] = 100 изменяет первый элемент слайса v, который является ссылкой на слайс a. Изначальный слайс a имел длину 5 и емкость также 5, поэтому операция append приведет к выделению нового массива для слайса v, но это уже не повлияет на исходный слайс a.

## 14. Что выведет данная программа и почему?
```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
Программа выведет:
```
[b b a][a a]
```
slice имеет длину 2 и емкость тоже 2, внутри анонимной функции при добавлении третьего элемента создается новый массив, и slice теперь указывает на новый слайс. Первый и второй элемент нового слайса меняется на "b". Оригинальный слайс slice остается неизменным в основной функции.