## ECHO를 통한 웹 통신

### install echo

```terminal
$ go get github.com/labstack/echo
```

### 프로젝트에 echo 추가

`main.go`

```go
import "github.com/labstack.echo"
```

```terminal
$ go mod tidy
```

### echo 오브젝트 생성 및 GET 처리

```go
func handleHome(c echo.Context) error {
    return .String(http.StatusOK, /* Put value here*/ )
}
...
e := echo.New()
e.GET("/", handleHome)
```

### echo 서버 시작

```go
e.Logger.Fatal(e.Start(":PORT_NO"))
```
