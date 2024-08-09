package database

import (
	"database/sql"
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
