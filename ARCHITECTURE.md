# Arquitetura do MayCMS

## Visão Geral

O MayCMS segue o padrão de **Arquitetura Hexagonal** (também conhecida como Ports and Adapters), combinado com as melhores práticas da comunidade Go. Este design garante uma aplicação altamente testável, mantível e desacoplada de frameworks específicos.

## Princípios Arquiteturais

### 1. Arquitetura Hexagonal (Ports & Adapters)

A arquitetura é organizada em camadas concêntricas:

```
┌─────────────────────────────────────────┐
│         EXTERNAL WORLD                  │
│  (Gin, PostgreSQL, HTTP Clients, etc)   │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│  ADAPTERS (External Interfaces)         │
│  - API REST (Gin)                       │
│  - Repository (Database)                │
│  - External Services                    │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│  PORTS (Interfaces)                     │
│  - Repository Interfaces                │
│  - Service Interfaces                   │
└──────────────┬──────────────────────────┘
               │
┌──────────────▼──────────────────────────┐
│  APPLICATION CORE (Business Logic)      │
│  - Use Cases / Services                 │
│  - Domain Models                        │
│  - Business Rules                       │
└─────────────────────────────────────────┘
```

### 2. Inversão de Dependência

- O domínio **nunca** depende de adaptadores
- Adaptadores dependem de interfaces (ports) definidas pelo domínio
- Injeção de dependência é utilizada para desacoplar componentes

## Estrutura de Diretórios

```
maycms/
├── cmd/                          # Entry points da aplicação
│   └── maycms/
│       └── main.go              # Main function
│
├── internal/                     # Código privado (não exportável)
│   ├── domain/                   # Camada de Domínio (núcleo)
│   │   ├── entities/             # Modelos de domínio (estruturas principais)
│   │   │   ├── post.go
│   │   │   ├── user.go
│   │   │   └── comment.go
│   │   ├── value_objects/        # Objetos de valor (immutáveis)
│   │   │   ├── email.go
│   │   │   └── slug.go
│   │   └── errors/               # Erros de domínio específicos
│   │       └── errors.go
│   │
│   ├── application/              # Camada de Aplicação (Use Cases)
│   │   ├── dto/                  # Data Transfer Objects
│   │   │   ├── post_dto.go
│   │   │   └── user_dto.go
│   │   ├── services/             # Orquestração de lógica de negócio
│   │   │   ├── post_service.go
│   │   │   ├── user_service.go
│   │   │   └── interfaces.go     # Interfaces (ports)
│   │   └── mappers/              # Conversão entre camadas
│   │       ├── post_mapper.go
│   │       └── user_mapper.go
│   │
│   ├── infrastructure/           # Camada de Infraestrutura (Adaptadores)
│   │   ├── persistence/          # Implementações de Repository
│   │   │   ├── postgres/
│   │   │   │   ├── post_repository.go
│   │   │   │   └── user_repository.go
│   │   │   └── migrations/       # Migrações do banco de dados
│   │   │       └── 001_initial.sql
│   │   ├── http/                 # Handlers HTTP (Gin)
│   │   │   ├── handlers/
│   │   │   │   ├── post_handler.go
│   │   │   │   └── user_handler.go
│   │   │   ├── middleware/       # Middlewares
│   │   │   │   ├── logger.go
│   │   │   │   ├── error_handler.go
│   │   │   │   └── auth.go
│   │   │   └── router.go         # Definição de rotas
│   │   ├── config/               # Configurações
│   │   │   └── config.go
│   │   └── logger/               # Implementação de logging
│   │       └── logger.go
│   │
│   └── shared/                   # Código compartilhado (utilitários)
│       ├── utils/
│       │   ├── pagination.go
│       │   └── validator.go
│       └── constants/
│           └── constants.go
│
├── tests/                        # Testes (fora de internal)
│   ├── integration/              # Testes de integração
│   │   └── post_test.go
│   └── fixtures/                 # Dados de teste
│       └── seed.go
│
├── go.mod                        # Módulo Go
├── go.sum                        # Dependências
├── Dockerfile                    # Containerização
├── docker-compose.yml            # Ambiente local
├── .env.example                  # Variáveis de ambiente
├── README.md                     # Documentação geral
└── ARCHITECTURE.md               # Este arquivo
```

## Camadas Explicadas

### 1. **Domain Layer** (Camada de Domínio)

O coração da aplicação, contendo a lógica de negócio pura, independente de qualquer tecnologia.

**Responsabilidades:**
- Definir entidades (entities) e objetos de valor (value objects)
- Expressar regras de negócio
- Lançar exceções de domínio específicas

**Exemplo:**
```go
// internal/domain/entities/post.go
package entities

type Post struct {
    ID        int64
    Title     string
    Content   string
    Author    User
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (p *Post) Validate() error {
    if len(p.Title) == 0 {
        return errors.New("título é obrigatório")
    }
    return nil
}
```

### 2. **Application Layer** (Camada de Aplicação)

Orquestra o domínio através de use cases e serviços, sem conhecer detalhes de implementação.

**Responsabilidades:**
- Definir interfaces (ports)
- Implementar casos de uso
- Converter dados entre camadas (DTOs e Mappers)
- Coordenar fluxos de negócio

**Exemplo:**
```go
// internal/application/services/interfaces.go
package services

import "context"
import "maycms/internal/domain/entities"

type PostRepository interface {
    GetByID(ctx context.Context, id int64) (*entities.Post, error)
    Save(ctx context.Context, post *entities.Post) error
    Delete(ctx context.Context, id int64) error
}

// internal/application/services/post_service.go
package services

type PostService struct {
    repo PostRepository
}

func (s *PostService) CreatePost(ctx context.Context, dto *PostDTO) (*PostDTO, error) {
    post := &entities.Post{
        Title:   dto.Title,
        Content: dto.Content,
    }
    
    if err := post.Validate(); err != nil {
        return nil, err
    }
    
    if err := s.repo.Save(ctx, post); err != nil {
        return nil, err
    }
    
    return toDTO(post), nil
}
```

### 3. **Infrastructure Layer** (Camada de Infraestrutura)

Implementa as interfaces definidas pela camada de aplicação, adaptando o domínio ao mundo externo.

**Responsabilidades:**
- Implementar repositories (persistência)
- Implementar handlers HTTP
- Gerenciar configurações
- Integração com bibliotecas externas (Gin, PostgreSQL, etc.)

**Exemplo:**
```go
// internal/infrastructure/persistence/postgres/post_repository.go
package postgres

import (
    "context"
    "database/sql"
    "maycms/internal/domain/entities"
)

type PostRepository struct {
    db *sql.DB
}

func (r *PostRepository) GetByID(ctx context.Context, id int64) (*entities.Post, error) {
    // Implementação específica do PostgreSQL
    var post entities.Post
    err := r.db.QueryRowContext(ctx, 
        "SELECT id, title, content FROM posts WHERE id = $1", id).
        Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
        return nil, err
    }
    return &post, nil
}
```

### 4. **Shared Layer** (Camada Compartilhada)

Utilitários e código reutilizável que não pertence a nenhuma camada específica.

## Fluxo de Requisição HTTP

```
HTTP Request
    │
    ▼
┌─────────────────────────┐
│ Handler (Adapter)       │ Recebe HTTP
├─────────────────────────┤
    │
    ▼ Converte para DTO
┌─────────────────────────┐
│ Service (Application)   │ Executa lógica
├─────────────────────────┤
    │
    ▼ Usa domínio
┌─────────────────────────┐
│ Entity (Domain)         │ Valida regras
├─────────────────────────┤
    │
    ▼ Persiste
┌─────────────────────────┐
│ Repository (Adapter)    │ Salva no BD
├─────────────────────────┤
    │
    ▼ Retorna DTO
┌─────────────────────────┐
│ Handler (Adapter)       │ Serializa JSON
├─────────────────────────┤
    │
    ▼
HTTP Response
```

## Melhores Práticas Go Aplicadas

### 1. **Naming Conventions**

- Nomes de pacotes: minúsculos, sem underscores
- Tipos públicos: PascalCase
- Funções públicas: PascalCase
- Variáveis privadas: camelCase

```go
package post        // ✓ Bom
package post_service // ✗ Ruim

type PostService struct{}   // ✓ Bom
func (s *PostService) Save() {} // ✓ Bom
```

### 2. **Error Handling**

Sempre retorne erros como último valor. Use tipos de erro customizados para domínio.

```go
// Bom
func (s *Service) Create(ctx context.Context, data *Data) error {
    if err := validate(data); err != nil {
        return fmt.Errorf("validação falhou: %w", err)
    }
    return nil
}

// Evitar
func (s *Service) Create(data *Data) (error, *Result) {} // ✗ Ordem errada
```

### 3. **Context Propagation**

Use context para propagar prazos, cancelamentos e valores entre camadas.

```go
func (s *Service) GetPost(ctx context.Context, id int64) (*Post, error) {
    post, err := s.repo.GetByID(ctx, id) // Passa context
    if err != nil {
        return nil, fmt.Errorf("falha ao buscar post: %w", err)
    }
    return post, nil
}
```

### 4. **Interface Segregation**

Interfaces pequenas e focadas, não grandes e genéricas.

```go
// ✓ Bom - Específico
type PostRepository interface {
    GetByID(ctx context.Context, id int64) (*Post, error)
}

// ✗ Ruim - Muito genérico
type Repository interface {
    Get(ctx context.Context, id string) (interface{}, error)
    Save(ctx context.Context, data interface{}) error
    Delete(ctx context.Context, id string) error
}
```

### 5. **Dependency Injection**

Use injeção de dependência através de construtores.

```go
// ✓ Bom
func NewPostService(repo PostRepository) *PostService {
    return &PostService{
        repo: repo,
    }
}

// ✗ Ruim - Acoplamento com implementação
func NewPostService() *PostService {
    return &PostService{
        repo: &PostgresRepository{}, // Acoplado
    }
}
```

### 6. **Testabilidade**

Estruture o código para facilitar testes unitários:

```go
// internal/application/services/post_service_test.go
package services

import (
    "context"
    "testing"
)

type MockRepository struct {
    GetByIDFunc func(ctx context.Context, id int64) (*Post, error)
}

func (m *MockRepository) GetByID(ctx context.Context, id int64) (*Post, error) {
    return m.GetByIDFunc(ctx, id)
}

func TestPostService_GetPost(t *testing.T) {
    mockRepo := &MockRepository{
        GetByIDFunc: func(ctx context.Context, id int64) (*Post, error) {
            return &Post{ID: 1, Title: "Test"}, nil
        },
    }
    
    service := NewPostService(mockRepo)
    post, err := service.GetPost(context.Background(), 1)
    
    if err != nil {
        t.Fatalf("erro inesperado: %v", err)
    }
    if post.ID != 1 {
        t.Errorf("esperado ID 1, obteve %d", post.ID)
    }
}
```

### 7. **Logging Estruturado**

Use logging estruturado para melhor rastreabilidade.

```go
logger.WithFields(map[string]interface{}{
    "post_id": post.ID,
    "action":  "create",
}).Info("Post criado com sucesso")
```

## Padrões Complementares

### Repository Pattern

Abstrai a persistência de dados, permitindo trocar implementações facilmente.

```go
// internal/application/services/interfaces.go
type PostRepository interface {
    GetByID(ctx context.Context, id int64) (*Post, error)
    GetAll(ctx context.Context) ([]*Post, error)
    Save(ctx context.Context, post *Post) error
    Delete(ctx context.Context, id int64) error
}
```

### DTO (Data Transfer Object)

Transfere dados entre camadas sem expor modelos internos.

```go
// internal/application/dto/post_dto.go
type PostDTO struct {
    ID        int64  `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    CreatedAt string `json:"created_at"`
}
```

### Mapper

Converte entre diferentes representações (Entity ↔ DTO).

```go
// internal/application/mappers/post_mapper.go
func ToDTO(post *entities.Post) *PostDTO {
    return &PostDTO{
        ID:        post.ID,
        Title:     post.Title,
        Content:   post.Content,
        CreatedAt: post.CreatedAt.Format(time.RFC3339),
    }
}

func ToEntity(dto *PostDTO) *entities.Post {
    return &entities.Post{
        ID:      dto.ID,
        Title:   dto.Title,
        Content: dto.Content,
    }
}
```

## Ciclo de Vida da Aplicação

```
main.go
  │
  ├─► config.Load()           # Carrega configurações
  │
  ├─► persistence.Connect()   # Conecta ao banco
  │
  ├─► NewRepository()          # Instancia repositories
  │
  ├─► NewService()             # Instancia serviços (com repos injetados)
  │
  ├─► NewRouter()              # Configura rotas (com serviços injetados)
  │
  └─► router.Run()             # Inicia servidor
```

## Exemplo de Inicialização

```go
// cmd/maycms/main.go
package main

import (
    "maycms/internal/infrastructure/config"
    "maycms/internal/infrastructure/persistence/postgres"
    "maycms/internal/infrastructure/http"
    "maycms/internal/application/services"
)

func main() {
    cfg := config.Load()
    
    db, err := postgres.Connect(cfg.Database)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Instancia repositories
    postRepo := postgres.NewPostRepository(db)
    
    // Instancia serviços
    postService := services.NewPostService(postRepo)
    
    // Configura rotas
    router := http.NewRouter(postService)
    
    // Inicia servidor
    router.Run(cfg.Server.Port)
}
```

## Testes

### Testes Unitários
- Testam camadas isoladas (domain, application)
- Usam mocks para dependências
- Devem ter cobertura > 80%

### Testes de Integração
- Testam fluxos entre camadas
- Usam banco de dados real ou in-memory
- Validam comportamento end-to-end

### Testes de Aceitação
- Testam via API HTTP
- Validam respostas completas
- Simulam uso real

## Considerações de Deploy

1. **Docker**: Use multi-stage build para reduzir tamanho
2. **Configuração**: Use variáveis de ambiente, nunca hardcode
3. **Migrações**: Execute migrações antes de iniciar a aplicação
4. **Health Checks**: Implemente endpoints de health check
5. **Graceful Shutdown**: Feche conexões adequadamente

## Referências

- [Hexagonal Architecture - Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Clean Code in Go - Dustin Krysak](https://www.youtube.com/watch?v=sO0HzcjyGiM)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

---

**Versão**: 1.0  
**Última Atualização**: 2025-05-25  
**Mantido por**: @mgarcia88
