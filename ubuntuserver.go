package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
)

func getFile(res http.ResponseWriter, req *http.Request) {
	
	files, _ := ioutil.ReadDir("./publicFolder/")
	 res.Header().Set("Content-Type", "text/html; charset=utf-8")

	myHtml := `
	<!DOCTYPE html>
	<html>
		<head>
		<title>
			Home Network
		</title>
		</head>
		<body>
			<ul>`
	for _, f := range files	{
		fmt.Println(f.Name())
		myHtml += (`
			<li>
			 <a href="./publicFolder/` + f.Name() + `" download="`+f.Name()+`" >` + f.Name() + `</a>
			</li>`)

	}
	//
	myHtml += (
		`
		</ul>
		</body>
		</html>
		`)	

	io.WriteString(res, myHtml)
}

func main() {
	

	http.HandleFunc("/getFile/", getFile)
	http.Handle("/getFile/publicFolder/", http.StripPrefix("/getFile/publicFolder/", http.FileServer(http.Dir("./publicFolder"))))
	http.ListenAndServe(":8008", nil)
}
