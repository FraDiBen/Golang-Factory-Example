package database

import "fmt"

type MySQLClient struct {
	store []interface{}
}

func (c *MySQLClient) Find() interface{} {
	fmt.Println("MySQL Client Find")
	x, store := pop(c.store)
	c.store = store
	return x
}

func (c *MySQLClient) Save(in interface{}) error {
	fmt.Println("MySQL Client Save")
	c.store = append(c.store, in)
	return nil
}
