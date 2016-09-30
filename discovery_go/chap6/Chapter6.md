
# 웹 애플리케이션 작성하기

## Hello, 새계!

## 할 일 목록 관리 웹 앱 만들기

### RESTful API
  * 클라이언트/서버, Stateless
  * net/url package: URL 경로 및 쿼리 추출

### Data Access Object
  * 데이터베이스에 필요한 연산을 추상 인터페이스로 만들어서 사용하는 것
    * 데이터베이스와 비지니스 로직을 구분하기 위한 목적

### RESTful API 핸들러 구현

## 코드 리팩토링
  * 코드가 일단 동작하면 다시 지금까지 작성한 코드를 읽어보고 수정하는 것이 좋습니다.

### 통일성 있게 파일 나누기
  * 코드 나누기
    * 데이터 액세스 인터페이스 (accessor.go)
    * 메모리 데이터 엑세스 구현 (mem_accessor.go)
    * 응답 자료형 및 구현 (response.go)
    * 핸들러 (handlers.go)
  * 이름 곱씹어보기
    * 연관된 모듈 찾아 위치 이동

### 라우터 사용하기
  * Gorilla Web Toolkit mux

## 추가 주제

### HTTP 파일 서버
  * ```http.FileServer``` 사용
  * 실제 경로 명 숨기기, 실행 위치에서 상대 경로 지정

  ```go
  http.Handle(
    "/css/",
    http.StripPrefix(
      "/css/",
      http.FileServer(http.Dir("path/to/cssfiles")),
    ),
  )
  ```

### 몽고디비와 연동하기
  * mgo

### 에러 처리
  * 에러를 반환하고 직접 검사하는 방법 사용

  * 에러에 추가 정보 실어서 보내기
    * 에러는 interface 자료형임
      - error 는 Error() 메서드만 있으면 만족

    ```go
    type error interface{
      Error() string
    }
    ```

    * 에러 생성
      - ```errors.New(...)``` or ```fmt.Errorf(...)```
    * 그러나 에러는 interface이므로 새로운 자료형을 정의하고, Error() 메서드를 구현해 주면 됨.
      - 복잡한 경우에는 구조체로 만들어 사용할 수 있다. (필드를 더 포함할 수 있다.)
