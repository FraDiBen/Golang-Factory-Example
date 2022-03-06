package internal

import (
	db "factoryExampleGo/internal/database"
)

// Run is the real main, can be tested
func Run(dbFactory db.DBFactory) {
	// pass the factory to the Constructor of teh Service (svc)
	svc := NewSvc(dbFactory)
	// svc doesn't know what DB it will act upon, the Factory will provide the right object
	// note that DoSomethingWithDB() is not irectly dependent of the DB type itself, this is encapsulated
	// by the factory
	svc.DoSomethingWithDB()
}
