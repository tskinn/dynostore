

## How to create a cluster for dynostore
```
aws dynamodb create-table --table-name dynostore --attribute-definitions AttributeName=key,AttributeType=S --key-schema AttributeName=key,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```
