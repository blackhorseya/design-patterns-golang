package factorymethod

// Operator declare actual operator interface
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory declare factory interface
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase declare struct based on Operator interface
type OperatorBase struct {
	a int
	b int
}

// SetA serve caller set a value
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB serve caller set b value
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator declare plus operator
type PlusOperator struct {
	*OperatorBase
}

// Result serve caller get a result
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// PlusOperatorFactory declare plus operator factory
type PlusOperatorFactory struct {
}

// Create serve caller to create an Operator
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator declare minus operator
type MinusOperator struct {
	*OperatorBase
}

// Result serve caller get a result
func (o MinusOperator) Result() int {
	return o.a - o.b
}

// MinusOperatorFactory declare minus operator factory
type MinusOperatorFactory struct {
}

// Create serve caller to create an Operator
func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
