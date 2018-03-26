package object

var (
	nilClass RubyClassObject = newClass(
		"NilClass", objectClass, nilMethods, nilClassMethods, notInstantiatable,
	)
	// NIL represents the singleton object nil
	NIL RubyObject = &nilObject{}
)

func init() {
	classes.Set("NilClass", nilClass)
}

type nilObject struct{}

func (n *nilObject) Inspect() string  { return "nil" }
func (n *nilObject) Type() Type       { return NIL_OBJ }
func (n *nilObject) Class() RubyClass { return nilClass }

var nilClassMethods = map[string]RubyMethod{}

var nilMethods = map[string]RubyMethod{
	"to_i": withArity(0, publicMethod(notImplemented)),
	"to_f": withArity(0, publicMethod(notImplemented)),
	"to_s": withArity(0, publicMethod(notImplemented)),
	"to_a": withArity(0, publicMethod(notImplemented)),
	"to_h": withArity(0, publicMethod(notImplemented)),
	// "inspect": withArity(0, publicMethod(notImplemented)),
	"&": withArity(1, publicMethod(notImplemented)),
	"|": withArity(1, publicMethod(notImplemented)),
	"^": withArity(1, publicMethod(notImplemented)),
	// "===": withArity(1, publicMethod(rbEqual)),

	"nil?": withArity(0, publicMethod(nilIsNil)),
}

func nilIsNil(context CallContext, args ...RubyObject) (RubyObject, error) {
	return TRUE, nil
}
