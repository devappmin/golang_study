## Concurrency

### Goroutines

단순히 함수를 호출할 때 앞에다가 `go`를 붙여주면 됨.

```go
func main() {
    go hello("kim")
    hello("lee")
}

func hello(name string) {
    fmt.Println("Hello," name)
}
```

하지만 `main function`이 종료가 되면 병렬처리가 되고 있던 작업들도 전부 종료됨.

### Channels

`pipe`를 통해서 값을 전달하는 것을 `channel`이라고 한다.

```go
c := make(chan bool) // chan type
go myFunc(..., c)
```

위와 같은 방식으로 channel을 만들고 `goroutine` 함수에 인자값으로 넣어주면 된다.

그 후에는 함수 안에서 인자값을 처리해주면 된다.

```go
func myFunc(..., c chan bool) {
    // Your codes here..
    c <- true
}
```

```go
// chan 대신 chan<-으로 적어서 send only로 만들 수 있음.
func myFunc(..., c chan<- bool) {
    // Your codes here..

    fmt.Prinlnt(<-c) // error!
}
```

값을 넣어줄 때에는 `<-` 화살표를 통해서 넣어주면 된다.

```go
c := make(chan bool)
...
result := <-c
// 바로 출력할 경우 fmt.Println(<-c)
```

동일하게 `channel`의 값을 변수로 넣어줄 때도 `<-`를 활용하면 된다.<br>
이 떄는 호출을 한 함수(예제의 경우 `main`)에서 값을 가져올 때까지 기다린다.<br>
`C`의 `pthread_join`과 `Dart`의 `await`와 비슷하게 다음 코드로 넘어가지 않는 것 같음.

```go
c := make(chan bool)

people := [2]string{"kim", "lee"}

for _, person := range people {
    go isAwesome(person, c)
}

fmt.Println(<-c)
fmt.Println(<-c)
```

위와 같이 `channel`을 두 번 호출했으면 값을 두 번 받을 수 있음. 만약에 **두 번 호출**하고 **세 번 이상** 값을 받으려고 하면 `deadlock`에 걸림.
