package newmember

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerNewMemberAdapter interface {
		Get(*NewMember) error
		Search(*[]NewMember, NewMemberParam) (int, error)
		Post(*NewMember) error
		Patch(NewMember) error
		Delete(*NewMember) error
	}

	RestNewMember struct {
		managerNewMember ManagerNewMemberAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerNewMember(sl)
	hl := NewRestNewMember(ml)
	hl.LoadNewMemberRoutes(eg)
}

func NewRestNewMember(mnew ManagerNewMemberAdapter) *RestNewMember {
	return &RestNewMember{managerNewMember: mnew}
}

func (h *RestNewMember) LoadNewMemberRoutes(eg *echo.Group) {
	eg.GET("/newmember/:id", h.Get)
	eg.POST("/newmember/search", h.Search)
	eg.POST("/newmember", h.Post)
	eg.PATCH("/newmember", h.Patch)
	eg.DELETE("/newmember/:id", h.Delete)
}

func (h *RestNewMember) Get(c echo.Context) error {

	newMember := &NewMember{}
	if err := h.managerNewMember.Get(newMember); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *newMember, nil, nil))
}

func (h *RestNewMember) Search(c echo.Context) error {
	param := NewMemberParam{}
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
	newMembers := &[]NewMember{}
	totalCount, err := h.managerNewMember.Search(newMembers, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *newMembers, nil, &totalCount))
}

func (h *RestNewMember) Post(c echo.Context) error {
	new := &NewMember{}
	if err := c.Bind(new); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerNewMember.Post(new); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *new, nil, nil))
}

func (h *RestNewMember) Patch(c echo.Context) error {
	new := NewMember{}
	if err := c.Bind(&new); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerNewMember.Patch(new); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestNewMember) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	newMember := &NewMember{Id: int(id)}
	if err := h.managerNewMember.Delete(newMember); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
