package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/lareii/siker.im/internal/models"
	"github.com/lareii/siker.im/internal/repository"
	"github.com/lareii/siker.im/internal/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) CreateURL(ctx context.Context, req *models.CreateURLRequest) (*models.URLResponse, int) {
	var slug string
	var err error

	if req.Slug != "" {
		exists, err := s.repo.ExistsBySlug(ctx, req.Slug)
		if err != nil {
			return nil, fiber.StatusInternalServerError
		}
		if exists {
			return nil, fiber.StatusConflict
		}
		slug = req.Slug
	} else {
		slug, err = utils.GenerateSlug()
		if err != nil {
			return nil, fiber.StatusInternalServerError
		}

		for {
			exists, err := s.repo.ExistsBySlug(ctx, slug)
			if err != nil {
				return nil, fiber.StatusInternalServerError
			}
			if !exists {
				break
			}
			slug, err = utils.GenerateSlug()
			if err != nil {
				return nil, fiber.StatusInternalServerError
			}
		}
	}

	url := &models.URL{
		TargetURL: req.TargetURL,
		Slug:      slug,
	}

	if err := s.repo.Create(ctx, url); err != nil {
		return nil, fiber.StatusInternalServerError
	}

	return s.toResponse(url), fiber.StatusCreated
}

func (s *URLService) GetURLByID(ctx context.Context, id bson.ObjectID) (*models.URLResponse, error) {
	url, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("slug not found")
		}
		return nil, err
	}

	return s.toResponse(url), nil
}

func (s *URLService) GetURLBySlug(ctx context.Context, slug string) (*models.URLResponse, error) {
	url, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("slug not found")
		}
		return nil, err
	}

	return s.toResponse(url), nil
}

func (s *URLService) GetTargetURL(ctx context.Context, slug string) (string, int) {
	url, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fiber.StatusNotFound
		}
		return "", fiber.StatusInternalServerError
	}

	if !url.IsActive {
		return "", fiber.StatusForbidden
	}

	if err := s.repo.IncrementClickCount(ctx, slug); err != nil {
		fmt.Printf("Failed to increment click count: %v\n", err)
	}

	return url.TargetURL, fiber.StatusFound
}

// func (s *URLService) DeleteURL(ctx context.Context, id string) error {
// 	objectID, err := bson.ObjectIDFromHex(id)
// 	if err != nil {
// 		return errors.New("invalid url id")
// 	}

// 	return s.repo.Delete(ctx, objectID)
// }

func (s *URLService) toResponse(url *models.URL) *models.URLResponse {
	return &models.URLResponse{
		ID:         url.ID.Hex(),
		TargetURL:  url.TargetURL,
		Slug:       url.Slug,
		CreatedAt:  url.CreatedAt,
		ClickCount: url.ClickCount,
		IsActive:   url.IsActive,
	}
}
