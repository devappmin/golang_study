# GOLANG STUDY

## Getting Started.

### formatting package

```go
import "fmt"
```

### Print something

```go
import "fmt"
fmt.Println("Hello world!")
```

### Private / Public

`Public` - 변수 혹은 함수 명을 대문자로 시작할 경우<br>
`Private` - 변수 혹은 함수 명을 소문자로 시작할 경우

### Constant / Variable

`const` 상수<br>
`var`변수

```go
// 둘은 동일
// :=은 새로 만들 때만 사용
var name string = "hello"
name := "hello" // 함수 내부에서만 사용 가능
```

## Function

### General

파라미터 및 리턴 값은 타입을 지정해야 함.<br>

```go
func multiply(a int, b int) int {
    return a * b
}
```

or

```go
func multiply(a, b int) int {
    return a * b
}
```

**리턴 값은 두 개 이상일 수도 있다!!**

```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
    // _ means ignore
    totalLength, upperName := lenAndUpper("Kim Seung Hwan")
    second, _ := lenAndUpper("Kim Seung Hwan")
}
```

**`...`을 통해서 여러 개의 값을 파라미터로 가져올 수 있다.**

```go
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("Kim", "Lee", "Han", "Hong", "Kyu")
}
```

### Naked return

아래와 같이 리턴 값을 미리 정해둘 수 있다.

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func main() {
	// _ means ignore
	totalLength, upperName := lenAndUpper("Kim Seung Hwan")
	fmt.Println(totalLength, upperName)
}
```

### defer

Really cool feature<br>
defer로 입력된 부분은 함수가 끝난 다음에 호출이 됨.

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}
```

## Loop

Go에는 for 이외에 loop를 도는 방법이 없음.

### For each 느낌

```go
// Index only
for number := range numbers {
    fmt.Println(number)
}
```

```go
// Index and number
for index, number := range numbers {
    fmt.Println(index, number)
}
```

### 전통적인 for 느낌

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

## if else

기본 모습

```go
if age < 18 {
    return false
} else {
    return true
}
```

go에서는 if문 내에서도 변수를 선언할 수 있다.

```go
// Variable expression
koreanAge := age + 2
if koreanAge < 18 {
    return false
}
// Is equal as below
if koreanAge := age + 2; koreanAge < 18 {
    return false
}
```

## Switch

```go
switch age {
case 16:
    return false
case 18:
    return true
}
```

go에서는 switch안에 expression이 들어가도 됨.

```go
switch {
case age < 16:
    return false
case age == 18:
    return true
case age > 50:
    return false
}
```

또한 `if else`와 동일하게 `Variable expression`이 가능하다.

```go
switch koreanAge:= age + 2; koreanAge {
    // Put case here..
}
```

## Pointer

기본적으로 C와 동일

```go
func pointerExample() {
    a := 5
    b := &a
    fmt.Println(&a) // Address of a
    fmt.Println(b)  // Same as &a
    fmt.Println(*b) // point value of b
}
```

## Array

### 정적 할당 배열

```go
func arrayExample() {
    // [length]type{init values}
    names := [5]string{"hello", "world", ":)"}
    names[3] = "sup"
    names[4] = "have a nice day!"
}
```

### Slices

동적으로 크기를 변경하는 배열을 생성할 수도 있다.

```go
func slicesExample() {
    // 정적 배열을 생성할 때에서 배열의 크기 부분만 없애면 됨.
    names := []string{"hello", "world", ":)"}

    // 값 추가는 append 함수를 사용해서 추가하면 됨.
    names = append(names, "sup")
}
```

## Map

```go
variable := map[string]string{"Hello": "World", "Nice": "day!"}
// map[key]value{init values}
```

다른 언어들과 같이 `map`을 `for loop`에 돌릴 수 있음.

```go
for key, value := range variable {
    fmt.Println(key, "is", value)
}
```

## struct

`C`나 `C++`처럼 `struct`를 사용하여 구조체를 생성할 수 있음.

```json
{
  "name": "kim",
  "age": 18
}
```

위와 같이 value 부분에 dynamically하게 object가 들어가게 하는 방법은 `struct`를 사용하는 것.

구조체에는 메소드를 생성할 수 있음.

### 구조체 선언

만드는 방식은 다른 언어들과 비슷하나 `,`를 추가해서는 안됨.

```go
type person struct {
	name string
	age int
	favFood []string
}
```

### 구조체 생성

생성하는 것도 다른 언어랑 비슷하나 map과 같이 key값을 명시하여 출력하는 것을 선호.

```go
favFood := []string{"sushi", "kimbap", "ramen", "rice", "eggplant"}
// 명확하지 않음.
kim := person{"kim seung hwan", 26, favFood}

// key 값을 명시해줌.
kim := person{name:"kim seung hwan", age:26, favFood: favFood}

// 전체 출력
fmt.Println(kim)

// 일부 출력
fmt.Println(kim.name)
```

## After basic course.

`go`에는 클래스가 없다.
