package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello Chris!"
        assertCorrectMessage(t, got, want)
    })
    
    t.Run("saying 'Hello world!' when an empty string is passed", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello world!"
        assertCorrectMessage(t, got, want)
    })
    
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola Elodie!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Jaque", "French")
        want := "Bonjour Jaque!"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Dude", func(t *testing.T) {
        got := Hello("Brad", "Dude")
        want := "Wassssup Brad!"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got string, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
