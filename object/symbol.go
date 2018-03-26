package object

import (
	"hash/fnv"
)

var symbolClass RubyClassObject = newClass(
	"Symbol",
	objectClass,
	symbolMethods,
	symbolClassMethods,
	func(RubyClassObject, ...RubyObject) (RubyObject, error) {
		return NewSymbol(""), nil
	},
)

func init() {
	classes.Set("Symbol", symbolClass)
}

var allSymbols = make(map[string]*symbol)

// TODO make symbol private and always use this
func NewSymbol(value string) *symbol {
	if s, found := allSymbols[value]; found {
		return s
	}
	s := &symbol{value}
	allSymbols[value] = s
	return s
}

// A symbol represents a symbol in Ruby
type symbol struct {
	Value string
}

// Inspect returns the value of the symbol
func (s *symbol) Inspect() string { return ":" + s.Value }

// Type returns SYMBOL_OBJ
func (s *symbol) Type() Type { return SYMBOL_OBJ }

// Class returns symbolClass
func (s *symbol) Class() RubyClass { return symbolClass }

func (s *symbol) hashKey() hashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return hashKey{Type: s.Type(), Value: h.Sum64()}
}

var symbolClassMethods = map[string]RubyMethod{
	"all_symbols": withArity(0, publicMethod(symbolAllSymbols)),
}

var symbolMethods = map[string]RubyMethod{
	"to_s": withArity(0, publicMethod(symbolToS)),
}

func symbolToS(context CallContext, args ...RubyObject) (RubyObject, error) {
	if sym, ok := context.Receiver().(*symbol); ok {
		return &String{Value: sym.Value}, nil
	}
	return nil, nil
}

func symbolAllSymbols(context CallContext, args ...RubyObject) (RubyObject, error) {
	elements := make([]RubyObject, len(allSymbols))
	i := 0
	for _, s := range allSymbols {
		elements[i] = s
		i++
	}
	return &Array{elements}, nil
}
