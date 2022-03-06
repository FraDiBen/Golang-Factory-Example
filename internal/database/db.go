package database

import "fmt"

//cachedSession is a session representing a DB connection. A session is not dependent from a DB, and can be compose by
// sub interfaces so that it better adjust to different databases
var cachedSession *Session

//DBFactory is an interface for a factory: a factory builds concrete implementations of various DB
type DBFactory interface {
	Create() *Session
}

//dbClient is the common interface for a basic DB
type dbClient interface {
	Save(in interface{}) error
	Find() interface{}
}

type Session struct {
	client dbClient
}

func (s Session) Read() interface{} {
	x := s.client.Find()
	if x == nil {
		x = 0
	}
	fmt.Printf("got %v\n", x)
	return x
}

func (s Session) Write(x interface{}) error {
	err := s.client.Save(x)
	fmt.Println("saved")
	return err
}

func pop(s []interface{}) (interface{}, []interface{}) {
	if len(s) == 0 {
		return nil, []interface{}{}
	}
	// we use s[0] because we want to pop the first element in the array
	s[len(s)-1], s[0] = s[0], s[len(s)-1]
	return s[len(s)-1], s[:len(s)-1]
}
