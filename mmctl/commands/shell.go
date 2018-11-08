package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Starts a shell to run commands interactively",
	RunE:  shellCmdF,
}

func init() {
	RootCmd.AddCommand(ShellCmd)
}

func shellCmdF(command *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("The shell command doesn't expect any arguments")
	}

	fmt.Println("Welcome to the mmctl shell. Type \"help\" to get a list of the available commands and \"quit\" to exit the shell.")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("mmctl> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Exiting...")
				return nil
			}
			fmt.Println("There was a problem while reading the input. Error: " + err.Error())
		}
		cleanInput := strings.TrimSuffix(input, "\n")
		inputArgs := strings.Split(cleanInput, " ")
		switch inputArgs[0] {
		case "quit":
			fmt.Println("Exiting...")
			return nil
		case "shell":
			fmt.Println("Cannot run a shell within a shell")
		default:
			Run(inputArgs)
		}
	}
}
