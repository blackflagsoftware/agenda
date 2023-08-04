package role

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerRoleAdapter interface {
		Get(*Role) error
		Search(*[]Role, RoleParam) (int, error)
		Post(*Role) error
		Patch(Role) error
		Delete(*Role) error
	}

	RestRole struct {
		managerRole ManagerRoleAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerRole(sl)
	hl := NewRestRole(ml)
	hl.LoadRoleRoutes(eg)
}

func NewRestRole(mrol ManagerRoleAdapter) *RestRole {
	return &RestRole{managerRole: mrol}
}

func (h *RestRole) LoadRoleRoutes(eg *echo.Group) {
	eg.GET("/role/:id", h.Get)
	eg.POST("/role/search", h.Search)
	eg.POST("/role", h.Post)
	eg.PATCH("/role", h.Patch)
	eg.DELETE("/role/:id", h.Delete)
}

func (h *RestRole) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		parse := ae.ParamError("id", err)
		return c.JSON(http.StatusBadRequest, util.NewOutput(c, nil, &parse, nil))
	}
	role := &Role{Id: int(id)}
	if err := h.managerRole.Get(role); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *role, nil, nil))
}

func (h *RestRole) Search(c echo.Context) error {
	param := RoleParam{}
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
	roles := &[]Role{}
	totalCount, err := h.managerRole.Search(roles, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *roles, nil, &totalCount))
}

func (h *RestRole) Post(c echo.Context) error {
	rol := &Role{}
	if err := c.Bind(rol); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRole.Post(rol); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *rol, nil, nil))
}

func (h *RestRole) Patch(c echo.Context) error {
	rol := Role{}
	if err := c.Bind(rol); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRole.Patch(rol); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestRole) Delete(c echo.Context) error {

	role := &Role{}
	if err := h.managerRole.Delete(role); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
