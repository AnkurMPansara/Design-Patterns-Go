# Structural Design Patterns

Kind of meh section, just read this if your OCD is triggered or you have too much free time, in which case play video games or something. If Creational patterns are about *making* the bricks, Structural patterns are about *stacking* them without the whole building falling over. These patterns help you realize that inheritance isn't always the answer (spoiler: it usually isn't) and that **composition** is your best friend. 

Basically, this section is about how to wrap objects inside other objects so your code doesn't look like a plate of spaghetti.

---

## Included Patterns

| Pattern | Summary | The "Real Talk" Use Case |
| :--- | :--- | :--- |
| **[Adapter](./adapter)** | Converts one interface to another. | When you bought a third-party library that doesn't fit your code's "plug." |
| **[Bridge](./bridge)** | Decouples an abstraction from its implementation. | When you have a "Product" and "Color" and don't want to create 500 subclasses. |
| **[Composite](./composite)** | Treats individual objects and compositions uniformly. | When you're building a Tree structure (like a File System or UI components). |
| **[Decorator](./decorator)** | Adds behavior to objects without inheritance. | When you want to add "features" (like Scrollbars or Encryption) dynamically. |
| **[Facade](./facade)** | Provides a simple entry point to a complex mess. | When a library has 50 classes but you only need to call `doWork()`. |
| **[Flyweight](./flyweight)** | Shares state to support large numbers of small objects. | When your RAM is crying because you created 1,000,000 Tree objects. |
| **[Proxy](./proxy)** | Provides a placeholder to control access. | Lazy loading or adding a security layer before hitting the real object. |

---

## Why use Structural Patterns?

1. **Composition > Inheritance**: These patterns help you follow the "Favor composition over inheritance" rule, which stops your class hierarchy from becoming a recursive nightmare.
2. **Maintenance**: You can change how parts of the system interact without breaking the objects themselves.
3. **Efficiency**: Patterns like Flyweight or Proxy can literally save your server from crashing under heavy memory load.

---

## How to Run Examples

Assuming you've actually set up your Java environment and aren't just "vibing" through a tutorial, run these from the root:

```bash
java structural/[patternName].java