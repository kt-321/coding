package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/tetete", nil)
	res := httptest.NewRecorder()

	//実行したいハンドラ
	myHandler(res, req)

	t.Log(res.Code)
	t.Log(res.Body)

	if res.Code != http.StatusOK {
		t.Errorf("but %d", res.Code)
	}
}

func TestPostHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/post", nil)
	res := httptest.NewRecorder()

	//TODO この書き方であってるのか
	PostHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("but %d", res.Code)
	} else {
		fmt.Println("success")
	}
}

