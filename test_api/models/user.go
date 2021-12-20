package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	// "strconv"
	// "time"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	// u := User{"Akib", "Abdullah", "q.aqib.abdullah@gmail.com", "01677734142", "1234", "29-04-1998"}
	u := User{}
	UserList["user_11111"] = &u
}

type User struct {
	Firstname string
	Lastname  string
	Email     string `valid:"Email; MaxSize(100)"` // Need to be a valid Email address and no more than 100 characters.
	Phone     string `valid:"Phone"`               // Must be a valid mobile number
	Password  string `valid:"MinSize(6)"`
	DoB       string `valid:"Match(/^\d{4}-\d{2}-\d{2}$/)"`
}

func AddUser(u User) string {

	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
		// 	if err != nil {
		// 		panic(err)
	}

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
			// return "Data not Inserted"
		}
	} else {
		const (
			host     = "localhost"
			port     = 5432
			user     = "postgres"
			password = "newPassword"
			dbname   = "user_db"
		)

		// connection string
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)

		// close database
		defer db.Close()
		passwor := u.Password
		hash, _ := HashPassword(passwor)

		data := "SELECT * FROM information_schema.tables WHERE table_name = 'post_data';"
		table, _ := db.Exec(data)
		result, _ := (table.RowsAffected())

		if result == 0 {
			Createtable := "CREATE TABLE post_data(Firstname VARCHAR(255),Lastname VARCHAR (255),Email VARCHAR (255) PRIMARY KEY,Phone VARCHAR (255),Password VARCHAR (255) NOT NULL,DoB VARCHAR(255));"
			db.Exec(Createtable)
			sql := `INSERT INTO "post_data"("firstname", "lastname", "email", "phone", "password", "dob") VALUES ($1, $2, $3, $4, $5, $6)`
			_, e := db.Exec(sql, u.Firstname, u.Lastname, u.Email, u.Phone, hash, u.DoB)
			CheckError(e)

			// check db
			err = db.Ping()
			CheckError(err)
			fmt.Println("Connected!")

		} else {

			isEmail := "SELECT * FROM post_data WHERE Email=" + "'" + u.Email + "';"

			a, _ := db.Exec(isEmail)
			email_exist, _ := a.RowsAffected()

			if email_exist < 1 {
				sql := `INSERT INTO "post_data"("firstname", "lastname", "email", "phone", "password", "dob") VALUES ($1, $2, $3, $4, $5, $6)`
				_, e := db.Exec(sql, u.Firstname, u.Lastname, u.Email, u.Phone, hash, u.DoB)
				CheckError(e)

				// check db
				err = db.Ping()
				CheckError(err)
				fmt.Println("Connected!")
			} else {
				return "User Already Exist"
			}

		}

	}

	return "Succesfully Inserted"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}
