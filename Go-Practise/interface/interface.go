package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// our struct here is not dependent on a concrete type(like stripe or razorpay)
	// rather dependent on an interface.(paymenter)
	// So any type which implements that interface(paymenter) can use(be converted to line 22) our type and
	// then call the required method(common method for all types like (strip, razorpay)) like make_payment on it.
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func main() {
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Make payment using razorpay", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Make payment using paypal", amount)
}
