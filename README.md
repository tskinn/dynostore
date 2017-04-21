

## How to create a cluster for dynostore
```
aws dynamodb create-table --table-name dynostore --attribute-definitions AttributeName=key,AttributeType=S --key-schema AttributeName=key,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```

```
NAME:
   dynostore

USAGE:
   dynostore --put --key <key> --value <value>
   dynostore --get --key <key>
   dynostore --getall --key <key>

VERSION:
   0.0.0

DESCRIPTION:
   A simple key-value store cli for dynamodb

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value  aws region in which the dynamodb table resides (default: "us-east-1")
   --table value   name of the dynamodb table used for dynostore (default: "dynostore")
   --key value     key to store in dynamodb
   --value value   value of the value to be stored
   --get           get the value of the given key
   --getall        get all the keys with the prefix of key
   --put           put the key value pair in the cluster
   --help, -h      show help
   --version, -v   print the version
```
