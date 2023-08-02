package roles

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerRolesAdapter interface {
		Get(*Roles) error
		Search(*[]Roles, RolesParam) (int, error)
		Post(*Roles) error
		Patch(Roles) error
		Delete(*Roles) error
	}

	RestRoles struct {
		managerRoles ManagerRolesAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerRoles(sl)
	hl := NewRestRoles(ml)
	hl.LoadRolesRoutes(eg)
}

func NewRestRoles(mrol ManagerRolesAdapter) *RestRoles {
	return &RestRoles{managerRoles: mrol}
}

func (h *RestRoles) LoadRolesRoutes(eg *echo.Group) {
	eg.GET("/roles/:id", h.Get)
	eg.GET("/roles", h.Search)
	eg.POST("/roles", h.Post)
	eg.PATCH("/roles", h.Patch)
	eg.DELETE("/roles/:id", h.Delete)
}

func (h *RestRoles) Get(c echo.Context) error {

	roles := &Roles{}
	if err := h.managerRoles.Get(roles); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *roles, nil, nil))
}

func (h *RestRoles) Search(c echo.Context) error {
	param := RolesParam{}
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
	roless := &[]Roles{}
	totalCount, err := h.managerRoles.Search(roless, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *roless, nil, &totalCount))
}

func (h *RestRoles) Post(c echo.Context) error {
	rol := &Roles{}
	if err := c.Bind(rol); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRoles.Post(rol); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *rol, nil, nil))
}

func (h *RestRoles) Patch(c echo.Context) error {
	rol := Roles{}
	if err := c.Bind(rol); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRoles.Patch(rol); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestRoles) Delete(c echo.Context) error {

	roles := &Roles{}
	if err := h.managerRoles.Delete(roles); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
