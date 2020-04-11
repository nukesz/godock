package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	var image string
	flag.StringVar(&image, "image", "", "Docker image")
	flag.Parse()

	if len(image) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a Docker image name")
		os.Exit(1)
	}

	cmd := exec.Command("docker", "build", "-f template/Dockerfile-go", ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully built docker image.")
}
