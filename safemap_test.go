package xtypes

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewSafeMap(t *testing.T) {
	sm := NewSafeMap()
	if sm == nil {
		t.Fatalf("failed to create map")
	}
}

func TestSafeMap_Set(t *testing.T) {
	safeMap := NewSafeMap()

	k, v := "Hello", "World"
	safeMap.Set(k, v)

	a, ok := safeMap.Get(k)
	if !ok {
		t.Fatalf("key not found")
	}

	if a != v {
		t.Fatalf("expected %s, got %s", v, a)
	}
}

func TestSafeMap_Get(t *testing.T) {
	TestSafeMap_Set(t)
}

func TestSafeMap_Keys(t *testing.T) {
	safeMap := NewSafeMap()

	testData := []struct {
		key, value string
	}{
		{"Hello", "World"},
		{"Lord", "Of The Rings"},
		{"Star", "Wars"},
	}

	expected := make([]string, 0, len(testData))
	for _, td := range testData {
		safeMap.Set(td.key, td.value)
		expected = append(expected, td.key)
	}

	// NOTE: SafeMap.Keys() returns a sorted slice of keys!
	sort.Strings(expected)

	actual := safeMap.Keys()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func TestSafeMap_Len(t *testing.T) {
	safeMap := NewSafeMap()

	testData := []struct {
		key, value string
	}{
		{"Hello", "World"},
		{"Lord", "Of The Rings"},
		{"Star", "Wars"},
	}

	for _, td := range testData {
		safeMap.Set(td.key, td.value)
	}

	expected := len(testData)
	actual := safeMap.Len()

	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func TestSafeMap_Del(t *testing.T) {
	safeMap := NewSafeMap()

	testData := []struct {
		key, value string
	}{
		{"Hello", "World"},
		{"Lord", "Of The Rings"},
		{"Star", "Wars"},
	}

	testKey := "Lord"

	expected := make([]string, 0, len(testData))
	for _, td := range testData {
		safeMap.Set(td.key, td.value)
		if td.key != testKey {
			expected = append(expected, td.key)
		}
	}

	// NOTE: SafeMap.Keys() returns a sorted slice of keys!
	sort.Strings(expected)

	safeMap.Del(testKey)
	actual := safeMap.Keys()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func TestSafeMap_Drain(t *testing.T) {
	safeMap := NewSafeMap()

	testData := []struct {
		key, value string
	}{
		{"Hello", "World"},
		{"Lord", "Of The Rings"},
		{"Star", "Wars"},
	}

	expectedValues := make([]string, 0, len(testData))
	for _, td := range testData {
		safeMap.Set(td.key, td.value)
		expectedValues = append(expectedValues, td.value)
	}

	expectedLen := len(expectedValues)

	values := safeMap.Drain()

	if valuesLen := len(values); valuesLen != expectedLen {
		t.Fatalf("expected %d items, got %d items", expectedLen, valuesLen)
	}

	// Drain returns a slice of values sorted by key.
	// Also Drain returns a slice of interfaces - []interface{}.
	// We need to convert it to slice of strings since we know the original type.
	// If we didn't know the type we would use type switching.
	actualValues := make([]string, 0, len(values))
	for _, v := range values {
		switch v := v.(type) {
		case string:
			actualValues = append(actualValues, v)
		default:
			t.Errorf("expected type string, got another")
		}
	}

	if actualLen := len(actualValues); actualLen != expectedLen {
		t.Fatalf("expected %d items, got %d items", expectedLen, actualLen)
	}

	if !reflect.DeepEqual(expectedValues, actualValues) {
		t.Fatalf("expected %#v, got %#v", expectedValues, actualValues)
	}
}
