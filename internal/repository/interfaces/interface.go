package interfaces

import (
	"Bucket/internal/model"
)

type IBucket interface {
	addAdvertToBucket(advert model.Advert)
	deleteAdvertToBucket(advertId int)
}
