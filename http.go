package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/wake", wakeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Default handler only shows a link to wake up the desired computer
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/wake\">Wake Computer?</a>")
}

// Execute the wol command if the URI /wake is visited
func wakeHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/bin/wol", mymac)
	err := cmd.Run()
	fmt.Fprintf(w, "wol command finished with error: %v", err)
}
