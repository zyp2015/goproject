package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"regexp"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func sayHelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	if r.Method == "GET" {
		t,_:=template.ParseFiles("./src/index.html")
		t.Execute(w,nil)
	}else {

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		curtime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(curtime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))

		t,_:=template.ParseFiles("./src/login.html")
		t.Execute(w,token)
	} else {
		fmt.Println(len(r.Form["username"]))
		fmt.Println(r.Form["password"])
		fmt.Println(r.Form)
		fmt.Println(r.Form["interest"][0])
		fmt.Fprintf(w, "username:%s\npassword:%s\n", r.Form["username"], r.Form["password"])
		if re, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("cardid")); re {
			fmt.Fprintf(w, "验证通过")
		}
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curtime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,string(curtime))
		token:=fmt.Sprintf("%x",h.Sum(nil))

		t,_:=template.ParseFiles("./src/upload.html")
		t.Execute(w,token)
	}else {
		r.ParseMultipartForm(1<<20)
		file,hander,err:=r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",hander.Header)
		f,err:=os.OpenFile("./"+hander.Filename,os.O_CREATE|os.O_WRONLY,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}
func main() {
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)

	http.Handle("/css/",http.StripPrefix("/css/",http.FileServer(http.Dir("./css"))))
	http.Handle("/img/",http.StripPrefix("/img/",http.FileServer(http.Dir("./img"))))

	err :=http.ListenAndServe(":80", nil)
	if err==nil{
		log.Fatal(err)
	}
}
