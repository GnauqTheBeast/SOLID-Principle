## Dependency Inversion Principle (DIP)

The **Dependency Inversion Principle (DIP)** is one of the core principles of object-oriented design, and it ensures that high-level modules do not depend on low-level modules but on abstractions. This helps achieve a flexible and maintainable system.

### Benefits of Following DIP

1. **Loose Coupling**:
   - The `UserService` class depends on the `Database` interface, not on specific implementations like `MySQLDatabase` or `PostgresDatabase`.
   - This allows changing the underlying database implementation without modifying the `UserService`.

2. **Easy to Extend**:
   - Adding a new database system (like MongoDB) only requires creating a new struct that implements the `Database` interface.
   - No changes are needed in the existing `UserService` code, keeping it **open for extension but closed for modification** (Open/Closed Principle).

3. **Better Testing**:
   - The use of interfaces allows easy mocking of the `Database` in unit tests.
   - This decouples the tests from the actual database implementation, making tests faster and more reliable.

4. **Adheres to the Open/Closed Principle**:
   - The `UserService` is open for extension (new database types can be added) but closed for modification (no need to change existing code).
   - This ensures that adding new features does not break existing functionality.
