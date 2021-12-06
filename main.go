//go:generate swag init -o docs/docs

package main

import (
	"fmt"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "goswagger/docs/docs"
)

type UserHandler struct {
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Enabled bool   `json:"enabled"`
	Roles   []Role `json:"roles"`
}

type UserListResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Role struct {
	Title string `json:"title"`
}

func main() {
	userHandler := &UserHandler{}
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Get("/users", userHandler.getUsers)
	r.Post("/users", userHandler.saveUser)
	r.Get("/users/{id}", userHandler.getUserByID)

	log.Fatal(http.ListenAndServe(":8080", r))
}

// Return all users
// @Summary      Get all users
// @Description  Get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  UserListResponse
// @Router       /users [get]
func (u *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all users")
	users := []User{
		{
			Name:    "Kevin",
			Age:     18,
			Enabled: true,
			Roles: []Role{
				{
					Title: "admin",
				},
				{
					Title: "user",
				},
			},
		},
		{
			Name:    "John",
			Age:     20,
			Enabled: false,
			Roles: []Role{
				{
					Title: "user",
				},
			},
		},
	}
	render.JSON(w, r, users)
}

// Save a user
// @Summary      Save a user
// @Description  Save a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param         user body User true "User"
// @Success      200  {object}  User
// @Router       /users [post]
func (u *UserHandler) saveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("saving user")

}

// Get a user by ID
// @Summary      Get a user by ID
// @Description  Get a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param         id path string true "User ID"
// @Success      200  {object}  User
// @Router       /users/{id} [get]
func (u *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting user by id")
	user := User{
		Name:    "FakeUser",
		Age:     18,
		Enabled: true,
	}
	render.JSON(w, r, user)
}
