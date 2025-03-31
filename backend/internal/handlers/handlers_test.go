package handlers

import (
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHomeHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    HomeHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "Welcome") {
        t.Errorf("unexpected body: %s", string(body))
    }
}

func TestUserHandlerGet(t *testing.T) {
    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()

    UserHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "List of users") {
        t.Errorf("unexpected body: %s", string(body))
    }
}

func TestUserHandlerPost(t *testing.T) {
    req := httptest.NewRequest("POST", "/users", nil)
    w := httptest.NewRecorder()

    UserHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "Create a new user") {
        t.Errorf("unexpected body: %s", string(body))
    }
}

func TestPostHandlerGet(t *testing.T) {
    req := httptest.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()

    PostHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "List of posts") {
        t.Errorf("unexpected body: %s", string(body))
    }
}

func TestPostHandlerPost(t *testing.T) {
    req := httptest.NewRequest("POST", "/posts", nil)
    w := httptest.NewRecorder()

    PostHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Errorf("expected status 200 OK, got %d", res.StatusCode)
    }

    body, _ := io.ReadAll(res.Body)
    if !strings.Contains(string(body), "Create a new post") {
        t.Errorf("unexpected body: %s", string(body))
    }
}
