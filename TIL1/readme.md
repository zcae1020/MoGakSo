# go 개발환경 설정

- vs code를 이용한 go 개발환경 설정(windows) : https://code.visualstudio.com/docs/languages/go
- https://www.youtube.com/watch?v=KBdz5c-0t1w&list=PLy-g2fnSzUTBHwuXkWQ834QHDZwLx6v6j&index=1 
을 보는 것을 추천(한국어)

# go 실행 명령어

- go run main.go: main.go 파일 실행
- go mod init example: main module 파일 만들기
- go build: 실행파일 만들기(module을 만들고 난 이후에 build 가능)

# helloworld 찍기

- package main: main package 불러오기
- import 'fmt': fmt 불러오기
- func main() : main함수 생성

# GO 특징
- 클래스, 상속, 제너릭 프로그래밍, 네임스페이스 기능 x
- 메서드, 인터페이스, 익명함수, 가비지 컬렉터, 포인터 기능 o
- 클래스 대신 메서드를 가지는 구조체 지원
- 익명함수는 함수 리터럴이라는 이름으로 제공

# Go type
- 숫자
    - 정수
        - uint8, uint16, uint32, uint64
        - int8, int16, int32, int64
        - byte = uint8
        - rune = int32
        - int, 32bit -> int32, 64bit -> int64
        - uint, 32bit -> uint32, 64bit -> uint64
    - 실수
        - float32, float64
    - 복소수
        - complex64: 8byte, complext128: 16byte
- bool
- string
- array
- slice
- struct
- pointer
- function
- map
- interface
- channel

### 숫자, bool, string 타입별 기본값은 c와 동일하지만 나머지는 nil이라는 정의되지않은 메모리 주소를 나타내는 것으로 기본값이 됨

# Go는 casting이 굉장히 중요하다.
- 타입이 다르면 연산이 안됨 cf) int와 int64끼리도 안됨
- casting 방법
    - ex) int64(variable)
- casting 시 주의사항
    - 실수->정수: 소수점 짤림
    - 큰 메모리를 사용하는 수에서 작은 메모리를 사용하는 type으로 변경시 수가 이상해질 가능성 존재

### 실수를 가지고 계산할 때 소수부를 생각해야함
- float32는 7자리, float64는 15자리를 표현가능함
- ex) 1234567.8901234라는 숫자를 float32를 통해 표현한다고 하면 1.234567e6으로 표현됨 => 총 7자리만 표현되므로 소수점 뒤의 정보를 사라짐

# Go lang 문법
- 함수 정의
    - func name() {}
- 주석
    - //
    - /* */
- 변수선언
    - var a int = 20
    - var a int // a는 0임
    - var a = 10 //type을 안정할시 선언시 초기값이 있어야함
    - a:=10