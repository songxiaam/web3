package main

import (
	"context"
	"fmt"
	"gin-study/routers"
	"gin-study/test"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.ForceConsoleColor()
	//fmt.Println("start....")
	//database.InitDB()
	//
	//routers := routers.SetupRouter()
	//routers.Run(":8080")
	//fmt.Println("服务器启动成功：http://localhost:8080/user/login")

	//testBindCustomStruct()
	//testBindSrarchForm()
	//testBindUri()
	//testCookie()
	//testLog()
	//testMiddleWare()
	//testValidator()
	//testRouterLogFormat()
	//testGoroutine()
	//testRestart()
	//testPush()
	//testRouterParam()
	//testPureJSON()
	//testRedirect()
	//testRender()
	//testMultiServer()
	//testSecureJSON()
	//testHeader()
	//testReader()
	//testLetsEncrypt()
	//testUpload()
	//testSecret()
	testLogFile()
}

func testBindCustomStruct() {
	router := gin.Default()
	router.GET("/getb", test.GetDataB)
	router.GET("/getc", test.GetDataC)
	router.GET("/getd", test.GetDataD)

	router.Run()
}

func testBindSrarchForm() {
	router := gin.Default()
	router.GET("/testing", test.StartPage)
	router.Run(":8085")
}

func testBindUri() {
	router := gin.Default()
	router.GET("/:name/:id", test.BindUri)
	router.Run(":8088")
}

func testCookie() {
	router := gin.Default()
	router.GET("/cookie", test.TestCookie)
	router.Run(":8080")
}

func testLog() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		fmt.Println("log.....\n")
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
	}))
	router.Use(gin.Recovery())
	router.GET("/testing", test.StartPage)
	router.Run(":8080")
}

func testMiddleWare() {
	router := gin.New()
	router.Use(test.Logger())
	router.GET("/testing", test.StartPage)
	router.Run(":8080")
}

func testValidator() {
	router := gin.Default()
	router.Use(test.Logger())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("bookabledate", test.BookableDate)
		if err != nil {
			fmt.Println(err)
		}
	}
	router.GET("/valid", test.GetBookable)
	router.Run(":8080")
}

func testRouterLogFormat() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})
	router.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	router.Run(":8080")
}

func testGoroutine() {
	router := gin.Default()
	router.GET("/long_async", test.LongAsync)
	router.GET("/long_sync", test.LongSync)
	router.Run(":8080")
}

func testRestart() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(7 * time.Second)
		c.String(http.StatusOK, "hello world")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func testPush() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(test.PushHtml)
	router.GET("/", test.TestPush)
	router.RunTLS(":8080", "./cert/cert.pem", "./cert/key.pem")
}

func testRouterParam() {
	router := gin.Default()
	router.GET("/user/:name", test.TestUserName)
	router.GET("/user/:name/*action", test.TestUserNameAction)
	router.Run(":8080")
}

func testPureJSON() {
	router := gin.Default()

	// 提供 unicode 实体
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

func testRedirect() {
	router := gin.Default()
	router.GET("/test", test.TestRedirect1)
	router.POST("/test2", test.TestRedirect2)
	router.GET("/test3", func(c *gin.Context) {
		test.TestRedirect3(c, router)
	})
	router.GET("/test4", test.TestRedirect4)
	router.Run(":8080")
}

func testRender() {
	router := gin.Default()
	router.GET("/some_json", test.RenderJSON)
	router.GET("/some_more_json", test.RenderMoreJSON)
	router.GET("/some_xml", test.RenderXML)
	router.GET("/some_yaml", test.RenderYAML)
	router.GET("/some_proto_buf", test.RenderProtoBuf)
	router.Run(":8080")
}

var g errgroup.Group

func testMultiServer() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      test.Router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      test.Router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func testSecureJSON() {
	router := gin.Default()
	router.GET("/names", test.TestSecureJSON)
	router.Run(":8080")
}

func testHeader() {
	router := gin.Default()
	router.Use(test.UseHeader)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}

func testReader() {
	router := gin.Default()
	router.GET("/from_reader", test.TestReader)
	router.Run(":8080")

}

func testLetsEncrypt() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	err := router.RunTLS(":8443", "/Users/mikezhao/certs/server.crt", "/Users/mikezhao/certs/server.key")
	if err != nil {
		panic(err)
	}
}

func testUpload() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8的二进制左移20位 = 8*2^20 ≈ 8M
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./files/" + file.Filename
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./files/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}

func testSecret() {
	router := gin.Default()
	test.GetSecretUser(router)
	router.Run(":8080")
}

func testLogFile() {
	gin.DisableConsoleColor()

	if err := os.MkdirAll("./logs", 0755); err != nil {
		panic(err)
	}

	f, _ := os.Create("./logs/test.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8080")
}

func testTest() {
	router := routers.SetupRouter()
	router.Run(":8080")
}
