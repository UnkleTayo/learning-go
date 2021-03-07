# Interface

Interfaces are types that just declare behavior. This behavior is never implemented by the interface type directly but instead by user-defined types via methods.

In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

Interface specifies what methods a type should have and the type decides how to implement these methods.

If the interface type contains only one method, the name of the interface ends with the er suffix. This is a common naming convention in Go


For a user-defined type to implement an interface, the type in question needs to implement all the methods that are declared within that interface type

 Unlike when you call methods directly from values and pointers, when you call a method via an interface type value, the rules are different. Methods declared with pointer receivers can only be called by interface type values that contain pointers.
 
Methods declared with value receivers can be called by interface type values that contain both values and pointers.
