# Projeto Go de um jogo de cartas

Esse é um projeto escrito em Go que implementa um embaralhador de cartas. O projeto consiste em um arquivo `main.go` e outro `deck.go` que possuem o código principal da aplicação, e um arquivo de teste `deck_test.go` que testa algumas funcionalidades do código.

# Como rodar

## Dev

Para rodar a aplicação, basta rodar o arquivo `main.go` e `deck.go` no terminal:

```
* go run main.go deck.go
```

## Teste

Para rodar a aplicação como teste, basta rodar o comando no terminal:

```
go test
```

## Build

Para rodar fazer o build do projeto utilize o comando

```
go build
```

E então será gerado um arquivo no mesmo folder, no caso de windows com extensão `.exe`

---

Ao rodar o código, uma nova instância de um baralho é criada, embaralhada, impressa na tela e salva em um arquivo chamado cards.txt. Em seguida, o código lê o baralho a partir do arquivo cards.txt e o imprime novamente na tela.

Funcionalidades
O projeto possui as seguintes funcionalidades:

- Criar um novo baralho
- Embaralhar o baralho
- Imprimir o baralho na tela
- Salvar o baralho em um arquivo
- Ler um baralho a partir de um arquivo
- O projeto possui também testes automatizados que garantem o correto funcionamento das funcionalidades.
