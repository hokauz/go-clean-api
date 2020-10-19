package sample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokauz/go-clean-api/api/pkg/router"
	"github.com/hokauz/go-clean-api/api/pkg/validation"
	"github.com/hokauz/go-clean-api/core/entity"
	coreSample "github.com/hokauz/go-clean-api/core/sample"
)

var service coreSample.Service

// Start -
func Start(group *gin.RouterGroup, s coreSample.Service, point string) {
	service = s
	mapEnd := []*router.EndPoint{
		{
			Name:    point,
			Method:  "POST",
			Handler: create,
			Group:   group,
		},
		{
			Name:    point + "/:id",
			Method:  "PUT",
			Handler: update,
			Group:   group,
		},
		{
			Name:    point + "/:id",
			Method:  "GET",
			Handler: readOne,
			Group:   group,
		},
		{
			Name:    point,
			Method:  "GET",
			Handler: readAll,
			Group:   group,
		},
		{
			Name:    point + "/:id",
			Method:  "DELETE",
			Handler: delete,
			Group:   group,
		},
	}

	for _, point := range mapEnd {
		router.EnableHandlers(point)
	}
}

// TODO melhorar respostas
func create(ctx *gin.Context) (int, *router.Response) {
	var data entity.Sample
	err := ctx.BindJSON(&data)
	if err != nil {
		// TODO add response correta
		return http.StatusBadRequest, router.NewResposeError(err.Error(), "unexpected")
	}

	err = validation.Test(data)
	if err != nil {
		fmt.Println("Erro ao criar. Faltando parametros requeridos", err.Error())
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), "unexpected-match")
	}

	res, errCode, err := service.Create(&data)
	if err != nil {
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), errCode)
	}

	return http.StatusCreated, router.NewResponseSuccess(res)
}

func update(ctx *gin.Context) (int, *router.Response) {
	var data *entity.Sample
	id := ctx.Param("id")
	if id == "" {
		return http.StatusBadRequest, router.NewResposeError("Não possível reconhecer o objeto da operação", "invalid-id")
	}

	err := ctx.BindJSON(&data)
	if err != nil {
		// TODO add response correta
		return http.StatusBadRequest, router.NewResposeError(err.Error(), "place-body-error")
	}

	err = validation.Test(data)
	if err != nil {
		fmt.Println("Erro ao atualizar. Faltando parametros requeridos", err.Error())
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), "unexpected-match")
	}

	res, errCode, err := service.Update(id, data)
	if err != nil {
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), errCode)
	}

	return http.StatusOK, router.NewResponseSuccess(res)
}

func readOne(ctx *gin.Context) (int, *router.Response) {
	id := ctx.Param("id")

	if id == "" {
		// TODO add response correta
		return http.StatusBadRequest, nil
	}

	res, errCode, err := service.ReadOne(id)
	if err != nil {
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), errCode)
	}

	return http.StatusOK, router.NewResponseSuccess(res)
}

func readAll(ctx *gin.Context) (int, *router.Response) {
	res, errCode, err := service.ReadAll()
	if err != nil {
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), errCode)
	}

	return http.StatusOK, router.NewResponseSuccess(res)
}

func delete(ctx *gin.Context) (int, *router.Response) {
	id := ctx.Param("id")
	if id == "" {
		return http.StatusBadRequest, router.NewResposeError("Não possível reconhecer o objeto da operação", "invalid-id")
	}

	msg, err := service.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, router.NewResposeError(err.Error(), msg)
	}

	return http.StatusOK, router.NewResponseSuccess(nil)
}
