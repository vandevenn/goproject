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
	UserInfo User //기존 구조체
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
	fmt.Printf("Vip유저: %s ID: %s 나이: %d VIP레벨: %d VIP 가격: %d만원\n", user.Name, user.ID, user.Age, vip.VipLevel, vip.Price)
}
