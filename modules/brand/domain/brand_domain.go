package brand_domain

import (
	"context"

	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/modules/brand"
)

type BrandDomain struct {
	brandRepository	brand.Repository
}

func NewBrandDomain(br brand.Repository) brand.Domain {
	return &BrandDomain {
		brandRepository: br,
	}
}

func (d *BrandDomain) GetAll(ctx context.Context) (error, []*brand.Entity) {
	err, brands := d.brandRepository.GetAll(ctx)

	if err != nil {
		return err, nil
	}

	return nil, brands
}

func (d *BrandDomain) Create(ctx context.Context, data *brand.Entity) error {
	err, nameExist := d.brandRepository.GetByName(ctx, data.Name)

	if err != nil {
		return err
	}

	if nameExist != nil {
		return response.BadRequest("Brand name already taken")
	}

	errInsert := d.brandRepository.Insert(ctx, data)

	if errInsert != nil {
		return err
	}

	return nil
}