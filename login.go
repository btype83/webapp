package main

import (
    "html/template"
    "net/http"
    "log"
)

type Product struct {
    Name string
    Price string
}

func main(){
    http.HandleFunc("/", Index)
    http.ListenAndServe(":8080", nil)
    

}


func check(err error){
    if err != nil {
        log.Fatal(err)
    }

}


func Index(w http.ResponseWriter, r *http.Request){
    
    if r.Method == "GET" {
        t,err := template.ParseFiles("Index.html")
        check(err)
        t.Execute(w, nil)
    }else {
        r.ParseForm()
        myPro := Product{}
        myPro.Name = r.Form.Get("Name")
        myPro.Price = r.Form.Get("Price")
        t,err := template.ParseFiles("Login.html")
        check(err)        
        t.Execute(w, myPro)
        

    }
}


  