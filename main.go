package main

import "fmt"

const (
	MyMoney = 23.00
	Question1           = "Скільки грошей треба, щоб купити 9 яблук та 8 груш?"
	Question2           = "Скільки груш ми можемо купити?"
	Question3           = "Скільки яблук ми можемо купити?"
	Question4           = "Чи можемо ми купити 2 груші та 2 яблука"
)

var (
	ApplePrice          = 5.99
	PearPrice           = 7.00
	ApplesCount         = 9
	PearsCount          = 8
	TwoPiecces  float64 = 2
	Boolean             = true
)

func main() {
	fmt.Printf("%v \n Нам потрібно %v грн\n", Question1, float64(ApplesCount)*ApplePrice+float64(PearsCount)*PearPrice)
	fmt.Printf("%v \n Маючи %v грн, ми можемо купити %v груш\n", Question2, MyMoney, int64(MyMoney)/int64(PearPrice))
	fmt.Printf("%v \n Маючи %v грн, ми можемо купити %v яблук\n", Question3, MyMoney, int64(MyMoney)/int64(ApplePrice))
	if MyMoney >= ApplePrice*TwoPiecces+float64(PearsCount)*TwoPiecces {
		Boolean = true
	} else {
		Boolean = false
	}
	fmt.Printf("%v \n Відповідь: %v", Question4, Boolean)
}
