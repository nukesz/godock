package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fileName := getDockerfile()
	runCommand(fileName)
}

func getDockerfile() string {
	return "template/Dockfile-go"
	// dat, err := ioutil.ReadFile("template/Dockfile-go")
	// if err != nil {
	// log.Fatalf("Could not open template file %v \n", err)
	// }
	// _ = dat
}

func runCommand(fileName string) {
	command := fmt.Sprintf("docker -f %v %v", fileName, strings.Join(os.Args[1:], " "))
	log.Printf("Executing command %v", command)

	cmd := exec.Command("docker", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully built docker image.")
}
