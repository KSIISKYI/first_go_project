package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func hashString(value string) string {
    hash := md5.New()
    io.WriteString(hash, value)
    return hex.EncodeToString(hash.Sum(nil))
}

func login() *User {
	var username string
	var password string
	var user *User
	var err error

	for {
		fmt.Print(makeTextColor("Enter Name: ", TextColorYellow))
		fmt.Scan(&username)

		user, err = findUserByLogin(username)

		if err != nil {
			fmt.Println(makeTextColor(err.Error(), TextColorRed))
		} else if user == nil {
			fmt.Println(makeTextColor("user not found", TextColorRed))
		} else {
			break
		}
	}

	for {
		fmt.Print(makeTextColor("Enter password: ", TextColorYellow))
		fmt.Scan(&password)

		if user.Password != hashString(password) {
			fmt.Println(makeTextColor("invalid password", TextColorRed))
		} else {
			break
		}
	}

	fmt.Println(makeTextColor("Welcome, " + user.Name, TextColorGreen))

	return user
}

func register() (*User, error) {
	var username string
	var password string

	for {
		fmt.Print(makeTextColor("Enter Name: ", TextColorYellow))
		fmt.Scan(&username)

		if err := validateUserName(username); err != nil {
			fmt.Println(makeTextColor(err.Error(), TextColorRed))
		} else {
			break
		}
	}

	for {
		fmt.Print(makeTextColor("Enter password: ", TextColorYellow))
		fmt.Scan(&password)

		if err := validatePassword(password); err != nil {
			fmt.Println(makeTextColor(err.Error(), TextColorRed))
		} else {
			break
		}
	}

	user := &User{
		Name: username,
		Password: hashString(password),
	}

	err := createUser(user)

	if err != nil {
		return nil, err
	}

	fmt.Println(makeTextColor("Welcome, " + user.Name, TextColorGreen))

	return user, nil
}

func validateUserName(username string) error {
	if len(username) < 8 {
		return errors.New("Name must be at least 8 characters long")
	}

	if len(username) > 20 {
		return errors.New("Name must be less than 20 characters")
	}

	if !isOnlyLetters(username) {
		return errors.New("Name must contain only letters")
	}

	user, err := findUserByLogin(username)

	if user != nil {
		return errors.New("user already exists")
	}

	return err
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > 20 {
		return errors.New("password must be less than 20 characters")
	}

	if !isOnlyLetters(password) {
		return errors.New("password must contain only letters")
	}

	return nil
}
