// golangでプログラムを作成する場合、必ずmainパッケージが存在する必要がある
// mainパッケージ中にmain関数が定義される必要がある
// 参考：https://dev-yakuza.posstree.com/golang/package/
package main

import "fmt" // Println関数を含むパッケージをインポート

func Hello() string {
    return "Hello, world" // 文字列を返す
}

func main() {
    fmt.Println(Hello()) // fmtが持っているPrintlnメソッドで文字列出力
}