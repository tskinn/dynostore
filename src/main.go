package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/tskinn/dynostore/store"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "cluster",
			Value: "dynostore",
			Usage: "name of the dynamodb table used for dynostore",
		},
		cli.StringFlag{
			Name: "key",
			Value: "",
			Usage: "key to store in dynamodb",
		},
		cli.StringFlag{
			Name: "value",
			Value: "",
			Usage: "value of the value to be stored",
		},
		cli.BoolFlag{
			Name: "get",
			Usage: "get the value of the given key",
		},
		cli.BoolFlag{
			Name: "put",
			Usage: "put the key value pair in the cluster",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("get") && c.String("key") != "" {
			//
		}
		if c.Bool("put") && c.String("value") != "" && c.String("key") != "" {
			//
		}

		fmt.Println("get: ", c.Bool("get"))
		fmt.Println("put: ", c.Bool("put"))
		fmt.Println("cluster: ", c.String("cluster"))
		fmt.Println("key: ", c.String("key"))
		fmt.Println("value: ", c.String("value"))
		return nil
	}

	app.Run(os.Args)
}
