package users

import (
	"football/cmd/util"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"primaryKey"`
	Password string
}

// サインアップ処理
func Signup(r *http.Request) (*User, error) {

	user := User{}

	if validateResult := Validation(r); validateResult != nil {
		return nil, validateResult
	}

	// ユーザー入力パラメータ（メールアドレス、パスワード）
	email := r.FormValue("email")
	password := r.FormValue("password")

	// DB接続
	db, err := util.DbConnect()
	if err != nil {
		return nil, errors.WithStack(util.ERR_USER_SYSTEM_ERROR)
	}

	// DBマイグレーション
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}) // nolint

	// メールアドレス二重登録チェック
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return nil, util.ERR_USER_EMAIL_REGISTERED
	}

	// パスワード暗号化
	encryptPw, err := util.PasswordEncrypt(password)
	if err != nil {
		return nil, util.ERR_USER_SYSTEM_ERROR
	}

	// ユーザー登録
	user = User{Email: email, Password: encryptPw}
	db.Create(&user)

	return &user, nil
}

// ログイン
func Login(email, password string) (*User, error) {

	// DB接続
	db, err := util.DbConnect()
	if err != nil {
		return nil, err
	}

	// メールアドレスチェック
	user := User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return nil, util.ERR_USER_EMAIL_NOT_EXIST
	}

	// パスワードチェック
	err = util.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return nil, util.ERR_USER_PASSWORD_MISMATCH
	}
	return &user, nil
}

// ログインチェック
func CheckLogin(r *http.Request) (bool, error) {

	// CookieKey取得
	cookieKey := os.Getenv("FOOTBALL_REDIS_COOKIE")
	// ログインチェック
	id, err := util.GetSession(r, cookieKey)
	if err != nil {
		return false, err // 未ログインかつセッション取得時エラー
	}
	if id == "" {
		return false, nil // 未ログイン
	}
	err = util.ExtendSession(r, cookieKey)
	if err != nil {
		return false, err //　未ログインかつセッション延長時エラー
	}
	return true, nil // ログイン中
}
