package controller

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	UPLOAD_DIR = "./uploads"
	RandomStr  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type UploadContoller struct {
}

func (u UploadContoller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		str := `
			<html>
				<head>
					<title>上传</title>
				</head>
				<body>
					<div>
						<form method="post" action="upload" enctype="multipart/form-data">
								<input type="file" name="file" />
								<input type="submit" value="上传" />
						</form>
					</div>
				</body>
			</html>
        `
		io.WriteString(w, str)
	}

	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(h.Filename)
		defer f.Close()
		filename := fmt.Sprintf("zc%s%d%s", getRandomStr(6), time.Now().UnixNano(), path.Ext(h.Filename))
		targetFile, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer targetFile.Close()
		_, err = io.Copy(targetFile, f)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}

func getRandomStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	res := ""
	sli := []byte(RandomStr)
	for i := 0; i < n; i++ {
		res += string(sli[rand.Intn(len(RandomStr))])
	}
	return res
}
