package object

import (
	"testing"
)

func TestBasicObjectMethodMissing(t *testing.T) {
	context := &callContext{receiver: NIL}
	result, err := basicObjectMethodMissing(context, &symbol{"foo"})

	checkResult(t, result, nil)

	expected := NewNoMethodError(NIL, "foo")

	checkError(t, err, expected)
}

func TestBasicObjectInitialize(t *testing.T) {
	context := &callContext{
		receiver: &Self{
			RubyObject: &classInstance{class: basicObjectClass},
			Name:       "BasicObject",
		},
	}

	result, err := basicObjectInitialize(context)

	checkError(t, err, nil)

	checkResult(t, result, context.Receiver())
}
