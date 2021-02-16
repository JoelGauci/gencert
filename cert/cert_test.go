package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2021-01-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}

	if c == nil {
		t.Error("Cert should be a valid reference. got=nil")
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2021-01-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "aazertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklm"
	_, err := New(course, "Bob", "2021-01-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2021-01-31")
	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "aazertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklm"
	_, err := New("azertyuiopq", name, "2021-01-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (name=%s)", "name")
	}
}
