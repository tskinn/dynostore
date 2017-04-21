package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/tskinn/dynostore/store"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "region",
			Value: "us-east-1",
			Usage: "aws region in which the dynamodb table resides",
		},
		cli.StringFlag{
			Name: "table",
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
			Name: "getall",
			Usage: "get all the keys with the prefix of key",
		},
		cli.BoolFlag{
			Name: "put",
			Usage: "put the key value pair in the cluster",
		},
	}

	app.Action = func(c *cli.Context) error {
		var err error
		sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(c.String("region"))}))
		db := dynamodb.New(sess)
		if c.Bool("get") && c.String("key") != "" {
			// get the key!
			fmt.Println(store.Get(db, c.String("table"), c.String("key")))
		}
		if c.Bool("getall") && c.String("key") != "" {
			
			keys, values := store.GetAll(db, c.String("table"), c.String("key"))
			for i := 0; i < len(keys); i++ {
				fmt.Printf("%s=%s\n", keys[i], values[i])
			}
		}
		if c.Bool("put") && c.String("value") != "" && c.String("key") != "" {
			// put the key-value!
			err = store.Put(db, c.String("table"), c.String("key"), c.String("value"))
			if err != nil {
				fmt.Println("failed to put key-value")
			} else {
				fmt.Printf("Successfully put key-value:\n\t%s:\t%s\n", c.String("key"), c.String("value"))
			}
		}

		// fmt.Println("get:     ", c.Bool("get"))
		// fmt.Println("put:     ", c.Bool("put"))
		// fmt.Println("table:   ", c.String("table"))
		// fmt.Println("key:     ", c.String("key"))
		// fmt.Println("value:   ", c.String("value"))
		return err
	}

	app.Run(os.Args)
}