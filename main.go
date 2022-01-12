package main

import (
	"fmt"
	A "fundamental/tasks/Task_A"
	AAsterisk "fundamental/tasks/Task_A_asterisk"
	B "fundamental/tasks/Task_B"
	C "fundamental/tasks/Task_C"
	D "fundamental/tasks/Task_D"
	E "fundamental/tasks/Task_E"
	"log"
	"net/http"
)

func main() {
	HTTPServerPort := ":8081"
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)

	a := A.Check()
	http.HandleFunc("/a",
		func(w http.ResponseWriter, _ *http.Request) {
			a.Render(w)
		})

	aAsterisk := AAsterisk.Check()
	http.HandleFunc("/a-asterisk",
		func(w http.ResponseWriter, _ *http.Request) {
			aAsterisk.Render(w)
		})

	b := B.Draw()
	http.HandleFunc("/b",
		func(w http.ResponseWriter, _ *http.Request) {
			b.Render(w)
		})

	c := C.Draw()
	http.HandleFunc("/c",
		func(w http.ResponseWriter, _ *http.Request) {
			c.Render(w)
		})

	d := D.Draw()
	http.HandleFunc("/d",
		func(w http.ResponseWriter, _ *http.Request) {
			d.Render(w)
		})

	http.HandleFunc("/e", func(w http.ResponseWriter, _ *http.Request) {
		E.Process(w)
	})

	fmt.Printf("Server running @ http://localhost%v\n", HTTPServerPort)
	if err := http.ListenAndServe(HTTPServerPort, nil); err != nil {
		log.Fatal(err)
	}
}
