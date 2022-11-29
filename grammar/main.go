package main

import "time"

// 채널(channel)
/*
데이터를 주고 받는 통로로, make()를 통해 미리 생성되어야 하며 채널 연산자 "<-"를 통해 데이터를 주고 받음
흔히 goroutine들에서 데이터를 주고 받는 데 사용되며 상대방을 채널에서 기다림으로써
별도의 lock을 걸지 않고 데이터를 동기화
*/
/*func main() {
	// 정수형 채널을 생성한다.
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	// Go 채널 수신, 송신은 서로 기다리기 때문에 별도로 대기코드를 추가하지 않음
	var i int
	i = <-ch
	println(i)
}*/
// Go 채널 속성을 이용해서 아래와 같이 적용 가능
/*func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	<-done
}*/
/*
Buffered Channel 은 수신자가 데이터를 받을 준비가 되지 않았더라도 버퍼의 수 만큼 데이터를 보내고 다른일을 할 수 있다.
원래 아래코드는 데드락이 작동해야하지만, 버퍼 채널이므로 괜찮음
*/
/*func main() {
	ch := make(chan int, 1)
	ch <- 101 // 원래는 수신 루틴이 없어서 데드락이 발생해야함
	fmt.Println(<-ch)
}*/
// 채널 닫아도 송신은 불가능할 뿐 수신은 가능하다
/*func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더 이상 데이터 없음")
	}
}*/
// 채널 range 문
/*func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// 채널을 닫아야 range 문으로 수신 가능
	close(ch)
	for i := range ch {
		println(i)
	}

}*/
/*
 채널 select 문
복수 채널들을 기다리면 준비된 채널을 실행하는 기능을 제공
=> 즉 selet문은 여러 개의 case 문에서 각각 다른 채널을 기다리다가 준비된 채널에 맞춘 case를 실행
*/
func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

// Go 루틴 (goroutine)
/*
GO 런타임이 고ㅓㅏㄴ리하는 Lightweight 경량 논리적 가상 쓰레드이다.
"go" 키워드를 통해 사용, 비동기적으로 함수루틴이 실행됨, 기본 적으로 OS 스레드보다
더 가벼움, Go 채널을 통해 Go 루틴 간의 통신이 자유료움
*/
/*func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	//// 동기적
	//say("Sync")
	//// 비동기적
	//go say("Async 1")
	//go say("Async 2")
	//go say("Async 3")

	// 익명함수 goroutine
	// Waitgroup 생성, 2개의 go루틴 기다림
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	// Wait 기다림
	wait.Wait()
	// 메인 종료시 고루틴도 종료이므로 대기
	time.Sleep(time.Second * 3)

	// Go는 디폴트로 1개의 CPU를 사용하여 이를 시분할로하여 Go를 처리
	// 만약 복수의 CPU가 있는 머신이라면
	//runtime.GOMAXPROCS(2)

}
*/

// defer, panic
/*
Go 언어의 defer 키워드는 특정 문장 혹은 함수를 나중에 실행하는 것으로
defer를 호출한 함수가 리턴되기 직전에 실행함

panic은 함수를 즉시 멈추고 리턴 작업을 진행, panic 모드 실행은 상위 함수에도 적용
계속 콜스택을 타고 올라가며 적용되고 마지막에는 프로그램이 에러내고 종료
*/
/*func main() {
	openFile("Invalid.txt")
	println("Done")

	f, err := os.Open("1.txt")
	if err != nil {
		println("error")
		panic(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}

func openFile(fn string) {
	// panic에도 호출
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}*/

// Go Error
/*
Go는 내장 타입으로 Error 라는 Interface 타입을 갖는다.
Interface는 메서드 하나를 가지며 이 인터페이스를 구현하여 커스텀 에러 타입을 만듦
*/
/*type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C:\\temp\\t.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	_, err := otherFunc()

	// err.(type) => 이런 식으로 에러 타입으로 다르게 처리 가능
	switch err.(type) {
	default:
		println("ok")
	case MyError:
		log.Println("Log my error")
	case error:
		log.Fatal(err.Error())

	}
}*/

// 인터페이스
/*
구조체가 필드들의 집합이라면, Interface는 메서드들의 집합체.
Interface는 타입이 구현해야하는 메서드 원형을 정의한다. 하나는 사용자 정의 타입이 Interface를
구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현
=> 타 언어랑 개념은 동일 한듯

빈 인터페이스 개념이 존재하는 데, 모든 타입을 담을 수 있음
=> Java에서 Object가 할 수 있는 역할이랑 비슷
*/
/*type Shape interface {
	area() float64
	perimeter() float64
}

// 구조체 선언
type Rect struct{ width, height float64 }
type Circle struct{ radius float64 }

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

// Circle 타입에 대한 Shpale 인터페이스 구현
func (c Circle) area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.radius }

func main() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"

	printIt(x) //Tom

	var a interface{} = 1

	i := a
	j := a.(int)

	println(i) // 포인터 주소
	println(j) // 1
}

func printIt(v interface{}) {
	fmt.Println(v)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}*/

// 메서드(Method)
/*
타 언어 OOP는 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go는 Struct가 필드만을 가지며
메서드는 따로 분리되어 정의

Go 메서드는 특별한 형태의 func
함수 정의에서 func 키워드와 함수명 사이에 어떤 struct를 위한 메서드인지를 표시해야함

Value vs 포인터 receiver
포인터를 사용해서 변경 값을 적용 가능
*/
/*
type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
	area = rect.area2()
	println(rect.width, area)
}*/

// 구조체(Struct)
/*
Custom Type을 정의 Type문을 사용
*/
/*type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	p := person{}

	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p1)
	fmt.Println(p2)

	// new 로 필드를 Zero value 로 초기화
	// new 로 선언시 Person 객체의 포인터를 리턴하는데 이 경우에도 dot 으로 필드 액세스
	p3 := new(person)
	p3.name = "Lee"
	fmt.Println(p3)

	dic := newDict() // 생성자 호출, 이렇게 함으로써 사용자가 map을 초기화 안해도 됨
	dic.data[1] = "A"
}*/

// 패키지(Package)
/*
Go는 패키지를 통해 코드의 모듈화 코드의 재사용 기능을 제공
Go에서 실제프로그램 개발에 필요한 많은 패키지들은 표준 라이브러리로 제공
GOROOT/pkg 안에 존재함, GOROOT는 Go 설치  폴더

일반적으로 패키지는 라이브러리로서 사용되지만 "main"이라고 명뎡된 패키지는 GO Compiler 에 의해
다르게 인식하여 패키지를 공유 라이브러리가 아닌 실행 프로그램으로 만듦
그리고 main이 시작점 entry Point가 됨
그러므로 공유 라이브러리를 만들 때는 main 패키지나 main 함수를 사용해서는 안됨

패키지의 함수들으의 이름이 첫 글자가 대문자면 public 으로 사용됨, 첫 글자가 소문자면 non-public
패키지에서 init()는 패키지 로드시 자동 호출

패키지 호출시 _으로 받아서 init()만 사용하게 할 수 있고 다른 방법으로는
alias 지정도 가능하다
*/
/*func main() {
	// cd /testlib; go install
	println(testlib.GetMusic("Adele"))
}
*/

// 컬렉션 Map
/*
Map => Hash Table 자료 구조
Map[Key Type]Value Type으로 선언
*/
/*func main() {
	// 아래처럼 선언
	var _ map[int]string
	_ = make(map[int]string)

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}
	println(tickers)
	fmt.Println(tickers)

	var m map[int]string
	m = make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	println(str)

	noData := m[999]
	println(noData)

	delete(m, 777)
	fmt.Println(m)

	val, exists := tickers["MMM"]
	if !exists {
		println("No MMM ticker")
	} else {
		println(val)
	}

	for key, val := range tickers {
		println(key, val)
	}
}*/

// 슬라이드(slice)
/*
배열에서 수정한 것으로, 크기가 동적으로 늘어나며 발췌 용이한 등
여러가지 기능 등이 추가됨

s := make([]int, 5, 10) 로 Lenght 및 Capacity를 지정 가능
*/
/*func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 6, 10)
	println(len(s), cap(s))

	var c []int
	if c == nil {
		println("Nil Slice")
		println(len(c), cap(c))
	}

	s = []int{0, 1, 2, 3, 4, 5}
	// slice 인덱스 시작은 1부터, 1 => 0인덱스
	fmt.Println(s[2:5])
	fmt.Println(s[1:])
	fmt.Println(s[:4])
	fmt.Println(s)

	s2 := []int{9, 10}
	s = append(s, 6)
	s = append(s, 7, 8)
	s = append(s, s2...)
	fmt.Println(s)
	fmt.Println(s)
	// [11/12]주소 => [길이/용량]주소
	println(s)
}*/

// 클로저(closure)
/*
함수가 종료되도 변수가 사리지지 않고 계속 존재함 사용할 수 있도록 해주는 것
- 함수가 inline화 되어 코드 최적화가 일어남
- 사용하는 변수는 closure 형태로 바뀌면서 stack -> heap으로 이동
- 불필요한 전역 변수 제거 및 함수 캡슐화
*/
/*func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()
	println(next())
	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}*/

// 일급 함수
/*함수 자체를 입력 파라미터 혹은 리턴 파라미터로 할 수 있음
함수 => 기본 타입과 동일
type 문으로 구조체, 인터페이스 등 커스텀 타입을 정의 하는데 함수 원형도 정의할 수 있음*/
/*type calculator func(int, int) int

func main() {
	add := func(i int, j int) int {
		return i + j
	}
	r1 := calc(add, 10, 20)
	println(r1)
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f calculator, a int, b int) int {
	return f(a, b)
}*/

// 익명 함수 Anonymous Function
/*func main() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5, 6, 7)
	println(result)
}*/
