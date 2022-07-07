package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	data "github.com/blaine-t-bush/gopar/data"
)

func getPostBySlug(slug string) data.Post {
	for _, post := range data.Posts {
		if post.Slug == slug {
			return post
		}
	}

	return data.Post{}
}

func getProjectBySlug(slug string) data.Project {
	for _, project := range data.Projects {
		if project.Slug == slug {
			return project
		}
	}

	return data.Project{}
}

func handleRouteAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about", gin.H{"page_title": "About"})
}

func handleRouteContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", gin.H{"page_title": "Contact"})
}

func handleRouteCV(c *gin.Context) {
	c.HTML(http.StatusOK, "cv", gin.H{"page_title": "CV"})
}

func handleRouteHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{"page_title": "Home"})
}

func handleRoutePosts(c *gin.Context) {
	c.HTML(http.StatusOK, "posts", gin.H{"page_title": "Posts", "Posts": data.Posts})
}

func handleRoutePost(c *gin.Context) {
	slug := c.Param("slug")
	post := getPostBySlug(slug)
	c.HTML(http.StatusOK, "post", gin.H{"page_title": post.Title, "post": post})
}

func handleRouteProjects(c *gin.Context) {
	c.HTML(http.StatusOK, "projects", gin.H{"page_title": "Projects", "Projects": data.Projects})
}

func handleRouteProject(c *gin.Context) {
	slug := c.Param("slug")
	project := getProjectBySlug(slug)
	c.HTML(http.StatusOK, "project", gin.H{"page_title": project.Name, "project": project})
}

func main() {
	// Load variables from .env
	err := godotenv.Load(".env.development") // .env.development .env.production
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()                   // Creates a router without any middleware
	router.HTMLRender = ginview.Default() // Use goview with gin support for template rendering
	router.Use(gin.Logger())              // Writes logs to gin.DefaultWriter, by default os.Stdout
	router.Use(gin.Recovery())            // Recovers from panics and instead delivers a 500
	router.Static("/assets", "./assets")  // Serve static files from the assets project directory in the assets route

	router.GET("/", handleRouteHome)
	router.GET("/about/", handleRouteAbout)
	router.GET("/contact/", handleRouteContact)
	router.GET("/cv/", handleRouteCV)
	router.GET("/posts/", handleRoutePosts)
	router.GET("/posts/:slug/", handleRoutePost)
	router.GET("/projects/", handleRouteProjects)
	router.GET("/projects/:slug/", handleRouteProject)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // listen and serve on http://0.0.0.0:8080, or http://localhost:8080 on windows
}
