package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Post(c *fiber.Ctx) error {

	// json情報の取得
	postData := c.FormValue("postsData")
	spew.Dump(postData)
	var data map[string]interface{}
	json.Unmarshal([]byte(postData), &data)

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

	// _, ok := data["title"].(string)
	// if ok != true {
	// 	c.Status(400)
	// 	return c.JSON(fiber.Map{
	// 		"message": "お店の名前は必須項目です",
	// 	})
	// }

	post := models.Post{
		ImagePath:  file_path,
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
