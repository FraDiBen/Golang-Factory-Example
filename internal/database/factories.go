package database

type MySQLFactory struct{}

func (m MySQLFactory) Create() *Session {
	if cachedSession != nil {
		return cachedSession
	}
	cachedSession = &Session{client: &MySQLClient{store: []interface{}{}}}
	return cachedSession
}

type MongoFactory struct{}

func (m MongoFactory) Create() *Session {
	if cachedSession != nil {
		return cachedSession
	}
	cachedSession = &Session{client: &MongoClient{store: []interface{}{}}}
	return cachedSession
}
