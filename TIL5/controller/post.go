package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func postRouter(post *gin.RouterGroup) {
	post.POST("/new", createPost)
	post.GET("/", getAllPost)
	post.GET("/:ID", getPost)
	post.PUT("/:ID", modifyPost)
	post.DELETE("/:ID", delPost)
}

func createPost(c *gin.Context) {
	post := domain.Post{}

	if err := c.Bind(&post); err != nil {
		log.Printf("[controller:post] error create post : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	_, _, err := service.JoinPost(post)

	if err != nil {
		log.Printf("[controller:post] error create post : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.Status(http.StatusOK)
}

func getAllPost(c *gin.Context) {
	posts, err := service.FindPosts()

	if err != nil {
		log.Printf("[controller:post] error get all post : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:post] error get post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	posts, err := service.FindPostsByID(ID.ID)

	if err != nil {
		log.Printf("[controller:post] error get post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, posts)
}

func modifyPost(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:post] error modify post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	post := domain.Post{}

	if err := c.Bind(&post); err != nil {
		log.Printf("[controller:post] error modify post : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	_, err = service.ModifyPost(ID.ID, post)

	if err != nil {
		log.Printf("[controller:post] error modify post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}

func delPost(c *gin.Context) {
	ID := domain.UriParameter{}
	err := c.ShouldBindUri(&ID)

	if err != nil {
		log.Printf("[controller:post] error delete post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	_, err = service.DelPost(ID.ID)

	if err != nil {
		log.Printf("[controller:post] error delete post : %v\n", err)
		c.JSON(http.StatusNotFound, err)
	}

	c.Status(http.StatusOK)
}
