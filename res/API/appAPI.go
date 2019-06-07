package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// color variables (in bytecode form)
var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// e.get("/alphabets/remote/", getAllAlpRemote)
func getAllAlpRemote(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"data": "all alphabets with remote routes for image assets",
	})
}

// e.get("/alphabets/local/", getAllAlpLocal)
func getAllAlpLocal(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"data": "all alphabets with local routes for image assets",
	})
}

// e.GET("/alphabet/", getOneAlp)
func getOneAlp(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"data": "the data of alphabet you've requested",
	})
}

// e.GET("/img/for/:id", getAlpImg)
func getAlpImg(c echo.Context) error {
	imgID := c.Param("id")
	imgForLocation := fmt.Sprintf("assets/img/for_img/%s.png", strings.ToLower(imgID))
	return c.File(imgForLocation)
}

// e.GET("/img/write/:id", getAlpImgWrite)
func getAlpImgWrite(c echo.Context) error {
	imgWID := c.Param("id")
	imgWforLocation := fmt.Sprintf("assets/img/writer_img/%s.gif", strings.ToUpper(imgWID))
	return c.File(imgWforLocation)
}

// route to access audio resource of pronouncation (local)
// e.GET("/aud/norm/:id", getAlpAud)
func getAlpAud(c echo.Context) error {
	// return c.JSON(200, map[string]string{
	// 	"data": "audio resource for pronounciation (local data)",
	// })
	audNormID := c.Param("id")
	audNormAddr := fmt.Sprintf("assets/audio/norm_aud/%s.mp3", strings.ToLower(audNormID))
	return c.File(audNormAddr)
}

// route to access audio resource for sample word usage (local)
// e.GET("/aud/for/:id", getAlpAudFor)
func getAlpAudFor(c echo.Context) error {
	// return c.JSON(200, map[string]string{
	// "data": "audio resource for word usage (local data)",
	// })
	audForID := c.Param("id")
	audForAddr := fmt.Sprintf("assets/audio/for_aud/%s.mp3", strings.ToLower(audForID))
	return c.File(audForAddr)
}

func main() {

	// block to confirm if the runtime enviroment is for DEVELOPMENT or PRODUCTION
	var ENVConfig string
	CLIConfig := os.Args
	// if args supplied
	if len(CLIConfig) > 1 {
		switch CLIConfig[1] {
		case "DEV":
			ENVConfig = "DEV"
		case "PROD":
			ENVConfig = "PROD"
		default:
			fmt.Println("Invalid arguments supplied.\nExiting the program.")
			os.Exit(0)
		}
	} else {
		// if no args supplied
		ENVConfig = "DEV"
	}

	e := echo.New()

	// debug mode
	e.Debug = true // optional
	// just to hide the echo framework commercial
	e.HideBanner = true
	// name definition for the runtime application (along with the runtime enviroment variant)
	name := fmt.Sprintf("R&D-%s", ENVConfig)
	// Adding trailing slash to request URI
	// e.Pre(middleware.AddTrailingSlash())

	// tailored logger adapting to the different runtime environment
	switch ENVConfig {
	case "DEV":
		// Debug version of LOG
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			var reqMethod string
			var resStatus int
			var statusColor, methodColor, resetColor string
			// request and response object
			req := c.Request()
			res := c.Response()
			// rendering variables for response status and request method
			resStatus = res.Status
			reqMethod = req.Method
			// for response status
			switch {
			case resStatus >= http.StatusOK && resStatus < http.StatusMultipleChoices:
				statusColor = green
			case resStatus >= http.StatusMultipleChoices && resStatus < http.StatusBadRequest:
				statusColor = white
			case resStatus >= http.StatusBadRequest && resStatus < http.StatusInternalServerError:
				statusColor = yellow
			default:
				statusColor = red
			}
			// for request method
			switch reqMethod {
			case "GET":
				methodColor = blue
			case "POST":
				methodColor = cyan
			case "PUT":
				methodColor = yellow
			case "DELETE":
				methodColor = red
			case "PATCH":
				methodColor = green
			case "HEAD":
				methodColor = magenta
			case "OPTIONS":
				methodColor = white
			default:
				methodColor = reset
			}
			// reset to return to the normal terminal color variables (kinda default)
			resetColor = reset
			// print formatting the custom logger tailored for DEVELOPMENT environment
			fmt.Printf("\n[%s] %v |%s %3d %s| %8s | %10s |%s %-7s %s %s",
				name, // name of server (APP) with the environment
				time.Now().Format("2006/01/02 - 15:04:05"), // TIMESTAMP for route access
				statusColor, resStatus, resetColor, // response status
				req.Proto,                          // protocol
				c.RealIP(),                         // client IP
				methodColor, reqMethod, resetColor, // request method
				req.URL, // request URI (path)
			)
		}))
	default:
		// Production version of LOG
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${bytes_in} bytes_in | ${bytes_out} bytes_out | ${method} | ${uri} ",
				name,
			),
			CustomTimeFormat: "2006/01/02 15:04:05", // custom readable time format
			Output:           os.Stdout,             // output method
		}))
	}

	// route for API info
	e.GET("/", func(c echo.Context) (err error) {
		reqM := c.Request()
		resM := c.Response()
		return c.JSON(200, map[string]string{
			"name":        "Children_app_api",
			"developer":   "Yuil Tripathee",
			"version":     "v1.0",
			"status_code": fmt.Sprintf("%d", resM.Status),
			"time":        time.Now().Format("2006/01/01 - 15:04:05"),
			"protocol":    reqM.Proto,
			"ip":          c.RealIP(),
			"method":      reqM.Method,
			"url":         fmt.Sprintf("%s", reqM.URL),
			"bytes_out":   fmt.Sprintf("%d", resM.Size),
		})
	})

	// route to share entire data (remote resources)
	e.GET("/alphabets/remote", getAllAlpRemote)

	// route to share entire data (local resources)
	e.GET("/alphabets/local", getAllAlpLocal)

	// route to share individual data
	e.GET("/alphabet", getOneAlp)

	// route to access image resources (local resources)
	e.GET("/img/for/:id", getAlpImg)

	// route to access image resources for writing (local resources)
	e.GET("/img/write/:id", getAlpImgWrite)

	// route to access audio resource of pronouncation (local)
	e.GET("/aud/norm/:id", getAlpAud)

	// route to access audio resource for sample word usage (local)
	e.GET("/aud/for/:id", getAlpAudFor)

	// test
	// e.GET("/monkey", func(c echo.Context) error {
	// 	name := c.QueryParam("name")
	// 	return c.String(http.StatusOK, name)
	// })

	// e.GET("/users/:name", func(c echo.Context) error {
	// 	name := c.Param("name")
	// 	return c.String(http.StatusOK, name)
	// })

	// static route for dummy landing page
	e.Static("/", "static")

	data, err := json.MarshalIndent(e.Routes(), "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("routes.json", data, 0644)
	// firing up the server
	e.Logger.Fatal(e.Start(":8000"))

}
