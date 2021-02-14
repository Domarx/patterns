```text
In object-oriented programming, the decorator pattern is a design pattern that allows behavior to be added to an
individual object, dynamically, without affecting the behavior of other objects from the same class. The decorator
pattern is often useful for adhering to the Single Responsibility Principle, as it allows functionality to be divided
between classes with unique areas of concern. Decorator use can be more efficient than subclassing, because an objects
behavior can be augmented without instantiating an entirely new object.

Which problems it solves?

    [x] Responsibilities should be added to (and removed from) an object dynamically at run-time.

    [x] A flexible alternative to subclassing for extending functionality should be provided.

Solution:

    [x] Implement interface of the extended (decorated) object transparently by forwarding all requests to it

    [x] Perform additional functionality before/after forwarding a request
```