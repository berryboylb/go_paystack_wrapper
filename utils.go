package paystack

import (
	"reflect"
	"strings"
	"strconv"
	"fmt"
	"net/url"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

func isValidTimestamp(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, err := time.Parse(time.RFC3339, value)
	return err == nil
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("timestamp", isValidTimestamp)
}

// validate struct

func Validate(structure interface{}) error {
	err := validate.Struct(structure)
	if err != nil {
		verr := err.(validator.ValidationErrors)
		var errs []string
		for _, e := range verr {
			errs = append(errs, fmt.Sprintf("Field '%s' failed on the '%s' tag", e.Field(), e.Tag()))
		}
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}

// filter empty variables
func (fb *FilterBanks) FilterEmptyFields() map[string]interface{} {
	inputValue := reflect.ValueOf(fb).Elem()
	inputType := inputValue.Type()
	filtered := make(map[string]interface{})

	for i := 0; i < inputType.NumField(); i++ {
		field := inputValue.Field(i)
		fieldType := inputType.Field(i)

		// Get the JSON tag from the struct field
		jsonTag := fieldType.Tag.Get("json")
		// Split the tag on comma to get the key
		jsonKey := strings.Split(jsonTag, ",")[0]

		// Check if the field is a pointer and if it's nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// Check if the field is a string and if it's empty
		if field.Kind() == reflect.String && field.String() == "" {
			continue
		}

		// Add the field to the filtered map using the JSON key
		filtered[jsonKey] = field.Interface()
	}

	return filtered
}

func (lt *ListTransactions) FilterFields() map[string]interface{} {
	inputValue := reflect.ValueOf(lt).Elem()
	inputType := inputValue.Type()
	filtered := make(map[string]interface{})

	for i := 0; i < inputType.NumField(); i++ {
		field := inputValue.Field(i)
		fieldType := inputType.Field(i)

		// Get the JSON tag from the struct field
		jsonTag := fieldType.Tag.Get("json")
		// Split the tag on comma to get the key
		jsonKey := strings.Split(jsonTag, ",")[0]

		// Check if the field is a pointer and if it's nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// Check if the field is a string and if it's empty
		if field.Kind() == reflect.String && field.String() == "" {
			continue
		}

		// Check if the field is an integer and if it's zero
		if field.Kind() == reflect.Int && field.Int() == 0 {
			continue
		}

		// Add the field to the filtered map using the JSON key
		filtered[jsonKey] = field.Interface()
	}

	return filtered
}


// encode 

func encodeFilteredFields(filtered map[string]interface{}) (string, error) {
	queryValues := url.Values{}

	for key, value := range filtered {
		// Convert the value to a string
		var strValue string
		switch v := value.(type) {
		case *bool:
			if v != nil {
				strValue = strconv.FormatBool(*v)
			}
		case string:
			strValue = v
		case int:
			strValue = strconv.Itoa(v)
		// Add other types as needed
		default:
			return "", fmt.Errorf("unsupported type for key %s: %T", key, v)
		}

		// Add the key-value pair to the query values
		if strValue != "" {
			queryValues.Add(key, strValue)
		}
	}

	return queryValues.Encode(), nil
}