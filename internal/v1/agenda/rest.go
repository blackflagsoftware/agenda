package agenda

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ae "github.com/blackflagsoftware/agenda/internal/api_error"
	util "github.com/blackflagsoftware/agenda/internal/util"
)

type (
	ManagerAgendaAdapter interface {
		Get(*Agenda) error
		Search(*[]Agenda, AgendaParam) (int, error)
		Post(*Agenda) error
		Patch(Agenda) error
		Delete(*Agenda) error
		Print(string) error
		Publish(string) error
	}

	RestAgenda struct {
		managerAgenda ManagerAgendaAdapter
	}
)

func InitializeRest(eg *echo.Group) {
	sl := InitStorage()
	ml := NewManagerAgenda(sl)
	hl := NewRestAgenda(ml)
	hl.LoadAgendaRoutes(eg)
}

func NewRestAgenda(mage ManagerAgendaAdapter) *RestAgenda {
	return &RestAgenda{managerAgenda: mage}
}

func (h *RestAgenda) LoadAgendaRoutes(eg *echo.Group) {
	eg.GET("/agenda/:date", h.Get)
	eg.GET("/agenda", h.Search)
	eg.POST("/agenda", h.Post)
	eg.PATCH("/agenda", h.Patch)
	eg.DELETE("/agenda/:date", h.Delete)
	eg.GET("/agenda/print/:date", h.Print)
	eg.GET("/agenda/publish/:date", h.Publish)
}

func (h *RestAgenda) Get(c echo.Context) error {
	date := c.Param("date")
	agenda := &Agenda{Date: date}
	if err := h.managerAgenda.Get(agenda); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *agenda, nil, nil))
}

func (h *RestAgenda) Search(c echo.Context) error {
	param := AgendaParam{}
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
	agendas := &[]Agenda{}
	totalCount, err := h.managerAgenda.Search(agendas, param)
	if err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *agendas, nil, &totalCount))
}

func (h *RestAgenda) Post(c echo.Context) error {
	age := &Agenda{}
	if err := c.Bind(age); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	if err := h.managerAgenda.Post(age); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.JSON(http.StatusOK, util.NewOutput(c, *age, nil, nil))
}

func (h *RestAgenda) Patch(c echo.Context) error {
	age := Agenda{}
	if err := c.Bind(&age); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, util.NewOutput(c, bindErr.BodyError(), &bindErr, nil))
	}
	fmt.Printf("**** Patch **** %+v\n", age)
	if err := h.managerAgenda.Patch(age); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestAgenda) Delete(c echo.Context) error {
	date := c.Param("date")
	agenda := &Agenda{Date: date}
	if err := h.managerAgenda.Delete(agenda); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestAgenda) Print(c echo.Context) error {
	date := c.Param("date")
	if err := h.managerAgenda.Print(date); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}

func (h *RestAgenda) Publish(c echo.Context) error {
	date := c.Param("date")
	if err := h.managerAgenda.Publish(date); err != nil {
		apiError := err.(ae.ApiError)
		be := apiError.BodyError()
		return c.JSON(be.StatusCode, util.NewOutput(c, nil, &apiError, nil))
	}
	return c.NoContent(http.StatusOK)
}
