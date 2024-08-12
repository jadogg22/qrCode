package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the global database connection pool
var DB *sql.DB

// Initialize initializes the SQLite database and creates the necessary tables
func init() {
	Initialize()
}
func Initialize() {
	var err error
	DB, err = sql.Open("sqlite3", "./sites.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS Sites (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL,
		key TEXT NOT NULL,
		site_name TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		salt TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
		email TEXT NOT NULL UNIQUE
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// InsertSite inserts a new site into the database
func InsertSite(user, key, siteName string) error {
	log.Println("Inserting site: ", siteName)
	insertSQL := `INSERT INTO Sites (user, key, site_name) VALUES (?, ?, ?)`
	_, err := DB.Exec(insertSQL, user, key, siteName)
	if err != nil {
		log.Println("Failed to insert site: ", err)
	}
	return err
}

// GetSites retrieves all sites from the database
func GetSites() ([]Site, error) {
	rows, err := DB.Query("SELECT id, user, key, site_name FROM Sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []Site
	for rows.Next() {
		var site Site
		err := rows.Scan(&site.ID, &site.User, &site.Key, &site.SiteName)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sites, nil
}

func GetSite(key string) (Site, error) {
	log.Println("Getting site with key: ", key)
	row := DB.QueryRow("SELECT id, user, key, site_name FROM Sites WHERE key = ?", key)

	var site Site
	err := row.Scan(&site.ID, &site.User, &site.Key, &site.SiteName)
	if err != nil {
		return Site{}, err
	}

	return site, nil
}

// Site represents a site record in the database
type Site struct {
	ID       int
	User     string
	Key      string
	SiteName string
}

func IncrementSite(site *Site) {
	if site == nil {
		log.Println("Site is nil")
		return
	}
	log.Println("Incrementing site: ", site.SiteName)
	return
}

func IsUnique(key string) bool {
	log.Println("Checking for unique key: ", key)
	row := DB.QueryRow("SELECT id FROM Sites WHERE key = ?", key)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return true
	}

	return false
}

func GetSitesByUser(user string) ([]Site, error) {
	// select all the sites for the user
	rows, err := DB.Query("SELECT id, user, key, site_name FROM Sites WHERE user = ?", user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []Site

	for rows.Next() {
		var site Site
		err := rows.Scan(&site.ID, &site.User, &site.Key, &site.SiteName)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sites, nil

}

// ErrUserExists is returned when a user already exists
var ErrUserExists = errors.New("user already exists")
var ErrEmailExists = errors.New("email already exists")

func UserExists(username, email string) error {
	rows, err := DB.Query("SELECT id, username, email FROM users WHERE username = ? OR email = ?", username, email)
	if err != nil {
		return err
	}
	defer rows.Close()

	// check if the user or email already exists

	var id int
	var user string
	var mail string

	for rows.Next() {
		err := rows.Scan(&id, &user, &mail)
		if err != nil {
			return err
		}
		if user == username {
			return ErrUserExists
		}
		if mail == email {
			return ErrEmailExists
		}
	}

	return nil
}

var (
	ErrUsernameTaken = errors.New("username already taken")
	ErrEmailTaken    = errors.New("email already taken")
	ErrDBInsert      = errors.New("failed to insert user into database")
)

func AddUser(username, passwordHash, salt, email string) error {
	// Insert the new user into the database
	insertSQL := `INSERT INTO users (username, password, salt, email) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(insertSQL, username, passwordHash, salt, email)
	if err != nil {
		if isUniqueConstraintError(err) {
			if isUsernameConstraintError(err) {
				return ErrUsernameTaken
			}
			if isEmailConstraintError(err) {
				return ErrEmailTaken
			}
		}
		return fmt.Errorf("%w: %v", ErrDBInsert, err)
	}
	return nil
}

// isUniqueConstraintError checks if the error is due to a unique constraint violation
func isUniqueConstraintError(err error) bool {
	// SQLite error codes for unique constraint violations are generally:
	// - "UNIQUE constraint failed" (error code 19)
	// This can be checked via error message or error code
	return err != nil && (err.Error() == "UNIQUE constraint failed" || err.Error() == "UNIQUE constraint failed: users.username" || err.Error() == "UNIQUE constraint failed: users.email")
}

// isUsernameConstraintError checks if the error is related to username uniqueness
func isUsernameConstraintError(err error) bool {
	// Specific handling for username constraint errors
	return err != nil && err.Error() == "UNIQUE constraint failed: users.username"
}

// isEmailConstraintError checks if the error is related to email uniqueness
func isEmailConstraintError(err error) bool {
	// Specific handling for email constraint errors
	return err != nil && err.Error() == "UNIQUE constraint failed: users.email"
}
