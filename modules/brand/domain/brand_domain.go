package brand_domain

import (
	"context"

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