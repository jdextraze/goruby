package object

import (
	"fmt"
)

var basicObjectClass RubyClassObject = newClass(
	"BasicObject",
	nil,
	basicObjectMethods,
	basicObjectClassMethods,
	func(RubyClassObject, ...RubyObject) (RubyObject, error) { return &basicObject{}, nil },
)

func init() {
	classes.Set("BasicObject", basicObjectClass)
}

// basicObject represents a basicObject object in Ruby
type basicObject struct {
	_ int // for uniqueness
}

// Inspect returns empty string. BasicObjects do not have an `inspect` method.
func (b *basicObject) Inspect() string {
	fmt.Println("(Object doesn't support #inspect)")
	return ""
}

// Type returns the ObjectType of the array
func (b *basicObject) Type() Type { return BASIC_OBJECT_OBJ }

// Class returns the class of BasicObject
func (b *basicObject) Class() RubyClass { return basicObjectClass }

var basicObjectClassMethods = map[string]RubyMethod{}

var basicObjectMethods = map[string]RubyMethod{
	"initialize":                 privateMethod(basicObjectInitialize),
	"method_missing":             privateMethod(basicObjectMethodMissing),
	"==":                         withArity(1, publicMethod(basicObjectEqual)),
	"equal?":                     withArity(1, publicMethod(basicObjectEqual)),
	"!":                          withArity(0, publicMethod(basicObjectNot)),
	"!=":                         withArity(1, publicMethod(basicObjectNotEqual)),
	"singleton_method_added":     withArity(1, privateMethod(dummyObj)),
	"singleton_method_removed":   withArity(1, privateMethod(dummyObj)),
	"singleton_method_undefined": withArity(1, privateMethod(dummyObj)),
}

func basicObjectMethodMissing(context CallContext, args ...RubyObject) (RubyObject, error) {
	if len(args) < 1 {
		return nil, NewWrongNumberOfArgumentsError(1, 0)
	}
	method, ok := args[0].(*symbol)
	if !ok {
		return nil, NewImplicitConversionTypeError(method, args[0])
	}
	return nil, NewNoMethodError(context.Receiver(), method.Value)
}

func basicObjectInitialize(context CallContext, args ...RubyObject) (RubyObject, error) {
	return context.Receiver(), nil
}

func basicObjectEqual(context CallContext, args ...RubyObject) (RubyObject, error) {
	if context.Receiver() == args[0] {
		return TRUE, nil
	}
	return FALSE, nil
}

// TODO !
func basicObjectNot(context CallContext, args ...RubyObject) (RubyObject, error) {
	return FALSE, nil
}

func basicObjectNotEqual(context CallContext, args ...RubyObject) (RubyObject, error) {
	if context.Receiver().(*basicObject) != args[0].(*basicObject) {
		return TRUE, nil
	}
	return FALSE, nil
}

// TODO Move. used by Object, Kernel, Class and Module.
func dummyObj(context CallContext, args ...RubyObject) (RubyObject, error) {
	return NIL, nil
}

func notImplemented(context CallContext, args ...RubyObject) (RubyObject, error) {
	return nil, NewException("Not implemented")
}

func rbEqual(context CallContext, args ...RubyObject) (RubyObject, error) {
	if rubyEqual(context.Receiver(), args[0]) {
		return TRUE, nil
	}
	return FALSE, nil
}

func rubyEqual(obj1, obj2 RubyObject) bool {
	switch obj1 := obj1.(type) {
	case *symbol:
		obj2t, ok := obj2.(*symbol)
		if ok {
			return obj1.Value == obj2t.Value
		}
	case *String:
		obj2t, ok := obj2.(*String)
		if ok {
			return obj1.Value == obj2t.Value
		}
	case *Integer:
		obj2t, ok := obj2.(*Integer)
		if ok {
			return obj1.Value == obj2t.Value
		}
	default:
	}

	// Nil, True, False
	if obj1 == obj2 {
		return true
	}
	return false
}

// RTEST return true except for NIL

// RUBY_Qnil = 0x08 // 0000 1000
// 1111 0111

// 0000 1000
// 0000 0000
// false

// 1000 0000
// 1000 0000
// true
