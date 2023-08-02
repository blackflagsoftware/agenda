package announcement

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerAnnouncementAdapter interface {
		Get(*Announcement) error
		Search(*[]Announcement, AnnouncementParam) (int, error)
		Post(*Announcement) error
		Patch(Announcement) error
		Delete(*Announcement) error
	}

	RestAnnouncement struct {
		managerAnnouncement ManagerAnnouncementAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerAnnouncement(sl)
	hl := NewRestAnnouncement(ml)
	hl.LoadAnnouncementRoutes(eg)
}

func NewRestAnnouncement(mann ManagerAnnouncementAdapter) *RestAnnouncement {
	return &RestAnnouncement{managerAnnouncement: mann}
}

func (h *RestAnnouncement) LoadAnnouncementRoutes(eg *echo.Group) {
	eg.GET("/announcement/:id", h.Get)
	eg.POST("/announcement/search", h.Search)
	eg.POST("/announcement", h.Post)
	eg.PATCH("/announcement", h.Patch)
	eg.DELETE("/announcement/:id", h.Delete)
}

func (h *RestAnnouncement) Get(c echo.Context) error {

	announcement := &Announcement{}
	if err := h.managerAnnouncement.Get(announcement); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *announcement, nil, nil))
}

func (h *RestAnnouncement) Search(c echo.Context) error {
	param := AnnouncementParam{}
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
	announcements := &[]Announcement{}
	totalCount, err := h.managerAnnouncement.Search(announcements, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *announcements, nil, &totalCount))
}

func (h *RestAnnouncement) Post(c echo.Context) error {
	ann := &Announcement{}
	if err := c.Bind(ann); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerAnnouncement.Post(ann); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *ann, nil, nil))
}

func (h *RestAnnouncement) Patch(c echo.Context) error {
	ann := Announcement{}
	if err := c.Bind(&ann); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerAnnouncement.Patch(ann); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestAnnouncement) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, nil, &bindErr, nil))
	}
	announcement := &Announcement{Id: int(id)}
	if err := h.managerAnnouncement.Delete(announcement); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
