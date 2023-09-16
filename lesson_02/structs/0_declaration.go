package structs

// Private struct, visible inside this package only
type user struct {
	// Contains different field names with different types
	Name    string
	Surname string
	Age     int
}

// Public struct
type Post struct {
	// Public field
	User user
	// Private field
	text string
	// Private field
}

type ConsecutiveFields struct {
	age, money    int
	name, surname string
}

type TypeAlias = struct {
	helloFromInside string
}

type AnonymousStructInDeclaration struct {
	user     user
	settings struct {
		vip      bool
		loggedIn bool
	}
}

func anonymousStruct() {
	config := struct {
		key   string
		value string
	}{
		key:   "hello",
		value: "world",
	}

	config.key = "good bye"
}

func lookMaStructInside() {

	type hiddenType struct {
		ringToRuleThemAll string
	}

	var theOne hiddenType

	theOne.ringToRuleThemAll = "O"

}

// var notTheOne hiddenType
