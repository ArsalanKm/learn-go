package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func new(cfg Config) (*mongo.Database,error){
	opts:=options.Client()
	opts.ApplyURI(cfg.URL)

	client,err:=mongo.NewClient(opts)
	if err!=nil{
		return nil,fmt.Errorf("db New client errir %s",err)
	}
{

	ctx ,done‌:=context.WithTimeout(context.Background(),cfg.ConnectionTimeout)
	defer done()
	if err := client.Connect(ctx):err !=nil{
		return nil,fmt.Errorf("db connection error %w",err)
	}
}


{

	ctx ,done‌:=context.WithTimeout(context.Background(),cfg.ConnectionTimeout)
	defer done()
	if err := client.Ping(ctx,readpref.Primary()):err !=nil{
		return nil,fmt.Errorf("db connection error %w",err)
	}
}
return client.Database(cfg.URL),nil

}