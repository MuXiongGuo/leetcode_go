// 模拟客户端上传文件
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{} // 是缓冲区,具有读写方法和可变大小的字节存储功能
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("err writing to buffer")
		return err
	}

	// 打开文件句柄
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("err opening file")
		return err
	}

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

// sample usage
func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./astaxie.pdf"
	postFile(filename, target_url)
}
