package utils

import "testing"

func TestUploadFile(t *testing.T) {
	txt, err := UploadFile("test.txt")
	if err != nil {
		t.Errorf("UploadFile() error = %v", err)
		return
	}
	if txt != "File uploaded successfully!" {
		t.Errorf("UploadFile() = %v, want %v", txt, "File uploaded successfully!")
	}
}
