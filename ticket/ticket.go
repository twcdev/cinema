package ticket

import "fmt"

func init() {
	fmt.Println("init: ticket")
}

func BuyTicket(movie string) {
	fmt.Println("I bought ticket to ", movie)
}
