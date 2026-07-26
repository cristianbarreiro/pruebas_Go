package main

import "fmt"

func main() {
	var phonenumber = []uint64{5554123, 6665534, 768444, 9897654}
	var name = []string{"iraPohl", "randPill", "bettyDodo", "danielEuclid"}
	phone := make(map[string]uint64)
	var nameInput string
	for i := 0; i < len(name); i++ {
		phone[name[i]] = phonenumber[i]
	}
	fmt.Println("Enter name to lookup")
	fmt.Scanf("%s\n", &nameInput)
	fmt.Println("phone is", phone[nameInput])

}
