package controller

import (
	modelv1 "project/model/v1"
)

func (c *Controller) GetUser(id string) (*modelv1.User, error) {
	// TODO: Get user FROM DATABASE

	return &modelv1.User{}, nil
}
