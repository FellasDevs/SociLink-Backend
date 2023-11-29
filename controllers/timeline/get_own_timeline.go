package timeline

import (
	"SociLinkApi/dto"
	likerepository "SociLinkApi/repository/like"
	postrepository "SociLinkApi/repository/timeline"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetOwnTimeline(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if posts, err := postrepository.GetOwnTimeline(userId, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetMainTimelineResponseDto{
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			likes, _ := likerepository.CountPostLikes(post.ID, db)
			response.Posts[i] = dto.PostToPostResponseDto(post, likes)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts recuperados com sucesso",
			"data":    response,
		})
	}
}
