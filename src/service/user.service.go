package service

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/gdsc-ys/fluentify-server/src/model"
)

func GetUser(client *auth.Client, uid string) (model.User, error) {
	userRecord, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return model.User{}, err
	}

	user := convertRecordToUser(userRecord)
	return user, nil
}

func UpdateUser(client *auth.Client, updateUserDTO map[string]interface{}) (model.User, error) {

	ctx := context.Background()
	params := (&auth.UserToUpdate{})
	customClaims := make(map[string]interface{})

	for field, value := range updateUserDTO {
		switch field {
		case "name":
			params = params.DisplayName(value.(string))
		case "age":
			if value.(int) < 0 {
				return model.User{}, &model.UserValidationError{Message: "Age must be greater than 0"}
			}
			customClaims["age"] = value.(int)
		case "disorderType":
			customClaims["disorderType"] = value.(int)
		}
	}
	params = params.CustomClaims(customClaims)

	userRecord, err := client.UpdateUser(ctx, updateUserDTO["uid"].(string), params)
	if err != nil {
		return model.User{}, err
	}

	user := convertRecordToUser(userRecord)
	return user, nil
}

func convertRecordToUser(record *auth.UserRecord) model.User {
	getAge := func() int {
		if age, ok := record.CustomClaims["age"].(float64); ok {
			return int(age)
		}
		return 0
	}

	user := model.User{
		Id:           record.UID,
		Name:         record.DisplayName,
		Age:          getAge(),
		DisorderType: 0,
	}

	return user
}
