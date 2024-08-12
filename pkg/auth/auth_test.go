package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Testing the Hashing function before implementing the database
	// tests I need to write:
	// 1. Test a wrong password
	// 2. Test a correct password * 10
	// 3. test a correct password with a wrong salt

	tests := []struct {
		name             string
		password         string
		providedPassword string
		salt             string
		providedSalt     string
		expectHashMatch  bool
		expectError      bool
	}{
		{
			name:             "bad Password",
			password:         "testpassword1",
			providedPassword: "testpassword",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  false,
			expectError:      true,
		},
		{
			name:             "Bad Password2",
			password:         "testpassword",
			providedPassword: "not the same",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  false,
			expectError:      true,
		},
		{
			name:             "Correct Password and Salt",
			password:         "testpassword",
			providedPassword: "testpassword",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  true,
			expectError:      false,
		},
		{
			name:             "Correct Password and Salt",
			password:         "password",
			providedPassword: "password",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  true,
			expectError:      false,
		},

		{
			name:             "Correct Password and Salt",
			password:         "test",
			providedPassword: "test",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  true,
			expectError:      false,
		},
		{
			name:             "Correct Password and Salt",
			password:         "test123",
			providedPassword: "test123",
			salt:             "salt",
			providedSalt:     "salt",
			expectHashMatch:  true,
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, salt, err := SavePassword(tt.password)
			// this should pass for now
			if err != nil {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.expectError)
			}
			// checks if the password matches the hash and if expactHashMatch is oposite
			if !CheckPasswordHash(tt.providedPassword, salt, hashedPassword) == tt.expectHashMatch {
				t.Errorf("HashPasword() = %v, want %v", hashedPassword, tt.expectHashMatch)
			}
		})
	}
}

// user in the database
// id | username | password | salt | created_at | updated_at | email
func TestAddUser(t *testing.T) {
	// now we are getting the db going first we need to test a couple of things,
	// mostly just if the user is already in the database.

	tests := []struct {
		name        string
		username    string
		password    string
		email       string
		expectError bool
	}{
		{
			name:        "Add User",
			username:    "testuser",
			password:    "testpassword",
			email:       "test1@gmail.com",
			expectError: false,
		},
		{
			name:        "Add Duplicate User",
			username:    "testuser",
			password:    "testpassword",
			email:       "test1@gmail.com",
			expectError: true,
		},
		{
			name:        "Add User with duplicate email",
			username:    "testuser2",
			password:    "testpassword",
			email:       "test1@gmail.com",
			expectError: true,
		},
		// these should pass just adding a bunch of users to the database for fun
		{
			name:        "Good user",
			username:    "testuser3",
			password:    "testpassword",
			email:       "tester3@gmail.com",
			expectError: false,
		},
		{
			name:        "Good user",
			username:    "testuser4",
			password:    "testpassword",
			email:       "tester4@gmail.com",
			expectError: false,
		},
		{
			name:        "Good user",
			username:    "SwagDaddy69",
			password:    "dahBoi",
			email:       "SwagDaddyBoi69@gmail.com",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddUser(tt.username, tt.password, tt.email)
			if err != nil && tt.expectError == false {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.expectError)
			}

			if err == nil && tt.expectError == true {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.expectError)
			}
		})
	}

}
