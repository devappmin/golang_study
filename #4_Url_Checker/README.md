## http

### GET

`get` 메소드를 실행하기 위해서는 `http`를 이용해야 한다.

```go
resp, err := http.Get(url)
```

이렇게 얻은 `response`의 값에서 `status code`는 아래 방식으로 얻을 수 있다.

```go
resp.StatusCode
```

## make

`map`과 같은 것은 초기화를 해야하고 초기화를 하지 않았을 시에 오류가 발생한다.

```go
// OK
results := map[string]string{}
var results = make(map[string]string)

// Panic
results := map[string]string
var results map[string]string
```

위에서 `Panic`부분은 값을 추가하려고 하면 정상적으로 진행되지 않는다.
