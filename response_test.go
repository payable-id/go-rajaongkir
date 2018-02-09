package ro

import (
	"testing"
)

func TestResponse(t *testing.T) {

	t.Run("Test Response Construct", func(t *testing.T) {
		status := Status{Code: 200, Description: "Ok"}

		if status.Code != 200 {
			t.Error("Status Code invalid")
		}

		if status.Description != "Ok" {
			t.Error("Status Description invalid")
		}

		province := Province{ProvinceID: "001", Province: "Jawa Tengah"}

		if province.ProvinceID != "001" {
			t.Error("Province Id nvalid")
		}

		if province.Province != "Jawa Tengah" {
			t.Error("Province Name invalid")
		}

	})

}
