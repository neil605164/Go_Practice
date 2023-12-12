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

	uo := ProviderBasicInfo()
	uo.SetName("neil")
	fmt.Println(uo.GetName())

	u.SetSex("female")
	fmt.Println(u.GetSex())
}

type IUser interface {
	IBasicUser
}

func ProviderUser(base IBasicUser) IUser {

	return &user{
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

type IBasicUser interface {
	GetName() string
	SetName(name string)
	GetSex() string
	SetSex(sex string)
}

type basicUser struct {
	name string
	sex  string
}

func ProviderBasicInfo() IBasicUser {
	return &basicUser{}
}

func (u *basicUser) GetName() string {
	return u.name
}

func (u *basicUser) SetName(name string) {
	u.name = name + "_other"
}

func (u *basicUser) GetSex() string {
	return u.sex
}

func (u *basicUser) SetSex(sex string) {
	u.sex = sex
}
