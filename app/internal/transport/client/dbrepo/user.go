package dbrepo

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/laterius/service_architecture_hw3/app/internal/domain"
	"github.com/laterius/service_architecture_hw3/app/internal/repo"
	"github.com/pkg/errors"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Get(userId domain.UserId) (user *domain.User, err error) {
	user = &domain.User{}
	err = r.db.Get(user, "SELECT * FROM users WHERE id = ? LIMIT 1", userId)
	if err != nil {
		return &domain.User{}, errors.Wrap(convertSQLError(err), "failed to get user")
	}

	return
}

// ByRemember looks up a user with the given remember token and returns that user.
// These methods expect the remember token to be already hashed.
// Errors handled as the same done by the ByEmail.
func (r *userRepo) ByRemember(rememberHash string) (user *domain.User, err error) {
	user = &domain.User{}
	err = r.db.Get(user, "SELECT * FROM users WHERE remember_hash = ?  LIMIT 1", rememberHash)
	if err != nil {
		return &domain.User{}, errors.Wrap(convertSQLError(err), "failed to get user")
	}

	return
}

func (r *userRepo) Login(username domain.Username, _ domain.Password) (user *domain.User, err error) {
	user = &domain.User{}
	err = r.db.Get(user, "SELECT * FROM users WHERE username = ?  LIMIT 1", username)
	if err != nil {
		return user, errors.Wrap(convertSQLError(err), "failed to get user")
	}

	return
}

func (r *userRepo) Exists(userId domain.UserId) (exists bool, err error) {
	res := domain.User{}
	err = r.db.Get(&res, "SELECT * FROM users WHERE id = ?  LIMIT 1", userId)
	if err != nil {
		return false, errors.Wrap(convertSQLError(err), "failed to get user")
	}

	return res.Id != 0, nil
}

func (r *userRepo) Create(u *domain.User) (*domain.User, error) {
	_, err := r.db.Exec(
		"INSERT INTO users (id, username, first_name, last_name, email, phone, hobby, age, gender, city, password_hash, remember_hash, remember) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		u.Id, u.Username, u.FirstName, u.LastName, u.Email, u.Phone, u.Hobby, u.Age, u.Gender, u.City, u.PasswordHash, u.RememberHash, u.Remember,
	)

	if err != nil {
		return u, errors.Wrap(convertSQLError(err), "failed to insert into users")
	}

	return u, nil
}

func (r *userRepo) Update(id domain.UserId, user *domain.User) (*domain.User, error) {
	user.Id = id

	_, err := r.db.Exec(
		"UPDATE users SET"+
			" username = ?,"+
			"first_name = ?,"+
			"last_name = ?,"+
			"hobby = ?,"+
			"age = ?,"+
			"gender = ?,"+
			"email = ?,"+
			"phone = ?, "+
			"city = ?,"+
			"password_hash = ?,"+
			"remember_hash = ?,"+
			"remember = ?",
		user.Username,
		user.FirstName,
		user.LastName,
		user.Hobby,
		user.Age,
		user.Gender,
		user.Email,
		user.Phone,
		user.City,
		user.PasswordHash,
		user.RememberHash,
		user.Remember,
	)

	if err != nil {
		return user, errors.Wrap(convertSQLError(err), "failed to insert into users")
	}

	return user, err
}

func (r *userRepo) PartialUpdate(id domain.UserId, data *domain.UserPartialData) (u *domain.User, err error) {
	upd := make(map[string]interface{})
	for k, v := range data.All() {
		upd[k] = v
	}

	_, err = r.db.Exec(
		"UPDATE users SET username = ?, first_name = ?, last_name = ?, hobby = ?, gender = ?, city = ?, email = ?, phone = ?",
		upd["username"],
		upd["first_name"],
		upd["last_name"],
		upd["hobby"],
		upd["gender"],
		upd["city"],
		upd["email"],
		upd["phone"],
	)

	if err != nil {
		return nil, errors.Wrap(convertSQLError(err), "failed to insert into users")
	}

	return r.Get(id)
}

func (r *userRepo) Delete(id domain.UserId) error {
	_, err := r.db.Exec(
		"DELETE FROM users WHERE id = ?", id,
	)

	if err != nil {
		return errors.Wrap(convertSQLError(err), "failed to insert into users")
	}

	return nil
}

func convertSQLError(err error) error {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return domain.ErrUserAlreadyExists
		}
	}

	switch {
	case err == sql.ErrNoRows:
		return domain.ErrUserNotFound
	}

	return err
}
