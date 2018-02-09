package ro

import (
	"testing"
	"time"
)

func TestRajaOngkir(t *testing.T) {

	t.Run("Test Construct", func(t *testing.T) {
		raja := New("123456", 10*time.Second)

		if raja == nil {
			t.Error("Cannot construct raja ongkir")
		}
	})

	t.Run("Test Default Env Value", func(t *testing.T) {
		raja := New("123456", 10*time.Second)

		if raja.Env != Starter {
			t.Error("Default Env is not Starter")
		}
	})

	t.Run("Test Change Default Env to Basic", func(t *testing.T) {
		raja := New("123456", 10*time.Second)

		if raja.Env != Starter {
			t.Error("Default Env is not Starter")
		}

		raja.Env = Basic

		if raja.Env != Basic {
			t.Error("Env is not Basic")
		}
	})

	t.Run("Test Change Default Env to Pro", func(t *testing.T) {
		raja := New("123456", 10*time.Second)

		if raja.Env != Starter {
			t.Error("Default Env is not Starter")
		}

		raja.Env = Pro

		if raja.Env != Pro {
			t.Error("Env is not Pro")
		}
	})
}
