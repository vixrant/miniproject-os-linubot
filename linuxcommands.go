package main

import (
	"fmt"
	"os"
	"os/exec"
)

func GrepSearch(res NLPResponse) error {
	query := res.Entities["query"].(string)
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	command := exec.Command("grep", "-i", query, home, "-r")
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Grep executed")
	return nil
}

func TldrSearch(res NLPResponse) error {
	query := res.Entities["query"].(string)
	command := exec.Command("tldr", query)
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
