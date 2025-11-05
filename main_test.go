package main

import "testing"

func TestBegruessung(t *testing.T) {
    erwartet := "Hallo, Welt!"
    erhalten := Greeting("Welt")
    if erhalten != erwartet {
        t.Errorf("Begruessung('Welt') = %s; erwartet %s", erhalten, erwartet)
    }
}
