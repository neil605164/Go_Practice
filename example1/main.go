package main

import "fmt"

func main() {
	basicUser := ProviderBasicInfo()

	u := ProviderUser(basicUser)
	u.SetName("neil")
	fmt.Println(u.GetName())

	u.SetSex("male")
	fmt.Println(u.GetSex())

	//////////////////////

	uo := ProviderUserOther(basicUser)
	uo.SetName("neil")
	fmt.Println(uo.GetName())

	u.SetSex("female")
	fmt.Println(u.GetSex())
}

type IUser interface {
	IBasicUser
	GetName() string
	SetName(name string)
}

func ProviderUser(base IBasicUser) IUser {

	return &user{
		IBasicUser: base,
	}
}

func ProviderUserOther(base IBasicUser) IUser {
	return &otherUser{
		IBasicUser: base,
	}
}

type user struct {
	IBasicUser
	name string
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

type otherUser struct {
	IBasicUser
	name string
}

func (u *otherUser) GetName() string {
	return u.name
}

func (u *otherUser) SetName(name string) {
	u.name = name + "_other"
}

type IBasicUser interface {
	GetSex() string
	SetSex(sex string)
}

type basicUser struct {
	sex string
}

func ProviderBasicInfo() IBasicUser {
	return &basicUser{}
}

func (u *basicUser) GetSex() string {
	return u.sex
}

func (u *basicUser) SetSex(sex string) {
	u.sex = sex
}
