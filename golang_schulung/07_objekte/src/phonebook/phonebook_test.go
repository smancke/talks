package phonebook

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Phonebook_Save(t *testing.T) {
	err := NewPhonebook(storeMock{
		t:     t,
		key:   "42",
		value: `{"Id":"42","Name":"Sebastian","Phonenumber":"01234"}`,
	}).Save(User{"42", "Sebastian", "01234"})
	assert.NoError(t, err)

	err = NewPhonebook(storeMock{
		t:     t,
		key:   "42",
		value: `{"Id":"42","Name":"Sebastian","Phonenumber":"01234"}`,
		err:   errors.New("some error"),
	}).Save(User{"42", "Sebastian", "01234"})
	assert.Error(t, err)
}

func Test_Phonebook_Get(t *testing.T) {
	expected := User{"42", "Sebastian", "01234"}
	result, exist := NewPhonebook(storeMock{
		t:     t,
		key:   "42",
		value: `{"Id":"42","Name":"Sebastian","Phonenumber":"01234"}`,
		exist: true,
	}).Get(expected.Id)
	assert.True(t, exist)
	assert.Equal(t, expected, result)
}

func Test_Phonebook_GetNotFound(t *testing.T) {
	_, exist := NewPhonebook(storeMock{
		t:     t,
		key:   "42",
		value: `{"Id":"42","Name":"Sebastian","Phonenumber":"01234"}`,
		exist: false,
	}).Get("42")
	assert.False(t, exist)
}

type storeMock struct {
	t     *testing.T
	key   string
	value string
	exist bool
	err   error
}

func (s storeMock) Set(key, value string) error {
	assert.Equal(s.t, s.key, key)
	assert.Equal(s.t, s.value, value)
	return s.err
}

func (s storeMock) Get(key string) (string, bool) {
	assert.Equal(s.t, s.key, key)
	return s.value, s.exist
}
