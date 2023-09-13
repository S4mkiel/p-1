package controller

import (
	"fmt"
	"net/http"

	"github.com/S4mkiel/p-1/domain/entity"
	"github.com/S4mkiel/p-1/domain/service"
	"github.com/S4mkiel/p-1/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SexoController struct {
	Logger      *zap.SugaredLogger
	SexoService *service.SexoService
}

func NewSexoController(logger *zap.SugaredLogger, sexoService *service.SexoService) *SexoController {
	return &SexoController{
		Logger:      logger,
		SexoService: sexoService,
	}
}

func (c *SexoController) RegisterRoutes(app fiber.Router) {
	sexo := app.Group("/sexo")
	sexo.Post("/create", c.Create)
	sexo.Post("/update", c.Update)
	sexo.Get("/get", c.Get)
}

func (c *SexoController) Create(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//fazer payload
	type Payload struct {
		Sexo string `json:"sexo"`
	}

	//validar payload
	body := new(Payload)
	if err := ctx.BodyParser(body); err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao fazer o parse do body")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//verificações das regras de negocio
	if body.Sexo == "" {
		responseFalse.Error = append(responseFalse.Error, "Sexo é obrigatório")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//criando a entidade
	var s *entity.Sexo = &entity.Sexo{
		Sexo: body.Sexo,
	}

	//chamar service com o banco de dados
	sexo, err := c.SexoService.Create(s)
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Erro ao criar o sexo")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//retornar response
	responseTrue.Message = "Sexo criado com sucesso"
	responseTrue.Data = sexo

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}

func (c *SexoController) Update(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//fazer payload
	type Payload struct {
		ID   uint   `json:"id"`
		Sexo string `json:"sexo"`
	}

	//validar payload
	body := new(Payload)
	if err := ctx.BodyParser(body); err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao fazer o parse do body")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//verificações das regras de negocio
	if body.ID < 1 || body.Sexo == "" {
		responseFalse.Error = append(responseFalse.Error, "Id e Sexo são obrigatórios")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//criando a entidade
	var s *entity.Sexo = &entity.Sexo{
		Model: gorm.Model{
			ID: body.ID,
		},
		Sexo: body.Sexo,
	}

	//chamar service com o banco de dados
	sexo, err := c.SexoService.Update(s)
	if err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao criar o sexo")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//retornar response
	responseTrue.Message = "Sexo atualizado com sucesso"
	responseTrue.Data = sexo

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}

func (c *SexoController) Get(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//chamar service com o banco de dados
	sexo, err := c.SexoService.Get()
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Erro ao vizualizar os sexos")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//retornar response
	responseTrue.Message = "Sucesso"
	responseTrue.Data = sexo

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}
