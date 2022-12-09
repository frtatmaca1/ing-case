package controller

import (
	"net/http"

	"github.com/frtatmaca/case/service"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	personService *service.PersonService
}

func NewPersonController(personService *service.PersonService) *PersonController {
	return &PersonController{personService: personService}
}

func (p *PersonController) Person(ctx *gin.Context) {
	data, err := p.personService.Person()

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
