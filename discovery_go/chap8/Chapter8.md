# 실무 패턴
  * 어떠한 문제를 풀 것인가가 중요하다.
  * 어떠한 문제가 풀고 싶은가가 중요하다.

## 오버로딩
  * 오버로딩: 같은 이름의 함수 및 메서드를 인자의 자료형이나 개수에 따라 여러 개 둘 수 있는 기능
    * Go 언어에서는 지원 안됨
  * 어떤 문제를 풀기 위해 오버로딩이 필요한가?
    * 자료형에 따른 다른 이름 붙이기: 오버로딩이 반드시 필요하진 않음
    * 동일한 자료형의 자료 개수에 따른 오버로딩: 가변인자를 사용하여 해결
    * 자료형 스위치 활용하기: 오버로딩을 반드시 해야하는 경우는 인터페이스로 인자를 받고, 메서드 내에서 자료형 스위치로 다른 자료형에 맞추어 다른 코드가 수행되게 할 수 있다.
    * 다양한 인자 넘기기: 여러 설정을 넘길 경우 모두를 묶은 구조체를 넘기는 것을 고려하자
    * 인터페이스 활용하기: 인터페이스를 활용하는 경우가 더 나은 경우도 있다.

### 연산자 오버로딩
  * 연산자 오버로딩: 새로 정의한 자료형에 대하여 기본적인 연산자를 구현
    * Go 언어에서는 지원 안됨
  * 편의성을 위한 기능이므로 인터페이스를 이용하여 해결한다.

## 템플릿 및 제네릭 프로그래밍
  * 제네릭: 알고리즘을 표현하면서 자료형을 배제할 수 있는 프로그래밍 패러다임.
    * Go 언어에서는 지원 안됨

### 유닛 테스트
  * 여러 자료형에 따른 비교를 위해 ```reflect.DeepEqual``` 사용

  ```go
  func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
    if !reflect.DeepEqual(expected, actual) {
      t.Errorf("%v != %v", expected, actual)
    }
  }
  ```

  * 테이블 기반 테스트: input data를 테이블 기반으로 만들고, 반복문을 돌면서 정합성 체크

  ```go
  func Test(t *testing.T) {
    examples := []struct {
      desc, expected, input string
    }{{
      desc: "empty case",
      expected: "",
      input: "",
    }, {
      ...
    }}
    for _, ex := range examples {
      actual := f(ex.input)
      if ex.expected != actual {
        t.Errorf("%s: %s != %s", ex.desc, ex.expected, ex.actual)
      }
    }
  }
  ```

### 컨테이너 알고리즘
  * 정렬과 힙 알고리즘에서 자료형과 독립적으로 구현하기 위해, 두 자료의 대소 비교를 하는 부분을 인터페이스를 이용하여 두 인덱스를 주고 자료를 비교함.

### 자료형 메타 데이터
  * 어떤 자료형이 넘어 왔는지에 따라 다른 코드가 동작하게 하려면, 자료형 스위치 사용
  * ```reflect``` 패키지 이용: 구조체를 받아서 구조체에 필드 이름과 자료형 알아내기
  * 자료형을 받아 키와 값의 자료형으로 하는 맵 만들기
    * 메타데이터: 자료에 대한 자료
    * NewMap 함수 사용시 리턴 값이 interface{} 이므로 형단언 사용

  ```go
  func NewMap(key, value interface{}) interface{} {
    keyType := reflect.TypeOf(key)
    valueType := reflect.TypeOf(value)
    mapType := reflect.MapOf(keyType, valueType)
    mapValue := reflect.MakeMap(mapType)
    return mapValue.Interface()
  }

  m := NewMap("", 0).(map[string]int)
  ```

### go generate
  * 임의의 명령을 수행하여 프로그램 코드를 생성함
  * 코드에 ``` //go:generate ``` 뒤에 명령을 넣어 코드에 넣은 뒤, go generate를 수행

## 객체지향
  * Go는 객체지향을 완전히 지원하지 않음

### 다형성
  * 객체지향의 핵심
  * 객체에 메서드가 호출되었을 때, 그 객체가 메서드에 대한 다양한 구현을 가질 수 있다.
    * 객체가 호출 되었을 때 어떤 자료형이냐에 따라 다른 구현을 할 수 있게 함
  * Go의 인터페이스로 구현
    * 같은 유형의 메서드만 구현하고 있으면 다형성에 사용될 수 있음

### 인터페이스
  * 인터페이스는 같은 메서드만 구현하고 있으면 그 인터페이스를 구현한 것이 됨.

### 상속
  * 객체지향에서 상속은 어떤 클래스의 구현들을 재사용하기 위하여 사용됨.
  * 관계 종류: IsA, HasA 관계 성립
    * HasA 관계: 상속보다는 객체 구성(object composition)이 더 나음
    * IsA 관계
      - 추상 클래스 상속: 인터페이스 구현
      - 구현 클래스 상속: 인터페이스 구현 + 구조체 내장

#### 메서드 추가
  * 기존에 있던 코드를 재사용하면서 기능 추가하고 싶은 경우에 상속 사용 가능
  * 구조체 내장: 구조체 안에 다른 구조체를 넣어서 필드 참조나 메서드 호출을 위해 불필요한 코드 작성을 피함

  ```go
  type Rectangle struct {
    Width, Height float32
  }

  func (r Rectangle) Area() float32 {
    return r.Width * r.Height
  }

  type RectangleCircum struct {
    Rectangle
  }

  func (r RectangleCircum) Circum() float32 {
    return 2 * (r.Width + r.Height)
  }
  ```

#### 오버라이딩
  * 기존에 있던 구현을 다른 구현으로 대체하고자 하는 경우에도 상속 사용 가능
  * 구조체 내장을 이용하여 구현
    * 구조체를 내장하고, 같은 이름의 메서드를 정의함

#### 서브 타입
  * 기존 객체가 쓰이던 곳에 상속받은 객체를 쓰고자 상속 사용
  * 인터페이스와 구조체 내장을 모두 사용함

### 캡슐화
  * 패키지 단위의 public과 private이 존재
  * public: 대문자로 시작, 다른 패키지에서 참조 가능
  * private: 소문자로 시작, 같은 패키지에서만 참조 가능
  * internal: 접근 할 수 있는 패키지를 지정하는 방법
  * naming

    ```go
    Len()
    SetLen(x)
    ```

## 디자인 패턴
  * Go로 구현한 디자인 패턴

### 반복자 패턴
  * 클로저를 이용하여 호출하는 반복자(4장)
  * 콜백을 넘겨주어서 이 함수가 모든 원소에 대해 호출되게 하는 반복자(4장)
  * 인터페이스를 이용한 반복자(5장)
  * 채널을 이용한 반복자(7장)
    * 채널을 이용한 반복자 중 중간에 중단하기 위해서는 ```done``` 채널이나 ```context.Context``` 를 받아서 처리함

### 추상 팩토리 패턴
  * 팩토리를 여럿 묶어놓은 팩토리를 추상화하는 패턴
  * 인터페이스를 활용하여 구현

### 비지터 패턴
  * 알고리즘을 객체 구조에서 분리시키기 위한 디자인 패턴
  * 인터페이스를 이용하여 구현
  * 인터페이스를 구현한 객체를 인자로 넣어주면, 원래 기존 구조가 각 인터페이스의 구현을 방문하면서 실행한다.








