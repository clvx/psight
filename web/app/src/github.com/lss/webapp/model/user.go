package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const passwordSalt = "noadmin"

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	/*
		hasher := sha512.New()
		hasher.Write([]byte(passwordSalt))
		hasher.Write([]byte(email))
		hasher.Write([]byte(password))
		pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	*/
	// QueryRow executes a query that is expected to return at most one row.
	pwd := "password"
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname
		FROM public.users
		WHERE email = $1
			AND password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	t := time.Now()
	//ignoring the result object
	_, err = db.Exec(`
	UPDATE public.users
	SET lastlogin = $1
	WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	} else {
		log.Printf("new time %v: ", t)
	}
	return result, nil
}
