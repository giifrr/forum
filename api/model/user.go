package model

import (
	"errors"
	"html"
	"os"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/giifrr/forum/api/security"
	"gorm.io/gorm"
)

type User struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Email      string    `gorm:"size:100;not null:unique" json:"email"`
	Password   string    `gorm:"size:100;not null" json:"password"`
	AvatarPath string    `gorm:"size:255;null" json:"avatar_path"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := security.Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) AfterFind() (err error) {
	if err != nil {
		return err
	}

	if u.AvatarPath != "" {
		u.AvatarPath = os.Getenv("DO_SPACES_URL") + u.AvatarPath
	}

	// dont return the user password
	u.Password = ""
	return nil
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "login":
		if u.Email == "" {
			err = errors.New("Required email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if u.Username == "" {
			err = errors.New("Required Password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Username == "" {
			err = errors.New("Required username")
			errorMessages["Required_username"] = err.Error()
		}
		if u.Password != "" && len(u.Password) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("Required email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}

	return errorMessages
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (u *User) FindUserById(db *gorm.DB, id uint32) (*User, error) {
	var err error

	err = db.Debug().Model(User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		err = gorm.ErrRecordNotFound
		return &User{}, err
	}

	return u, nil
}
