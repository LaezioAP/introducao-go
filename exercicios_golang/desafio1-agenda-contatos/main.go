package main

import "fmt"

type Contato struct {
	Nome     string
	Telefone string
	Email    string
}

type Printar interface {
	mostrarContatos()
}

type Agenda struct {
	Contatos map[string]Contato
}

func (a Agenda) mostrarContatos() {
	for _, valor := range a.Contatos {
		fmt.Println(valor)
	}
}

func adicionarContatos(valores map[string]Contato, name, tel, email string) {
	valores[name] = Contato{Nome: name, Telefone: tel, Email: email}
}

func removerContatos(valores map[string]Contato, chave string) {
	delete(valores, chave)
}

func buscarContato(valores map[string]Contato, chave string) {
	fmt.Println(valores[chave])
}

func main() {
	contatos := make(map[string]Contato)

	adicionarContatos(contatos, "Laezio", "981382227", "laezio@fakemail.com")
	adicionarContatos(contatos, "Fulano", "3333", "fula@fakemail.com")
	adicionarContatos(contatos, "Joice", "22222", "Joice@fakemail.com")

	agenda := Agenda{Contatos: contatos}
	agenda.mostrarContatos()

	fmt.Println("--------------------------")
	removerContatos(contatos, "Laezio")
	agenda.mostrarContatos()

	fmt.Println("--------------------------")
	buscarContato(contatos, "Joice")

}
