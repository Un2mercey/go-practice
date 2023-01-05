package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup")
	result := m.Run()
	fmt.Println("Tear down")

	os.Exit(result)
}

func TestMultiple(t *testing.T) {
	// setup: insert test data within bd
	t.Run("group:A", func(t *testing.T) {
		t.Log("A")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 2, 2, 4

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 222, 222, 49284

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()

			var x, y, result = -2, 4, -8

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
	})

	t.Run("group:B", func(t *testing.T) {
		t.Log("B")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 2, 2, 4

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 222, 222, 49284

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()

			var x, y, result = -2, 4, -8

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
	})

	t.Run("group:C", func(t *testing.T) {
		t.Log("C")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 2, 2, 4

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 222, 222, 49284

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()

			var x, y, result = -2, 4, -8

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
	})

	t.Run("group:D", func(t *testing.T) {
		t.Log("D")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 2, 2, 4

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("medium", func(t *testing.T) {
			t.Parallel()

			var x, y, result = 222, 222, 49284

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel()

			var x, y, result = -2, 4, -8

			testResult := Multiple(x, y)

			if testResult != result {
				t.Errorf("expected result: %d != %d :tested result", result, testResult)
			}
		})
	})
	// tearDown: remove test data within bd
}
