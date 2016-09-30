
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

  * 반복된 에러 처리 피하기
    * 에러처리를 단순화하는 테크닉
    * 에러 발생 시 프로그램 종료

    ```go
    if err := f(); err != nil {
      panic(err)
    }
    ```

    * 함수로 만들어 중복 제거

    ```go
    func Must(err error) {
      if err != nil {
        panic(err)
      }
    }
    ```

  * 추가 정보와 함께 반복된 처리 피하기
    * error를 struct로 정의하고 추가 정보를 넣은 후 함수로 만들어 반복을 제거

  * panic과 recover
    * panic 발생 시 호출 스택을 타고 역순으로 올라가서 프로그램을 종료한다.
    * defer 도 호출 스택을 역순으로 타고 갈때 처리되는데, 이때 defer 안에 recover를 사용하면 panic이 전파되지 않는다.
      - recover는 defer 안에서만 효력이 발생한다.
      - 반환값을 변경하여 지정할 수 있다.

    ```go
    func f() (i int) {
      defer func() {
        if r:= recover(); r != nil {
          fmt.Println("Recovered in f", r)
          i = -1  // panic 시 recover할때 반환값 
        }
      }()
      g() // This function panics.
      return 100
    }

    func g() {
      panic("I panic!")
    }
    ```
