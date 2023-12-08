package homework01

func VerifyIfCanDrive() {
	license := true
	age := 14

	if license && age >= 15 {
		println("You can drive")
	} else {
		println("You can't drive")
	}
}