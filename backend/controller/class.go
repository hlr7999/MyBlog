package controller

import (
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"

	"MyBlog/model"
	"MyBlog/app"
)

func InitClassNotAuth(e *echo.Echo) {
	e.GET("/classes", getAllClasses)
	e.GET("/classes/first", getFirstClasses)
	e.GET("/classes/second/:id", getSecondClasses)
}

func InitClassAuth(g *echo.Group) {
	g.POST("/classes/first", newFirstClass)
	g.POST("/classes/second", newSecondClass)
	g.DELETE("/classes/:id", deleteClass)
}

type classTree struct {
	ID    string               `bson:"_id" json:"_id"`
	Name  string               `bson:"name" json:"name"`
	Child []model.ClassPartial `bson:"child" json:"child"`
}

func getAllClasses(c echo.Context) error {
	collection := app.DB().C(model.ClassC)

	var classes []model.Class
	err := collection.Find(bson.M{"level": model.FirstLevel}).All(&classes)
	if err != nil {
		return app.ServerError(c, err)
	}

	allClass := make([]classTree, len(classes))
	var childClass model.Class
	for i, item := range classes {
		allClass[i].ID = item.ID.Hex()
		allClass[i].Name = item.Name
		allClass[i].Child = make([]model.ClassPartial, len(item.Child))
		for j, sItem := range item.Child {
			err = collection.Find(bson.M{"_id": bson.ObjectIdHex(sItem)}).One(&childClass)
			if err != nil {
				return app.ServerError(c, err)
			}
			allClass[i].Child[j] = childClass.ToPartial()
		}
	}

	return app.Ok(c, allClass)
}

func getFirstClasses(c echo.Context) error {
	collection := app.DB().C(model.ClassC)

	var classes []model.Class
	err := collection.Find(bson.M{"level": model.FirstLevel}).All(&classes)
	if err != nil {
		return app.ServerError(c, err)
	}
	
	return app.Ok(c, classes)
}

func getSecondClasses(c echo.Context) error {
	collection := app.DB().C(model.ClassC)

	id := c.Param("id")
	var father model.Class
	err := collection.FindId(bson.ObjectIdHex(id)).One(&father)
	if err != nil {
		return app.ServerError(c, err)
	}

	classes := make([]model.Class, len(father.Child))
	for i, item := range father.Child {
		err := collection.FindId(bson.ObjectIdHex(item)).One(&classes[i])
		if err != nil {
			return app.ServerError(c, err)
		}
	}

	return app.Ok(c, classes)
}

type newClass struct {
	Name string `bson:"name" json:"name"`
}

func newFirstClass(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	collection := app.DB().C(model.ClassC)

	nClass := new(newClass)
	err := c.Bind(nClass)
	if err != nil {
		return app.BadRequest(c, "Bad Request")
	}
	
	n, err := collection.Find(bson.M{"name": nClass.Name}).Count()
	if err != nil {
		return app.ServerError(c, err)
	}
	if n != 0 {
		return app.BadRequest(c, "1")
	}

	class := new(model.Class)
	class.ID = bson.NewObjectId()
	class.Level = model.FirstLevel
	class.Name = nClass.Name
	err = collection.Insert(class)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, nClass)
}

type newChildClass struct {
	Name string `bson:"name" json:"name"`
	Father string `bson:"father" json:"father"`
}

func newSecondClass(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	collection := app.DB().C(model.ClassC)

	nClass := new(newChildClass)
	err := c.Bind(nClass)
	if err != nil {
		return app.BadRequest(c, "Bad Request")
	}
	
	n, err := collection.Find(bson.M{"name": nClass.Name}).Count()
	if err != nil {
		return app.ServerError(c, err)
	}
	if n != 0 {
		return app.BadRequest(c, "1")
	}

	class := new(model.Class)
	class.ID = bson.NewObjectId()
	class.Level = model.SecondLevel
	class.Name = nClass.Name
	class.Child = []string{nClass.Father}
	err = collection.Insert(class)
	if err != nil {
		return app.ServerError(c, err)
	}

	err = collection.Update(
		bson.M{"_id": bson.ObjectIdHex(nClass.Father)},
		bson.M{"$addToSet": bson.M{"child": class.ID.Hex()}},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, nClass)
}

func deleteClass(c echo.Context) error {
	token := app.GetToken(c)
	if token.Role != model.AdminRole {
		return app.BadRequest(c, "Bad Request")
	}

	collection := app.DB().C(model.ClassC)

	// get class
	id := c.Param("id")
	var class model.Class
	err := collection.FindId(bson.ObjectIdHex(id)).One(&class)
	if err != nil {
		return app.ServerError(c, err)
	}

	// delete first or second class
	if class.Level == model.FirstLevel {
		for _, item := range class.Child {
			err = collection.Remove(
				bson.M{"_id": bson.ObjectIdHex(item)},
			)
			if err != nil {
				return app.ServerError(c, err)
			}
		}
	} else if class.Level == model.SecondLevel {
		err = collection.Update(
			bson.M{"_id": bson.ObjectIdHex(class.Child[0])},
			bson.M{"$pull": bson.M{"child": class.ID.Hex()}},
		)
	}
	
	// delete class
	err = collection.Remove(
		bson.M{"_id": bson.ObjectIdHex(id)},
	)
	if err != nil {
		return app.ServerError(c, err)
	}

	return app.Ok(c, map[string]string{
		"message": "delete success",
	})
}
