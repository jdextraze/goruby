package object

import "testing"

func TestSymbol_hashKey(t *testing.T) {
	hello1 := &symbol{Value: "Hello World"}
	hello2 := &symbol{Value: "Hello World"}
	diff1 := &symbol{Value: "My name is johnny"}
	diff2 := &symbol{Value: "My name is johnny"}

	if hello1.hashKey() != hello2.hashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.hashKey() != diff2.hashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.hashKey() == diff1.hashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestSymbolToS(t *testing.T) {
	context := &callContext{
		receiver: &symbol{Value: "foo"},
	}

	result, err := symbolToS(context)

	checkError(t, err, nil)

	expected := &String{Value: "foo"}

	checkResult(t, result, expected)
}
