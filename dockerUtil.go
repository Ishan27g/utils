package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)
func logIt(path string) (string) {
	cmd := exec.Command("/bin/bash", "-c", path)
	op, err := cmd.Output()
	if err != nil {
		return "error"
	}
	return string(op)
}
func listContainers(){
	fmt.Println(logIt(`sudo docker ps -a --format 'table {{.Names}}''\t''{{.Status}}'`))
}
func stopContainers(){
	fmt.Println(logIt(`sudo docker stop $(sudo docker ps -a -q)`))
	fmt.Println(logIt(`sudo docker rm $(sudo docker ps -a -q)`))
	fmt.Println("\nFollowing containers are running - ")
	fmt.Println(logIt(`sudo docker ps -a --format 'table {{.Names}}''\t''{{.Status}}'`))
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("-> list, stop, exit ?")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("list", text) == 0 {
			listContainers()
		}
		if strings.Compare("stop", text) == 0 {
			stopContainers()
		}
		if strings.Compare("exit", text) == 0 {
			return
		}

	}
}
