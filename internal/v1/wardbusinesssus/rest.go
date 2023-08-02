package wardbusinesssus

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerWardBusinessSusAdapter interface {
		Get(*WardBusinessSus) error
		Search(*[]WardBusinessSus, WardBusinessSusParam) (int, error)
		Post(*WardBusinessSus) error
		Patch(WardBusinessSus) error
		Delete(*WardBusinessSus) error
	}

	RestWardBusinessSus struct {
		managerWardBusinessSus ManagerWardBusinessSusAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerWardBusinessSus(sl)
	hl := NewRestWardBusinessSus(ml)
	hl.LoadWardBusinessSusRoutes(eg)
}

func NewRestWardBusinessSus(mwa ManagerWardBusinessSusAdapter) *RestWardBusinessSus {
	return &RestWardBusinessSus{managerWardBusinessSus: mwa}
}

func (h *RestWardBusinessSus) LoadWardBusinessSusRoutes(eg *echo.Group) {
	eg.GET("/wardbusinesssus/:id", h.Get)
	eg.POST("/wardbusinesssus/search", h.Search)
	eg.POST("/wardbusinesssus", h.Post)
	eg.PATCH("/wardbusinesssus", h.Patch)
	eg.DELETE("/wardbusinesssus/:id", h.Delete)
}

func (h *RestWardBusinessSus) Get(c echo.Context) error {

	wardBusinessSus := &WardBusinessSus{}
	if err := h.managerWardBusinessSus.Get(wardBusinessSus); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *wardBusinessSus, nil, nil))
}

func (h *RestWardBusinessSus) Search(c echo.Context) error {
	param := WardBusinessSusParam{}
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
	wardBusinessSuss := &[]WardBusinessSus{}
	totalCount, err := h.managerWardBusinessSus.Search(wardBusinessSuss, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *wardBusinessSuss, nil, &totalCount))
}

func (h *RestWardBusinessSus) Post(c echo.Context) error {
	wa := &WardBusinessSus{}
	if err := c.Bind(wa); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerWardBusinessSus.Post(wa); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *wa, nil, nil))
}

func (h *RestWardBusinessSus) Patch(c echo.Context) error {
	wa := WardBusinessSus{}
	if err := c.Bind(&wa); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerWardBusinessSus.Patch(wa); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestWardBusinessSus) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	fmt.Println(id)
	wardBusinessSus := &WardBusinessSus{Id: int(id)}
	if err := h.managerWardBusinessSus.Delete(wardBusinessSus); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
