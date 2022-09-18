package main

import (
	"fmt"
	"net/http"

	// go get -u github.com/gorilla/muxでインストールできる 参考：https://github.com/gorilla/mux
	// ルーティングなどに使用するもの 参考：https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a
	"github.com/gorilla/mux" 
)

func main() {
	// ルーティング用の関数
	r := mux.NewRouter()

	// ルートにgetリクエストでアクセスするとhello()関数を実行する
	r.HandleFunc("/", hello).Methods(http.MethodGet)

	// http://localhost:8085 で実行する
	http.ListenAndServe(":8085", r)
}

// この関数内容を変更して反映する場合もdocker-compose buildしてから、docker-compose upする
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}