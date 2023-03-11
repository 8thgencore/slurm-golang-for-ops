package mongo
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Number struct {
	Name string `bson:"name"` //для парсинга значений в/из монги, мы пользуемся аннотациями bson
	Value float64 `bson:"value"`
}

func MongoInsert() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))//коннектимся с опциями клиента, здесь я просто задаю URI

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		} //не забываем закрыть соединение
	}()

	collection := client.Database("testing").Collection("numbers") //указываем коллекцию с которой будем работать

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}}) //здесь данные уже запишутся в монгу
	id := res.InsertedID //печатаем ID

	fmt.Println(id)
}

func MongoInsertObject() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	num := Number{
		Name: "epsilon",
		Value: 2.789,
	}

	res, err := collection.InsertOne(ctx, num) //вставляем обьект вместо bson словаря на этот раз
	id := res.InsertedID

	fmt.Println(id)
}


func MongoGetList() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	var num Number

	res, err := collection.Find(ctx, bson.D{{}}) //синтаксис ПРАКТИЧЕСКИ аналогичен сырым запросам в монгу, легко запомнить
	for res.Next(ctx) { //стандартный итератор, пока есть следующее значение в списке продолжаем читать
		res.Decode(&num) //запихиваем значение в наш num
		elems,_ :=res.Current.Elements() //получаем сырой элемент
		fmt.Println(elems[0])
		fmt.Println(num)
	}
}

func MongoGetOne() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("testing").Collection("numbers")

	var num Number

	res := collection.FindOne(ctx, bson.D{{"name", "pi"}}) //фильтруем значения
	res.Decode(&num) //FindOne избавляет нас от необходимости итерироваться
}
