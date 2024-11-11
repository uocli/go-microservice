package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/uocli/go-microservice/internal/dberrors"
	"github.com/uocli/go-microservice/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Where(models.Product{
		VendorID: vendorId,
	}).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
	}
	return product, nil
}

func (c Client) GetProductByID(ctx context.Context, ID string) (*models.Product, error) {
	product := &models.Product{}
	result := c.DB.WithContext(ctx).
		Where(&models.Product{ProductID: ID}).
		First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "product",
				ID:     ID,
			}
		}
		return nil, result.Error
	}
	return product, nil
}
