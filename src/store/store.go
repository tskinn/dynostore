package store

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamodbItem struct {
	Key: string
	Value: string
}

func Put(db *dynamodb.DynamoDB, table, key, value string) error {
	dynamodbItem := &DynamodbItem{
		Key: key,
		Value: value,
	}
	item, err := dynamodbattribute.MarshalMap(dynamodbItem)
	if err != nil {
		return err
	}
	params := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(table),
	}
	_, err = db.PutItem(params)
	return err
}

func Get(db *dynamodb.DynamoDB, cluster, key string) string {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(key),
			},
		},
		TableName: aws.String(cluster),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		return ""
	}
	if resp.Item["value"].S != nil {
		return *resp.Item["value"].S
	}
	return ""
}
