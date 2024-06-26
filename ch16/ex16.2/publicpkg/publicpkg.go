package publicpkg

import "fmt"

const (
	PI = 3.1415   //공개
	pi = 3.141516 //비공개
)

var ScreenSize int = 1080 //공개
var screenHeight int      //비공개

func PublicFunc() { //공개 함수
	const MyConst = 100 //비공개 상수
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() { //비공개
	fmt.Println("This is a private function")
}

type MyInt int       //공개
type myString string //비공개

type MyStruct struct { //공개 구조체
	Age  int    //공개 구조체 필드
	name string //비공개 구조체 필드
}

func (m MyStruct) PublicMethod() { //공개 메소드
	fmt.Println("This is public method")
}

func (m MyStruct) privateMethod() { //비공개 메소드
	fmt.Println("This is a private method")
}

type myPrivateStruct struct { //비공개 구조체
	Age  int
	name string
}

func (m myPrivateStruct) PrivateMethod() { //비공개 메소드
	fmt.Println("This is a private method")
}
