# Java Design Patterns Library

A comprehensive collection of **Design Patterns** implemented in Java. This repository serves as a practical reference for understanding the "Gang of Four" (GoF) patterns and other modern architectural styles through clean, documented, and executable code.

![Java](https://img.shields.io/badge/Java-ED8B00?style=for-the-badge&logo=openjdk&logoColor=white)
![JDK](https://img.shields.io/badge/JDK-21-ED8B00?style=for-the-badge&logo=openjdk&logoColor=white)

---

## Project Structure

Each pattern is a standalone file that can be run directly using JDK:
`design-patterns-java/[category]/[pattern_name]`

- **Interfaces**: Define the contracts for the patterns.
- **Concrete Classes**: Real-world implementations of the logic.
- **Demo/App**: A class with `main` method in every package to demonstrate the pattern in action.

---

## Patterns Catalog

### 1. Creational Patterns
*Deals with object creation mechanisms, trying to create objects in a manner suitable to the situation.*

| Pattern | Goal |
| :--- | :--- |
| **Singleton** | Ensure a class has only one instance with global access. |
| **Factory Method** | Delegate instantiation logic to subclasses. |
| **Abstract Factory** | Produce families of related objects without specifying concrete classes. |
| **Builder** | Construct complex objects step-by-step. |
| **Prototype** | Create new objects by cloning an existing one. |

### 2. Structural Patterns
*Deals with class and object composition.*

| Pattern | Goal |
| :--- | :--- |
| **Adapter** | Bridge incompatible interfaces. |
| **Decorator** | Attach new behaviors to objects dynamically. |
| **Facade** | Provide a simplified interface to a complex subsystem. |
| **Proxy** | Provide a placeholder for another object to control access. |

### 3. Behavioral Patterns
*Deals with communication between objects and the assignment of responsibilities.*

| Pattern | Goal |
| :--- | :--- |
| **Observer** | Notify multiple objects about state changes (Pub/Sub). |
| **Strategy** | Switch between different algorithms at runtime. |
| **State** | Allow an object to change behavior when its internal state changes. |
| **Command** | Encapsulate a request as an object. |

---

## Getting Started

### Prerequisites
* **JDK 21** or higher (Or might work in 1.8 I dont really know).
* **RTX 5090** (optional, if you want to play video games).

### How to Run
To run a specific example (e.g., the Observer pattern), navigate to the root and use:

```bash
# Using JDK
java [category]/[pattern_name].java