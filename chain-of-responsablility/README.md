## Chain Of Responsability Design Pattern

### Definition
Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it along the chain.

### Problem
Imagine that you’re working on an online ordering system. You want to restrict access to the system so only authenticated users can create orders. Also, users who have administrative access should have additional permissions, such as the ability to cancel existing orders.

### Solution
The Chain of Responsibility pattern suggests that you build a chain of objects that can process a command. Each object in the chain either processes the command or passes it to the next object in the chain. The number of handlers and their order are decided at runtime.\

### How to implement
1. Declare the handler interface and describe the signature of a method for handling requests.
2. Create concrete handler classes that implement the handler interface. Each handler has a field for storing a reference to the next handler in the chain. In most cases, handlers are self-contained and immutable, accepting all necessary data just once via the constructor.
3. Implement the default chaining behavior inside a base handler class. This behavior should only consist of a field for storing a reference to the next handler. The base handler should delegate the work of handling a request to the next handler.
4. The client is responsible for building the actual chains, usually in the configuration stage of the application.

### Pros and Cons
#### Pros
1. You can control the order of request handling.
2. Single Responsibility Principle. You can decouple classes that invoke operations from classes that perform operations.
3. Open/Closed Principle. You can introduce new handlers into the app without breaking the existing client code.
4. You can implement different variants of processing chains and select between them at runtime.
5. The code becomes easier to extend when new request types are introduced.
6. You can add or remove responsibilities from an object at runtime.
7. You can add or remove handlers from a chain at runtime.

#### Cons
1. Some requests may end up unhandled.
2. The application may become more complicated because you’re introducing a lot of new classes.
3. You may run into difficulties debugging the code, since it’s not always clear where the request will end up.
4. The chain can introduce a lot of additional processing overhead.