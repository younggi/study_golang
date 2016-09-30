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
