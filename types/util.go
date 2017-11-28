package types

var builtinTypes = map[string]bool{
	"bool":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"float32":    true,
	"float64":    true,
	"complex64":  true,
	"complex128": true,
	"string":     true,
	"int":        true,
	"uint":       true,
	"uintptr":    true,
	"byte":       true,
	"rune":       true,
}

var builtinFunctions = map[string]bool{
	"append":  true,
	"copy":    true,
	"delete":  true,
	"len":     true,
	"cap":     true,
	"make":    true,
	"new":     true,
	"complex": true,
	"real":    true,
	"imag":    true,
	"close":   true,
	"panic":   true,
	"recover": true,
	"print":   true,
	"println": true,
}

// Checks is type is builtin type.
func IsBuiltin(t Type) bool {
	if t.TypeOf() != T_Name {
		return false
	}
	for {
		switch tt := t.(type) {
		case TName:
			return IsBuiltinTypeString(tt.TypeName)
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return false
			}
			t = next.NextType()
		}
	}
}

func IsBuiltinTypeString(t string) bool {
	return builtinTypes[t]
}

func IsBuiltinFuncString(t string) bool {
	return builtinFunctions[t]
}

func IsBuiltinString(t string) bool {
	return IsBuiltinTypeString(t) || IsBuiltinFuncString(t)
}

// Returns name of type if it has it.
// Raw maps and interfaces do not have names.
func TypeName(t Type) *string {
	for {
		switch tt := t.(type) {
		case TName:
			return &tt.TypeName
		case TInterface:
			return nil
		case TMap:
			return nil
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return nil
			}
			t = next.NextType()
		}
	}
}

// Returns Import of type or nil.
func TypeImport(t Type) *Import {
	for {
		switch tt := t.(type) {
		case TImport:
			return tt.Import
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return nil
			}
			t = next.NextType()
		}
	}
}

// Returns first array entity of type.
// If array not found, returns nil.
func TypeArray(t Type) Type {
	for {
		switch tt := t.(type) {
		case TArray:
			return tt
		case TInterface:
			return nil
		case TMap:
			return nil
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return nil
			}
			t = next.NextType()
		}
	}
}

func TypeMap(t Type) Type {
	for {
		switch tt := t.(type) {
		case TInterface:
			return nil
		case TMap:
			return tt
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return nil
			}
			t = next.NextType()
		}
	}
}

func TypeInterface(t Type) Type {
	for {
		switch tt := t.(type) {
		case TInterface:
			return tt
		case TMap:
			return nil
		default:
			next, ok := tt.(LinearType)
			if !ok {
				return nil
			}
			t = next.NextType()
		}
	}
}

func IsType(f func(Type) Type) func(Type) bool {
	return func(t Type) bool {
		return f(t) != nil
	}
}

// Checks, is type contain some type.
// Generic checkers.
var (
	// Checks, is type contain array.
	IsArray = IsType(TypeArray)
	// Checks, is type contain map.
	IsMap = IsType(TypeMap)
	// Checks, is type contain interface.
	IsInterface = IsType(TypeInterface)
)
