# 연산자
- +,-,*,/,%
- 비트연산: &,|,^,&^,<<,>>
    - &^: 비트 클리어, 10&^2는 ^2를 먼저계산을 하고 10과 &연산을 한다
    - ^2: ...111101, 10&^2: 1000 => 2에대한 비트만 clear됨
- 논리연산: &&,||,!
- 값교환: a,b=b,a
    - a, b := function(c,d)와 같이 자동으로 return값이 match되는 것도 가능
- 복합 대입연산자: +=,-=, ...
- 증감 연산자: ++, --
    - c언어에서의 전위, 후위 증감연산자는 없음
    - b = a++과 같이사용하면 a++은 return 값이 없기에 값이 안들어감
- 그외 연산자
    - 배열 요소 접근: []
    - 구조체나 패키지 요소 접근: .
    - 메모리 주솟값: &
    - 메모리 주소 접근: *
    - 슬라이스 요소 접근 or 가변인수 생성: ...
    - 배열의 일부분 가져오기: :
    - 채널에서 값을 빼거나 넣을때 사용: <-

# overflow
- 모든 type은 메모리크기가 정해져있기에 숫자에 맞는 type을 사용해야한다.

# 실수 오차
- 컴퓨터는 100% 자연상태의 실수를 표현할 수 없다.
- 0.1 + 0.2 == 0.3은 false
    - 이를 해결하느 방법
        - math.Nextafter(a,b) == b: Nextafter은 b를 항해 1비트만 바뀐다. 즉 a와 b가 1비트만 차이난다면 equal이라고 할 수 있는 것이다.

# 함수
- func Add(a int, b int) int {  return a + b }
    - type이 같은 parameter끼리 묶어서, Add(a,b int)로 작성도 가능
    - a int, b float, c int와 a, c int, b float은 parameter로 보면 똑같은 의미지만, parameter를 넣을때 type순서가 다름
    - return값이 2개 이상일때 (int, bool)과 같이 작성 가능, cf) return시 return 100, true
    - return 값에 변수명을 붙여 할당 후 단지 return만 해주면 됨: funcExample.go의 divide1 func
- 메모리 차원에서의 함수
    - 함수를 부른 명령어에서 stack에 함수를 부른 명령어의 다음 명령어 위치, 파라미터1, 파라미터2, ... 순서로 넣음
        - push 51(호출한 위치), push 3, push 9, jump 1000(함수 위치)
        - 점프 후 pop해서 parameter, 9, 3사용
        - return시 return값도 마찬가지로 push하고 호출한 위치로 jump해서 다시 pop하여 변수에 return값 할당
    - 함수를 가서 파라미터 개수만큼 stack pop해서 사용하고 return 시에 함수를 부른 명령어의 다음 명령어 위치로 다시 점프해서 진행

# 상수
- const a int = 10
- iota
    - iota는 한 const 변수마다 +1이 되어 할당이 된다. ex) constExample.go
    - 1 << iota로 사용하여 비트 계산도 쉽게 가능하다
- type없는 상수는 type을 설정안해줘도 자동으로 casting 됨