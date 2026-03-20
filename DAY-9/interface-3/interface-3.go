package main

import "fmt"

type PaymentPolicyInterface interface {
	makeRefund(int)
}

// PaymentInterface extends PaymentPolicyInterface
type PaymentInterface interface {
	makePayment()
	PaymentPolicyInterface
}

type Payment struct {
	gateway PaymentInterface
}

// 1st gateway
type PayPal struct {
	amount float32
}

func (p PayPal) makePayment() {
	fmt.Println("Payment was successful form paypal of ₹", p.amount)
}
func (p PayPal) makeRefund(userId int) {
	fmt.Println("Refund is successfully form paypal")
}

// 1st gateway
type PhonePe struct {
	amount float32
}

func (p PhonePe) makeRefund(userId int) {
	fmt.Println("Refund is successfully")
}

func (p PhonePe) makePayment() {
	fmt.Println("Payment was successful from phone pay of ₹", p.amount)
}

func main() {
	payment := Payment{
		gateway: PayPal{
			amount: 100,
		},
	}

	payment.gateway.makePayment()
	payment.gateway.makeRefund(100)
}
