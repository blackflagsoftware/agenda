package wardbusinessrel

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerWardBusinessRelAdapter interface {
		Get(*WardBusinessRel) error
		Search(*[]WardBusinessRel, WardBusinessRelParam) (int, error)
		Post(*WardBusinessRel) error
		Patch(WardBusinessRel) error
		Delete(*WardBusinessRel) error
	}

	RestWardBusinessRel struct {
		managerWardBusinessRel ManagerWardBusinessRelAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerWardBusinessRel(sl)
	hl := NewRestWardBusinessRel(ml)
	hl.LoadWardBusinessRelRoutes(eg)
}

func NewRestWardBusinessRel(mwar ManagerWardBusinessRelAdapter) *RestWardBusinessRel {
	return &RestWardBusinessRel{managerWardBusinessRel: mwar}
}

func (h *RestWardBusinessRel) LoadWardBusinessRelRoutes(eg *echo.Group) {
	eg.GET("/wardbusinessrel/:id", h.Get)
	eg.POST("/wardbusinessrel/search", h.Search)
	eg.POST("/wardbusinessrel", h.Post)
	eg.PATCH("/wardbusinessrel", h.Patch)
	eg.DELETE("/wardbusinessrel/:id", h.Delete)
}

func (h *RestWardBusinessRel) Get(c echo.Context) error {

	wardBusinessRel := &WardBusinessRel{}
	if err := h.managerWardBusinessRel.Get(wardBusinessRel); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *wardBusinessRel, nil, nil))
}

func (h *RestWardBusinessRel) Search(c echo.Context) error {
	param := WardBusinessRelParam{}
	numberStr := c.QueryParam("page[number]")
	if numberStr != "" {
		var errNumber error
		param.Number, errNumber = strconv.Atoi(numberStr)
		if errNumber != nil {
			parse := ae.ParamError("page[number]", nil)
			return c.JSON(parse.StatusCode, util.NewOutput(c, nil, &parse, nil))
		}
	}
	sizeStr := c.QueryParam("page[size]")
	if sizeStr != "" {
		var errSize error
		param.Size, errSize = strconv.Atoi(sizeStr)
		if errSize != nil {
			parse := ae.ParamError("page[size]", nil)
			return c.JSON(parse.StatusCode, util.NewOutput(c, nil, &parse, nil))
		}
	}
	param.Sort = c.QueryParam("sort")
	if err := c.Bind(&param); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	wardBusinessRels := &[]WardBusinessRel{}
	totalCount, err := h.managerWardBusinessRel.Search(wardBusinessRels, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *wardBusinessRels, nil, &totalCount))
}

func (h *RestWardBusinessRel) Post(c echo.Context) error {
	war := &WardBusinessRel{}
	if err := c.Bind(&war); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerWardBusinessRel.Post(war); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *war, nil, nil))
}

func (h *RestWardBusinessRel) Patch(c echo.Context) error {
	war := WardBusinessRel{}
	if err := c.Bind(&war); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerWardBusinessRel.Patch(war); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestWardBusinessRel) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	fmt.Println(id)
	wardBusinessRel := &WardBusinessRel{Id: int(id)}
	if err := h.managerWardBusinessRel.Delete(wardBusinessRel); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
