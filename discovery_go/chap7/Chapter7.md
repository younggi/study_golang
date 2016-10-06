# 동시성

## 고루틴
  * 가벼운 스레드 처럼 현재 수행 흐름과 메모리를 공유하는 논리적으로 별개의 흐름을 만듬
  * ```go``` 로 호출 (유닉스 쉘에서 &를 명령의 마지막에 붙이는 것과 유사)

  ```go
  go f(x, y, z)
  ```

### 병렬성(Parallelism)과 병행성(Concurrency)
  * 병렬성: 물리적으로 별개의 흐름(동시에 각각의 흐름이 수행되는 경우)
  * 병행성: 물리적으로 두 흐름이 있지는 않지만, 동시에 두 가지를 하는 것, 순서 중요하지 않음
    * 동시성이 있는 두 루틴은 서로 의존 관계가 없음

  ```go
  func main() {
    go func() {
      fmt.Println("In goroutine")
    }()
    fmt.Println("In main routine")
  }
  ```

### 고루틴 기다리기
  * 수행 순서가 필요한 것들을 제어하기 위해 고루틴 기다리기가 필요
  * ```sync.WaitGroup```: 0이 되면 끝나는 카운터
    * 시작하기 전에 main routine에서 wg.Add()에 끝나기를 기다리는 go routine 개수를 넣는다.
    * 각각 go routine이 끝날때 wg.Done()을 실행하게 하고
    * 원래 main routine에서는 wg.Wait()으로 기다림 (카운터가 0이 될때 까지)

  * 공유 메모리와 병렬 최소값 찾기
    * 고루틴 들간 메모리 공유
    * 고루틴이 변수 포인터를 받아서 해당 변수에 원하는 값을 넣어 줄수도 있음
    * 주로 서로 소통할 필요가 없는 종류의 문제: ex) 최소값 찾기

## 채널
  * 고루틴 끼리 서로 통신하기 위한 방법(공유 메모리를 사용하는 것 보다 좀 더 나은 방법)
  * 채널(channel): 넣은 데이터를 뽑아낼 수 있는 파이프 라인 같은 형태의 자료 구조
    * 기본 자료형으로 제공하므로 채널도 일급 시민(first class citizen)임
    * 단방향 채널과 양방향 채널이 있음
    * 양방향 채널은 단방향 채널로 변환 가능
    * 맵처럼 make로 생성해야 쓸 수 있음
      * 채널 복사 시 동일한 채널을 가리킴: 레퍼런스 형(포인터와 비슷))  

    ```go
    chan 자료형
    ```

    ```go
    c1 := make(chan int)      // channel 생성
    var chan int c2 = c1      // 동일한 채널
    var <-chan int c3 = c1    // 자료를 뺄 수 만(받을 수만) 있는 채널: receive only
    var chan<- int c4 = c1    // 자료를 넣을 수 만(보낼 수만) 있는 채널: send only
    ```

    ```go
    c <- 100      // sned
    data := <- c  // receive
    ```

### 일대일 단방향 채널 소통
  * 채널에 보낸 데이터와 받은 데이터 숫자가 맞지 않으면 고루틴이 채널에 의해 멈출 수 있음
  * 채널을 보내지지 않거나, 받아지자 않을때, 다른 고루틴으로 문맥 전환(context switch)함
  * 채널을 close로 닫아 주면 더이상 채널에 데이터를 보내지 않겠다는 의미
  * 주로 쓰는 패턴
    * 함수가 채널을 만들어 반환하는 패턴

### 생성기 패턴
  * 생성기를 채널을 이용하여 만듬: 채널로 생성된 데이터를 보냄
  * 생성된 데이터를 스트림되는 것처럼 사용할 수 있음.
  * 채널 사용 장점
    * 생성하는 쪽에서는 상태 저장 방법을 복잡하게 고민할 필요가 없다.
    * 받는 쪽에서는 for의 range를 이용할 수 있다.
    * 채널 버퍼를 이용하면 멀티 코어를 활용하거나 입출력 성능상의 장점을 이용할 수 있다.

### 버퍼 있는 채널
  * 버퍼가 없을 경우 채널에 값을 보낼때, 받는 쪽도 준비가 되어 있어야 함
    * 버퍼가 없을 경우 동기적으로 작동함
    * 버퍼가 있을 경우 버퍼가 가득 차기 전까지 비동기로 동작함 (성능상 잇점)
    * 보내는 쪽과 받는 쪽의 속도가 동일하지 않는 경우가 많으므로 버퍼 사용
    * 꼭 고루틴이 아니어도 비동기도 동작 가능
    * 버그를 막기 위해서는 먼저 버퍼 없이 만든 후 성능을 위하여 추후 버퍼를 설정 추천

    ```go
    c := make(chan int, 10)
    ```

### 닫힌 채널
  * 채널이 닫히면
    * for range를 이용할때 반복이 종료
    * 채널이 열린 상태면 ok가 true임, 닫히면 기본값과 ok에는 false가 넘어옴
    * 채널이 닫힌 상태이므로 기다리지 않음
    * 닫은 채널을 다시 닫으면 panic 발생

    ```go
    val, ok := <- c
    ```

## 동시성 패턴

### 파이프라인 패턴
  * 파이프라인은 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조
  * 파이프라인 패턴은 생성기 패턴의 일종
  * 받기 전용 채널을 넘겨 받아서 입력으로 활용: 사슬처럼 연결된 파이프라인 구성 가능
  * 생성기 패턴과 동일하게 데이터를 보내는 쪽에서 채널을 닫아야 함
  * 바로 사용하지 않고 다른 곳으로 넘길때에는 Chain 고계 함수를 이용한다.

  ```go
  type IntPipe func(<-chan int) <-chan int

  func Chain(ps ...IntPipe) IntPipe {
  	return func(in <-chan int) <-chan int {
  		c := in
  		for _, p := range ps {
  			c = p(c)
  		}
  		return c
  	}
  }
  ```

### 채널 공유로 팬아웃하기
  * 팬아웃: 하나의 출력이 여러 입력으로 들어가는 경우(하나의 입력 값은 그 중 하나만 사용함)
  * 채널 하나를 여럿에게 공유하면 됨
  * for 안에 있는 제어 변수를 사용할 때는 고루틴에 파라미터로 넘겨서 사용 주의
    * 매개변수로 넘길 경우 넘길때의 값이 복사 및 고정됨.

  ```go
  for i := 0; i < 3; i++ {
    go func(i int) {
      ...
    }(i)
  ```

### 팬인하기
  * 팬인: 하나의 입력에 여러개의 입력이 들어가는 경우
  * 채널 닫을때 주의
    * 보내는 고루틴에서 채널을 닫을 경우 여러번 닫히므로 패닉 발생 가능
    * 채널을 닫기 위한 하나의 고루틴을 만들고 이 고루틴을 보내는 고루틴들이 모두 종료된 뒤에 채널을 닫고 종료하게 구성(sync.WaitGroup 사용)

### 분산처리
  * 팬아웃해서 파이프라인을 통과시키고 다시 팬인시키면 분산처리가 됨

  ```go
  func Distribute(p IntPipe, n int) IntPipe {
    return func(in <-chan int) <-chan int {
      cs := make([]<-chan int, n)
      for i:=0; i<n; i++ {
        cs[i] = p(in)
      }
      return FanIn(cs...)
    }
  }

  // Cut 하고 Draw, Paint, Decorate를 분산처리 하고, 다시 합쳐서 Box 처리함
  out := Chain(Cut, Distribute(Chain(Draw, Paint, Decorate), 10), Box)(in)
  ```

### select
  * select는 동시에 여러 채널과 통신할 수 있음.
  * 사용 형태는 switch문과 비슷하지만 동시성 프로그래밍에 사용됨.
    * 모든 case가 실행됨
    * 하나라도 실행가능한 case가 있으면 막히지 않고 실행됨
    * default가 없을 경우 모든 case 입출력이 불가할 경우 대기한다.

    ```go
    select {
    case n:= <-c1:
      fmt.Println(n, "is from c1")
    case n:= <-c2:
      fmt.Println(n, "is from c2")
    case c3 <- f():
      fmt.Println("1 is sent to c3")
    default:
      fmt.Println("No channel is ready")
    }
    ```

#### 팬인하기
  * select를 이용하면 고루틴을 여러 개 이용하지 않고 팬인을 할 수 있다.
  * c1, c2, c3 중 어느 채널이라도 준비되어 있으면 동작함
  * c1, c2, c3 중 닫힌 채널이 있을 경우 계속 해서 기본값을 받아 갈 수 있으므로 추가 처리 필요함.
    * channel의 닫힘 여부도 함께 받아서 처리 ```n, ok := <-in``` 하여 모두 nil로 닫음
    * 닫힌 채널을 nil로 바꾸어줌: nil 채널에는 보내기 및 받기가 모두 막힘.

  ```go
  select {
  case n:= <-c1: c <- n
  case n:= <-c2: c <- n
  case n:= <-c3: c <- n
  }
  ```

#### 채널을 기다리지 않고 받기
  * 채널값이 있으면 받고, 없으면 그냥 스킵하는 경우 select 사용

  ```go
  select {
  case n := <-c:
    fmt.Println(n)
  default:
    fmt.Println("Data is not ready. Skipping...")
  }
  ```

#### 시간 제한
  * 일정 시간이 지나면, select가 종료됨
    * ```return``` 이 select만 종료 시킴

  ```go
  select {
  case n := <-recv:
    fmt.Println(n)
  case send <- 1:
    fmt.Println("sent 1")
  case <-time.After(5 * timeSecond)
    fmt.Println("No send and receive communication for 5 seconds")
    return
  }
  ```

  * 실행 시 5초 동안만 실행하도록 반복하는 방법

  ```go
  timeout := time.After(5 * time.Second)
  for {
    select {
    case n:= <-recv:
      fmt.Println(n)
    case send <- 1:
      fmt.Println("sent 1")
    case <-timeout
      fmt.Println("communication wasn't finished in 5 sec")
      return
    }
  }
  ```

### 파이프라인 중단하기
  * 채널이 닫힐때 까지 자료를 모두 빼 주어야 고루틴이 종료될 수 있음
    * 남아 있는 고루틴은 메모리 누수 원인이 될 수 있음
    * 채널을 받는 쪽에서 강제로 닫을 수 없음
  * 해결책
    * 보내는 쪽에 done 채널을 하나 더 주어, done이 보내는 쪽에서 채널을 close 함

### 컨텍스트(context.Context) 활용하기
  * 위와 같이 done을 사용해도 좋으나, 복잡한 상황을 다루기 위해 context 패턴 사용
    * 고루틴 종료 신호 외에 다른 공유되어야 하는 정보가 있는 경우
    * ex) 인증 정보나 요청 정보 마감 등  
  * 설치필요: ```go get golang.org/x/net/context```
  * 관례상 함수의 맨 첫번째 인자로 넘겨주고 받음
  * 계층 구조
    * context.Background() 아래 트리 구조로 context를 붙일 수 있고, 상위 구조가 취소되면 그 하위에 있는 모든 컨텍스트도 취소된다.
    * 종류: ```WithCancel, WithDeadline, WithTimeout, WithValue```

    ```go
    // context.Background() 아래 context.WithCancel을 붙임
    ctx, cancel := context.WithCancel(context.Background())
    ```

### 요청과 응답 짝짓기
  * 요청한 채널 값과 다른 채널로 응답 받은 채널 값이 같은 값끼리 짝짓기
  * 채널에 직접 정의한 자료형을 넘기고, 그 안에 고유한 요청 ID를 포함시키면 요청에 대한 응답 channel인지 확인가능
  * 다른 고루틴이 받아가지 않고 요청한 고루틴이 응답을 받기 위한 테크닉
    * 요청을 보낼때 결과를 받고 싶은 채널을 함께 실어서 보내는 방법
  * 채널은 보내는 쪽에서 보내고 나면 닫는다.

  ```go
  type Request struct {
  	Num  int
  	Resp chan Response
  }

  type Response struct {
  	Num      int
  	WorkerID int
  }
  ```

### 동적으로 고루틴 이어 붙이기
  * 동적으로 채널을 통하여 고루틴을 이어붙이는 방법
  * prime의 배수 걸러내기 예제
    * 이미 출력된 숫자의 배수가 되는 숫자들을 걸러내는 필터 고루틴
    * 생성기 패턴으로 필터를 만듬
    * 조건이 될때만 리턴하면 그때 고루틴 함수가 동적으로 생성된다.

### 주의점
  * 자료를 보내는 채널은 보내는 쪽에서 닫는다.
  * 보내는 쪽에서 반복문 등을 활용해서 보내다가 중간에 return을 할 수 있으므로 닫을 때는 defer를 이용하는 것이 좋다. 그렇지 않으면 중간에 return했을 때 채널을 닫지 않고 종료할 수 있다.
  * 받는 쪽이 끝날때까지 기다리는 것이 모든 자료의 처리가 끝나는 시점까지 기다리는 방법으로 더 안정적이다.
  * 특별한 이유가 없다면 받는 쪽에서는 range를 이용하는 것이 좋다. 생산자가 채널을 닫은 경우에 반복문을 빠져 나오게 되기 때문에 편리하다.
  * 루틴이 끝났음을 알리고 다른 쪽에서 기다리는 것은 sync.WaitGroup을 이용하는 것이 더 나은 경우가 많다.
  * 끝났음을 알리는 done 채널은 자료를 보내는 쪽에서 결정할 사항이 아니다.
  * done 채널에 자료를 보내어 신호를 주는 많은 예제가 있는데, close(done)으로 채널을 닫는 것이 더 나은 방법인 경우가 많다.

## 경쟁 상태
  * 경쟁 상태: 어떤 공유된 자원에 둘 이상의 프로세스가 동시에 접근하여 잘못된 결과가 나올 수 있는 상태
  * sync, atomic library 활용 방법

### 동시성 디버그
  * 경쟁 상태 탐지 기능: ```-race``` option 사용
  * runtime library 활용

    ```go
    runtime.NumGoroutine()
    ```

  * panic 발생을 통한 고루틴 스택 추적

### atomic 과 sync.WaitGroup
  * 연산이 원자성을 띄지 않기 때문에 경쟁 상태 발생 가능
  * atomic 사용

    ```go
    atomic.AddInt64(&cnt, -1)
    atomic.LoadInt64(&cnt)
    ```

  * 채널을 이용한 다양한 동시성 문제 해결 가능하나, 때로는 WaitGroup을 사용하는 것이 코드 가독성 측면에서 나을수 있음.

  ```go
  func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
      wg.Add(1)
      go func() {
        defer wg.Done()
        // do something
      }()
    }
    wg.Wait()
  }
  ```

### sync.Once
  * 한 번만 어떤 코드를 수행하고자 할 때 쓸 수 있는 것

  ```go
  func main() {
    var once sync.Once
    var wg sync.WaitGroup
    for i := 0; i< 3; i++ {
      wg.Add(1)
      go func(i int) {
        defer wg.Done()
        once.Do(func() {
          fmt.Println("Initialized")
        })
        fmt.Println("Goroutine:", i)
      }(i)
    }
    wg.Wait()
  }
  ```

### Mutex와 RWMutex
  * Mutex: 상호 배타 잠금 기능(동시에 둘 이상의 고루틴에서 코드의 흐름 제어 가능)
    * 외부 자원에 접근하는 경우 효과적인 경우가 있다.
    * sync.Mutex

    ```go
    type Accessor struct {
      R *Resource
      L *sync.Mutex
    }

    Accessor{&Resource, &sync.Mutex{}}

    func (acc *Accessor) Use() {
      // do something
      acc.L.Lock()
      // Use acc.R
      acc.L.Unlock()
      // Do something else
    }
    ```

  * sync.RWMutex
    * 쓰기가 많을 경우 상대적으로 Mutex에 비하여 성능 저하
    * 여러 프로세스가 동시에 읽기 가능
    * 하나의 프로세스가 쓰는 동안은 모두 읽기 불가
    * Map 에서 사용하기 적합함

    ```go
    type ConcurrentMap struct {
      M map[string]string
      L *sync.RWMutex
    }

    func (m ConcurrentMap) Get(key string) string {
      m.L.RLock()
      defer m.L.RUnlock()
      return m.M[key]
    }

    func (m ConcurrentMap) Set(key, value string) {
      m.L.Lock()
      m.M[key] = value
      defer m.L.Unlock()
    }

    func main() {
      m := ConcurrentMap(map[string]string{}, &sync.RWMutex{})
    }
    ```

## 문맥 전환
  * 여러 프로세스 혹은 스레드에서 동작할 때 기존에 하던 작업을 메모리에 보관해두고 다른 작업을 시작하는 것
    * 병행으로 수행 가능하게 되지만 비용 발생
  * Go 컴파일러가 문맥 전환 코드를 생성하는 경우
    * 파일이나 네트워크 연산처럼 시간이 오래 걸리는 입출력 연산이 있을때
    * 채널에 보내거나 받을때
    * go로 고루틴이 생성될 때
    * 가비지 컬랙션 사이클이 지난뒤
  * 강제 문맥 전환

    ```go
    time.Sleep(0)
    ```

























