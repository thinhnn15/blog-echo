package main
import (
	"github.com/labstack/echo"
	"github.com/thinhnn15/blog-echo/db"
	"github.com/thinhnn15/blog-echo/handler"
	"github.com/thinhnn15/blog-echo/repository/repo_impl"
	"github.com/thinhnn15/blog-echo/router"
)
func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "thinhnn",
		Password: "secret",
		DbName:   "blog",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	postHandler := handler.PostHandler{
		PostRepo: repo_impl.NewPostRepo(sql),
	}
	
	api := router.API{
		Echo:        e,
		PostHandler: postHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))

}
