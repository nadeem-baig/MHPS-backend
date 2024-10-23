package user

import (
	"database/sql"

	"github.com/nadeem-baig/MHPS-backend/types"
	"github.com/nadeem-baig/MHPS-backend/utils/logger"
)

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	CreateUser(user types.User) error
	GetUserByID(id string) (*types.User, error)
}

type Store struct {
	db *sql.DB
}

// Ensure Store implements UserStore
var _ UserStore = (*Store)(nil)

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, logger.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, logger.Errorf("user not found")
	}
	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName,lastName,email,password) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}



// Ensure Store implements UserStore

type MembersStore interface {
	CheckMemberExistsByAadhaar(AadhaarNumber string) (bool, error)
	InsertMember(member *types.Member) error
}

// Ensure Store implements MembersStore
var _ MembersStore = (*Store)(nil)

// CheckMemberExistsByAadhaar checks if a member with the given Aadhaar number exists
func (s *Store) CheckMemberExistsByAadhaar(AadhaarNumber string) (bool, error) {
	// Use the correct PostgreSQL positional parameter syntax with $1
	query := "SELECT 1 FROM members WHERE aadhaar_number = $1 LIMIT 1"

	// Use QueryRow for efficiency and speed
	var exists int
	err := s.db.QueryRow(query, AadhaarNumber).Scan(&exists)

	// If no row is found, return false (user does not exist)
	if err == sql.ErrNoRows {
		return false, nil
	}

	// If there's any other error, return the error
	if err != nil {
		return false, err
	}

	// User exists
	return true, nil
}

// InsertMember inserts a new member into the members table
func (s *Store) InsertMember(member *types.Member) error {
	// Prepare the SQL query to insert a new member
	query := `
		INSERT INTO members (aadhaar_number, address, blood_group, contact_number, date_of_birth, education, email, father_name, marital_status, name, std_pin)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	// Execute the SQL insert statement
	_, err := s.db.Exec(query, member.AadhaarNumber, member.Address, member.BloodGroup, member.ContactNumber, member.DateOfBirth, member.Education, member.Email, member.FatherName, member.MaritalStatus, member.Name, member.StdPin)
	if err != nil {
		// Log and return the error if the insert fails
		return logger.Errorf("failed to insert member: %v", err)
	}

	// Return nil if the insert was successful
	return nil
}