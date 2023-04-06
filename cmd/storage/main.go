package main

import (
	"fmt"
	"github.com/suzibill/storage/internal/storage"
	"log"
)

func main() {
	fmt.Println("Aboba")
	st := storage.NewStorage()
	fmt.Println(st)

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	file, err = st.GetByID(file.ID)

	fmt.Println(file)

}
