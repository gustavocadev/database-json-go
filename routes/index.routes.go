package routes

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

func ReadJSON(name string) ([]*User, error) {
	var users []*User

	dataBytes, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, errors.New("hay un error, a lo mejor el archivo no existe :(")
	}
	// convierte de sintaxis a json a sintaxis de Go
	json.Unmarshal(dataBytes, &users)
	return users, nil
}

func WriteJSON(name string, mySlice []*User) error {
	// convierte de sintaxis Go a sintaxis JSON
	dataBytes, _ := json.MarshalIndent(mySlice, "", "\t")

	err := ioutil.WriteFile(name, dataBytes, 0744)

	if err != nil {
		return errors.New("hay un error, a lo mejor el archivo no existe :(")
	}
	return nil
}

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Id   string `json:"id"`
}

func UseRoute(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {

		users, _ := ReadJSON("data.json")
		return c.Render("index", fiber.Map{
			"users": users,
		})
	})

	router.Post("/newUser", func(c *fiber.Ctx) error {

		users, _ := ReadJSON("data.json")

		type Request struct {
			Name string `json:"name"`
			Age  uint8  `json:"age"`
		}

		var body Request

		c.BodyParser(&body)

		newUser := &User{
			Name: body.Name,
			Age:  body.Age,
			Id:   xid.New().String(),
		}

		users = append(users, newUser)

		WriteJSON("data.json", users)

		return c.Redirect("/")
	})

	router.Get("/delete/:id", func(c *fiber.Ctx) error {

		users, _ := ReadJSON("data.json")

		id := c.Params("id")

		for idx, user := range users {
			if user.Id == id {
				users = append(users[:idx], users[idx+1:]...)
				break
			}
		}

		WriteJSON("data.json", users)
		return c.Redirect("/")
	})

	router.Get("/update/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		users, _ := ReadJSON("data.json")

		var dataUser *User

		for _, user := range users {
			if user.Id == id {
				dataUser = user
				break
			}
		}

		return c.Render("editForm", fiber.Map{
			"user": dataUser,
		})
	})

	router.Post("/update/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		users, _ := ReadJSON("data.json")

		type Request struct {
			Name *string `json:"name"`
			Age  *uint8  `json:"age"`
		}

		var body Request

		c.BodyParser(&body)

		for _, u := range users {
			if u.Id == id {
				u.Name = *body.Name
				u.Age = *body.Age
				break
			}
		}

		WriteJSON("data.json", users)

		// fmt.Println(users)
		return c.Redirect("/")
	})

}
