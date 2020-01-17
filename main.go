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
var conf1Flag = flag.String("conf1", "config", "first config file")

//init sets up the flags
func init() {

	flag.StringVar(userFlag, "u", "", "a user")
	flag.StringVar(passwordFlag, "p", "", "a password")
	flag.StringVar(ipFlag, "i", "", "device ip")
	flag.StringVar(conf1Flag, "c1", "", "first config file")
}

func main() {

	flag.Parse()

	ipPort := *ipFlag + ":22"

	//get the switch brand(vendor), include h3c,huawei and cisco
	brand, err := ssh.GetSSHBrand(*userFlag, *passwordFlag, ipPort)
	if err != nil {
		fmt.Println("GetSSHBrand err:\n", err.Error())
	}
	fmt.Println("Device brand is:\n", brand)

	//run the cmds in the switch, and get the execution results
	cmds := make([]string, 0)
	cmds = append(cmds, "sh run int vlan 5")
	//cmds = append(cmds, "dis vlan")
	result, err := ssh.RunCommands(*userFlag, *passwordFlag, ipPort, cmds...)
	if err != nil {
		fmt.Println("RunCommands err:\n", err.Error())
	}
	fmt.Println("RunCommands result:\n", result)

	firstFile, err := readConfig(*conf1Flag)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range firstFile {
		fmt.Println(line)
	}
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

}
