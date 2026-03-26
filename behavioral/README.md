# Behavioral Design Patterns

To be honest, only Strategy and Observer patterns matter, rest are either too niche or applied at low level where you dont need to touch the code if you need to read these examples. This is the "brain" of your application. If Creational is about *making* objects and Structural is about *connecting* them, Behavioral is about **how they talk to each other**. This is where you stop using `if-else` chains that are 500 lines long and start acting like a Senior Dev.

---

## Included Patterns

| Pattern | Summary | The "Real Talk" Use Case |
| :--- | :--- | :--- |
| **[Chain of Responsibility](./chain-of-responsibility)** | Passes a request along a chain of handlers. | Middleware, logging levels, or multi-stage form validation. |
| **[Command](./command)** | Turns a request into a stand-alone object. | Implementation of "Undo/Redo" buttons or task queuing. |
| **[Iterator](./iterator)** | Accesses elements of a collection without exposing its logic. | When you want to loop through a custom Data Structure without breaking it. |
| **[Mediator](./mediator)** | Restricts direct communications between objects. | Chat rooms or Air Traffic Control—stop objects from talking to everyone at once. |
| **[Memento](./memento)** | Saves and restores the previous state of an object. | Checkpoints in a game or "Save State" functionality. |
| **[Observer](./observer)** | Notifies multiple objects about events. | The classic Pub/Sub. When a YouTube creator uploads, the subscribers get hit. |
| **[State](./state)** | Lets an object change behavior when its state changes. | When your `VendingMachine` acts differently if it's `OUT_OF_STOCK`. |
| **[Strategy](./strategy)** | Switches algorithms at runtime. | Payment processing: Switch between `PayPal`, `CreditCard`, or `Crypto` instantly. |
| **[Template Method](./template-method)** | Defines the skeleton of an algorithm in a superclass. | Data miners: `Open` -> `Extract` -> `Close`. Only the `Extract` part changes. |
| **[Visitor](./visitor)** | Separates an algorithm from the object structure. | Running a "Report Generator" over a bunch of different XML nodes. |

---

## Why use Behavioral Patterns?

1. **Responsibility**: It stops your classes from trying to do everything. One class handles the logic, another handles the communication.
2. **Readability**: Instead of a giant "God Object" managing the whole app, you have small, specialized objects that know exactly what their job is.
3. **Open/Closed Principle**: You can add new behaviors (like a new `DiscountStrategy`) without touching the code that actually uses it.

---

## How to Run Examples

If you've made it this far into the repo and still haven't figured out how to run a Java file, I admire your persistence. Run this from the root:

```bash
java behavioral/[patternName].java