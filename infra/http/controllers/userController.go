package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/S4mkiel/p-1/domain/entity"
	"github.com/S4mkiel/p-1/domain/service"
	"github.com/S4mkiel/p-1/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController struct {
	Logger      *zap.SugaredLogger
	UserService *service.UserService
	SexoService *service.SexoService
}

func NewUserController(logger *zap.SugaredLogger, userService *service.UserService, sexoService *service.SexoService) *UserController {
	return &UserController{
		Logger:      logger,
		UserService: userService,
		SexoService: sexoService,
	}
}

func (c *UserController) RegisterRoutes(app fiber.Router) {
	user := app.Group("/user")
	user.Post("/create", c.Create)
	user.Post("/update", c.Update)
	user.Get("/get", c.Get)
	user.Post("/getById", c.GetByID)
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//fazer payload
	type Payload struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Name       string `json:"name"`
		Telefone   string `json:"telefone"`
		Email      string `json:"email"`
		SexoId     uint   `json:"sexo_id"`
		Nascimento string `json:"nascimento"`
	}

	//validar payload
	body := new(Payload)
	if err := ctx.BodyParser(body); err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao fazer o parse do body")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//verificações das regras de negocio
	if body.Name == "" || body.Username == "" || body.Password == "" {
		responseFalse.Error = append(responseFalse.Error, "Nome, username e senhas são obrigatórios")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//cripografia da senha
	hashPassword := md5.Sum([]byte(body.Password))
	hashPasswordString := hex.EncodeToString(hashPassword[:])

	//criando a entidade
	var u *entity.User = &entity.User{
		Username:   body.Username,
		Password:   hashPasswordString,
		Name:       body.Name,
		Telefone:   body.Telefone,
		Email:      body.Email,
		SexoId:     body.SexoId,
		Nascimento: body.Nascimento,
	}

	//chamar service com o banco de dados
	user, err := c.UserService.Create(u)
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao criar o usuário")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}
	sexo, sErr := c.SexoService.Get()
	if sErr != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao pegar os sexos")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	for _, value := range *sexo {
		if value.ID == user.SexoId {
			user.Sexo = value
		}
	}

	//retornar response
	responseTrue.Message = "Usuário criado com sucesso"
	responseTrue.Data = user

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//fazer payload
	type Payload struct {
		Id       uint   `json:"id"`
		Password string `json:"password"`
		Telefone string `json:"telefone"`
		Email    string `json:"email"`
		SexoId   uint   `json:"sexo_id"`
	}

	//validar payload
	body := new(Payload)
	if err := ctx.BodyParser(body); err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao fazer o parse do body")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//verificações das regras de negocio
	if body.Telefone == "" || body.Email == "" || body.Password == "" {
		responseFalse.Error = append(responseFalse.Error, "Telefone, email e senhas são obrigatórios")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//cripografia da senha
	hashPassword := md5.Sum([]byte(body.Password))
	hashPasswordString := hex.EncodeToString(hashPassword[:])

	//criando a entidade
	var u *entity.User = &entity.User{
		Model: gorm.Model{
			ID: body.Id,
		},
		Password: hashPasswordString,
		Telefone: body.Telefone,
		Email:    body.Email,
		SexoId:   body.SexoId,
	}

	//chamar service com o banco de dados
	user, err := c.UserService.Update(u)
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao criar o usuário")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	sexo, sErr := c.SexoService.Get()
	if sErr != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao pegar os sexos")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	for _, value := range *sexo {
		if value.ID == user.SexoId {
			user.Sexo = value
		}
	}

	//retornar response
	responseTrue.Message = "Usuário atualizado com sucesso"
	responseTrue.Data = user

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}

func (c *UserController) Get(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//chamar service com o banco de dados
	user, err := c.UserService.Get()
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao pegar os usuários")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//retornar response
	responseTrue.Message = "Usuário atualizado com sucesso"
	responseTrue.Data = user

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}

func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	responseTrue := dto.APIResponseTrue{}
	responseFalse := dto.APIResponseFalse{}

	//fazer payload
	type Payload struct {
		ID uint `json:"id"`
	}

	//validar payload
	body := new(Payload)
	if err := ctx.BodyParser(body); err != nil {
		fmt.Println(err)
		responseFalse.Error = append(responseFalse.Error, "Erro ao fazer o parse do body")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	//chamar service com o banco de dados
	user, err := c.UserService.GetByID(body.ID)
	if err != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao criar o usuário")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	sexo, sErr := c.SexoService.GetByID(user.SexoId)
	if sErr != nil {
		responseFalse.Error = append(responseFalse.Error, "Error ao pegar os sexos")
		return ctx.Status(http.StatusBadRequest).JSON(responseFalse)
	}

	type UserResponse struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Telefone string `json:"telefone"`
		Email    string `json:"email"`
		Sexo     struct {
			ID   uint   `json:"id"`
			Sexo string `json:"sexo"`
		}
		Nascimento string `json:"nascimento"`
	}

	var resp *UserResponse = &UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		Name:       user.Name,
		Telefone:   user.Telefone,
		Email:      user.Email,
		Nascimento: user.Nascimento,
		Sexo: struct {
			ID   uint   `json:"id"`
			Sexo string `json:"sexo"`
		}{
			ID:   sexo.ID,
			Sexo: sexo.Sexo,
		},
	}

	//retornar response
	responseTrue.Message = "Usuário criado com sucesso"
	responseTrue.Data = resp

	return ctx.Status(http.StatusCreated).JSON(responseTrue)
}
