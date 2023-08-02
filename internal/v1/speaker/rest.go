package speaker

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerSpeakerAdapter interface {
		Get(*Speaker) error
		Search(*[]Speaker, SpeakerParam) (int, error)
		Post(*Speaker) error
		Patch(Speaker) error
		Delete(*Speaker) error
	}

	RestSpeaker struct {
		managerSpeaker ManagerSpeakerAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerSpeaker(sl)
	hl := NewRestSpeaker(ml)
	hl.LoadSpeakerRoutes(eg)
}

func NewRestSpeaker(mspe ManagerSpeakerAdapter) *RestSpeaker {
	return &RestSpeaker{managerSpeaker: mspe}
}

func (h *RestSpeaker) LoadSpeakerRoutes(eg *echo.Group) {
	eg.GET("/speaker/:id", h.Get)
	eg.POST("/speaker/search", h.Search)
	eg.POST("/speaker", h.Post)
	eg.PATCH("/speaker", h.Patch)
	eg.DELETE("/speaker/:id", h.Delete)
}

func (h *RestSpeaker) Get(c echo.Context) error {

	speaker := &Speaker{}
	if err := h.managerSpeaker.Get(speaker); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *speaker, nil, nil))
}

func (h *RestSpeaker) Search(c echo.Context) error {
	param := SpeakerParam{}
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
	speakers := &[]Speaker{}
	totalCount, err := h.managerSpeaker.Search(speakers, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *speakers, nil, &totalCount))
}

func (h *RestSpeaker) Post(c echo.Context) error {
	spe := &Speaker{}
	if err := c.Bind(spe); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerSpeaker.Post(spe); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *spe, nil, nil))
}

func (h *RestSpeaker) Patch(c echo.Context) error {
	spe := Speaker{}
	if err := c.Bind(&spe); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerSpeaker.Patch(spe); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestSpeaker) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	speaker := &Speaker{Id: int(id)}
	if err := h.managerSpeaker.Delete(speaker); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
