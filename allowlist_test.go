package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	allowlist := allowList{"pkg1", "pkg2"}

	if !allowlist.contains("pkg1") {
		t.Errorf("allowlist should contain pkg1")
	}

	if !allowlist.contains("pkg2") {
		t.Errorf("allowlist should contain pkg2")
	}

	if allowlist.contains("pkg3") {
		t.Errorf("allowlist should not contain pkg3")
	}
}

func TestMatch(t *testing.T) {
	var allow allowList

	if !allow.match("test", "test") {
		t.Error("test should exactly match test")
	}

	wildcard := "wildcard/test/*"
	subject := "wildcard/test/abc"
	if !allow.match(wildcard, subject) {
		t.Errorf("%s should match %s", wildcard, subject)
	}
}
