package main

// 메서드(Method)
/*
타 언어 OOP는 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go는 Struct가 필드만을 가지며
메서드는 따로 분리되어 정의

Go 메서드는 특별한 형태의 func
함수 정의에서 func 키워드와 함수명 사이에 어떤 struct를 위한 메서드인지를 표시해야함
*/

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
}

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
