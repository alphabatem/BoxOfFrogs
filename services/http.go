package services

import (
	"errors"
	"fmt"
	"github.com/cloakd/common/context"
	"github.com/cloakd/common/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

type HttpService struct {
	services.DefaultService
	BaseURL string
	Port    int
}

var ErrUnauthorized = errors.New("unauthorized")
var DeleteResponseOK = `{"status": 200, "error": ""}`

func (svc HttpService) Id() string {
	return "http"
}

func (svc *HttpService) Configure(ctx *context.Context) error {
	port := os.Getenv("HTTP_PORT")
	portFlag, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	svc.Port = portFlag

	return svc.DefaultService.Configure(ctx)
}

func (svc *HttpService) Start() error {
	//svc.imgSvc = svc.Service(IMAGE_SVC).(*ImageService)

	r := gin.Default()

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	//r.Static("static", "static")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Validation endpoints
	r.GET("/ping", svc.ping)

	v1 := r.Group("/v1")
	//docs.SwaggerInfo.BasePath = "/v1"

	//v1.POST("sites", svc.createSite)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r.Run(fmt.Sprintf(":%v", svc.Port))
}

type Pong struct {
	Message string `json:"message"`
}

//
// @Summary Ping services
// @Accept  json
// @Produce json
// @Router /ping [get]
func (svc *HttpService) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//HELPERS Below

func (svc *HttpService) paramErr(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}

func (svc *HttpService) httpError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
}
