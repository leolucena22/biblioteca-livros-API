# 📚 Desafio Intermediário: API de Biblioteca de Livros em Go
**Domine o básico e construa uma API robusta para gerenciar livros e autores!**

---

## 🎯 **Objetivo**
Criar uma API RESTful para gerenciar livros e autores, com autenticação JWT e relacionamentos entre entidades.

---

## 📋 **Nível 1: CRUD Básico de Livros**

### **Rotas**
| Método | Rota        | Descrição                     |
|--------|-------------|-------------------------------|
| POST   | `/books`    | Cadastrar livro               |
| GET    | `/books`    | Listar livros (com filtros)   |
| GET    | `/books/:id`| Buscar livro por ID           |
| PUT    | `/books/:id`| Atualizar livro               |
| DELETE | `/books/:id`| Excluir livro                 |

### **Modelo `Book`**
```go
type Book struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title" validate:"required,min=3"`
    ISBN      string    `json:"isbn" validate:"required,isbn"`
    Year      int       `json:"year" validate:"gte=1900"`
    AuthorID  uint      `json:"author_id"` // Para relacionamento
    CreatedAt time.Time `json:"created_at"`
}