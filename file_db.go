package main

import (
	"encoding/json"
	"io"
	"os"
)

const (
	filePath = "db.json"
)

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

func findUserByLogin(login string) (*User, error) {
	file, err := getFileDB()

	if err != nil {
		return nil, err
	}

	// Close file after function execution
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var users []User
	err = json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Name == login {
			return &user, nil
		}
	}

	return nil, nil
}

func createUser(user *User) error {
	file, err := getFileDB()

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	var users []User

	if len(data) != 0 {
		err = json.Unmarshal(data, &users)
	}

	if err != nil {
		return err
	}

	users = append(users, *user)

	data, err = json.Marshal(users)

	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0666)
}

func updateUser(user *User) error {
	file, err := getFileDB()

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	var users []User
	err = json.Unmarshal(data, &users)

	if err != nil {
		return err
	}

	for i, u := range users {
		if u.Name == user.Name {
			users[i] = *user
		}
	}

	data, err = json.Marshal(users)

	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0666)
}

func getFileDB() (file *os.File, err error) {
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)

	return file, err
}
