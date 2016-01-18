package objects

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

type Printable interface { // same as fmt.Stringer()
	String() string
}

type myInt int

func (i myInt) String() string {
	return strconv.Itoa(int(i))
}

func Test_Stringer(t *testing.T) {
	printable := []Printable{
		NewItem("Atom"),
		NewStarship("HerzAusGold", 42),
		myInt(42),
		os.ModeAppend | os.ModeSocket,
	}

	for _, p := range printable {
		fmt.Println(p.String())
	}
}
