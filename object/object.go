package object

var objectClass = newMixin(newClass(
	"Object",
	basicObjectClass,
	objectMethods,
	objectClassMethods,
	func(RubyClassObject, ...RubyObject) (RubyObject, error) {
		return &Object{}, nil
	},
), kernelModule)

func init() {
	classes.Set("Object", objectClass)
}

// Object represents an Object in Ruby
type Object struct {
	_ int // for uniqueness
}

// Inspect return ""
func (o *Object) Inspect() string { return "" }

// Type returns OBJECT_OBJ
func (o *Object) Type() Type { return OBJECT_OBJ }

// Class returns objectClass
func (o *Object) Class() RubyClass { return objectClass }

var objectClassMethods = map[string]RubyMethod{}

var objectMethods = map[string]RubyMethod{
	"send": publicMethod(objectSend),
	"public_send": publicMethod(objectPublicSend),
	"respond_to?": publicMethod(objectRespondTo),
}

func objectSend(context CallContext, args ...RubyObject) (RubyObject, error) {
	if len(args) < 1 {
		return nil, NewArgumentError("wrong number of arguments (given 0, expected at least 1)")
	}
	method, err := stringify(args[0])
	if err != nil {
		return nil, err
	}

	receiver := context.Receiver()
	class := receiver.Class()

	// search for the method in the ancestry tree
	for class != nil {
		fn, ok := class.Methods().Get(method)
		if !ok {
			class = class.SuperClass()
			continue
		}

		return fn.Call(context, args...)
	}

	methodMissingArgs := append(
		[]RubyObject{&symbol{method}},
		args...,
	)

	return methodMissing(context, methodMissingArgs...)
}

func objectPublicSend(context CallContext, args ...RubyObject) (RubyObject, error) {
	if len(args) < 1 {
		return nil, NewArgumentError("wrong number of arguments (given 0, expected at least 1)")
	}
	method, err := stringify(args[0])
	if err != nil {
		return nil, err
	}
	return Send(context, method, args[1:]...)
}

func objectRespondTo(context CallContext, args ...RubyObject) (RubyObject, error) {
	if len(args) < 1 {
		return nil, NewArgumentError("wrong number of arguments (given 0, expected at least 1)")
	}
	if len(args) > 2 {
		return nil, NewArgumentError("wrong number of arguments (given %d, expected at most 2)", len(args))
	}

	method, err := stringify(args[0])
	if err != nil {
		return nil, err
	}

	var includeAll *Boolean
	if len(args) == 2 {
		includeAll, _ = args[1].(*Boolean)
		if includeAll == nil {
			return nil, NewArgumentError("include_all: wrong argument type (expected Boolean)")
		}
	} else {
		includeAll = FALSE.(*Boolean)
	}

	receiver := context.Receiver()
	class := receiver.Class()

	// search for the method in the ancestry tree
	for class != nil {
		fn, ok := class.Methods().Get(method)
		if !ok {
			class = class.SuperClass()
			continue
		}

		if (fn.Visibility() == PRIVATE_METHOD || fn.Visibility() == PROTECTED_METHOD) && !includeAll.Value {
			break
		}

		return TRUE, nil
	}
	return FALSE, nil
}