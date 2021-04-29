package facade

import "fmt"

// API declare api interface
type API interface {
	Echo() string
}

type apiImpl struct {
	a AAPI
	b BAPI
}

// NewAPIImpl serve caller to create an API
func NewAPIImpl() API {
	return &apiImpl{
		a: NewAImpl(),
		b: NewBImpl(),
	}
}

func (i *apiImpl) Echo() string {
	a := i.a.EchoA()
	b := i.b.EchoB()
	return fmt.Sprintf("%s\n%s", a, b)
}

// AAPI declare an A API interface
type AAPI interface {
	EchoA() string
}

type aImpl struct {
}

// NewAImpl serve caller to create an AAPI
func NewAImpl() AAPI {
	return &aImpl{}
}

func (a *aImpl) EchoA() string {
	return "A running"
}

// BAPI declare a B API interface
type BAPI interface {
	EchoB() string
}

type bImpl struct {
}

// NewBImpl serve caller to create a BAPI
func NewBImpl() BAPI {
	return &bImpl{}
}

func (b *bImpl) EchoB() string {
	return "B running"
}
