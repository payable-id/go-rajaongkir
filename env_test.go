package ro

import (
	"testing"
)

// TestEnvironment, for Raja Ongkir Enviroment testing
func TestEnvironment(t *testing.T) {

	t.Run("Test Starter Env", func(t *testing.T) {
		starterEnv := Starter

		starterString := starterEnv.String()

		if starterString != "https://api.rajaongkir.com/starter" {
			t.Errorf("Env is not Starter Account %s", starterString)
		}
	})

	t.Run("Test Basic Env", func(t *testing.T) {
		basicEnv := Basic

		basicString := basicEnv.String()

		if basicString != "https://api.rajaongkir.com/basic" {
			t.Errorf("Env is not Basic Account %s", basicString)
		}
	})

	t.Run("Test Pro Env", func(t *testing.T) {
		proEnv := Pro

		proString := proEnv.String()

		if proString != "https://pro.rajaongkir.com/api" {
			t.Errorf("Env is not Pro Account %s", proString)
		}
	})

}
