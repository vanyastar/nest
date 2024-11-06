
# NestGo

NestGo is a powerful and flexible framework for building scalable server-side applications with Golang. Inspired by [NestJS](https://nestjs.com/), NestGo provides a modular architecture, dependency injection, and a rich set of features, making it ideal for building robust, maintainable, and production-ready applications.

## Key Features

- **Modular Architecture**: Organize your code in modules, services, and controllers to ensure scalability and maintainability.
- **Dependency Injection**: Simplify dependency management and testing with built-in dependency injection.
- **Decorator-based Routing**: Define routes using decorators for a clean and intuitive API structure.
- **Middleware Support**: Easily add custom middleware for request handling.
- **Guards, Interceptors, and Validation**: Enforce security, handle requests/responses, and validate input seamlessly.
- **Customizable**: Extensible design allows integration with a wide range of libraries and tools.
- **Built with Go's Performance in Mind**: Fast, efficient, and optimized for high-performance applications.

---

## Installation

To get started with NestGo, install the framework using Go modules:

```bash
go get github.com/yourusername/nestgo
```

---

## Project Structure

A typical `NestGo` project is structured as follows:

```
my-nestgo-app/
├── main.service..go               # Main entry point for the application
├── hello/
│   ├── hello.controller.go # Controller with route handlers
│   ├── hello.service.go  # Service containing business logic
│   ├── dto/              # Data Transfer Objects (DTOs) for request validation
│   ├── validators/       # Validators
│   └── guards/           # Guards for route protection
├── common/
│   └── middlewares.go    # Example custom middleware
└── config/
    └── config.go         # Configuration settings
```

## Core Concepts

### Controllers
Controllers handle incoming requests and return responses to the client. NestGo uses decorators to simplify route handling and make it more readable.

### Services
Services contain the business logic of your application and are injected into controllers as dependencies.

### Middleware
Middleware functions are executed before the request reaches the controller, making them ideal for logging, authentication, or modifying the request object.

### Guards
Guards are used to enforce authorization rules, deciding if a request should proceed to the next stage or be blocked.

### Interceptors
Interceptors transform the request or response, making it easier to apply logic like response formatting or error handling globally.

### DTO (Data Transfer Objects)
DTOs define the structure of data sent to and from the server. DTOs are used for validating and transforming incoming data in the controllers.

### Validation
NestGo includes built-in validation using DTOs, ensuring that incoming requests adhere to the specified structure and constraints.

---

## Quick Start

Follow these steps to create a new NestGo project:

1. **Initialize Your Project**

   ```bash
   mkdir my-nestgo-app
   cd my-nestgo-app
   go mod init my-nestgo-app
   ```

2. **Install NestGo**

   ```bash
   go get github.com/vanyastar/nest
   ```

3. **Create Your First Module**

   Organize your application by creating modules, controllers, and services.

---

## License

NestGo is open source and available under the [MIT License](LICENSE).

---

This README includes a file structure and setup details to help you get started with NestGo. For more information, check out the documentation or reach out with any questions!
