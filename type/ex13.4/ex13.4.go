package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VipUser struct {
	User  //기존 구조체 명 생략
	Price int
	Level int //내부 구조체 타입명 겹칠 시
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VipUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("Vip유저: %s ID: %s 나이: %d VIP레벨: %d 유저 레벨: %d\n", vip.Name, vip.ID, vip.Age, vip.Level, vip.User.Level) //포함된 구조체명 쓰고 접근
}
