package auth

import (
	"qrCode/pkg/database"
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

func TestSavePassword(t *testing.T) {
	// this is a simple test for the salt generation function
	// we just need to test if the salt is the correct length and if the salt is random
	tests := []struct {
		name        string
		pasword     string
		givenPass   string
		salt        string
		givenSalt   string
		expectError bool
	}{
		{
			name:        "good test",
			pasword:     "testpassword",
			givenPass:   "testpassword",
			expectError: false,
		},
		{
			name:        "good test",
			pasword:     "password",
			givenPass:   "password",
			expectError: false,
		},
		{
			name:        "good test",
			pasword:     "thePasword",
			givenPass:   "thePasword",
			expectError: false,
		},
		{
			name:        "bad password test",
			pasword:     "password",
			givenPass:   "password1",
			salt:        "12345",
			givenSalt:   "12345",
			expectError: true,
		},
		{
			name:        "bad password test",
			pasword:     "myPassword",
			givenPass:   "NotMyPassword",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, salt, err := SavePassword(tt.pasword)
			if err != nil && tt.expectError == false {
				t.Errorf("SavePassword() error = %v, wantErr %v", err, tt.expectError)
			}
			// checks if the password matches the hash and if expactHashMatch is oposite
			if CheckPasswordHash(tt.givenPass, salt, hashedPassword) == tt.expectError {
				t.Errorf("SavePassword() = %v, want %v", hashedPassword, tt.expectError)
			}
		})
	}
}

func TestRegisterUser(t *testing.T) {
	// this is a simple test for the salt generation function
	// we just need to test if the salt is the correct length and if the salt is random
	tests := []struct {
		name        string
		username    string
		password    string
		email       string
		expectError bool
	}{
		{
			name:        "good test",
			username:    "testuser",
			password:    "testpassword",
			email:       "testuser@gmail.com",
			expectError: false,
		},
		{
			name:        "bad test - same user",
			username:    "testuser",
			password:    "testpassword",
			email:       "testuser@gmail.com",
			expectError: true,
		},
		{
			name:        "good test",
			username:    "testuser2",
			password:    "testpassword",
			email:       "testuser2@gmail.com",
			expectError: false,
		},
		{
			name:        "good test",
			username:    "mikeOxLong",
			password:    "biggerThanAverage",
			email:       "mike@oxlong.university",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, salt, err := SavePassword(tt.password)
			if err != nil {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.expectError)
			}

			err = database.AddUser(tt.username, hashedPassword, salt, tt.email)
			if err != nil && tt.expectError == false {
				t.Errorf("RegisterUser() @AddUser error = %v, wantErr %v", err, tt.expectError)
			}

			err = CheckPassword(tt.username, tt.password)
			if err != nil {
				t.Errorf("RegisterUser() @CheckPassword error = %v, wantErr %v", err, tt.expectError)
			}

			err = CheckPassword(tt.username, "wrongPassword")
			if err == nil && tt.expectError == true {
				t.Errorf("RegisterUser() @CheckPassword @wrongPassword error = %v, wantErr %v", err, tt.expectError)
			}

		})
	}

}
