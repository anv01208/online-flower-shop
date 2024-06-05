# online-flower-shop
Clean Architecture
Key idea
Programmers realize the optimal architecture for an application after most of the code has been written.

A good architecture allows decisions to be delayed to as late as possible.

The main principle
Dependency Inversion (the same one from SOLID) is the principle of dependency inversion. The direction of dependencies goes from the outer layer to the inner layer. Due to this, business logic and entities remain independent from other parts of the system.

So, the application is divided into 2 layers, internal and external:

Business logic (Go standard library).
Tools (databases, servers, message brokers, any other packages and frameworks).
The inner layer with business logic should be clean. It should:

Not have package imports from the outer layer.
Use only the capabilities of the standard library.
Make calls to the outer layer through the interface (!).
The business logic doesn't know anything about Postgres or a specific web API. Business logic has an interface for working with an abstract database or abstract web API.

The outer layer has other limitations:

All components of this layer are unaware of each other's existence. How to call another from one tool? Not directly, only through the inner layer of business logic.
All calls to the inner layer are made through the interface (!).
Data is transferred in a format that is convenient for business logic (internal/entity).

![image](https://github.com/anv01208/online-flower-shop/assets/145286694/505be0ad-5d80-4ad9-a563-a7a9bfb2e629)
