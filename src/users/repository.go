package users

import "fmt"

var us = []*User{
	{
		Name: "John Doe", Description: "テストの説明_1",
	},
	{
		Name: "Jane Doe", Description: "説明_2",
	},
}

func GetUser(name *string) (u *User, err error) {
	for _, u := range us {
		if u.Name == *name {
			return u, nil
		}
	}

	return nil, fmt.Errorf("%s is not found", *name)
}
