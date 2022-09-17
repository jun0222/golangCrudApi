package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello() // テストするHello()関数を呼び出す
    want := "Hello, world" // 期待する結果

    // gotとwantが異なる場合 hello_test.go:10: got "Hello, world" want "Hello world" のように教えてくれる
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}