# firebase와 연동하여 인증방식 없는 로그인 구현

- 단지 id, pw를 통신하여 로그인을 구현

🔧 사용스택
- Go (framework : gin-gonic)
- DB: firebase의 firestore
- Develop Tool : Visual Studio Code

### mvc 패턴 사용
- controller, model, service, repository 구분

### firestore 사용
- nosql로 객체 그대로 쉽게 삽입 가능

### api 구현
- signup, login 두개의 api구현