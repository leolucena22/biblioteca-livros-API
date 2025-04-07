# 📚 API Básica de Livros em Go - Nível 1

Uma API simples para gerenciamento de livros com operações CRUD básicas.

## 🚀 Começando

### Pré-requisitos
- Go 1.20+
- PostgreSQL
- Git
- Ferramenta para testar APIs (Postman/curl)

### Instalação
1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/biblioteca.git
cd biblioteca
```
2. Instale as dependências:
```bash
go mod download
```
3. Configure o banco de dados:
```sql
CREATE DATABASE biblioteca;
```
4. Atualize a string de conexão em src/db/conn.go:
```go
dsn := "host=localhost user=postgres password=sua_senha dbname=biblioteca port=5432 sslmode=disable"
```
5. Execute as migrações:
```bash
go run cmd/main.go
```
### 📋 Endpoints da API
1. Criar Livro
```http request
POST /books
Content-Type: application/json

{
  "title": "Aprendendo Go",
  "isbn": "978-85-333-0227-3",
  "year": 2023
}
```
2. Listar Todos os Livros
```http request
GET /books
```
3. Buscar Livro por ID
```http request
GET /books/1
```
4. Atualizar Livro
```http request
PUT /books/1
Content-Type: application/json

{
  "title": "Go Avançado",
  "isbn": "978-85-333-0450-5",
  "year": 2024
}
```
5. Deletar Livro
```http request
DELETE /books/1
```
