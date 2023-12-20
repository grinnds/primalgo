package primalgo

type Input interface{}

type Test interface {
	Test(input Input) bool
}
