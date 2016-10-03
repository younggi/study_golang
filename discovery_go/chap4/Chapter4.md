# 함수
  * 스택으로 구현
    * 스택에 현재 프로그램의 카운티 및 인자값을 입력
    * 카운터 값을 변경 => 호출
  * 값에 의한 호출만 지원 (Call by value)
    * 변경이 필요한 경우는 주소값을 전달하여 처리

## 값 넘겨 주고 넘겨 받기
  * 슬라이스 => 배열에 대한 포인터, 길이, 용량을 가진 구조체
    * 배열 내용(값)만 변경 시 에는 그냥 인자로 넘김
    * 슬라이스의 값이 변경 될 경우에는 슬라이스의 구조체를 넘겨야 함
      - 슬라이스의 배열 크기 변경 시
  * 변수 앞에 &를 붙이면 해당 변수에 담겨 있는 값의 포인터를 얻을 수 있다.
  * 포인터로 넘어온 값은 \*를 앞에 붙여서 값을 참조
  * 중요 포인트 : 무엇의 값인지?
    * 주소 값이 넘어와서 함수 내에서는 포인터 변수에 담게 된다.

### 둘 이상의 반환값
  * ( , ) 형식으로 반환, 반환값도 명명 가능

### 에러값 주고 받기
  * 에러의 값
  * 사용 패턴
  * 필요시 에러는 다시 상위 함수로 리턴하여 전달  

  ```go
  if err := MyFunc(); err != nil {

  }
  ```

  * 에러 위치를 메세지에 포함 필요
    * ```error.New or fmt.Errorf``` 사용

### 명명된 결과 인자
  * 코드 가독성 시 효과가 있을 경우만 사용 권장

### 가변인자
  * 인자명 뒤에 ```...```를 추가한다.
    * 입력 시는 나열 값
    * 함수 내에서 받았을 때는 슬라이스 값

  ```go
    func WriteTo(w io.Writer, lines... string) (n int64, err error)
  ```

  * 슬라이스를 가변인자에 넘길때에서 ```...``` 사용

  ```go
  lines := []string("hello", "world", "Go language")
  WriteTo(w, lines...)
  ```

### 값으로 위급되는 함수
  * 함수: First-class citizen

### 함수 리터럴
  * 함수 이름을 없앰

  ```go
  func(a, b int) int {
    return a + b
  }
  ```

### 고계 함수 (High-order function)
  * 함수를 넘기고 받는 함수
  * 추상화에 사용시 유용

### 클로저 (Closure)
  * 닫힘이라는 의미의 쿨로저
  * 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근 할 수 있는 코드
    * 외부에서 사용하던 함수의 변수도 함께 전달된다

### 생성기
  * 클로저 이용하여 생성기 만들기
  * Function을 리턴한다. (함수 리터럴을 리턴)
  * 고계 함수
  * NewIntGenerator() 호출시 마다 next 변수는 별도로 할당됨
  * Lazy evaluation 이나 무한한 크기의 자료 구조 만들때 응용 가능

  ```go
  func NewIntGenerator() func() int {
  	var next int
  	return func() int {
  		next++
  		return next
  	}
  }

  func ExampleNewIntGenerator() {
  	gen := NewIntGenerator()
  	fmt.Println(gen(), gen(), gen(), gen(), gen())
  	fmt.Println(gen(), gen(), gen(), gen(), gen())
  	// Output:
  	// 1 2 3 4 5
  	// 6 7 8 9 10
  }
  ```

### 명명된 자료형 (Named Type)
  * 별칭 붙이기 가능

  ```go
  type rune int32
  ```

  * 자료형을 검사함으로써 프로그램이 직접 수행해보기 전에 컴파일 시점에서 버그를 어느정도 예방할 수 있다.
    * 실제 데이터 타입이 같은 변수를 명명하여 혼용하는 실수를 미연에 방지 가능함
  * 명명된 자료형 끼리는 호환 되지 않음: casting 필요
  * 명명되지 않은 자료형과 명명된 자료형은 표현이 같으면 호환 됨  
  * 자료형 일괄 변환시에도 편리함

### 명명된 함수형
  * 함수의 자료형 정의

  ```go
  type BinOp func(int, int) int
  ```

### 인자 고정
  * 함수를 리턴하는 고계 함수를 만들고 그것을 인자로 넣으면서 인자를 고정한다.  
  * 함수를 호출하는 시점에 파라미터 값으로 실행되는 함수에서는 고정됨. 
  * m 고정 => m을 ReadFrom 함수의 인자로 만든다.
    * m 이 Insert의 인자인데, 매번 같이 호출하지 않아도 되도록 고정한다.

  ```go
  m := NewMultiSet()
  ReadFrom(r, func(line, string) {
    Insert(m, line)
  })

  m := NewMultiSet()
  ReadFrom(r, InsertFunc(m))

  m := NewMultiSet()
  ReadFrom(r, BindMap(Insert, m))
  ```

### 패턴의 추상화
  * 변수: 어떤 값이나 연산 결과에 이름을 붙여 추상화
  * 함수: 코드를 값의 입출력으로 추상화
  * 반복되는 패턴은 추상화 한다.
  * 함수를 Functional programming처럼 작은 다른 함수들로 추상화 할 수도 있다.

### 자료구조에 담은 함수
  * 함수를 자료구조에 담을 수 있음
  * struct 에 func 를 담을 수 있다.

## 메서드
  * 리시버(Receiver)가 붙는 함수는 메서드라고 부름
  * 메서드 내에서 리시버를 참조하여 사용할 수 있다.

  ```go
  func (recv T) MethodName(p1 T1, p2 T2) R1
  ```

### 단순 자료형 메서드
  * 모든 명명된 자료형에 메서드를 정의 가능

  ```go
  type VertexID int

  func (id VertexID) String() string {
    return fmt.Sprintf("VertexId(%d)",id)
  }
  ```

### 문자열 다중 집합
  * 자료형에 이름을 붙인다.
  * 그다음 메서드 정의
  * 자료 추상화를 가능하게 함

  ```go
  type MultiSet map[string]int

  func (m MultiSet) Insert(val string) {
    m[val]++
  }

  func (m MultiSet) Erase(val string) {
    if m[val] <= 1 {
      delete(m, val)
    } else {
      m[val]--
    }
  }
  ```

  func (m MultiSet) Count(val string) {
    return m[val]
  }

  func (m MultiSet) String() string {
    s := "{ "
    for val, count := range m {
      s += strings.Repeat(val + " ", count)
    }
    return s + "}"
  }

### 포인터 리시버
  * 자료형이 포인터인 리시버
  * 메서드 인자로 값의 변경이 필요해서 포인터가 필요한 경우 포인터 리시버로 사용한다.
  * 리시버의 이름은 길게 붙이지 않는다.

  ```go
  type Graph [][]int

  func WriteTo(w io.Writer, adjList [][]int) error
  func ReadFrom(r io.Reader, adjList *[][]int) error

  func (adjList Graph) WriteTo(w io.Writer) error
  func (adjList *Graph) ReadFrom(r io.Reader) error
  ```

### 공개 및 비공개
  * 공걔: 이름이 대분자로 시작
  * 비공개: 이름이 소문자로 시작
  * 적용 대상: 메서드, 자료형, 변수, 상수, 함수

## 활용
  * 라이브러리 활용법

### 타이머 활용하기
  * ```time.Sleep``` : blocking timer
  * ```time.Timer*``` : non-blocking timer
  * 비동기 상황에서 사용되는 기술: 콜백
    * 어떤 조건이 만족될때 호출해 달라고 요청하는 것
  * 라이브러리중 함수를 인자로 받는 것은 콜백 호출을 하는 것임

  ```go
  time.AfterFunc(5*time.Second, func() {
    // 메세지를 없애는 코드
    // 비동기적으로 실행하고자 하는 코드
    fmt.Println("I am so excited!")
  })

  // 강제 타이머 종료
  time.Stop()
  ```

### path/filepath 패키지
  * [Walk](https://golang.org/pkg/path/filepath/#Walk)
    * 지정된 디렉터리 경로 아래에 있는 파일에 대해 어떤 일을 할 수 있는 함수
    * 고계함수
