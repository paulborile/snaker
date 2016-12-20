package snaker_test

import (
	"snaker"
	"testing"
)

func TestSnaker_CamelToSnake(t *testing.T) {
	camel := "AdvertisedDeviceOS"
	snake := "advertised_device_os"
	sna := snaker.CamelToSnake(camel)

	if sna != snake {
		t.Errorf("snake = %s, camel = %s\n", sna, camel)
	}
}

func TestSnaker_SnakeToCamel(t *testing.T) {
	camel := "AdvertisedDeviceOS"
	snake := "advertised_device_os"

	cam := snaker.SnakeToCamel(snake)

	if cam != camel {
		t.Errorf("snake = %s, camel = %s\n", snake, camel)
	}
}
