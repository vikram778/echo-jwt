package model

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"echo-jwt/app/errs"
	"time"
)

type Client struct {
	Db        *sqlx.DB    `db:"-" json:"-"`
	ID        null.Int    `db:"id" json:"id"`
	Email     null.String `db:"email" json:"email"`
	UserName  null.String `db:"user_name" json:"user_name"`
	Password  null.String `db:"password" json:"password"`
	CreatedAt null.Time   `db:"created_at" json:"created_at"`
	UpdatedAt null.Time   `db:"updated_at" json:"updated_at"`
}

// NewClient function ...
func NewClient(db *sqlx.DB) (*Client, error) {
	if db == nil {
		return nil, errors.New("No databse connection")
	}

	return &Client{Db: db}, nil
}

func (me *Client) GetClient() (err error) {

	query := `SELECT * FROM clients WHERE user_name = ? OR email = ?`
	err = me.Db.Get(me, query, me.UserName, me.Email)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New(errs.ErrInternalDBError)
		return
	}

	return nil
}

func (me *Client) GetClientByEmail() (err error) {

	query := `SELECT * FROM clients WHERE email = ?`
	err = me.Db.Get(me, query, me.Email)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New(errs.ErrInternalDBError)
		return
	}

	return nil
}

func (me *Client) GetClientByID() (err error) {

	query := `SELECT * FROM clients WHERE id = ?`
	err = me.Db.Get(me, query, me.ID)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New(errs.ErrInternalDBError)
		return
	}

	return nil
}

func (me *Client) InsertOrUpdate(ok bool) (err error) {
	if ok {
		me.CreatedAt.SetValid(time.Now())
		query := `INSERT INTO clients (email,user_name, password, created_at) VALUES (?,?,?,?)`
		result, errsql := me.Db.Exec(query, me.Email, me.UserName, me.Password, me.CreatedAt)
		if errsql != nil {
			err = errors.New(errs.ErrInternalDBError)
			return
		}

		id, _ := result.LastInsertId()
		me.ID.SetValid(id)
	} else {
		me.UpdatedAt.SetValid(time.Now())
		query := `UPDATE clients set  password = ? ,updated_at=? WHERE id =?`
		_, errsql := me.Db.Exec(query, me.Password, me.UpdatedAt, me.ID)
		if errsql != nil {
			err = errors.New(errs.ErrInternalDBError)
			return
		}
	}
	return nil
}

func (me *Client) DeleteClient() (err error) {

	query := `DELETE FROM clients WHERE id = ?`
	_, err = me.Db.Exec(query, me.ID)
	if err != nil {
		err = errors.New(errs.ErrInternalDBError)
		return
	}

	return nil
}
