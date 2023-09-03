package bootstrap

import "github.com/cp-Coder/khelo/mongo"

// Application struct defining application
type Application struct {
	Env   *Env
	Mongo mongo.Client
}

// App method to create new application
func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = *NewMongoDatabase(app.Env)
	return *app
}

// CloseDBConnection method to close database connection
func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
