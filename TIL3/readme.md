# 내부, 외부 package import

## GOPATH, GOROOT

- Go에는 여러 환경변수가 있다. go env를 하면 전체의 환경변수 목록을 볼 수 있다.
- 그중 GOPATH, GOROOT는 package와 src를 담당하는 path를 담고있다.

---

## 외부 라이브러리 추가

[Tutorial: Get started with Go - The Go Programming Language](https://go.dev/doc/tutorial/getting-started)

```bash
go mod init [example/hello]
```

다음의 명령어는 example/hello라는 이름의 mod파일을 생성한다.

이 이름은 나중에 mod파일이 있는 directory를 다른 코드가 불러올 때 사용되니 중복이 안되도록 작성하여야한다.

- [pkg.go.dev](https://pkg.go.dev/search?q=quote)여기에서 쓰고 싶은 패키지를 찾는다.
- 패키지 이름 위에 조그마하게 [rsc.io/quote/v4](http://rsc.io/quote/v4)이렇게 적혀있는데 버젼넘버인 v4를 빼고 “rsc.io/quote”를 import해주면 된다.
- import하고 quote에 해당하는 method를 쓰려면 quote를 객체라 생각하고 밑의 코드와 같이 작성한다.

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}
```

- 그전에 만들어 놓은 mod 파일을 새롭게 추가된 "rsc.io/quote"도 반영하기 위해 다음의 코드를 shell에서 사용한다.
    
    ```bash
    go mod tidy
    ```
    
- 실행하면 원하는 값이 나온다.

---

## 내부 라이브러리 추가(내가 만든 내부 패키지)

- 같은 디렉토리(패키지)내에 있다면 그냥 그대로 import도 필요없이 사용하면 된다.
- 다른 디렉토리에 있을 시
    - 일단 두 디렉토리에 mod파일을 둘다 만들어야한다.
    - 그 후 module을 가져올 파일에 module이름에 맞게 import를 해준다.
    - 다음의 맞게 코드를 작성하여 shell에 작성한다.
    
    ```bash
    go mod edit -replace [module이름]=[module까지의 상대주소]
    ex) go mod edit -replace module/config=./module/config
    ```
    
    - 그 후 go get [module이름]한 후 실행을 돌리면된다.
- 참고로, 대문자로 시작하면 public, 소문자로 시작하면 private이기에 공유하고 싶다면 대문자로 시작해야한다.

---

## JSON path 설정

- 다른 디렉토리에 있는 json파일을 가져올때 상대주소를 사용한다고 한다면,

**현재 코드를 작성하고 있는 작업공간이 아닌,** 

**mod 파일(main 파일…?, 정확한것은 아직 모르겠다.)이 있는 곳을 기준으로** 상대주소를 적어야한다.

---

### alias

import하는데 import할 것 앞에 alias를 붙여 별칭같이 사용가능하다.

ex) m "module"