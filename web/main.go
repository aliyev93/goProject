package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

func main() {

	f, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err2 := decoder.Decode(&cfg)

	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", cfg)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	type Employee struct {
		Id   string `json:"id"`
		Name string `json:"employee_name"`
		Age  string `json:"employee_age"`
	}
	type People struct {
		People []Employee `json:"employee"`
	}

	db, err := sql.Open("mysql", cfg.Database.User+":"+cfg.Database.Password+"@tcp("+cfg.Database.Host+":"+cfg.Database.Port+")/"+cfg.Database.Dbname)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	e.GET("/health", health)

	e.POST("/employee", func(c echo.Context) error {
		emp := new(Employee)
		if err := c.Bind(emp); err != nil {
			return err
		}
		//
		sql := "INSERT INTO employee(employee_name, employee_age) VALUES( ?, ?)"
		stmt, err := db.Prepare(sql)

		if err != nil {
			fmt.Print(err.Error())
		}
		defer stmt.Close()
		result, err2 := stmt.Exec(emp.Name, emp.Age)

		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(result.LastInsertId())

		return c.JSON(http.StatusCreated, emp.Name)
	})

	e.DELETE("/employee/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		sql := "Delete FROM employee Where id = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			fmt.Println(err)
		}
		result, err2 := stmt.Exec(requested_id)
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(result.RowsAffected())
		return c.JSON(http.StatusOK, "Deleted")
	})

	e.GET("/employee/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id)
		var name string
		var id string
		var age string

		err = db.QueryRow("SELECT id,employee_name, employee_age  FROM employee WHERE id = ?", requested_id).Scan(&id, &name, &age)

		if err != nil {
			fmt.Println(err)
		}

		response := Employee{Id: id, Name: name, Age: age}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
