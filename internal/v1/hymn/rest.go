package hymn

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerHymnAdapter interface {
		Get(*Hymn) error
		Search(*[]Hymn, HymnParam) (int, error)
		Post(*Hymn) error
		Patch(Hymn) error
		Delete(*Hymn) error
	}

	RestHymn struct {
		managerHymn ManagerHymnAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerHymn(sl)
	hl := NewRestHymn(ml)
	hl.LoadHymnRoutes(eg)
}

func NewRestHymn(mhym ManagerHymnAdapter) *RestHymn {
	return &RestHymn{managerHymn: mhym}
}

func (h *RestHymn) LoadHymnRoutes(eg *echo.Group) {
	eg.GET("/hymn/:id", h.Get)
	eg.GET("/hymn", h.Search)
	eg.POST("/hymn", h.Post)
	eg.PATCH("/hymn", h.Patch)
	eg.DELETE("/hymn/:id", h.Delete)
}

func (h *RestHymn) Get(c echo.Context) error {

	hymn := &Hymn{}
	if err := h.managerHymn.Get(hymn); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *hymn, nil, nil))
}

func (h *RestHymn) Search(c echo.Context) error {
	param := HymnParam{}
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
	hymns := &[]Hymn{}
	totalCount, err := h.managerHymn.Search(hymns, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *hymns, nil, &totalCount))
}

func (h *RestHymn) Post(c echo.Context) error {
	hym := &Hymn{}
	if err := c.Bind(hym); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerHymn.Post(hym); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *hym, nil, nil))
}

func (h *RestHymn) Patch(c echo.Context) error {
	hym := Hymn{}
	if err := c.Bind(hym); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerHymn.Patch(hym); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestHymn) Delete(c echo.Context) error {

	hymn := &Hymn{}
	if err := h.managerHymn.Delete(hymn); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
