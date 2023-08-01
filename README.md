### Clean Architecture Boilerplate(Golang)

A Golang project template adhering to Clean Architecture and Hexagonal Architecture principles. It offers a well-organized structure for scalable applications, separating business logic from infrastructure. Docker integration facilitates seamless deployment and scalability. With modularity, components can be developed and tested independently, ensuring flexibility and maintainability.

##### To run the service, execute the command:

```
make service
```

##### To stop the service, execute the command:

```
make down
```

- NOTE: you can customize the dockerized items on your own

#### The layers' description:

##### App:

- This layer handles the bootstrap process of the service, including setting up HTTP routes, HTTP middlewares, RPC entities, modules, and other components necessary for the service to function.

##### Cmd:

- This layer handles custom commands which could be(is offered) handled by the COBRA CLI.

##### Config:

- This layer handles any type of config files like those supported by the VIPER package. Also, any other config files like locale contain in this layer.

##### Database:

- This layer handles SQL queries (migration, tables alternatives, etc.) or other database configs which should be compatible with the GORM package.

##### Driver

- This layer handles any external functionalities, packages, or third-party services, should be implemented here, and should be added to the APP layer in CORE and APP files to be bootstrapped.

`NOTE: the driver layer items, could be wrapped as the custom packages and be used in the APP layer as the abstraction and interfaces.`

##### Helper

- This layer provides utility functions and reusable code snippets that assist other layers in the application.

##### Module

- This layer handles business logic and data manipulation. Any Module contains the layers of `abstraction`, `entity`, `repository`, and `usecase`, that finally bind to the delivery layer. Then the delivery layer also binds to HTTP or RPC wrapper in the `APP` layer by the `init` file which exists in any module.

