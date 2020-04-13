package handlers

import (
	"net/http"
	"strconv"

	"vscale-task/cmd/providers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	client providers.Client
}

func NewServer(port string, client providers.Client) *Server {
	return &Server{
		port:   port,
		client: client,
	}
}

func (s *Server) Run() error {
	r := gin.Default()

	r.POST("/:number", s.CreateServer)
	r.DELETE("/:group")

	return r.Run(":" + s.port)
}

func (s *Server) CreateServer(c *gin.Context) {
	var param = c.Param("number")

	number, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}

	var (
		servReq  = providers.DefaultCreateServerRequest()
		servResp providers.CreateServerResponse
	)

	for i := 0; i < number; i++ {
		go func() {
			s.client.CreateServer(servReq)
		}()
	}
}
