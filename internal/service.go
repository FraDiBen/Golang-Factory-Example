package internal

import db "factoryExampleGo/internal/database"

type Svc struct {
	factory db.DBFactory // reference to the factory of a DB used by this service
}

func (s Svc) DoSomethingWithDB() {
	// create a DB instance: this is totally independent of the DB itself, and can be a cached DB too (like in this case)
	// notice how cachedSession is private to the database package and therefore even if global can't be instantiated
	// the only way is to use the factories
	session := s.factory.Create()

	for i := 0; i < 5; i++ {
		x := session.Read()
		_ = session.Write(x.(int) + 1)
	}
}

// NewSvc constructs a Service
func NewSvc(factory db.DBFactory) Svc {
	return Svc{factory: factory}
}
