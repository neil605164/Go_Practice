package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	drive "google.golang.org/api/drive/v3"
)

// ServiceAccount 身份認證
func ServiceAccount(credentialFile string) *http.Client {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		log.Fatal(err)
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	_ = json.Unmarshal(b, &c)
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(oauth2.NoContext)
	return client
}

// GetDriveList 取 drive 清單
func GetDriveList(srv *drive.Service) {
	r, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}

// CreateToDrive 新增檔案至 google drive
func CreateToDrive(srv *drive.Service, filename string, parents ...string) {
	// 創建文檔的格式，若要創建資料夾請使用 application/vnd.google-apps.folder
	baseMimeType := "text/plain" // MimeType

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fileInf, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	f := &drive.File{
		Name:    filename,
		Parents: parents,
	}
	res, err := srv.Files.
		Create(f).
		ResumableMedia(context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()
	if err != nil {
		log.Fatalln("", err)
	}
	fmt.Printf("%s\n", res.Id)
}

func CreateDir(srv *drive.Service, dirname string, parents ...string) {
	baseMimeType := "application/vnd.google-apps.folder" // 創資料夾

	f := &drive.File{
		Name: dirname,
		// 可轉換為google driver檔案格式, 若不需轉換帶入空值即可
		MimeType: baseMimeType,
		Parents:  parents,
	}

	res, err := srv.Files.Create(f).SupportsAllDrives(true).Fields("id").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	fmt.Println("----->res", res.Id)
}

// DeleteToDrive 新增檔案至 google drive
func DeleteToDrive(srv *drive.Service, id string) {

	err := srv.Files.Delete(id).Do()
	if err != nil {
		log.Fatalln(err)
	}
}
