# 문자열 및 자료구조

### 스택
  * LIFO(Last In First Out)
  * 슬라이스를 사용
  * 함수 리터럴 : 익명 함수
  * switch 의 case 문은 다음 case로 넘어가지 않음 (break 필요 없음)

## 맵
  * 해시테이블로 구현됨
  * 키와 값으로 되어 있음.
  * 상수시간
  * 순서가 없음
  * 초기화 필요 (생성)

  ```go
  var m map[keyType]valueType
  m := make(map[keyType]valueType)
  m := map[keyType]valueType{}
  ```

  * 해당 키가 없으면 값의 자료형의 기본값을 반환
  * 존재 여부를 두번째 변수로 받음 (bool 형)

  ```go
  value, ok := m[key]
  m[key] = value
  ```
### 맵 사용하기
  * 슬라이스와 다르게 기본적으로 맵 변수 자체에 다시 할당하지 않음
    - 포인터 사용없이 변경
  * 맵 결과 비교 (순서 없음)
    - reflect.DeepEqual 사용

### 집합
  * 상수시간에 키의 존재를 확인할 수 있는 집합은 지원 안함.
    - 맵을 이욯하면서 값을 bool형으로 구성하여 사용
  * Tip: 불필요한 bool값 때문에 메모리를 차지하는 문제
    - 빈 구조체를 값으로 사용한다.
  * map 삭제 ```delete(m, key)```

### 맵의 한계
  * 같은 키가 여러번 들어갈 수 없음
  * 스레드 안전하지 않음
    - 락 사용 필요
  * 키에 변경될 수 있는 값은 들어가면 안됨

## 입출력
  * io package

### io.Reader와 io.Writer
  * 표준 입출력, 파일, 네트워크에 바이트를 읽고 쓸수 있는 인터페이스

### 파일 읽기, 쓰기
  1. File Open or Create : return File Handle  
  2. Read or Write : with File Handle & Data
  3. Close File Handle : use defer

### 텍스트 리스트 읽고 쓰기
  * io.Reader & io.Writer use like File Handle
  * Read from io.Reader
  * Write to io.Writer

### 그래프의 인접 리스트 읽고 쓰기
  * Data를 읽어서 이차원 슬라이스에 넣기
  * bytes.NewBuffer(nil) 과 strings.NewReader("...") 를 이용하면 테스트 용이
