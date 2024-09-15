package TestModel

import (
	"backend/Service/DB"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type testModel struct {
	id primitive.ObjectID `bson:"_id,omitempty"`
	Name string  `bson:"name"`
	Password string  `bson:"password"`

}


func InsirtTestModel(name string, password string) {
	fmt.Println("Insert Test Model")

	valx := testModel{
		id:       primitive.NewObjectID(),

		Name:name,Password:password}

		result, err :=DB.GetClient().Database("test").Collection("testuser").InsertOne(context.TODO(),valx)

		if err != nil {

			panic(err)

		}

		if(result.InsertedID==nil){

			panic("Insert Fail")

		}
	
	
}