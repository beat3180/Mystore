package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//お店の情報登録
func Post(c *fiber.Ctx) error {

	// json情報の取得
	postData := c.FormValue("postsData")
	var data map[string]interface{}
	//json情報をパースする
	if err := json.Unmarshal([]byte(postData), &data); err != nil {
		return err
	}
	//型をintとしてキャスト
	user_id := int(data["user_id"].(float64))
	//spew.Dump(postData)
	// //リクエストデータをパースする
	// if err := postData.BodyParser(&data); err != nil {
	// 	return err
	// }

	// ファイルを取得
	file, err := c.FormFile("file")
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "ファイルがありません",
		})
	}

	fp, err := (file.Open())
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "ファイルを開けません",
		})
	}
	defer fp.Close()

	buff := make([]byte, 512)
	_, err = fp.Read(buff)
	if err != nil {

		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "ファイルのアップロードに失敗しました",
		})
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "許可されていないファイルタイプです。JPEGかPNGをアップロードしてください",
		})
	}

	//uuidを作成
	uu, err := uuid.NewRandom()
	if err != nil {

		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "ファイルのアップロードに失敗しました",
		})
	}

	// ファイルパス取得
	file_path := fmt.Sprintf("./tmp/%s", uu.String()+"."+file.Filename)

	// ファイルを保存
	if err := c.SaveFile(file, file_path); err != nil {

		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "ファイルの保存に失敗しました",
		})
	}

	post := models.Post{
		ImagePath:  file_path,
		UserId:     user_id,
		Title:      data["title"].(string),
		Comment:    data["comment"].(string),
		Lng:        data["lng"].(float64),
		Lat:        data["lat"].(float64),
		Address:    data["address"].(string),
		PublicFlag: data["public_flag"].(bool),
	}

	c.SendString("Upload : " + file.Filename)

	// データ登録
	// CreateはGORMのメソッド
	database.DB.Create(&post)

	return c.JSON(post)

}

//お店の情報取得
// func GetImage(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}
// 	file, err := os.Open(data["image_path"])
// 	defer file.Close()
// 	if err != nil {
// 		c.Status(500)
// 		return c.JSON(fiber.Map{
// 			"message": "ファイルを開くのに失敗しました",
// 		})
// 	}
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		c.Status(500)
// 		return c.JSON(fiber.Map{
// 			"message": "ファイルのデコードに失敗しました",
// 		})
// 	}

// 	buffer := new(bytes.Buffer)

// 	if err := jpeg.Encode(buffer, img, nil); err != nil {
// 		c.Status(500)
// 		return c.JSON(fiber.Map{
// 			"message": "ファイルのエンコードに失敗しました",
// 		})
// 	}
// 	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
// 	dataImage := map[string]interface{}{"Image": str}
// 	return renderTemplate(c, dataImage)

// }

//お店の情報取得
func GetPostID(c *fiber.Ctx) error {
	// URL(post/get/content/:id)からidを取得する
	id := c.Params("id")

	var post models.Post

	// &postを指定することでDBから取得したデータを直接格納できる
	database.DB.Where("ID = ?", id).First(&post)
	if post.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "コンテンツが見つかりません",
		})
	}

	return c.JSON(post)
}

//お店の情報取得
func GetPost(c *fiber.Ctx) error {

	post := []models.Post{}
	database.DB.Find(&post)

	return c.JSON(post)
}

//自分のお店の情報取得
func GetMystoreID(c *fiber.Ctx) error {

	var data map[string]int
	post := []models.Post{}

	// リクエストデータをパースする
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// &postを指定することでDBから取得したデータを直接格納できる
	database.DB.Where("user_id = ?", data["user_id"]).Find(&post)
	// if post.UserId == 0 {
	// 	c.Status(404)
	// 	return c.JSON(fiber.Map{
	// 		"message": "コンテンツが見つかりません",
	// 	})
	// }

	return c.JSON(post)
}
