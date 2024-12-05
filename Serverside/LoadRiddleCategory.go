package Serverside

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Category struct {
	mathematic bool
	Logic      bool
	wordRiddle bool
	General    bool
}

func GetCategorySelected(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	category, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	if string(category) == "" {
		fmt.Println("hello")
	}
}
