package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	data "github.com/blaine-t-bush/gopar/data"
)

func getProjectBySlug(slug string) data.Project {
	for _, project := range data.Projects {
		if project.Slug == slug {
			return project
		}
	}

	return data.Project{}
}

func handleRouteHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func handleRouteAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func handleRouteContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{})
}

func handleRouteProjects(c *gin.Context) {
	c.HTML(http.StatusOK, "projects.html", gin.H{"Projects": data.Projects})
}

func handleRouteProject(c *gin.Context) {
	slug := c.Param("slug")
	project := getProjectBySlug(slug)
	c.HTML(http.StatusOK, "project.html", project)
}

func main() {
	router := gin.New()        // Creates a router without any middleware
	router.Use(gin.Logger())   // Writes logs to gin.DefaultWriter, by default os.Stdout
	router.Use(gin.Recovery()) // Recovers from panics and instead delivers a 500

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handleRouteHome)
	router.GET("/about/", handleRouteAbout)
	router.GET("/contact/", handleRouteContact)
	router.GET("/projects/", handleRouteProjects)
	router.GET("/projects/:slug/", handleRouteProject)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
