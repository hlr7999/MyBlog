package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"HBlog/controller/auth"
	"HBlog/controller/user"
	"HBlog/data"
)

type Server struct {
	Addr string
	e    *echo.Echo
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
		e:    echo.New(),
	}
}

func (s *Server) Init() (err error) {

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	err = auth.Initialize(s.e)
	if err != nil {
		return err
	}

	g := s.e.Group("")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &data.Claims{},
		SigningKey: []byte(data.Secret),
	}))
	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u := new(data.User)
			t := c.Get("user").(*jwt.Token)
			claims := t.Claims.(*data.Claims)
			collection, closeConn := db.GlobalDB.User()
			defer closeConn()
			err = collection.FindId(claims.UID).One(u)
			if err != nil {
				if err == mgo.ErrNotFound {
					return echo.ErrUnauthorized
				}
				return err
			}
			c.Set(data.ContextUser, u)
			return next(c)
		}
	})

	err = user.Initialize(g)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(s.Addr))
}