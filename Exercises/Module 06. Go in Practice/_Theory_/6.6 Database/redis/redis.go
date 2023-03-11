package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func RedisSimple() {
	c, err := redis.Dial("tcp", "localhost:6379") //в данном случае интерфейс Redis, куда как ближе к уже нам знакомому net, в итоге мы получаем нечто похожее на conn
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	c.Do("SET", "k1", 1) //ВСЕ действия с базой выполняются методом Do, с глаголами Redis
	n, _ := redis.Int(c.Do("GET", "k1")) //оборачиваем полученное значение
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(c.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)
}

func RedisString() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	c.Do("SET", "strVal", "This is a string value")
	n, _ := redis.String(c.Do("GET", "strVal"))
	fmt.Printf("%#v\n", n)
}

func RedisNonExisting() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	n, _ := redis.String(c.Do("GET", "noVal"))
	fmt.Printf("%#v\n", n)
}

func RedisScan() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
		// handle error
	}
	defer c.Close()

	// here we'll store our iterator value
	iter := 0

	// this will store the keys of each iteration
	var keys []string
	for {

		// we scan with our iter offset, starting at 0
		if arr, err := redis.Values(c.Do("SCAN", iter)); err != nil {
			panic(err)
		} else { // списки в редисе можно обернуть в Values, которая вернет слайс

			// получаем итератор и ключи из bulk reply
			iter, _ = redis.Int(arr[0], nil)
			keys, _ = redis.Strings(arr[1], nil)
		}

		fmt.Println(keys)

		// если итератор 0 то вываливаемся из цикла
		if iter == 0  {
			break
		}
	}
}
