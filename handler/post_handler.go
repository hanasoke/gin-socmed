package handler

import (
	"fmt"
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	// Validasi tweet tidak boleh kosong
	if strings.TrimSpace(post.Tweet) == "" {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Tweet cannot be empty"})
		return
	}

	// Handle file upload jika ada
	if post.Picture != nil {
		// Buat directory jika belum ada (path relative)
		dirPath := "./public/picture"
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Generate nama file unik
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Path lengkap untuk penyimpanan
		dst := filepath.Join(dirPath, newFileName)

		// Simpan file
		if err := c.SaveUploadedFile(post.Picture, dst); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}

		// Simpan nama file ke struct untuk diproses service
		post.Picture.Filename = newFileName
		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(int)

	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	// panggila service
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success post your tweet",
	})

	c.JSON(http.StatusCreated, res)
}
