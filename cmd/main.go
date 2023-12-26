package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL string  `yaml:"base_url"`
	Port    string  `yaml:"port"`
	Routes  []Route `yaml:"routes"`
}

type Route struct {
	Path     string   `yaml:"path"`
	Method   string   `yaml:"method"`
	Response Response `yaml:"response"`
}

type Response struct {
	Code int    `yaml:"code"`
	Body string `yaml:"body"`
}

func setupRouter(config Config) *gin.Engine {
	router := gin.Default()

	for _, rt := range config.Routes {
		// We assume the method is always uppercase (GET, POST, etc.)
		// and the provided methods are supported by Gin framework.
		switch rt.Method {
		case "GET":
			router.GET(config.BaseURL+rt.Path, func(c *gin.Context) {
				c.String(rt.Response.Code, rt.Response.Body)
			})
		case "POST":
			router.POST(config.BaseURL+rt.Path, func(c *gin.Context) {
				c.String(rt.Response.Code, rt.Response.Body)
			})
		// You can add other HTTP methods as needed.
		default:
			fmt.Printf("Unsupported method: %s for route: %s\n", rt.Method, rt.Path)
		}
	}

	return router
}

func readConfig(filename string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := readConfig("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to read config: %v", err))
	}

	router := setupRouter(*config)
	router.Run(":" + config.Port)
}
