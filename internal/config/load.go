package config

import (
	"car_project/configs"
	"fmt"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

var AppConfiguration configs.Configuration

func Load() {
	configs := loadAppConfig()
	if configs.API.Host == "" {
		log.Fatalf("Error loading configurations: EMPTY API HOST FROM CONFIGURATION FILE")
	}
	if configs.API.Port <= 0 {
		log.Fatalf("Error loading configurations: INVALID API PORT IN CONFIGURATION FILE")
	}
	ValidateConfig(configs)
	AppConfiguration = configs
}

func loadAppConfig() configs.Configuration {
	targetFileConfig := fmt.Sprintf("%s/configs/app.yaml", os.Getenv(ROOT_FOLDER_VAR))
	file, errO := os.Open(targetFileConfig)
	if errO != nil {
		log.Fatalf("Error opening App yaml config file [%s]: %s\n", targetFileConfig, errO.Error())
	}
	defer file.Close()

	var configs configs.Configuration
	decoder := yaml.NewDecoder(file)
	err2 := decoder.Decode(&configs)
	if err2 != nil {
		log.Fatalf("Error parsing App yaml config file [%s]: %s\n", targetFileConfig, err2.Error())
	}
	return configs
}

func ValidateConfig(config interface{}) {
	v := reflect.ValueOf(config)
	t := reflect.TypeOf(config)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			ValidateConfig(field.Interface())
		} else {
			if IsEmptyValue(field) {
				log.Fatalf("Error loading configurations: EMPTY %s FROM CONFIGURATION FILE", fieldType.Name)
			}
		}
	}
}

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}
