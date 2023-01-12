package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

// AddTableItem adds an item to an Amazon DynamoDB table
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	year is the year when the movie was released
//	table is the name of the table
//	title is the movie title
//	plot is a summary of the plot of the movie
//	rating is the movie rating, from 0.0 to 10.0
//
// Output:
//
//	If success, nil
//	Otherwise, an error from the call to PutItem
func AddTableItem(svc dynamodbiface.DynamoDBAPI, year *int, table, title, plot *string, rating *float64) error {
	item := Item{
		Year:   *year,
		Title:  *title,
		Plot:   *plot,
		Rating: *rating,
	}

	fmt.Printf("AddTableItem: %#v\n", item)

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: table,
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	// snippet-start:[dynamodb.go.create_new_item.args]
	table := flag.String("d", "", "The name of the database table")
	year := flag.Int("y", 0, "The year the movie debuted")
	title := flag.String("t", "", "The title of the movie")
	plot := flag.String("p", "", "The plot of the movie")
	rating := flag.Float64("r", -1.0, "The movie rating, from 0.0 to 10.0")
	flag.Parse()

	if *table == "" || *year == 0 || *title == "" || *plot == "" || *rating == -1.0 {
		fmt.Println("You must supply a database table name, year, title, plot and rating")
		fmt.Println("-d TABLE -y YEAR -t TITLE -r RATING")
		return
	}

	config := aws.NewConfig()
	if eu := os.Getenv("AWS_ENDPOINT_URL"); eu != "" {
		config = config.WithEndpoint(eu)
	}

	// snippet-start:[dynamodb.go.create_new_table.session]
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				SharedConfigState: session.SharedConfigEnable,
				Config:            *config,
			}))

	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.create_new_item.session]

	err := AddTableItem(svc, year, table, title, plot, rating)
	if err != nil {
		fmt.Println("Got an error adding item to table:")
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully added '"+*title+"' ("+strconv.Itoa(*year)+") to table "+*table+" with rating", *rating)
}
