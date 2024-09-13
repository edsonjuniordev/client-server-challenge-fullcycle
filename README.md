# Desafio Client-Server-API Full Cycle

Este projeto é parte de um desafio do curso Pós Go Expert Full Cycle, cujo objetivo era desenvolver um servidor capaz de consultar a cotação do dólar, salvar os dados em um banco de dados SQLite e retornar a resposta em formato JSON. O cliente, por sua vez, faria uma requisição HTTP ao servidor para obter a cotação e salvá-la em um arquivo local.

## 🚧 Como instalar e rodar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/edsonjuniordev/client-server-challenge-fullcycle.git
cd client-server-challenge-fullcycle
```

### 2. Rodar o server

```bash
go run server/main.go
```

### 3. Rodar o client

```bash
go run client/main.go
```