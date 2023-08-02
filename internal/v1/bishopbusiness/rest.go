package bishopbusiness

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerBishopBusinessAdapter interface {
		Get(*BishopBusiness) error
		Search(*[]BishopBusiness, BishopBusinessParam) (int, error)
		Post(*BishopBusiness) error
		Patch(BishopBusiness) error
		Delete(*BishopBusiness) error
	}

	RestBishopBusiness struct {
		managerBishopBusiness ManagerBishopBusinessAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerBishopBusiness(sl)
	hl := NewRestBishopBusiness(ml)
	hl.LoadBishopBusinessRoutes(eg)
}

func NewRestBishopBusiness(mbis ManagerBishopBusinessAdapter) *RestBishopBusiness {
	return &RestBishopBusiness{managerBishopBusiness: mbis}
}

func (h *RestBishopBusiness) LoadBishopBusinessRoutes(eg *echo.Group) {
	eg.GET("/bishopbusiness/:id", h.Get)
	eg.POST("/bishopbusiness/search", h.Search)
	eg.POST("/bishopbusiness", h.Post)
	eg.PATCH("/bishopbusiness", h.Patch)
	eg.DELETE("/bishopbusiness/:id", h.Delete)
}

func (h *RestBishopBusiness) Get(c echo.Context) error {

	bishopBusiness := &BishopBusiness{}
	if err := h.managerBishopBusiness.Get(bishopBusiness); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *bishopBusiness, nil, nil))
}

func (h *RestBishopBusiness) Search(c echo.Context) error {
	param := BishopBusinessParam{}
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
	bishopBusinesss := &[]BishopBusiness{}
	totalCount, err := h.managerBishopBusiness.Search(bishopBusinesss, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *bishopBusinesss, nil, &totalCount))
}

func (h *RestBishopBusiness) Post(c echo.Context) error {
	bis := &BishopBusiness{}
	if err := c.Bind(&bis); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerBishopBusiness.Post(bis); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *bis, nil, nil))
}

func (h *RestBishopBusiness) Patch(c echo.Context) error {
	bis := BishopBusiness{}
	if err := c.Bind(&bis); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerBishopBusiness.Patch(bis); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestBishopBusiness) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	bishopBusiness := &BishopBusiness{Id: int(id)}
	if err := h.managerBishopBusiness.Delete(bishopBusiness); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
