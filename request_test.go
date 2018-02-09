package ro

import (
	"testing"
)

func TestRequest(t *testing.T) {

	t.Run("Test Request Construct", func(t *testing.T) {
		query := &QueryRequest{CityID: "001", ProvinceID: "002"}

		if query == nil {
			t.Error("Cannot construct Query Request")
		}

		if query.CityID != "001" {
			t.Error("Invalid City Id")
		}

		if query.ProvinceID != "002" {
			t.Error("Invalid Province Id")
		}
	})

}
