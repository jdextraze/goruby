package object

func getMethods(class RubyClass, visibility MethodVisibility, addSuperMethods bool) *Array {
	var methodSymbols []RubyObject
	set := make(map[string]bool)
	for class != nil {
		methods := class.Methods().GetAll()
		for meth, fn := range methods {
			if fn.Visibility() == visibility && !set[meth] {
				methodSymbols = append(methodSymbols, &symbol{meth})
				set[meth] = true
			}
		}
		if !addSuperMethods {
			break
		}
		class = class.SuperClass()
	}

	return &Array{Elements: methodSymbols}
}
