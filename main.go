package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

type Song struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	Title          string     `json:"title"`
	Artist         string     `json:"artist"`
	MusicAge       int        `json:"musicAge"`
	Image          string     `json:"image"`
	Video          string     `json:"video"`
	Album          string     `json:"album"`
	Description    string     `json:"description"`
	SpotifyTrackId string     `json:"spotifyTrackId"`
	UserID         uint       `json:"userId"`
}

func main() {
	test1 := "hoge"
	test2 := "fuga"

	song1 := Song{
		ID: 1,
	}

	song2 := Song{
		ID: 1,
	}

	fmt.Printf("%T: %v\n", song1, song1)

	if diff := cmp.Diff(test1, test2); diff != "" {
		fmt.Println(diff)
		//以下のように出力される
		//string(
		//	-       "hoge",
		//	+       "fuga",
		//)

		fmt.Println("異なる")
	} else {
		fmt.Println("同じ")
	}

	if diff := cmp.Diff(song1, song2); diff != "" {
		fmt.Println(diff)
		fmt.Println("異なる")
	} else {
		fmt.Println("同じ")
	}

	//h := func (w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello")
	//}

	//http.HandleFunc("/", myHandler)

	//req := httptest.NewRequest("GET", "/api/songs", nil)
	//res := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/post", PostHandler).Methods("POST")
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hogehoge")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postHandler")
}

