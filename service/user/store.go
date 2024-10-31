package user

import (
	"database/sql"
	"fmt"

	"github.com/yann-fk-21/todo-platform/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string)(*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email= ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)	

	for rows.Next() {
		u, err = scanRowInToUse(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
   return u, nil
}

func (s *Store)CreateUser(u types.User) error {
	_, err := s.db.Exec(
		"INSERT INTO users(firstname, lastname, email, password, createdAt) VALUES (?, ?, ?, ?, ?)",
		u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt,
	)

	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) GetUserByID(ID int)(*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id= ?", ID)
	if err != nil {
		return nil, err
	}

	u := new(types.User)	

	for rows.Next() {
		u, err = scanRowInToUse(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
   return u, nil
}

func scanRowInToUse(rows *sql.Rows)(*types.User, error) {
    user := new(types.User)

	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
    if err != nil {
		return nil, err
	}

	return user, nil
}
