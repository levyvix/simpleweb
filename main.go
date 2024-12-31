package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection string
	dsn := "user:user_password@(127.0.0.1:3306)/my_database?parseTime=true"

	// Configure the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Verify the connection
	{
		err = db.Ping()
		if err != nil {
			log.Fatalf("Failed to ping the database: %v", err)
		}

		fmt.Println("Successfully connected to the database!")

		//execute query
		query := `
    CREATE TABLE if not exists users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

		// Executes the SQL query in our database. Check err to ensure there was no error.
		_, err = db.Exec(query)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
		fmt.Println("Successfully created table!")
	}
	// inserting new values

	// username := "johndoe"
	// password := "secret"
	// createdAt := time.Now()

	// // Inserts our data into the users table and returns with the result and a possible error.
	// // The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	// result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	// if err != nil {
	// 	log.Fatalf("Failed to insert data: %v", err)
	// }

	// fmt.Println("Successfully inserted data!")
	// last_id, err := result.LastInsertId()
	// if err != nil {
	// 	log.Fatalf("Failed to get last inserted id: %v", err)
	// }
	// rows_affected, err := result.RowsAffected()
	// if err != nil {
	// 	log.Fatalf("Failed to get rows affected: %v", err)
	// }
	// fmt.Printf("Last inserted ID: %v, rows affected: %v\n", last_id, rows_affected)

	// query users
	{
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := `select id, username, password, created_at from users where id = ?`
		err = db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
		if err != nil {
			log.Fatalf("Failed to query data: %v", err)
		}
		fmt.Println("Successfully queried data!: Here is the result: ")
		fmt.Printf("ID: %v, username: %v, password: %v, created_at: %v\n", id, username, password, createdAt)
	}
	// query all users

	{
		type User struct {
			id         int
			username   string
			password   string
			created_at time.Time
		}

		rows, err := db.Query(`select id, username, password, created_at from users`)

		if err != nil {
			log.Fatal("Error selecting all users: ", err)
		}
		defer rows.Close()

		var users []User

		for rows.Next() {
			var u User
			err := rows.Scan(&u.id, &u.username, &u.password, &u.created_at)
			if err != nil {
				log.Fatalf("Error retrieving data from database: %v", err)
			}
			users = append(users, u)

		}
		err = rows.Err()
		if err != nil {
			log.Fatal("Error retrieving data from database: ", err)
		}

		for _, u := range users {
			fmt.Printf("ID: %v, username: %v, password: %v, created_at: %v\n", u.id, u.username, u.password, u.created_at)
		}
	}

	{ // Deleting a user from the database
		query := `delete from users where id = ?`

		_, err = db.Exec(query, 1)
		if err != nil {
			log.Fatalf("Failed to delete data: %v", err)
		}
		fmt.Println("Successfully deleted data!")
	}
}
