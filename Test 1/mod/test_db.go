package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type About struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Group     string `bson:"group"`
}

type User struct {
	GithubID string `bson:"github_id"`
	TgId     string `bson:"tg_id"`
	Role     string `bson:"role"`
	About    About  `bson:"about"`
}

func main() {
	// Соединение с монгодб
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("UsersDB").Collection("user")

	//Выводит документ
	filter := bson.D{{"github_id", "123"}}

	var user1 User

	err = collection.FindOne(context.TODO(), filter).Decode(&user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user1)

	//Создание документа
	var question string
	fmt.Println("Вы хотите создать документ? y/n:")
	fmt.Scan(&question)
	if question == "y" {
		var user User
		fmt.Println("Введите github id:")
		fmt.Scan(&user.GithubID)

		fmt.Println("Введите tg id:")
		fmt.Scan(&user.TgId)

		fmt.Println("Введите роль:")
		fmt.Scan(&user.Role)

		fmt.Println("Введите Имя")
		fmt.Scan(&user.About.FirstName)

		fmt.Println("Введите Фамилию")
		fmt.Scan(&user.About.LastName)

		fmt.Println("Введите группу")
		fmt.Scan(&user.About.Group)

		insert_result, err1 := collection.InsertOne(context.TODO(), user)
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println("Документ создан ", insert_result)
	}

	//Удаление документа
	fmt.Println("Вы хотите удалить документ? y/n:")
	fmt.Scan(&question)

	if question == "y" {
		var delete_id string
		fmt.Println("Введите github id пользователя")
		fmt.Scan(&delete_id)

		filter = bson.D{{"github_id", delete_id}}

		delete_result, err2 := collection.DeleteOne(context.TODO(), filter)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("Документ удалён ", delete_result)
	}
}
