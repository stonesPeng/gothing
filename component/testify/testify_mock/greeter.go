/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package testify_mock

type db struct {
}

func (d *db) FetchMessage(lang string) (string, error) {
	if lang == "en" {
		return "hello", nil
	}
	if lang == "es" {
		return "holla", nil
	}
	return "bzzzz", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type greeter struct {
	database DB
	lang     string
}

func (g greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang)
	return "Message is: " + msg
}

func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage()
	return "Message is: " + msg
}

type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

func NewDB() DB {
	return new(db)
}

func NewGreeter(db DB, lang string) GreeterService {
	return greeter{db, lang}
}
