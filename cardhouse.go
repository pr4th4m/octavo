package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

// SupportedOS to build infrastructure
var SupportedOS = []string{"centos:7.0", "ubuntu:14.04"}

func isSupportedOS(distro string) (string, error) {

	for _, os := range SupportedOS {
		if distro == os {
			return distro, nil
		}
	}

	return distro, fmt.Errorf("%s distro not supported. Current supported distros %v", distro, SupportedOS)
}

func main() {

	app := cli.NewApp()
	app.Name = "CardHouse"
	app.Version = "0.0.1"
	app.Usage = "Dynamic infrastructure for ansible playbook"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "os",
			Value: &cli.StringSlice{},
			Usage: "Supported os images",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			command := c.Args()[0]
			fmt.Println(command)
		} else {
			cli.ShowAppHelp(c)
			fmt.Println("Ansible playbook command is required")
		}

		selectedOS := c.StringSlice("os")
		if len(selectedOS) <= 0 {
			fmt.Println("Use flag --os to specify os type")
		}

		for _, osType := range selectedOS {
			_, err := isSupportedOS(osType)

			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(selectedOS)

		return nil
	}

	app.Run(os.Args)

}
