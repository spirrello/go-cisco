package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	ssh "github.com/spirrello/switch-ssh-go"
)

var userFlag = flag.String("user", "svc_network", "a user")
var passwordFlag = flag.String("password", "password", "a password")
var ipFlag = flag.String("ip", "127.0.0.1", "device ip")
var confFlag = flag.String("conf", "config", "first config file")

//init sets up the flags
func init() {

	flag.StringVar(userFlag, "u", "", "a user")
	flag.StringVar(passwordFlag, "p", "", "a password")
	flag.StringVar(ipFlag, "i", "", "device ip")
	flag.StringVar(confFlag, "c", "", "first config file")
}

func main() {

	flag.Parse()

	sshCommands(*confFlag)
}

//readConfig reads the config files
func readConfig(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//sshCommands will run a list of commands stored in a file
func sshCommands(config string) {

	ipPort := *ipFlag + ":22"

	configList, err := readConfig(config)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//get the switch brand(vendor), include h3c,huawei and cisco
	brand, err := ssh.GetSSHBrand(*userFlag, *passwordFlag, ipPort)
	if err != nil {
		fmt.Println("GetSSHBrand err:\n", err.Error())
	}
	fmt.Println("Device brand is:\n", brand)

	result, err := ssh.RunCommands(*userFlag, *passwordFlag, ipPort, configList...)
	if err != nil {
		fmt.Println("RunCommands err:\n", err.Error())
	}
	fmt.Println("RunCommands result:\n", result)

}
