package models

import (
	"background_rabbitmq/config"
	"background_rabbitmq/entity"
)

func SaveMessage(message string) error {
	var result entity.MessageStruct
	result.Message = message
	tx := config.DB.Create(&result)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
