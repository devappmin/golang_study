## Comments

### 구조체 export 시

상단에 Comments를 생성해아하며 첫 문자는 구조체 명을 적어야 함.

```go
// Account Struct
type Account struct {

}
```

## go.mod

go 디렉토리 내부에서 코딩을 하는 것이 아니라 로컬 디렉토리에서 모듈을 새로 만드니까 `import`를 할 수 없는 오류가 생기면 `go.mod`를 통해서 해결해야 함.

1. 프로젝트 Root Directory에서 `go.mod` 생성

   ```terminal
   $ go mod init main
   ```

   go mod를 초기화 하여 현재 디렉토리를 main으로 한다.

2. 생성하고자 하는 Module Directory 내에서 `go.mod` 생성

   ```terminal
   $ cd banking
   $ go mod init banking
   ```

   go mod를 생성한 뒤에 다시 Root Directory로 이동한다.

3. 프로젝트 Root Directory에서 Module Directory 변경

   ```mod
   replace main.com/banking => ./banking
   ```

   `main.com/banking`을 import 할 경우 현재 directory 내에 있는 `banking` directory를 자동으로 import 하게 함.

4. Module Linking
   ```terminal
   $ go mod tidy
   ```
   해당 명령어를 입력하면 자동으로 연결이 된다.

## Constructor

`go`의 구조체는 constructor가 없으므로 함수로 직접 빼서 만들어야 한다.

```go
// NewAccount creates a new account.
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}
```

## Method

`go`에서 메소드를 생성하기 위해서는 `func`와 메소드 이름 사이에 구조체를 적어주면 된다.

```go
func (a Account) Deposit(amount int) {
    // (a Account) means receiver.
    // 보통은 receiver의 명을 struct 명의 첫 글자를 소문자로 한다.
}
...
account // Account type object.
account.Deposit() // 이렇게 method가 된다.
```

값을 변경하기 위해서는 receiver가 pointer야 한다.

```go
func (a *Account) Deposit(amount int) {
    // a를 포인터로 받아와야한다.
    // 이유는 C나 C++와 동일.
}
```

### String()

`String()` 메소드는 해당 구조체를 print할 때 출력하는 값을 의미한다. `python`의 `**str**``와 동일.

```go
func (a Account) String() string {
    return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
```

## Error handling

```go
// Withdraw x amount from your account.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw from account")
	}
	a.balance -= amount

	// nil means no error
	return nil
}
```

`errors.Error()` 혹은 `errors.New()`로 에러를 생성하면 된다. 에러가 없으면 `nil`을 리턴해주는 것으로 처리.

```go
err := account.Withdraw(20)
if err != nil {
    log.Fatalln(err)
}
```

`log.Fatalln`은 프로그램을 출력 후 종료시켜준다.

```go
var errNoMoney = errors.New("Can't withdraw from account")
```

에러 변수를 만들 때는 변수 이름을 `err`로 시작해야한다.
