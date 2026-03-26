# Creational Design Patterns

Most important pattern category, I think. If you dont want to get humiliated at system design interview, atleast memorize this section. Kind of easy section, just remember to use Factory atleast once in your code and you are golden. Creational design patterns are all about **object instantiation**. They provide various mechanisms to create objects in a way that increases flexibility and reuse of existing code, rather than instantiating objects directly with the `new` operator.

---

## Included Patterns

| Pattern | Summary | When to use? |
| :--- | :--- | :--- |
| **[Singleton](./Singleton.java)** | Restricts a class to a single instance. | When you need a global state (e.g., Database connection, Logging). |
| **[Factory Method](./Factory.java)** | Provides an interface for creating objects in a superclass. | When you don't know the exact types and dependencies of the objects beforehand. |
| **[Abstract Factory](./AbstractFactory.java)** | Produces families of related objects. | When your code needs to work with various families of related products. |
| **[Builder](./Builder.java)** | Construct complex objects step-by-step. | When you have a "telescoping constructor" problem (too many parameters). |
| **[Prototype](./Prototype.java)** | Copy existing objects without depending on classes. | When the cost of creating a new object is higher than cloning an existing one. |

---

## Why use Creational Patterns?

1. **Encapsulation**: They hide the logic of how objects are created and combined.
2. **Decoupling**: The system becomes independent of how its objects are created, composed, and represented.
3. **Flexibility**: You can swap out what gets created (e.g., switching from a `MySQL` connection to `PostgreSQL`) without changing the client code.

---

## How to Run Examples

For some reason, you did not read main README file. Dont worry, I will save you the trouble, just use command below when your terminal is set to root directory (I you dont know how to do that, lets meet again after you complete college):

```bash
java creational/[patternName].java