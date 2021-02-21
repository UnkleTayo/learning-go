# Interface

In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

Interface specifies what methods a type should have and the type decides how to implement these methods.

If the interface type contains only one method, the name of the interface ends with the er suffix. This is a common naming convention in Go


For a user-defined type to implement an interface, the type in question needs to
implement all the methods that are declared within that interface type
