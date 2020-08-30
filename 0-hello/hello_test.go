package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to a name", func(t *testing.T){
		got := Hello("Batata", "English")
		want := "Hello, Batata"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world' when the name is empty", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to a name in Spanish", func(t *testing.T){
		got := Hello("Batata", "Spanish")
		want := "Hola, Batata"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to a name in Portuguese", func(t *testing.T){
		got := Hello("Batata", "Portuguese")
		want := "Olá, Batata"
		assertCorrectMessage(t, got, want)
	})
}