package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

//noinspection ALL
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type news struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	BlogID   int    `json:"blog_id"`
}

type blog struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllNewsFromDB() []news {

	rows, err := db.Query("SELECT * FROM lab4.news")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()
	var _news []news

	for rows.Next() {
		var w news
		err := rows.Scan(&w.Id, &w.Url, &w.AuthorID, &w.BlogID, &w.Title, &w.Content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_news = append(_news, w)
	}
	return _news
}

func GetUserByEmail(email string) User {

	rows, err := db.Query("SELECT * FROM lab4.users where email == $1", email)
	if err != nil {
		fmt.Println(err)
		return User{}
	}

	defer rows.Close()
	var _user User

	if rows.Next() {
		err := rows.Scan(&_user.Id, &_user.Name, &_user.Email)
		if err != nil {
			fmt.Println(err)
		}
	}
	return _user
}

func getUserPasswordHashByEmail(email string) string {

	rows, err := db.Query("SELECT password FROM lab4.users where users.email = $1", email)
	if err != nil {
		fmt.Println(err)
		return "*"
	}

	defer rows.Close()
	var hash string

	if rows.Next() {
		err := rows.Scan(&hash)
		if err != nil {
			fmt.Println(err)
		}
	}
	return hash

}

//func GetSectionsFromDB() []section {
//
//	rows, err := db.Query("SELECT * FROM sections")
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	defer rows.Close()
//	var sections []section
//
//	for rows.Next() {
//		t := section{}
//		err := rows.Scan(&t.Id, &t.Name)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		sections = append(sections, t)
//	}
//
//	return sections
//
//}

func addUserToDB(user User) {

	_, err := db.Exec("insert into lab4.users VALUES (default, $1,$2,$3)",
		user.Name,
		user.Email,
		user.Password)

	if err != nil {
		fmt.Println(err)
	}

}

func addNewsToDB(news news) bool {

	_, err := db.Exec("insert into lab4.news VALUES (default,$1,$2,$3,$4,$5)",
		news.Url,
		news.AuthorID,
		news.BlogID,
		news.Title,
		news.Content)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func updateNews(news news) bool {

	_, err := db.Exec("update lab4.news set url = $2, authorid = $3, blogid = $4, title = $5, content = $6 where id=$1",
		news.Id,
		news.Url,
		news.AuthorID,
		news.BlogID,
		news.Title,
		news.Content)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

//func CheckUserWithLoginExists(login string) bool {
//
//	rows, err := db.Query("select count(*) from users where login = $1", login)
//
//	if err != nil {
//		fmt.Println(err)
//		return true
//	}
//
//	var res int
//
//	defer rows.Close()
//	rows.Next()
//	err = rows.Scan(&res)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//}

func DBConn() {

	err := error(nil)

	//connStr := "postgres://postgres:12345@172.18.0.1:54322/postgres"
	connStr := "postgres://postgres:12345@ec2-3-15-209-228.us-east-2.compute.amazonaws.com:54322/postgres"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

}
