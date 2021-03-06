// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// User represents a row from 'users'.
type User struct {
	ID       sql.NullInt64  `json:"id"`        // id
	Username sql.NullString `json:"username"`  // username
	Address  sql.NullString `json:"address"`   // address
	FullName sql.NullString `json:"full_name"` // full_name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO users (` +
		`username, address, full_name` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Username, u.Address, u.FullName)
	res, err := db.Exec(sqlstr, u.Username, u.Address, u.FullName)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = sql.NullInt64(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE users SET ` +
		`username = ?, address = ?, full_name = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Username, u.Address, u.FullName, u.ID)
	_, err = db.Exec(sqlstr, u.Username, u.Address, u.FullName, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM users WHERE id = ?`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UsersByFullName retrieves a row from 'users' as a User.
//
// Generated from index 'name_index'.
func UsersByFullName(db XODB, fullName sql.NullString) ([]*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, username, address, full_name ` +
		`FROM users ` +
		`WHERE full_name = ?`

	// run query
	XOLog(sqlstr, fullName)
	q, err := db.Query(sqlstr, fullName)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*User{}
	for q.Next() {
		u := User{
			_exists: true,
		}

		// scan
		err = q.Scan(&u.ID, &u.Username, &u.Address, &u.FullName)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

// UserByUsername retrieves a row from 'users' as a User.
//
// Generated from index 'sqlite_autoindex_users_1'.
func UserByUsername(db XODB, username sql.NullString) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, username, address, full_name ` +
		`FROM users ` +
		`WHERE username = ?`

	// run query
	XOLog(sqlstr, username)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&u.ID, &u.Username, &u.Address, &u.FullName)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByID retrieves a row from 'users' as a User.
//
// Generated from index 'users_id_pkey'.
func UserByID(db XODB, id sql.NullInt64) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, username, address, full_name ` +
		`FROM users ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Username, &u.Address, &u.FullName)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
