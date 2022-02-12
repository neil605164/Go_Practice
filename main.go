package main

import (
	"Go_Practice/app"
	"context"
	"log"

	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {

	// 帳號認證
	client := app.ServiceAccount("credential.json") // Please set the json file of Service account.

	// 初始化 drive
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	// 取 drive 清單內容
	app.GetDriveList(srv)

	// 創建資料
	app.CreateToDrive(srv, "sample.txt", "1n0UbDipl2SmuoiUNdszAdJhVyGI46HWs")
}
