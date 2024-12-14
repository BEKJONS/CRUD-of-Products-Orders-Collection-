package repo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ulab3/internal/entity"
	"ulab3/internal/usecase"
)

type orderRepo struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) usecase.OrderRepository {
	return &orderRepo{collection}
}

func (repo *orderRepo) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	order.ID = uuid.New().String()
	_, err := repo.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *orderRepo) FindAll(ctx context.Context) ([]entity.Order, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []entity.Order
	for cursor.Next(ctx) {
		var order entity.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *orderRepo) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	var order entity.Order
	err := repo.collection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *orderRepo) Update(ctx context.Context, id string, order *entity.Order) error {
	update := bson.M{"$set": order}
	_, err := repo.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	return err
}

func (repo *orderRepo) Delete(ctx context.Context, id string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
