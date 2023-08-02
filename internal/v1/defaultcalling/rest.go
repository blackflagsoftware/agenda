package defaultcalling

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerDefaultCallingAdapter interface {
		Get(*DefaultCalling) error
		Search(*[]DefaultCalling, DefaultCallingParam) (int, error)
		Post(*DefaultCalling) error
		Patch(DefaultCalling) error
		Delete(*DefaultCalling) error
	}

	RestDefaultCalling struct {
		managerDefaultCalling ManagerDefaultCallingAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerDefaultCalling(sl)
	hl := NewRestDefaultCalling(ml)
	hl.LoadDefaultCallingRoutes(eg)
}

func NewRestDefaultCalling(mdef ManagerDefaultCallingAdapter) *RestDefaultCalling {
	return &RestDefaultCalling{managerDefaultCalling: mdef}
}

func (h *RestDefaultCalling) LoadDefaultCallingRoutes(eg *echo.Group) {
	eg.GET("/defaultcalling/:id", h.Get)
	eg.GET("/defaultcalling", h.Search)
	eg.POST("/defaultcalling", h.Post)
	eg.PATCH("/defaultcalling", h.Patch)
	eg.DELETE("/defaultcalling/:id", h.Delete)
}

func (h *RestDefaultCalling) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	defaultCalling := &DefaultCalling{Id: int(id)}
	if err := h.managerDefaultCalling.Get(defaultCalling); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *defaultCalling, nil, nil))
}

func (h *RestDefaultCalling) Search(c echo.Context) error {
	param := DefaultCallingParam{}
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
	defaultCallings := &[]DefaultCalling{}
	totalCount, err := h.managerDefaultCalling.Search(defaultCallings, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *defaultCallings, nil, &totalCount))
}

func (h *RestDefaultCalling) Post(c echo.Context) error {
	def := &DefaultCalling{}
	if err := c.Bind(def); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerDefaultCalling.Post(def); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *def, nil, nil))
}

func (h *RestDefaultCalling) Patch(c echo.Context) error {
	def := DefaultCalling{}
	if err := c.Bind(def); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerDefaultCalling.Patch(def); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestDefaultCalling) Delete(c echo.Context) error {

	defaultCalling := &DefaultCalling{}
	if err := h.managerDefaultCalling.Delete(defaultCalling); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
