package topics

import (
	"fmt"
	"strings"
)

func Stringer() {
	fmt.Println("Working with Strings")

	name := "Abike"
	friend := &name
	*friend = "Ayinke"
	fmt.Println(name, friend)

	fmt.Println(strings.ToUpper("adeleke"))
	fmt.Println(strings.Index("first", "i"))
}
