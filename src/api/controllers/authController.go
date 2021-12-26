package controllers

import (
	"api/database"
	"api/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストデータをパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// パスワードチェック
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	// パスワードをエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	// データ登録
	// CreateはGORMのメソッド
	database.DB.Create(&user)

	return c.JSON(user)
}

// ログイン
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	// emailに紐づくユーザーを取得
	// &userを指定することでDBから取得したデータを直接格納できる
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// パスワードのチェック
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// JWT
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"jwt": token,
	})

}
