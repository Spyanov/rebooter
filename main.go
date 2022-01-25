package main

import (
	"bytes"
	"log"
	"net/http"
	"os/exec"
)

const IP = "127.0.0.1"
const PORT = ":3000"
const COMMAND = "/tools/reboot.sh"

func reboot(w http.ResponseWriter, r *http.Request) {
	runScript := exec.Command(COMMAND)
	var stdout bytes.Buffer
	runScript.Stdout = &stdout
	runScript.Run()
	log.Println("stdout: ", &stdout)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", reboot)
	http.ListenAndServe(IP+PORT, nil)
}
