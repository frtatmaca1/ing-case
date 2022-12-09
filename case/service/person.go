package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frtatmaca/case/errors/httperror"
	"github.com/frtatmaca/case/model/request"
	proxy "github.com/frtatmaca/case/proxy/person"
)

type PersonService struct {
	client proxy.Client
}

func NewPersonService(client proxy.Client) *PersonService {
	return &PersonService{client: client}
}

func (p *PersonService) Person() (request.Person, error) {
	var person request.Person
	res, err := p.client.GetPerson()

	if err != nil || res.StatusCode != http.StatusOK {
		return person, httperror.New(httperror.PersonNotFound)
	}

	err = json.NewDecoder(res.Body).Decode(&person)

	personGreeting := request.PersonGreeting{Person: person, Greeting: fmt.Sprintf("Hello %s (%d)", person.Name, person.Age)}
	_, err = p.client.CreatePerson(personGreeting)

	if err != nil || res.StatusCode != http.StatusOK {
		return person, httperror.New(httperror.PersonCreateError)
	}

	return person, nil
}
