package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Simple(t *testing.T) {

	t.Logf("This Test fails")
	t.Fail()

}

func Test_With_Testify(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1)
}
