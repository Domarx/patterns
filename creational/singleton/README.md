```text
Singleton pattern is a software design patter that restricts the instantiation of a class to one "single" instance.
This is useful when exactly one object is needed to coordinate actions across the system.

Common uses:
    
    [x] The abstract factory, factory method, builder and prototype can use singletons in their implementations.
    
    [x] Facade objects are often singletons because only one facade object is required.
    
    [x] State objects are often singletons
    
    [x] Singeltons are often preferred to global variables because:
    
        [x] They do not pollute the global namespace with unnecessary variables
        
        [x] They permit lazy allocation and initialization, whereas global variables in many languages will always 
            consume resources.
```