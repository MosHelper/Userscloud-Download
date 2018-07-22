package main

import (
	"os"

	"github.com/MosHelper/Userscloud-Download"
)

func main() {
	var basePath = "/home/" + os.Getenv("USER")
	userscloud.DownloadFile(basePath+"/test_download.rar", "https://userscloud.com/ztxyxgpo6nzp")
}
