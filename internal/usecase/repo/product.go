package repo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ulab3/internal/entity"
	"ulab3/internal/usecase"
)

type productRepo struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) usecase.ProductRepository {
	return &productRepo{collection}
}

func (repo *productRepo) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	product.ID = uuid.New().String()
	_, err := repo.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *productRepo) FindAll(ctx context.Context) ([]entity.Product, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entity.Product
	for cursor.Next(ctx) {
		var product entity.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *productRepo) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	var product entity.Product
	err := repo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *productRepo) Update(ctx context.Context, id string, product *entity.Product) error {
	update := bson.M{"$set": product}
	_, err := repo.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	return err
}

func (repo *productRepo) Delete(ctx context.Context, id string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
