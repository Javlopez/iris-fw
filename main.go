package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/mvc"
)

type Movie struct {
	Name string 	`json:"name"`
	Year int 		`json:"year"`
	Genre string	`json:"genre"`
	Poster string	`json:"poster"`
}

var movies = []Movie{
	{
		Name: "Casablanca",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	{
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	{
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	{
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
}

var basicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin":"password",
	},
})

func main(){
	app := iris.New()

	app.Use(basicAuth)
	app.Controller("/movies", new(MoviesController))
	app.Run(iris.Addr(":8081"))
}


type MoviesController struct {
	mvc.C
}


// Get returns list of the movies
// Demo:
// curl -i http://localhost:8080/movies
func (c *MoviesController) Get() []Movie{
	return movies
}

// curl -i http://localhost:8080/movies/1
func (c *MoviesController) GetBy(id int) Movie{
	return movies[id]
}

// PutBy updates a movie
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MoviesController) PutBy(id int) Movie{

	m := movies[id]
	file, info, err := c.Ctx.FormFile("poster")

	if err != nil{
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		return Movie{}
	}

	file.Close()

	poster := info.Filename
	genre := c.Ctx.FormValue("genre")

	// update the poster
	m.Poster = poster
	m.Genre = genre
	movies[id] = m
	return m
}


// DeleteBy deletes a movie
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MoviesController) DeleteBy(id int) iris.Map {
	deleted := movies[id].Name

	movies = append(movies[:id], movies[id+1:]...)
	// and return the deleted movie's name
	return iris.Map{"deleted": deleted}
}
