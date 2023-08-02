package ordinance

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v3"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerOrdinanceAdapter interface {
		Get(*Ordinance) error
		Search(*[]Ordinance, OrdinanceParam) (int, error)
		Post(*Ordinance) error
		Patch(Ordinance) error
		Delete(*Ordinance) error
	}

	RestOrdinance struct {
		managerOrdinance ManagerOrdinanceAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerOrdinance(sl)
	hl := NewRestOrdinance(ml)
	hl.LoadOrdinanceRoutes(eg)
}

func NewRestOrdinance(mord ManagerOrdinanceAdapter) *RestOrdinance {
	return &RestOrdinance{managerOrdinance: mord}
}

func (h *RestOrdinance) LoadOrdinanceRoutes(eg *echo.Group) {
	eg.GET("/ordinance/:date", h.Get)
	eg.GET("/ordinance", h.Search)
	eg.POST("/ordinance", h.Post)
	eg.PATCH("/ordinance", h.Patch)
	eg.DELETE("/ordinance/:id", h.Delete)
}

func (h *RestOrdinance) Get(c echo.Context) error {
	date := c.Param("date")
	ordinance := &Ordinance{Date: null.StringFrom(date)}
	if err := h.managerOrdinance.Get(ordinance); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *ordinance, nil, nil))
}

func (h *RestOrdinance) Search(c echo.Context) error {
	param := OrdinanceParam{}
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
	ordinances := &[]Ordinance{}
	totalCount, err := h.managerOrdinance.Search(ordinances, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *ordinances, nil, &totalCount))
}

func (h *RestOrdinance) Post(c echo.Context) error {
	ord := &Ordinance{}
	if err := c.Bind(ord); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerOrdinance.Post(ord); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *ord, nil, nil))
}

func (h *RestOrdinance) Patch(c echo.Context) error {
	ord := Ordinance{}
	if err := c.Bind(ord); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerOrdinance.Patch(ord); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestOrdinance) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	ordinance := &Ordinance{Id: int(id)}
	if err := h.managerOrdinance.Delete(ordinance); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
