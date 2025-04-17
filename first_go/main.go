package main

import (
	"context"
	//_ "first_go/pkg1"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	//testInt()
	//testFloat()
	//testString()
	//testTransfer()
	//testPoint()
	//testStruct()
	//testStruct2()
	//testCopy()
	//testGender()
	//testIota()
	//testSwitch()
	//testList()
	//testList2()
	//testSlice()
	//testSlice3()
	//testSliceRemove()
	//testAppend()
	//testMap()
	//testMap2()
	//testMap3()
	//testRune2()
	//testStrConv()
	//testInterfaceTransfer()
	//testInterfaceTransfer2()
	//testInterfaceTransfer3()
	//testGoRoutine()
	//testSafeCounter()
	//testChannel()
	//testChanTimeout()
	testChan()

	//for i := 1; i <= 5; i++ {
	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//fmt.Println("i =", 100/i)
	//}
}

func testInt() {
	var a uint = 250000046346345
	fmt.Println("a =", a)
}

func testFloat() {
	var f1 float32 = 3.1415926
	f2 := 3.1415926
	fmt.Printf("f1 = %f\n", f1)
	fmt.Println("f2 =", float32(f2))
	fmt.Println(float32(f2) == f1)
	fmt.Println(float32(f2) - f1)
}

func testRune() {
	var r1 rune = 'a'
	var i1 int32 = 'a'
	fmt.Println(r1 == i1)
}

func testString() {
	var s string = "hello\nworld"
	fmt.Println(s)
}

func testTransfer() {
	var s string = "Go语言"
	var bytes []byte = []byte(s)
	var runes []rune = []rune(s)

	fmt.Println(string(runes) == string(bytes))
	fmt.Println(len(s))
	fmt.Println(len(bytes))
	fmt.Println(len(runes))

	fmt.Println(s[0:8])
	fmt.Println(string(runes[0:3]))
}

var i int
var s string
var p1 *int = &i
var p2 *string = &s

func testPoint() {
	a := 2
	var p *int = &a
	fmt.Println(a)
	fmt.Println(p)
	*p = 3
	fmt.Println(a)
	fmt.Println(p)

	var pp **int = &p
	fmt.Println(*pp)
	fmt.Println(pp)
	**pp = 4
	fmt.Println(a)
}

type Person struct {
	Name    string
	Age     int
	Call    func() byte
	Map     map[string]string
	Ch      chan string
	Arr     [32]uint8
	Slice   []interface{}
	Ptr     *int
	once    sync.Once
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Person2 struct {
	Name  string            `json:"name" gorm:"column:<name>"`
	Age   int               `json:"age" gorm:"column:<name>"`
	Call  func()            `json:"-" gorm:"column:<name>"`
	Map   map[string]string `json:"map" gorm:"column:<name>"`
	Ch    chan string       `json:"-" gorm:"column:<name>"`
	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
	Ptr   *int              `json:"-"`
	O     Address           `json:"-"`
}

type Job struct {
	int
	string
	byte
	Name string
}

var Test = struct {
	Name string `json:"name" gorm:"column:<name>"`
	Age  int    `json:"age" gorm:"column:<name>"`
}{}

var ss1 = struct {
	Name string
	Age  int
	int
	Address
}{
	Name: "MikeZhao",
	Age:  21,
	int:  1000,
	Address: Address{
		Street: "sss",
		City:   "sss",
		State:  "sss",
		Zip:    "sss",
	},
}

var ss3 = struct {
	Name string
	Age  int
	int
}{}

func testStruct() {

	ss2 := struct {
		Name string
		Age  int
	}{
		Name: "Mike",
		Age:  2,
	}
	ss3.Name = "Mike"
	ss3.Age = 23
	ss3.int = 100
	fmt.Println(ss2.Name, ss2.Age)
	fmt.Println(ss1.Name, ss1.Age, ss1.int)
	fmt.Println(ss1)
	fmt.Println(ss3)
	fmt.Println(ss1.Street)
}

type Student struct {
	Name string
	Age  int
}

func (s Student) book() Student {
	s.Name = "Mike"
	s.Age = 23
	return s
}

func (s *Student) book2() *Student {
	s.Name = "Sam"
	s.Age = 29
	return s
}

func testStruct2() {
	s1 := Student{
		Name: "Tom",
		Age:  22,
	}
	s2 := s1.book()
	s3 := s1.book2()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(*s3)
}

func testCopy() {
	a := struct {
		a string
	}{a: "123"}

	b := a
	b.a = "345"
	fmt.Println(a)
	fmt.Println(b)

	c := &a
	(*c).a = "999"
	c.a = "888" //简写
	fmt.Println(a)
	fmt.Println(*c)
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func testGender() {
	var g = Male
	isMale := g.isMale()
	fmt.Println(isMale)
}
func (g *Gender) isMale() bool {
	return *g == Male
}

func (g *Gender) isFemale() bool {
	return *g == Female
}

type OrderStatus int

const (
	StatusInit       OrderStatus = 0
	StatusWaitingPay             = 2 + iota
	StatusPaid
	StatusFinished
	StatusAfterSale
	StatusEnd
)

func testIota() {
	fmt.Println(StatusInit)
	fmt.Println(StatusWaitingPay)
	fmt.Println(StatusPaid)
	fmt.Println(StatusFinished)
	fmt.Println(StatusAfterSale)
	fmt.Println(StatusEnd)
}

func testSwitch() {
	a := 1
	testSwitch2(a)
	b := &a
	testSwitch2(b)
	c := Address{Street: "street"}
	testSwitch2(c)
}

func testSwitch2(d interface{}) {

	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case *byte:
		fmt.Println("d is byte point type, ", t)
	case *int:
		fmt.Println("d is int type, ", t)
	case *string:
		fmt.Println("d is string type, ", t)
	case *Address:
		fmt.Println("d is CustomType pointer type, ", t)
	case Address:
		fmt.Println("d is CustomType type, ", t)
	case int:
		fmt.Println("d is int type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}
}

func testFor() {
	var stopped atomic.Bool
	stopped.Store(true)
	stopped.Load()
}

func testList() {
	var list = [...]int{1, 2, 3, 5, 7, 9}
	fmt.Println(list)
	list[0] = 100
	fmt.Println(list)

	marr := [2]map[string]string{make(map[string]string), make(map[string]string)}
	//marr[0] = make(map[string]string)
	//marr[1] = make(map[string]string)
	marr[0]["a"] = "3"
	marr[1]["b"] = "4"
	fmt.Println(marr)

	print := func(sarr [6]int) {
		fmt.Println(sarr)
	}
	print(list)

	for i, v := range list {
		fmt.Println(i, v)
	}
	fmt.Println(len(list))
}

func testList2() {
	a := [5]int{1, 2, 3, 4, 5}
	//a[6] = 9
	//a = append(a, 1, 2)
	fmt.Println(a)

	testList3(&a)
	fmt.Println(a)
}

func testList3(param *[5]int) {
	param[0] = 100
	fmt.Println(*param)
}

func testSlice() {
	var s0 []int
	s0 = append(s0, 0)
	fmt.Println(s0)
	s0[0] = 100
	fmt.Println(s0)

	s1 := []int{}
	s1 = append(s1, 0)
	fmt.Println(s1)

	var s2 = make([]int, 4, 10)
	fmt.Println(s2)

	list := []int{1, 2, 3, 4, 5}
	s3 := list[:4]
	fmt.Println(s3)

	fmt.Println("slice")
	arr := []int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]    // [2,3,4]
	slice2 := slice1[0:2] // [2,3]
	slice2[0] = 100
	arr[2] = 999
	fmt.Println(arr)    // [1,100,3,4,5]
	fmt.Println(slice1) // [100,3,4]
	fmt.Println(slice2) // [100,3]

	s10 := []int{1, 2, 3, 4, 5}
	testSlice2(s10)
	fmt.Println(s10)
}

func testSlice2(s []int) {
	s[1] = 100
	fmt.Println(s)
}

func testSlice3() {
	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:3], append([]int{3}, s4[2:]...)...)
	st := s4[2:]
	s4 = append(s4[:2], s4[2:]...)
	fmt.Println("s4 = ", s4)
	fmt.Println("st = ", st)
}

func testSliceRemove() {
	s4 := []int{1, 2, 3, 4, 5}
	s6 := s4[1:3]
	fmt.Println(s6)
	s6[0] = 1000
	fmt.Println(s6)
	fmt.Println(s4)
	s5 := append(s4, 2)
	s5[0] = 999
	fmt.Println(s4)
	fmt.Println(s5)
}

func testAppend() {
	fmt.Println("++++++++++++++")
	s := []int{1, 2}
	fmt.Println(s)
	s = append(s, 1000)
	s[0] = 100
	fmt.Println(s)
	//fmt.Println(s1)

	testAppend1(s)
	fmt.Println(s)
	//fmt.Println(s1)
	fmt.Println("++++++++++++++")
}

func testAppend1(s []int) {
	s = append(s, 6)
	//s[2] = 100
	s = append(s, 6)
	s[2] = 101
	s = append(s, 6)
	s[2] = 102
	s = append(s, 6)
	s[2] = 104
	//s = append(s, 6)
	//s = append(s, 6)

	s[2] = 100
	fmt.Println("----")
	//fmt.Println(a)
	fmt.Println(s)
	fmt.Println("-----")

}

func testMap() {
	map1 := map[string]interface{}{
		"a": "b",
		"b": 123,
	}
	map1["c"] = "d"
	map2 := map1
	map2["d"] = "e"
	map2["e"] = "f"
	map2["f"] = "f"
	map2["g"] = "f"
	map2["h"] = "f"
	map2["i"] = "f"
	map2["j"] = "f"
	map2["k"] = "f"

	fmt.Println(map1)
	fmt.Println(map2)

	map3 := make(map[string]string)
	map3["a"] = "b"
	map3["b"] = "c"
	fmt.Println(map3)

	var map4 map[string]string
	map4 = map3
	map4["a"] = "qqqqq"
	map4["z"] = "zzz"
	fmt.Println(len(map3))
	fmt.Println(map3)
	fmt.Println(len(map4))
	fmt.Println(map4)

	for k, v := range map4 {
		fmt.Println(k, v)
	}

}

func testMap2() {
	m := map[string]string{
		"a": "b",
		"b": "c",
		"c": "d",
	}
	a, exist_10 := m["b"]
	fmt.Println(a)
	fmt.Println("before delete, exist 10: ", exist_10)
	delete(m, "b")
	a, exist_10 = m["b"]
	fmt.Println(a)
	fmt.Println("after delete, exist 10: ", exist_10)
}

func testMap3() {
	m := make(map[string]int)
	m["a"] = 1
	var lock sync.Mutex

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()
	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	}
}

func testRune2() {
	str := "hello world, 你好世界"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes的值为: %v\n", bytes)
	fmt.Printf("runes 的值为: %v \n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2的值为: %v\n", str2)
	fmt.Printf("str3的值为: %v\n", str3)
}

func testStrConv() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	num1 := 1234
	str1 := strconv.Itoa(num1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(str1)

	ui64, err1 := strconv.ParseUint(str1, 10, 64)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(ui64)
}

func testInterfaceTransfer() {
	var i interface{} = 23
	a, ok := i.(int)
	if ok {
		fmt.Println(ok)
	} else {
		fmt.Println(ok)
	}
	fmt.Printf("'%d' is a int", a)
}

func testInterfaceTransfer2() {
	var i interface{} = complex(1, 2)
	switch i.(type) {
	case int:
		fmt.Println(i.(int))
	case string:
		fmt.Println(i.(string))
	case complex128:
		fmt.Println("complex128", i.(complex128))
	default:
		fmt.Println(i)
	}
}

type SubSupplier interface {
	GetValue() int
	SetValue(v int)
}
type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

func (i *DigitSupplier) GetValue() int {
	return i.value
}

func (i *DigitSupplier) SetValue(v int) {
	i.value = v
}

func testInterfaceTransfer3() {
	var a SubSupplier = &DigitSupplier{value: 1}
	fmt.Println(a.GetValue())

	b, ok := a.(*DigitSupplier)
	fmt.Println(b, ok)
}

// 协程
func testGoRoutine() {
	go func() {
		fmt.Println("1.run goroutine in closure")
	}()

	go func(s string) {
		fmt.Println(s)
	}("2.goroutine: closure params")
	go say("3.in goroutine: world")
	say("4.say hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func testSafeCounter() {
	counter := SafeCounter{count: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter.Inc()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println(counter.getCount())
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock() //函数返回时解锁
	c.count++
}

func (c *SafeCounter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func testChannel() {
	ch := make(chan int)
	go testChanSendOnly(ch)
	time.Sleep(5 * time.Second)
	go testChanReceiveOnly(ch)

	var timeOut <-chan time.Time = time.After(time.Millisecond * 10)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed\n ")
				return
			}
			fmt.Printf("主gorountine接收到:%d\n", v)
		case <-timeOut:
			fmt.Println("timeout\n")
			return
		default:
			fmt.Println("default\n")
			time.Sleep(time.Millisecond * 500)
			return
		}
	}
}

func testChanReceiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到%d\n", v)
	}
}

func testChanSendOnly(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("发送--->%d\n", i)
	}
}

func testChanTimeout() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // 超时机制
		fmt.Println("timeout!")
	}
}

func testChan() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()
	<-ch1
	<-ch2
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	for {
		select {
		case i := <-ch1:
			fmt.Println("ch1:", i)
			time.Sleep(3 * time.Second)
		case i := <-ch2:
			fmt.Println("ch2:", i)
		case i := <-ch3:
			fmt.Println("ch3:", i)
		case <-ctx.Done():
			fmt.Println("timeout")
			if err := ctx.Err(); err != nil {
				fmt.Println("err")
			}
			return
		default:
			break
		}
	}

}
