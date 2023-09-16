package main

import "fmt"

// ITERATOR INTERFACE
type Iterator[T any] interface {
	next() *T
	hasNext() bool
}

type Collection interface {
	createIterator() Iterator[any]
}

// IMPLEMENT ITERATOR INTERFACE FOR USERS

type User struct {
	name string
	age  int
}

type UserCollection struct {
	users []*User
}

// createIterator implementation for user struct
func (u *UserCollection) createIterator() Iterator[User] {
	return &UserIterator{
		users: u.users,
	}
}

type UserIterator struct {
	index int
	users []*User
}

func (ui *UserIterator) next() *User {
	if ui.hasNext() {
		user := ui.users[ui.index]
		ui.index++
		return user
	}

	return nil
}

func (ui *UserIterator) hasNext() bool {
	return ui.index < len(ui.users)
}

// CLIENT CODE
func main() {
	user1 := &User{name: "Eric", age: 36}
	user2 := &User{name: "Lana", age: 35}

	UserCollection := &UserCollection{users: []*User{user1, user2}}

	iterator := UserCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.next()
		fmt.Printf("User is %s\n", user.name)
	}
}
