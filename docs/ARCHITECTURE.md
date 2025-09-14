# ARCHITECTURE.md

## Event-Driven Architecture with Repository Pattern

### Current State Analysis

The existing monolithic application handles:
- Daily Bible quiz generation via LLM
- User management and authentication
- Quiz attempt tracking and scoring
- Leaderboard management
- Question/Answer content management

**Key Issues to Address:**
- Tight coupling between HTTP handlers and database
- No event-driven patterns for user actions
- No separation of concerns
- Single point of failure architecture

---

## Proposed Architecture

### High-Level Architecture Overview

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   API Gateway   │────│  Service Bus     │────│   Database      │
│  (Azure APIM)   │    │ (Azure Service   │    │ (Azure SQL)     │
└─────────────────┘    │     Bus)         │    └─────────────────┘
                       └──────────────────┘
                              │
                ┌─────────────┼─────────────┐
                │             │             │
        ┌───────▼───────┐ ┌───▼────┐ ┌─────▼─────┐
        │ User Service  │ │ Quiz   │ │Leaderboard│
        │   (Go)        │ │Service │ │ Service   │
        └───────────────┘ │ (Go)   │ └───────────┘
                         └────────┘
```

### Core Design Pattern: Repository Pattern

The Repository Pattern provides a clean abstraction layer between business logic and data access:

```go
// Domain Repository Interface
type UserRepository interface {
    Save(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id UserID) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
}

type QuizRepository interface {
    Save(ctx context.Context, quiz *Quiz) error
    FindDailyQuiz(ctx context.Context, date time.Time) (*Quiz, error)
    FindByID(ctx context.Context, id QuizID) (*Quiz, error)
}

type AttemptRepository interface {
    Save(ctx context.Context, attempt *Attempt) error
    FindByUserAndQuiz(ctx context.Context, userID, quizID string) (*Attempt, error)
}
```

---

## Azure Cloud Architecture

### Core Azure Services

- **Azure Container Apps** - Host microservices
- **Azure Service Bus** - Event messaging between services
- **Azure SQL Database** - Primary data storage
- **Azure OpenAI** - Quiz generation (replace local Ollama)
- **Azure AD B2C** - User authentication

---

## Microservices Architecture

### 1. User Service
- User registration and profile management
- Uses Repository pattern for data access
- Publishes `UserRegistered` events

### 2. Quiz Service
- Daily quiz generation via Azure OpenAI
- Question/answer management
- Publishes `DailyQuizGenerated` events

### 3. Attempt Service
- Quiz attempt submission and validation
- Answer correctness evaluation
- Publishes `QuizAttemptEvaluated` events

### 4. Leaderboard Service
- Score aggregation and ranking
- Listens to `QuizAttemptEvaluated` events
- Updates user rankings

---

## Event-Driven Implementation

### Domain Events

```go
// Core Events
type UserRegistered struct {
    UserID    string    `json:"user_id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Timestamp time.Time `json:"timestamp"`
}

type DailyQuizGenerated struct {
    QuizID      string    `json:"quiz_id"`
    Question    string    `json:"question"`
    Answer      string    `json:"answer"`
    Date        string    `json:"date"`
    Timestamp   time.Time `json:"timestamp"`
}

type QuizAttemptEvaluated struct {
    AttemptID   string    `json:"attempt_id"`
    UserID      string    `json:"user_id"`
    QuizID      string    `json:"quiz_id"`
    IsCorrect   bool      `json:"is_correct"`
    Score       int       `json:"score"`
    Timestamp   time.Time `json:"timestamp"`
}
```

### Event Flow

1. **User Registration**: User Service → `UserRegistered` → Leaderboard Service
2. **Daily Quiz**: Quiz Service → `DailyQuizGenerated` → All Services
3. **Quiz Attempt**: Attempt Service → `QuizAttemptEvaluated` → Leaderboard Service

---

## Service Structure with Repository Pattern

### Service Layer Structure
```go
// Service with Repository Injection
type UserService struct {
    userRepo   UserRepository
    eventBus   EventBus
}

func (s *UserService) RegisterUser(ctx context.Context, username, email string) error {
    user := &User{
        ID:       generateID(),
        Username: username,
        Email:    email,
    }
    
    // Save via repository
    if err := s.userRepo.Save(ctx, user); err != nil {
        return err
    }
    
    // Publish event
    event := UserRegistered{
        UserID:    user.ID,
        Username:  user.Username,
        Email:     user.Email,
        Timestamp: time.Now(),
    }
    
    return s.eventBus.Publish(ctx, "user.registered", event)
}
```

### Repository Implementation
```go
type SQLUserRepository struct {
    db *sql.DB
}

func (r *SQLUserRepository) Save(ctx context.Context, user *User) error {
    query := `INSERT INTO users (id, username, email) VALUES (?, ?, ?)`
    _, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Email)
    return err
}

func (r *SQLUserRepository) FindByID(ctx context.Context, id UserID) (*User, error) {
    // Implementation details...
}
```

---

## Implementation Steps

### Phase 1: Repository Pattern (Week 1-2)
1. Create repository interfaces for User, Quiz, Attempt
2. Implement SQL-based repositories
3. Refactor existing handlers to use repositories
4. Add dependency injection

### Phase 2: Event-Driven Architecture (Week 3-4)  
1. Set up Azure Service Bus
2. Create event publishing mechanism
3. Implement event handlers in each service
4. Test event flow between services

### Phase 3: Microservices Split (Week 5-6)
1. Split monolith into separate services
2. Deploy to Azure Container Apps
3. Configure inter-service communication
4. Add health checks and monitoring

---

## Benefits

- **Clean Code** - Repository pattern separates data access from business logic
- **Scalability** - Event-driven architecture allows independent service scaling  
- **Maintainability** - Clear boundaries between services and responsibilities
- **Testability** - Easy to mock repositories and test business logic in isolation