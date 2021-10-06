package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct{
	Title string `json:"title"`
}

func main(){
	http.HandleFunc("/", homePage)
	err:=http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	var dish Product
	data, errRead:=ioutil.ReadAll(r.Body)
	if errRead != nil{
		w.WriteHeader(400)
		return
	}
	r.Body.Close()
	err:= json.Unmarshal(data, &dish)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Write the name of dish"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	recipe:=callApi(dish)
	w.Write([]byte(recipe))
}

func callApi(dish Product) string{
	var recipe map[string][]map[string]string
	food:= dish.Title
	resp,_:=http.Get(fmt.Sprintf("https://www.themealdb.com/api/json/v1/1/search.php?s=%v",food))
	data,_:=ioutil.ReadAll(resp.Body)
	err:= json.Unmarshal(data,&recipe)
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(recipe["meals"][0]["strInstructions"])
	return recipe["meals"][0]["strInstructions"]
}
