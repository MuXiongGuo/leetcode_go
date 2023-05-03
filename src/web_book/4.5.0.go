// new:上传文件
// 防止表单重复提交
// 用http包建立web服务器
// 表单信息验证
package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie") // 写入到客户端的

}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token:", token)

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		// 请求的是登录数据
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证token的合法性
		} else {
			// 不存在则报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		fmt.Println("token:", token)
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	http.HandleFunc("/login", login)   // 设置访问的路由
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // 监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// 访问 http://localhost:9090

}

// 补充正则表达式在表单验证中的使用例子
// 1.针对数字
// if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
// return false
//}
//
// 2.针对中文
// if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {
// return false
//}
//
// 3.针对英文
// if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
// return false
//}
//
// 4.验证电子邮件地址是否正确
// if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`,
//r.Form.Get("email")); !m {
// fmt.Println("no")
//}else{
// fmt.Println("yes")
//}
//
// 5.手机号码验证
// if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !
//m {
// return false
//}
//
// 6.身份证号码 15位或18位 且末位可能有X
//验证 15 位身份证，15 位的是全部数字
//if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
// return false
//}
//验证 18 位身份证，18 位前 17 位为数字，最后一位是校验位，可能为数字或字符 X。
//if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
// return false
//}
// 7. 所有的单选、复选都要注意内部设一个哈希或数组来验证是不是在给定范围内的，防止出现问题
