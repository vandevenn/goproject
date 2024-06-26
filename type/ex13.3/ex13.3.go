package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   string
	Age  int
}

type VipUser struct {
	User     //기존 구조체 명 생략
	VipLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VipUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("Vip유저: %s ID: %s 나이: %d VIP레벨: %d VIP 가격: %d만원\n", vip.Name, vip.ID, vip.Age, vip.VipLevel, vip.Price) //한 구조체로 내부 구조체 접근 가능
}
