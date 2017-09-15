//go http上传文件
//学习http包 ioutil
//curl 127.0.0.1:9090/upload -F "uploadfile=@app.ini"
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const FormatDate = "2006-01-02"

func debugRequest(r *http.Request) {
	r.ParseForm()
	log.Println(r.Header)
	log.Println(r.Form)
	log.Println(r.PostForm)
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	log.Println(r.Header)
	log.Println(r.Method)
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Nginx 3.5")
	w.WriteHeader(200)
	debugRequest(r)
	fmt.Fprintln(w, "<h1>hello</h1>")
}
func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	savaPath := path.Join("upload", time.Now().Format(FormatDate))
	os.MkdirAll(savaPath, 0755)
	saveName := path.Join(savaPath, header.Filename)
	f, err := os.OpenFile(saveName, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, "%s", saveName)
}
func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil)
}
