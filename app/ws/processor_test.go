package ws

import (
	"c1pherten/yet-webapp2/appctx"
	"reflect"
	"testing"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	var p = NewProcessor(nil)
	type User struct {
		ID int
		Name string
	}
	type MsgStub struct {
		User User
		ID int
		Content string
		Users []User
	}
	p.Register(&User{})

	u := &User{
		ID:   12123,
		Name: "hshshs",
	}
	b, err := p.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	t.Log("marshaled raw", string(b))

	var r any
	r, err = p.Unmarshal(b)
	if err != nil {
		t.Error(err)
	}

	newu := r.(*User)
	t.Log("newu", newu)

	// MsgStubs
	p.Register(&MsgStub{})
	msgs := MsgStub{
			User:    User{22, "ffff"},
			ID:      0,
			Content: "",
			Users:   []User{{1, "aa"}},
	}
	t.Log("====", reflect.TypeOf(&msgs).Elem().Name())

	b, err = p.Marshal(&msgs)
	if err != nil {
		t.Error(err)
	}

	v, err := p.Unmarshal(b)
	if err != nil {
		t.Error(err)
	}
	newMsgs := v.(*MsgStub)
	t.Log(newMsgs)
}

func TestSetHandler(t *testing.T) {
	type User struct {
		ID int
		Name string
	}

	c := appctx.NewContainer(nil, nil, "")
	p := NewProcessor(c)
	var u *User
	p.SetHandler(u, func (v interface{}) (any, error) {
		data := v.(*User)
		t.Log("in handler data:", data, data.Name)
		return data, nil
	})

	u = &User{
		ID:   13123,
		Name: "fuckfuck",
	}
	p.call(u)
}