# Download File in userscloud.com
Golang libary download file in userscloud.com 
## Installation
```
go get github.com/MosHelper/Userscloud-Download
```
## Quick example
```
package main

import (
	"github.com/MosHelper/Userscloud-Download"
	"os"
)

func main()  {
	var basePath = "/home/"+os.Getenv("USER")
	userscloud.DownloadFile(basePath+"/test_download.rar","https://userscloud.com/ztxyxgpo6nzp")
}
```
