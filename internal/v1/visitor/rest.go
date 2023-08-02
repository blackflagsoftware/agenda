package visitor

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerVisitorAdapter interface {
		Get(*Visitor) error
		Search(*[]Visitor, VisitorParam) (int, error)
		Post(*Visitor) error
		Patch(Visitor) error
		Delete(*Visitor) error
	}

	RestVisitor struct {
		managerVisitor ManagerVisitorAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerVisitor(sl)
	hl := NewRestVisitor(ml)
	hl.LoadVisitorRoutes(eg)
}

func NewRestVisitor(mvis ManagerVisitorAdapter) *RestVisitor {
	return &RestVisitor{managerVisitor: mvis}
}

func (h *RestVisitor) LoadVisitorRoutes(eg *echo.Group) {
	eg.GET("/visitor/:id", h.Get)
	eg.POST("/visitor/search", h.Search)
	eg.POST("/visitor", h.Post)
	eg.PATCH("/visitor", h.Patch)
	eg.DELETE("/visitor/:id", h.Delete)
}

func (h *RestVisitor) Get(c echo.Context) error {

	visitor := &Visitor{}
	if err := h.managerVisitor.Get(visitor); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *visitor, nil, nil))
}

func (h *RestVisitor) Search(c echo.Context) error {
	param := VisitorParam{}
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
	vistorss := &[]Visitor{}
	totalCount, err := h.managerVisitor.Search(vistorss, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *vistorss, nil, &totalCount))
}

func (h *RestVisitor) Post(c echo.Context) error {
	vis := &Visitor{}
	if err := c.Bind(vis); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerVisitor.Post(vis); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *vis, nil, nil))
}

func (h *RestVisitor) Patch(c echo.Context) error {
	vis := Visitor{}
	if err := c.Bind(&vis); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerVisitor.Patch(vis); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestVisitor) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	visitor := &Visitor{Id: int(id)}
	if err := h.managerVisitor.Delete(visitor); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
