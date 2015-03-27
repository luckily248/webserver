package main

import (
	"fmt"
	"gopkg.in/mgo.v2-unstable"
	"gopkg.in/mgo.v2-unstable/bson"
	"html/template"
	"log"
	"net/http"
	"strings"
	//"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie") //写到输出
}
func testMongoDB() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	http.HandleFunc("/", sayhelloName) //访问的路径
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //端口
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}

}
