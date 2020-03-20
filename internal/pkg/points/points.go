package points

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

var (
	errPointsInvalid = "points invalid: %s"
)

var (
	PointsMasterTable = os.Getenv("DDB_POINTS_TABLE_NAME")
)

// PointsService models a ddb table and methods to add
type PointsService struct {
	dynamo *dynamodb.DynamoDB
}

type User struct {
	Username     string
	Uuid         string
	Points       float64
	Transactions []PointsInput
}

type PointsInput struct {
	Username  string
	From      string
	Points    float64
	Timestamp time.Time
}

func NewPointsService(sess *session.Session) *PointsService {
	return &PointsService{
		dynamo: dynamodb.New(sess),
	}
}

func (p *PointsService) Put(input *PointsInput) error {
	getInput := dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(input.Username),
			},
		},
		TableName: aws.String(PointsMasterTable),
	}

	response, err := p.dynamo.GetItem(&getInput)
	if err != nil {
		return err
	}
	var userStruct User

	err = dynamodbattribute.UnmarshalMap(response.Item, &userStruct)
	if err != nil {
		return err
	}
	userStruct.Transactions = append(userStruct.Transactions, *input)
	userStruct.Points += input.Points

	serialized, err := dynamodbattribute.Marshal(userStruct)
	putInput := dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"username": serialized,
		},
		TableName: aws.String(PointsMasterTable),
	}
	putOutput, err := p.dynamo.PutItem(&putInput)
	if err != nil {
		return err
	}
	logrus.Debug(putOutput)
	return nil
}
