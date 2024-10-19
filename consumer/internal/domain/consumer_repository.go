package domain


import "context"

type ConsumerRepository interface{
	Save(ctx context.Context, consumer *Consumer)error
	Find(ctx context.Context, consumerID string)(*Consumer,error)
	Update(ctx context.Context,consumer *Consumer)error
}