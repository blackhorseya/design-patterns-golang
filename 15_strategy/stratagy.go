package strategy

import "fmt"

// PaymentStrategy declare payment strategy interface
type PaymentStrategy interface {
	Pay(ctx *PaymentContext) string
}

// PaymentContext declare payment context
type PaymentContext struct {
	CardID string
	Name   string
	Money  int
}

// Payment declare payment
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

// NewPayment serve caller to create a payment
func NewPayment(name, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			CardID: cardID,
			Name:   name,
			Money:  money,
		},
		strategy: strategy,
	}
}

// Pay serve caller to pay something
func (p *Payment) Pay() string {
	return p.strategy.Pay(p.context)
}

// Cash declare cash
type Cash struct {
}

// Pay serve caller to pay money to someone
func (*Cash) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

// Bank declare bank
type Bank struct {
}

// Pay serve caller to pay money to someone
func (*Bank) Pay(ctx *PaymentContext) string {
	return fmt.Sprintf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
