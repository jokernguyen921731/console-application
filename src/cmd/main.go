/*
 *	Execution code and command line for end-user
 *  create_at: 15 Sep 2021 by Weifen
 */

package main

import (
	"console-application/src/service"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	cmd, ok := commands[flag.Arg(0)]
	if !ok {
		os.Exit(1)
	}
	cmd.do(flag.Args()[1:]...)
}

var commands = map[string]struct {
	name string
	do   func(...string)
}{
	"help": {
		name: "help",
		do:   help,
	},
	"start": {
		name: "start",
		do:   start,
	},
	"send": {
		name: "send",
		do:   send,
	},
	"version": {
		name: "version",
		do:   version,
	},
}

func help(args ...string) {
	if len(args) != 0 {
		log.Fatal("You mean './email-application help'?")
	}
	fmt.Println("\tThe application uses Golang language and Restful API for sending an email" +
		" to customer. \n\tThis software is created for interview purposes has a license" +
		" under PI.EXCHANGE company: https://www.pi.exchange/")

	var usageCommands = `
	Usage: ./email-application COMMAND [Params]
	Example: ./email-application help
	Commands:
    		start [port]                                                running the application as a server using Restful API
    		send [input_json] [input_csv] [output_folder] [error_csv]   sending an email to the customer using input template
    		version                                                     show current version`
	fmt.Println(usageCommands)
}

func start(args ...string) {
	fmt.Println("The application will implement starting service soon!")
}

func send(args ...string) {
	if len(args) != 4 {
		log.Fatal("command must contain 4 parameters should be like this: [input_json] [input_csv] [output] [error]")
	}
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		log.Fatal("directory of the json input file is not exist!")
	}
	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		log.Fatal("directory of the csv input file is not exist!")
	}
	if _, err := os.Stat(args[2]); os.IsNotExist(err) {
		log.Fatal("directory of the output folder is not exist!")
	}
	if _, err := os.Stat(args[3]); os.IsNotExist(err) {
		log.Fatal("directory of the csv error file is not exist!")
	}
	if err := service.MergeEmailHandler(args[0], args[1], args[2], args[3]); err != nil {
		log.Fatal("fatal error merging: ", err)
	}
}

func version(args ...string) {
	if len(args) != 0 {
		log.Fatal("You mean './email-application version'?")
	}
	fmt.Println("console-application@v0.0.1")
}
