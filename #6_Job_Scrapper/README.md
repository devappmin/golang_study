## Go query

`jQuery`와 비슷한데 `Go`버전의 라이브러리임.

```terminal
$ go get github.com/PuerkitoBio/goquery
```

### 페이지에서 읽기

```go
res, err := http.Get(/* Your URL Here.. */)

doc, err := goquery.NewDocumentFromReader(res.Body)

doc.Find("Class name of website")
```

해당 클래스의 값을 갖는 오브젝트를 생성하는 방법이다.

```go
doc.Find("Class name of website").Each(func(i int, s *goquery.Selection) {
    // Each values..
    // i => index
    // s => each values
})
```

얻은 오브젝트에서 값을 하나씩 돌아가면서 loop를 도는 방법이다.

`Find`가 아니라 `Attribute`를 통해서도 값을 얻을 수 있다.

```go
val, exists = doc.Attr("Attributes")
```

## String

### String Conversion

`strconv`를 사용하여 값을 String으로 변경할 수 있음.

```go
strconv.Itoa(5000)
```

### strings.TrimSpace

`TrimSpace`를 이용해서 string 앞뒤에 있는 빈 공간을 제거할 수 있음.

```go
strings.TrimSpace("Hello\n\nWor    ld")
```

### strings.Fields

텍스트를 텍스트의 배열로 나눠줌.

```go
// Return value is string[]
strings.Fields(strings.TrimSpace("Hello\n\nWor     ld"))
```

### strings.Join

텍스트 배열을 하나로 합치면서 `sep`을 넣어줌.

```go
strings.Join(strings.Fields(strings.TrimSpace("Hello\n\nWor     ld")), " ")
```

### TrimSpace + Fields + Join

```go
strings.Join(strings.Fields(strings.TrimSpace("   Hello\n\nWor     ld   ")), " ")
```

를 실행할 경우 아래와 같이 변경 됨.

1. TrimSpace
   ```go
   "Hello\n\nWor     ld"
   ```
2. Fields
   ```go
   "Hello", "Wor", "ld"
   ```
3. Join
   ```go
   "Hello Wor ld"
   ```

## CSV

### create file

```go
file, err := os.Create()
```

### create New CSV / Flush it

```go
w := csv.NewWriter(file)
defer w.Flush()
```

### CSV에 쓰기

```go
header := []string{"id", "title", "location", "company", "salary", "summary", "upload"}
wErr := w.Write(header)
```
