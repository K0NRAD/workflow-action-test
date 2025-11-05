package main

import "fmt"

func main() {
    fmt.Println("Hallo von GitHub Actions!")
}

// Eine einfache Funktion zum Testen
func Greeting(name string) string {
    if name == "" {
        return "Hallo!"
    }
    return fmt.Sprintf("Hallo, %s!", name)
}
