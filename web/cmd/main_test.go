package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)
	// pass the fake HTTP request to the handler
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v, expected %v", status, http.StatusOK)
	}
	expected := "Hello sir, My name is Ian burns. I'll mostly choose the car rental system due to its complexity and usability. Also it would be useful for a family member who owns a car rental.\n"
	got := rr.Body.String()
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestAboutHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/about", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(about)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got status %v, expected %v", status, http.StatusOK)
	}

	// Check response body contains expected content
	got := rr.Body.String()
	if !strings.Contains(got, "Ian Burns") {
		t.Errorf("expected response to contain 'Ian Burns', got %v", got)
	}
	if !strings.Contains(got, "cybersecurity") {
		t.Errorf("expected response to contain 'cyber security', got %v", got)
	}
}

func TestContactHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/contact", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(contact)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got status %v, expected %v", status, http.StatusOK)
	}

	// Check response body contains expected content
	got := rr.Body.String()
	if !strings.Contains(got, "ianburns344@gmail.com") {
		t.Errorf("expected response to contain email, got %v", got)
	}
	if !strings.Contains(got, "6259908") {
		t.Errorf("expected response to contain phone number, got %v", got)
	}
}

func TestQouteHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/qoute", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(qoute)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got status %v, expected %v", status, http.StatusOK)
	}

	// Check response body contains expected content
	got := rr.Body.String()
	if !strings.Contains(got, "Walt Disney") {
		t.Errorf("expected response to contain 'Walt Disney', got %v", got)
	}
	if !strings.Contains(got, "The best way to get started") {
		t.Errorf("expected response to contain quote text, got %v", got)
	}
}
