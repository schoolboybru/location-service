package cache

import (
	"context"

	"github.com/schoolboybru/location-service/model"
)

var ctx = context.Background()

func GetLocation(name string) (*model.Location, error) {
	var location = new(model.Location)

	c, err := NewCache()

	if err != nil {
		return nil, err
	}

	err = c.Get(ctx, name).Scan(&location)

	if err != nil {
		return nil, err
	}

	return location, nil
}

func SetLocation(l model.Location) error {
	c, err := NewCache()

	if err != nil {
		return err
	}

	err = c.Set(ctx, l.Name, l, 0).Err()

	if err != nil {
		return err
	}

	return nil

}
