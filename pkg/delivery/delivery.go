package delivery

import (
	odontologoUCI "github.com/0xiroh/go-crud-api/pkg/interface/odontologoUCInterface"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// HTTPOdontologo - struct for http odontologo book
type HTTPOdontologo struct {
	odontologoUCI odontologoUCI.OdontologoUCI
}

type requestOdontologo struct {
	Nombre  string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula  string `json:"matricula"`
}

type reqUpdate struct {
	requestOdontologo
	ID int64 `json:"id"`
}

// GetAll - handler for find all data
func (handler *HTTPOdontologo) GetAll(c *gin.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	list, err := handler.odontologoUCI.FindAll()
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	if len(list) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusNoContent,
				"message": "success",
				"data":    &[]string{},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    &list,
		},
	)
}

// AddData - handler for insert data odontologo_books
func (handler *HTTPOdontologo) AddData(c *gin.Context) {
	var reqJSON requestOdontologo
	c.BindJSON(&reqJSON)

	if reqJSON.Nombre == "" || reqJSON.Apellido == "" || reqJSON.Matricula == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.odontologoUCI.CreateData(reqJSON.Nombre, reqJSON.Apellido, reqJSON.Matricula)
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)

}

// GetByID - handler for find by id data
func (handler *HTTPOdontologo) GetByID(c *gin.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	myID := c.Param("id")
	id, _ := strconv.ParseInt(myID, 10, 64)
	if id == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	data, err := handler.odontologoUCI.FindByID(id)
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    &data,
		},
	)
}

// EditData - handler for update data odontologo_books
func (handler *HTTPOdontologo) EditData(c *gin.Context) {
	var reqJSON reqUpdate
	c.BindJSON(&reqJSON)

	if reqJSON.Nombre == "" || reqJSON.ID == 0 || reqJSON.Apellido == "" || reqJSON.Matricula == ""  {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.odontologoUCI.UpdateData(reqJSON.ID, reqJSON.Nombre, reqJSON.Apellido, reqJSON.Matricula)
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)

}

// DeleteData - handler for delete data odontologo_books
func (handler *HTTPOdontologo) DeleteData(c *gin.Context) {
	type req struct {
		ID int64 `json:"id"`
	}

	var Req req

	c.BindJSON(&Req)

	if Req.ID == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": "Bad Request",
			},
		)
		return
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := handler.odontologoUCI.DeleteData(Req.ID)
	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  http.StatusBadGateway,
				"message": "System Error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "success",
		},
	)
}

// NewOdontologoHTTPHandler - initial http handler odontologo book
func NewOdontologoHTTPHandler(r *gin.Engine, odontologoUCI odontologoUCI.OdontologoUCI) {
	handler := &HTTPOdontologo{odontologoUCI}

	api := r.Group("/odontologo")
	{
		api.GET("/list", handler.GetAll)
		api.GET("/data/:id", handler.GetByID)
		api.POST("/add", handler.AddData)
		api.POST("/edit", handler.EditData)
		api.POST("/delete", handler.DeleteData)
	}
}