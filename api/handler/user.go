package handler

import (
	"log"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"

	"app/models"
	"app/storage"
)

func (h *HandlerV1) Create(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.Create(h.db, user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create"))
		return
	}


	userId, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

	c.JSON(http.StatusOK, userId)
}

func (h *HandlerV1) GetById(c *gin.Context) {

	id := c.Param("id")

	user, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HandlerV1) GetList(c *gin.Context) {

	users, err := storage.GetList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list"))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *HandlerV1) Delete(c *gin.Context) {

	id := c.Param("id")

	err := storage.Delete(h.db, id)
	if err != nil {
		log.Printf("error whiling delete by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id"))
		return
	}

}

func (h *HandlerV1) Update(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err= storage.Update(h.db, user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update"))
		return
	}


}

func (h *HandlerV1) Patch(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err= storage.Update_patch(h.db, user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update"))
		return
	}


}


// update
// patch
// delete

// github push
