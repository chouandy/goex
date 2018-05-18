package stringsex

import "testing"

func TestIn(t *testing.T) {
	stra := []string{"en", "zh-TW"}
	str1 := "en"
	str2 := "jp"

	if !In(str1, stra) {
		t.Fatal("Expected en is in [\"en\", \"zh-TW\"]")
	}

	if In(str2, stra) {
		t.Fatal("Expected jp is not in [\"en\", \"zh-TW\"]")
	}
}

func TestInclude(t *testing.T) {
	stra := []string{"en", "zh-TW"}
	str1 := "en"
	str2 := "jp"

	if !Include(stra, str1) {
		t.Fatal("Expected [\"en\", \"zh-TW\"] include en")
	}

	if Include(stra, str2) {
		t.Fatal("Expected [\"en\", \"zh-TW\"] do not include jp")
	}
}

func TestExclude(t *testing.T) {
	stra := []string{"en", "zh-TW"}
	str1 := "en"
	str2 := "jp"

	if Exclude(stra, str1) {
		t.Fatal("Expected [\"en\", \"zh-TW\"] do not exclude en")
	}

	if !Exclude(stra, str2) {
		t.Fatal("Expected [\"en\", \"zh-TW\"] exclude jp")
	}
}
