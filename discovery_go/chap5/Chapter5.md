# 구조체 및 인터페이스
  * 구조체는 필드들을 묶어 놓은 것
    * 더 복잡한 자료형을 정의 가능

  * 인터페이스는 메서드의 집합
    * 구현이 없는 메서드
    * 외부 의존성을 줄이는데 사용
    * 빈 인터페이스를 사용하여 어떤 자료형도 받을 수 있는 와일드 카드로 사용

## 구조체
  * 필드들의 모음 or 묶음
  * keyword: ```struct```

### 구조체 사용법
  * 정의 및 선언

  ```go
  type Task struct {
    title string
    done bool
    due *time.Time
  }

  var myTask Task

  var myTask = Task{"laundry", false, nil}
  var myTask = Task{
    title: "laundry",
    done: true,
  }
  ```

### const와 iota
  * 확장성 Tip
    * bool형 쓸 곳에 enum형을 사용한다. : 더 많은 자료형 추가
    * go에서는 enum 대신 const로 정의하여 사용한다.
    * 0 과 같이 필요없는 값은 dummy로 두는 것도 용이함
    * 값을 자동으로 붙이려고 할때 itoa 를 사용


  ```go
  type Task struct {
    title   string
    status  status
    due     *time.Time
  }

  type status int

  const (
    UNKNOWN status = 0
    TODO    status = 1
    DONE    status = 2
  )

  const (
    UNKNOWN status = iota   // 0
    TODO                    // 1
    DONE                    // 2
  )
  ```

### 테이블 기반 테스트
  * 테스트 시 assert가 지원되지 않으므로 if 나 테이블 기반 테스트를 이용한다.
    * struct 를 이용하여 input, output 슬라이스 데이터를 만들어 체크한다.

  ```go
  func TestFib(t *testing.T) {
    cases := []struct {
      in, want int
    }{
      {0, 0},
      {5, 5},
      {6, 8},
    }
    for _, c := range cases {
      got := seq.Fib(c.in)
      if got != c.want {
        t.Error("Fib(%d) == %d, want %d", c.in, got, c.want)
      }
    }
  }
  ```

### 구조체 내장
  * 구조체 포함 관계
    * 상속과 달리 실제로 내부에 필드를 내장하고 있으면서 바로 메소드를 호출할 수 있는 편의 제공
  * 구조체 내장
    * 구조체 내에 구조체가 포함되어 있을 경우 외부 구조체에서 내부 구조체를 모두 같은 메소드 이름으로 호출하는 불편함을 없애기 위한 기능
    * 내부 구조체의 필드 이름을 없앰
    * 자동으로 자료형과 같은 이름의 필드가 생성되고, 그 필드를 호출하는 메서드가 자동 생성됨
    * 직렬화 시 주의
      - 내장된 필드가 구조체 전체의 직렬화 결과를 바꿔 버리는 문제
    * 내부 객체에서 노출한 (대문자로 시작하는) 모든 것에 바로 접근할 수 있는 편의가 제공된다.

  ```go
  type Task struct {
    Title string
    Status status
    *Deadline
  }
  ```

## 직렬화와 역직렬화
  * 직렬화: 객체의 상태를 보관이나 전송 가능한 상태로 변환하는 것을 말합니다.
  * 역직렬화: 객체로 복원

### JSON
  * JSON(JavaScript Object Notation)

  * JSON 직렬화 및 역직렬화
    * 직렬화
      - ```json.Marshal(t)```
      - struct에서 대분자로 시작하는 필드만 직렬화 수행
    * 역직렬화
      - ```json.Unmarshal(b, &t)```
      - json string 설정시 \` 사용
  * JSON 태그
    * 구조체에 JSON 변환시 원하는 방식을 지정함
    * 64bit 정수를 직렬화할때 javascript의 오류를 제거하기 위해 string으로 변환
      - javascript는 실수형이므로 53bit가 넘어서면 정확도가 떨어짐

    ```go
    type MyStruct struct {
      Title     string `json:"title"`         // title을 JSON 필드로 사용
      Internal  string `json:"-"`             // JSON 변환시 무시됨
      Value     int64  `json:",omitempty"`    // 0 인 경우 JSON 변환시 무시됨
      ID        int64  `json:",string"`       // JSON에서는 string 값으로 출력
    }
    ```

  * JSON 직렬화 사용자 정의
    * 커스텀 코드 작성: struct에 json.Marshaler interface를 implementation

    ```go
    // MarshalJSON implements the json.Marshaler interface
    func (s status) MarshalJSON() ([]byte, error) {
      switch s {
      case UNKNOWN:
        return []byte(`"UNKNOWN"`), nil
      case TODO:
        return []byte(`"TODO"`), nil
      case DONE:
        return []byte(`"DONE"`), nil
      default:
        return nil, errors.New("status.MarshalJSON: unknown value")
      }
    }

    // UnmarshalJSON implements the json.Unmarshaler interface
    func (s *status) UnmarshalJSON(data []byte) error {
      switch string(data) {
      case `"UNKNOWN"`:
        *s = UNKNOWN
      case `"TODO"`:
        *s = TODO
      case `"DONE"`:
        *s = DONE
      default:
        return errors.New("status.MarshalJSON: unknown value")
      }
    }
    ```

  * 구조체가 아닌 자료형 처리
    * 맵, Key 순으로 정렬되어 JSON 변환됨
      * Key는 문자열형이어야 한다.
      * 임의의 자료형을 사용할 때 ```interface{}``` 자료형 사용

  * JSON 필드 조작하기
    * JSON 구조에 따라 struct의 구조가 제한되는 문제
    * 구조체 내장을 이용하면 원래 구조체를 고치지 않고, 원하는 필드들만 제외하거나 추가하여 직렬화 가능
    * json.Marshal 시 구조체 내장을 별도로 정의하여 사용
      * 원래 구조체의 필드와 동일 이름으로 정의할 경우 shadowing 될 수 있음
      * 두 구조체를 합쳐서 JSON 할 경우 구조체 내장으로 사용

    ```go
    type Fields struct {
      VisibleField    string
      InvisibleField  string
    }

    func ExampleOmitFields() {
      f := &Fields{"a","b"}
      b, _ := json.Marshal(struct {
        *Fields
        InvisibleField  string `json:",omitempty"`
        Additional      string
      }{Fields: f, Addtional: "c"})
      fmt.Println(string(b))
    }
    ```

  * Gob
    * Go 언어에서만 읽고 쓸수 있는 형태이고, 모두 Go로 되어 있을 경우 고려
    * gob.NewEncoder, gob.NewDecoder 생성하고 여기에 io.Writer, io.Reader를 넘긴다.

## 인터페이스
  * 인터페이스: 메서드 들의 묶음
    * 무언가를 할 수 있는 것을 의미
    * 인터페이스 명명: 인터페이스 메서드에 ```er```을 붙임
  * 구조체: 자료들의 묶음

### 인터페이스의 정의
  * 구조체와 유사

  ```go
  interface {
    Method1()
    Method2(i int) error
  }
  ```

  * 인터페이스 이름

  ```go
  type Loader interface {
    Load(filename string) error
  }
  ```

  * 인터페이스 합치기: 구조체 내장과 비슷한 방식

  ```go
  type ReadWriter {
    io.Reader
    io.Writer
  }
  ```

### 커스텀 프린터
  * 이름 붙인 자료형의 출력 양식 지정: String() 함수 정의
  * 프린트만을 위한 type을 별도로 지정하여 기존 type을 다른 방식으로 출력 지정 가능

  ```go
  type Task struct {
    ...
  }

  type IncludeSubTasks Task

  func (t IncludeSubTasks) String() string {
    ...
  }
  ```

### 정렬과 힙
  * Go의 sort.Sort에서 이용하는 방법은 비교 정렬이자 불안정 정렬(unstable sort)임.

  * 정렬 인터페이스의 구현
    * 제네릭은 지원하지 않지만, 인터페이스를 지원하여 다양한 형태 정렬 수행 가능
    * ```sort.Interface```를 따르면 정렬 가능
    * 새로운 이름을 타입을 정의하기만 하면 같은 자료형에 각각 다른 방식의 정렬을 구현할 수 있음

  * 정렬 알고리즘
    * 기본적으로 빠른 정렬(Quicksort)이고, 최악의 경우를 피하기 위하여 중위법(피벗 3개를 골라서 가운데 값을 고름)  
    * 너무 깊이 빠른 정렬에 빠지면 힙 정렬을 이용
    * 자료가 7개 이하일 경우 삽입 정렬

  * 힙
    * 자료 중에 가장 작은 값을 O(logN)의 시간 복잡도로 꺼낼 수 있는 자료구조
    * sort.Interface를 내장하고 있으므로 5개의 메서드 구현(3+2)
    * 정렬이 끝나지 않은 상황에서도 지금까지 정렬된 자료를 이용할 수 있음
      - 선택 정렬의 일종
    * 우선 순위 큐: 가장 낮은 값부터 소비  

### 외부 의존성 줄이기
  * 개발 시 인터페이스를 사용하도록 개발하면
  * 테스트 시 인터페이스를 활용하여 실제 대상 대신 테스트용 개체 사용가능

### 빈 인터페이스와 형 단언
  * 나열된 메서드들을 정의하고 있는 자료형은 인터페이스로 취급 될 수 있다.
  * 빈 인터페이스는 아무 자료형이나 취급할 수 있음
  * 형변환 (from interface{})
    * 자료형 변환과 다름 ```형(값)``` 형식
    * 형 단언(type assertion): 형변환이 실행 시간에 자료형이 맞는지 검사를 수행함
      - 실행시간에 내가 단언한 형이 아니면 패닉이 발생됨
      - ```popped.(string)``` : interface{} 형인 popped를 string으로 단언함
    * Panic 방지를 위해 변환 여부를 추가로 리턴 받음

    ```go
    var r io.Reader = NewReader()
    f, ok := r.(os.File)
    ```

### 인터페이스 변환 스위치
  * 인터페이스가 특정 자료형일 경우에 switch를 통하여 각 자료형 별 행위를 구현 할 수 있다.

  ```go
  switch x := a[i].(type) {
  case string:
    t[i] = x
  case int:
    t[i] = strconv.Itoa(x)
  case fmt.Stringer:
    t[i] = x.String()
  }
  ```
