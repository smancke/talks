package phonebook

import (
	"encoding/json"
)

type Store interface {
	Set(key, value string) error
	Get(key string) (value string, exist bool)
}

type User struct {
	Id          string
	Name        string
	Phonenumber string
}

type Phonebook struct {
	store Store
}

func NewPhonebook(store Store) *Phonebook {
	return &Phonebook{
		store: store,
	}
}

func (p *Phonebook) Save(u User) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return p.store.Set(u.Id, string(b))
}

func (p *Phonebook) Get(id string) (User, bool) {
	s, exist := p.store.Get(id)
	if !exist {
		return User{}, false
	}

	u := User{}
	err := json.Unmarshal([]byte(s), &u)
	return u, err == nil
}
