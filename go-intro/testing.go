// +build ignore, OMIT
package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func SomeTest(t *testing.T) {
    r, err := http.NewRequest("GET", "http://startup.com/", nil)
    if err != nil {
        t.Error(err.Error())
    }
    w := httptest.NewRecorder()
    myHandler(w, r)
}

