# **Object Oriented Programming**

## **Classes and Objects**

Classes and Objects are basic concepts of Object Oriented Programming which revolve around the real life entities.

### **Class**

A class is a user defined blueprint or prototype from which objects are created. It represents the set of properties or methods that are common to all objects of one type.

### **Object**

It is a basic unit of Object-Oriented Programming and represents the real life entities. A typical Java program creates many objects, which as you know, interact by invoking methods. An object consists of :

1. **State**: It is represented by attributes of an object. It also reflects the properties of an object.
2. **Behavior**: It is represented by methods of an object. It also reflects the response of an object with other objects.
3. **Identity**: It gives a unique name to an object and enables one object to interact with other objects.

## **Access Modifiers**

1. **Default – No keyword required**:

   - Can be accessed in `same class`, `same package subclass` and `same package non-subclass`

2. **Private**:

   - Can accessed in `same class`.

3. **Protected**:

   - Can be accessed in `same class`, `same package subclass`, `same package non-subclass` and `different package subclass`

4. **Public**:

   - Can be accessed in `same class`, `same package subclass`, `same package non-subclass`, `different package subclass` and `different package non-subclass`

## **Non-Access Modifiers**

1. **static**:
   - The static keyword means that the entity to which it is applied is `available outside` any particular instance of the class. That means the static methods or the attributes are a part of the class and not an object. The memory is allocated to such an attribute or method at the time of class loading. The use of a static modifier makes the program more efficient by saving memory. A static field exists across all the class instances, and without creating an object of the class, they can be called.
2. **final**:
   - The final keyword indicates that the specific class `cannot be extended` or a method cannot be overridden.
3. **abstract**:
   - Abstract keyword is used to declare a class as `partially implemented` means an object cannot be created directly from that class. Any subclass needs to be either `implement` all the methods of the abstract class, or it should also need to be an abstract class. The abstract keyword cannot be used with static, final, or private keywords because they prevent overriding, and we need to override methods in the case of an abstract class.
4. **synchronized**:
   - Synchronized keyword `prevents` a block of code from executing by `multiple threads at once`. It is very important for some critical operations.
5. **volatile**:
   - The volatile keyword is used to make the class thread-safe. That means if a variable is declared as volatile, then that can be modified by multiple threads at the same time without any issues. The volatile keyword is only applicable to a variable. A volatile keyword reduces the chance of memory inconsistency. The value of a volatile variable is always read from the main memory and not from the local thread cache, and it helps to improve thread performance. Let us understand by an example:
6. **transient**:

   - This needs prior knowledge of serialization in Java. You can refer to the following article for that:- serialization in java.

   - The transient keyword may be applied to member variables of a class to indicate that the member variable should not be serialized when the containing class instance is serialized. Serialization is the ​process of converting an object into a byte stream. When we do not want to serialize the value of a variable, then we declare it as transient. To make it more transparent, let’s take an example of an application where we need to accept UserID and Password. At that moment, we need to declare some variable to take the input and store the data, but as the data is susceptible, so we do not want to keep it stored after the job is done. To achieve this, we can use the transient keyword for variable declaration. That particular variable will not participate in the serialization process, and when we deserialize that, we will receive the default value of the variable.

7. **native**:
   - The native keyword may be applied to a method to indicate that the method is implemented in a language other than Java. Using this java application can call code written in C, C++, or assembler language. A shared code library or DLL is required in this case.

## **Pillars of OOPS**
