package model

import (
	"html"
	"os"
	"strings"
	"time"

	"github.com/giifrr/forum/api/security"
	log "github.com/sirupsen/logrus"
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

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Find(&users).Error
	if err != nil {
		log.Errorln(err.Error())
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

func (u *User) DeleteUser(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Delete(&User{})

	if db.Error != nil {
		log.Errorln(db.Error)
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
