package database

import "fmt"

type MongoClient struct {
	store []interface{}
}

func (c *MongoClient) Find() interface{} {
	fmt.Println("Mongo Client Find")
	x, store := pop(c.store)
	c.store = store
	return x
}

func (c *MongoClient) Save(in interface{}) error {
	fmt.Println("Mongo Client Store")
	c.store = append(c.store, in)
	return nil
}
