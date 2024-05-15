package handlers

import (
	"gobackend/pkg/core/model"
	"gobackend/pkg/core/service"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPosts(c *fiber.Ctx) error {
	defaultLimit := 10
	defaultPage := 1

	limit, err := strconv.Atoi(c.Query("limit", strconv.Itoa(defaultLimit)))
	if err != nil || limit <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	page, err := strconv.Atoi(c.Query("page", strconv.Itoa(defaultPage)))
	if err != nil || page <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page parameter",
		})
	}

	offset := (page - 1) * limit
	searchTitle := c.Query("title", "")
	articles, err := h.service.GetPostAll(limit, offset, searchTitle)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalArticles, err := h.service.GetTotalPostCount(searchTitle)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalPages := int(math.Ceil(float64(totalArticles) / float64(limit)))

	response := fiber.Map{
		"posts":      articles,
		"count":      len(articles),
		"limit":      limit,
		"page":       page,
		"total_page": totalPages,
	}

	return c.JSON(response)
}

func (h *Handler) GetPostID(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.IncrementViewCount(id)
	if err != nil {
		return err
	}
	post, err := h.service.GetPostID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get post",
		})
	}
	return c.JSON(post)
}

func (h *Handler) CreatePosts(c *fiber.Ctx) error {
	var post model.Posts
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	if post.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title is required",
		})
	}

	createdArticle, err := h.service.CreatePosts(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {

		response := fiber.Map{
			"id":         createdArticle.ID,
			"title":      createdArticle.Title,
			"content":    createdArticle.Content,
			"published":  createdArticle.Published,
			"created_at": createdArticle.CreatedAt.Format("2006-01-02T15:04:05"),
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateData model.Posts
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	existingPost, err := h.service.GetPostID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get existing post",
		})
	}

	if updateData.Title == "" {
		updateData.Title = existingPost[0].Title
	}
	if updateData.Content == "" {
		updateData.Content = existingPost[0].Content
	}
	if updateData.Published {
		updateData.Published = true
	} else {
		updateData.Published = false
	}

	updateData.UpdatedAt = time.Now()
	updateData.CreatedAt = existingPost[0].CreatedAt

	updatedPost, err := h.service.UpdatePost(id, updateData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	} else {
		response := fiber.Map{
			"id":         updatedPost.ID,
			"title":      updatedPost.Title,
			"content":    updatedPost.Content,
			"published":  updatedPost.Published,
			"created_at": updatedPost.CreatedAt.Format("2006-01-02T15:04:05"),
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func (h *Handler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, _ := h.service.GetPostID(id)
	if post == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}
	err := h.service.DeletePost(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete post",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
