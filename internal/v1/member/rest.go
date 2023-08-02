package member

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerMemberAdapter interface {
		Get(*Member) error
		Search(*[]Member, MemberParam) (int, error)
		Post(*Member) error
		Patch(Member) error
		Delete(*Member) error
		Splice() error
	}

	RestMember struct {
		managerMember ManagerMemberAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerMember(sl)
	hl := NewRestMember(ml)
	hl.LoadMemberRoutes(eg)
}

func NewRestMember(mmem ManagerMemberAdapter) *RestMember {
	return &RestMember{managerMember: mmem}
}

func (h *RestMember) LoadMemberRoutes(eg *echo.Group) {
	eg.GET("/member/:id", h.Get)
	eg.GET("/member", h.Search)
	eg.POST("/member", h.Post)
	eg.PATCH("/member", h.Patch)
	eg.DELETE("/member/:id", h.Delete)
	eg.PUT("/member", h.Splice)
}

func (h *RestMember) Get(c echo.Context) error {

	member := &Member{}
	if err := h.managerMember.Get(member); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *member, nil, nil))
}

func (h *RestMember) Search(c echo.Context) error {
	param := MemberParam{}
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
	members := &[]Member{}
	totalCount, err := h.managerMember.Search(members, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *members, nil, &totalCount))
}

func (h *RestMember) Post(c echo.Context) error {
	mem := &Member{}
	if err := c.Bind(mem); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerMember.Post(mem); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *mem, nil, nil))
}

func (h *RestMember) Patch(c echo.Context) error {
	mem := Member{}
	if err := c.Bind(&mem); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerMember.Patch(mem); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestMember) Delete(c echo.Context) error {

	member := &Member{}
	if err := h.managerMember.Delete(member); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestMember) Splice(c echo.Context) error {
	err := h.managerMember.Splice()
	return c.String(http.StatusOK, err.Error())
}
