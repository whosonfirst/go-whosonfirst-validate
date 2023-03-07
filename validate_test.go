package validate

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureValidGeoJSON(t *testing.T) {

	rel_path := "fixtures/102527513.geojson"

	abs_path, err := filepath.Abs(rel_path)

	if err != nil {
		t.Fatalf("Failed to derive absolute path for %s, %v", rel_path, err)
	}

	r, err := os.Open(abs_path)

	if err != nil {
		t.Fatalf("Failed to open %s for reading, %v", abs_path, err)
	}

	defer r.Close()

	_, err = EnsureValidGeoJSON(r)

	if err != nil {
		t.Fatalf("%s failed validation, %v", abs_path, err)
	}
}

func TestValidate(t *testing.T) {

	rel_path := "fixtures/102527513.geojson"

	abs_path, err := filepath.Abs(rel_path)

	if err != nil {
		t.Fatalf("Failed to derive absolute path for %s, %v", rel_path, err)
	}

	r, err := os.Open(abs_path)

	if err != nil {
		t.Fatalf("Failed to open %s for reading, %v", abs_path, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", abs_path, err)
	}

	err = Validate(body)

	if err != nil {
		t.Fatalf("%s failed validation, %v", abs_path, err)
	}
}
