package naming

import (
	"testing"
	"fmt"
)

func TestCamelcase(t *testing.T) {
	if CamelCase("foo_bar") != "FooBar" {
		t.Error("failed foo_bar")
	}

	if CamelCase("_foo_bar") != "FooBar" {
		t.Error("failed _foo_bar")
	}

	if UnderScore("FooBar") != "foo_bar" {
		t.Error("failed FooBar")
	}

	fmt.Println("FNV-1a Hashing:")
	fmt.Printf("%x\n", FNV1a("abc"))
	fmt.Printf("%x\n", FNV1a("工人小屋"))
}
