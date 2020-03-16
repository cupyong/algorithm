package main

import (
	"io"
	"log"
	"net/http"
	"fmt"
)



func main() {
    trie :=Makerouter(
        Get("/user",GetUser),
		Post("/user",PostUser),
	)
	HttpService(trie)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
func HttpService(tries *Trie){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		//获取method
		method := r.Method
		url := r.URL
		fmt.Println(url.Path)


		var httpFunc HandlerFunc

		for _,v:=range tries.routerList{
			fmt.Println(v.Method,method , v.Url,url.Path)
			if v.Method==method && v.Url == url.Path{
				httpFunc =v.Func
			}
		}
        if httpFunc!=nil{
			httpFunc(w,r)
			return
		}else {
			io.WriteString(w, "404")
			return
		}
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "get user")
}
func PostUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "post user")
}



//定义
type HandlerFunc func(http.ResponseWriter, *http.Request)
type Router struct {
	Method string
	Url    string
	Func   HandlerFunc
}
func Get(url string, handlerFunc HandlerFunc) *Router {
	return &Router{
		Method: "GET",
		Func:    handlerFunc,
		Url: url,
	}
}
func Post(url string, handlerFunc HandlerFunc) *Router {
	return &Router{
		Method: "GET",
		Func:    handlerFunc,
		Url: url,
	}
}

type Trie struct{
	routerList []*Router
}
func Makerouter(routes ...*Router)  *Trie{
	 routerList:= make([]*Router,0)
	for _,router :=range routes{
		routerList = append(routerList, router)
	}
	return  &Trie{
		routerList,
	}
}





