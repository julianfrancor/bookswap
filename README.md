# BookSwap

BookSwap is a project created by Julian Franco to apply Domain Driven Design (DDD) principles and enhance knowledge in the Go programming language. This project is designed to showcase the implementation of DDD concepts in building a simple book exchange application.

## Domain Driven Design (DDD)

**Domain Driven Design** is an approach to software development that places a strong emphasis on understanding and modeling the business domain. It involves collaboration between technical and domain experts to create a shared understanding of the problem domain, which is then reflected in the software design. DDD helps in building complex applications by focusing on the core business logic and creating a clear and maintainable codebase.

Key concepts of DDD include:

- **Ubiquitous Language**: Establishing a common, shared language between developers and domain experts to ensure clear communication and understanding of the domain.

- **Bounded Contexts**: Defining explicit boundaries within which a particular model is defined and applicable. This helps in avoiding ambiguity and conflicts in the domain model.

- **Entities and Value Objects**: Distinguishing between entities, which have a unique identity, and value objects, which are defined by their attributes. This helps in modeling the domain more accurately.

- **Aggregates and Aggregates Roots**: Grouping related entities and value objects into aggregates, with one entity designated as the aggregate root responsible for ensuring the consistency of the aggregate.

## BookSwap App

BookSwap is a simple book exchange application that allows users to interact with the system to manage books and users. The application is built with the goal of demonstrating how DDD principles can be applied to create a well-structured, maintainable, and scalable codebase.

### Components

- **cmd**: Contains the main application entry point and handler functions.
- **internal/application**: Implements application services and request/response structures.
- **internal/domain**: Defines the core domain entities and business logic.
- **internal/infrastructure/persistence**: Provides in-memory repositories for storing domain entities.

### How DDD is Applied

- **Ubiquitous Language**: The codebase reflects a shared language used by both developers and domain experts. Entities like `Book` and `User` directly map to their corresponding concepts in the business domain.

- **Bounded Contexts**: Clear boundaries are established between different components (e.g., `application`, `domain`, `persistence`) to ensure separation of concerns and maintainability.

- **Entities and Value Objects**: The distinction between entities (e.g., `Book`, `User`) and value objects is clearly defined. For example, `BookID` is a value object within the `Book` entity.

- **Aggregates and Aggregates Roots**: The concept of aggregates is demonstrated, with `User` being the aggregate root for the collection of `Book` entities associated with a user.

### Running the Application

To run the BookSwap application, follow these steps:

1. Clone the repository: `git clone https://github.com/julianfrancor/bookswap.git`
2. Change directory to the project folder: `cd bookswap`
3. Run the application: `go run cmd/main.go`

### Testing

To run tests for different components:

- For application services: `go test ./internal/application/...`
- For domain entities and logic: `go test ./internal/domain/...`
- For infrastructure/persistence: `go test ./internal/infrastructure/persistence/...`

Feel free to explore the codebase, experiment with DDD concepts, and use this project as a reference for your own DDD-based applications.

Note: This README provides an overview of the project and its main aspects. As the project evolves, more details and additional documentation will be added as needed.