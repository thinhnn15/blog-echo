package main
import (
	"github.com/labstack/echo"
	"github.com/thinhnn15/blog-echo/db"
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

}
