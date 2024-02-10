package category

import "testing"

func TestNewName(t *testing.T) {
	test := "    "

	name, err := NewName(test)
	if *name == "" && err != nil {
		t.Errorf("sdas")
	}
}
