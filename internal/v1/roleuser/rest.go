package roleuser

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v3"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerRoleUserAdapter interface {
		Login(RoleUser) (RoleLogin, error)
		Get(*RoleUser) error
		Search(*[]RoleUser, RoleUserParam) (int, error)
		Post(*RoleUser) error
		Patch(RoleUser) error
		Delete(*RoleUser) error
	}

	RestRoleUser struct {
		managerRoleUser ManagerRoleUserAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerRoleUser(sl)
	hl := NewRestRoleUser(ml)
	hl.LoadRoleUserRoutes(eg)
}

func NewRestRoleUser(mro ManagerRoleUserAdapter) *RestRoleUser {
	return &RestRoleUser{managerRoleUser: mro}
}

func (h *RestRoleUser) LoadRoleUserRoutes(eg *echo.Group) {
	eg.GET("/roleuser/login/:user/pwd/:pwd", h.Login)
	eg.GET("/roleuser/:id", h.Get)
	eg.GET("/roleuser", h.Search)
	eg.POST("/roleuser", h.Post)
	eg.PATCH("/roleuser", h.Patch)
	eg.DELETE("/roleuser/:id", h.Delete)
}

func (h *RestRoleUser) Login(c echo.Context) error {
	user := c.Param("user")
	pwd := c.Param("pwd")
	roleUser := RoleUser{Name: null.StringFrom(user), Pwd: null.StringFrom(pwd)}
	roleLogin, err := h.managerRoleUser.Login(roleUser)
	if err != nil {
		apiError := err.(ae.ApiError)
		fmt.Println(apiError.Error())
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, roleLogin, nil, nil))
}
func (h *RestRoleUser) Get(c echo.Context) error {

	roleUser := &RoleUser{}
	if err := h.managerRoleUser.Get(roleUser); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *roleUser, nil, nil))
}

func (h *RestRoleUser) Search(c echo.Context) error {
	param := RoleUserParam{}
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
	roleUsers := &[]RoleUser{}
	totalCount, err := h.managerRoleUser.Search(roleUsers, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *roleUsers, nil, &totalCount))
}

func (h *RestRoleUser) Post(c echo.Context) error {
	ro := &RoleUser{}
	if err := c.Bind(ro); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRoleUser.Post(ro); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *ro, nil, nil))
}

func (h *RestRoleUser) Patch(c echo.Context) error {
	ro := RoleUser{}
	if err := c.Bind(ro); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerRoleUser.Patch(ro); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestRoleUser) Delete(c echo.Context) error {

	roleUser := &RoleUser{}
	if err := h.managerRoleUser.Delete(roleUser); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
