package main

import (
	"fmt"
	"go-work/go-frame/gee/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {	// 我们可以看到，recover 捕获了 panic，程序正常结束。test_recover() 中的 after panic 没有打印，这是正确的，当 panic 被触发时，控制权就被交给了 defer 。就像在 java 中，try代码块中发生了异常，控制权交给了 catch，接下来执行 catch 代码块中的代码。而在 main() 中打印了 after recover，说明程序已经恢复正常，继续往下执行直到结束。
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}

func main() {
	test_recover()
	fmt.Println("after recover")

	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	/*r := gee.New()
	r.Static("/static", "./static/css")

	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")

	stu1 := &student{Name:"zhl", Age:22}
	stu2 := &student{Name:"jack", Age:20}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK,"arr.tmpl", gee.H{
			"title":"gee",
			"stuArr":[2]*student{stu1,stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title":"gee",
			"now":time.Date(2020,12,2,0,0,0,0, time.UTC),
		})
	})*/
	/*v1 := r.Group("v1")
	{
		v1.GET("/index", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Index Page</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("v2")
	v2.Use(onlyForV2())
	{

		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})*/

	r.Run(":8888")
}
