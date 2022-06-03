# go-design-patterns
all 24 GOF design patterns in go

#### Creational design patterns
| Name                                                                                                                  | Purpose                                                                                                                                           | Status |
|-----------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|--------|
| [Singleton](https://github.com/ishankhare07/go-design-patterns/tree/master/internal/pkg/creational/singleton)         | Having a unique instance of a type in the entire program                                                                                          | ✔      |
| [Builder](https://github.com/ishankhare07/go-design-patterns/tree/master/internal/pkg/creational/builder)             | Reusing an algorithm to create many implementations of an interface                                                                               | ✔      |
| [Factory Method](https://github.com/ishankhare07/go-design-patterns/tree/master/internal/pkg/creational/factory)      | Delegating the creation of different types                                                                                                        | ✔      |
| [Abstract factory](https://github.com/ishankhare07/go-design-patterns/tree/master/internal/pkg/creational/absfactory) | A factory of factories                                                                                                                            | ✔      |
| [Prototype](https://github.com/ishankhare07/go-design-patterns/tree/master/internal/pkg/creational/prototype)         | Have an object or set of objects that is already created at compile time. Serves as default template and can be cloned as many times as required. | ✔      |

#### Structural Patterns
| Name      | Purpose                                                                                                                   | Status |
|-----------|---------------------------------------------------------------------------------------------------------------------------|--------|
| Composite | Has a relationship vs is a relationship                                                                                   | ✘      |
| Adapter   | Allow us to use something that wasn't built for a specific task                                                           | ✘      |
| Bridge    | Decouple and abstraction from its implementation                                                                          | ✘      |
| Proxy     | Wrap an object to hide some of its characteristics                                                                        | ✘      |
| Facade    | Shield code from unwanted access and hides complexity scope from the user. A group of proxies can be considered a facade. | ✘      |
| Decorator | Decorate and already existing type with more functional features without actually touching it                             | ✘      |
| Flyweight | Allows sharing the state of a heavy object between many instances of some type                                            | ✘      |
