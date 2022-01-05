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
