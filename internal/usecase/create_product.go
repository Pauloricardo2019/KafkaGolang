package usecase

import "GoKafkaMessenger/internal/entity"

type CreateProductInputDto struct {
	Name  string
	Price float64
}

type CreateProductOutDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutDto, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductOutDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
