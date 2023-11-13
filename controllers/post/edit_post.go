package postcontroller

import (
	"SociLinkApi/dto"
	postrepository "SociLinkApi/repository/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func EditPost(context *gin.Context, db *gorm.DB) {
	var postData dto.EditPostRequestDto

	if err := context.ShouldBindJSON(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	postId, err := uuid.Parse(postData.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	post, err := postrepository.GetPost(postId, db)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	uid, _ := context.Get("userId")

	userId := uid.(uuid.UUID)
	if userId != post.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "você não tem permissão para editar esse post",
		})
		return
	}

	if postData.Images != nil {
		post.Images = postData.Images
	}
	if postData.Content != "" {
		post.Content = postData.Content
	}
	if postData.Visibility != "" {
		post.Visibility = postData.Visibility
	}

	err = postrepository.UpdatePost(&post, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post editado com sucesso!",
	})
}
