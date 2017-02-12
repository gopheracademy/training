name: Useful Interfaces
description: Useful Interfaces
level: Intermediate
topic: Go
class: center, middle
module-time: 45
# Creating Useful Interfaces


---
name: Useful Interfaces

# Useful Interfaces

Brian Ketelsen
bketelsen@gopheracademy.com
@bketelsen

---
name: Creating Useful Interfaces

# Creating Useful Interfaces 


In this section we’ll look at Go’s interfaces and learn how to use them for maximum flexibility and code reuse.  We’ll single out several interfaces from the standard library and explore what makes them so powerful.  You’ll learn how to make your abstractions count with small and concise interfaces.

- What Makes a Good Interface
- Standard Library Interface Examples
- Interface Size versus Interface Utility

---
name: What Makes a Good Interface

# What Makes a Good Interface

- Interfaces should represent behaviors
- Interfaces should wrap the behavior of a dependency
- Good interfaces represent the smallest behavior possible

---
name: Philosophy: Go and Interfaces

# Philosophy: Go and Interfaces

Proverb:
	The bigger the interface, the weaker the abstraction

> Rob_Pike_ at [Gopherfest SV, 2015](https://www.youtube.com/watch?v=PAAkCSZUG1c)

---
name: What Makes a Good Interface

# What Makes a Good Interface

To define your interfaces, start by defining the behaviors of your types.

<code src="usefulinterfaces/demos/inventory/product.go"> 

---
name:What Makes a Good Interface

# What Makes a Good Interface

An interface should represent one type of behavior.  In the previous example it was _Storage_.
An interface should not mix different types behaviors.  Don't have a _ProductInterface_ that mixes _Storage_ with _Display_ for example.


---
name:Practical Interfaces

# Practical Interfaces

We're going to use an Inventory application as an example that models something that you might write for your business.

*BusinessDomainInformation*

- Products -- things we sell
- Suppliers -- vendors that provide Products
- Orders -- requests from us to our supplier to get more Products

- We store our data in *postgres*
- We have one Supplier *Acme*, with an HTTP API for placing orders

---
name:Practical Interfaces - Dependencies

# Dependencies

Our fictional inventory application has two data sources, a PostgreSQL database and a vendor API.

Suppliers, Products, and Orders are stored in PostgreSQL, but the interactions with the supplier are done through the supplier's HTTP API.

That means we have two dependencies which should be in their own packages: 

- postgres
- acme

---
name:Practical Interfaces - Dependencies

# Dependencies

The *postgres* package will have implementations of the domain interfaces for each of the domain types that are stored in the PostgreSQL database.

The *acme* package will have an implementation of the domain interface for interacting with the supplier.

---
name:Transport Mechanisms

# Transport Mechanisms

For compatibility with other systems, the inventory application will offer consumers of the data two choices in network transport:

- HTTP/REST API
- net/rpc calls 

Therefore we'll have a package *transport* containing an *rpc* package and an *http* package.

---
name:Interface

# Interface

PRO TIP:

- Use https://github.com/josharian/impl to generate your implementation stubs
```shell
go get github.com/josharian/impl

USAGE:
impl 'r *Receiver' gopath/to/pkg.Interface

Example:
impl 's *Source' golang.org/x/oauth2.TokenSource > mocks/token.go
```

To make it easy to isolate and test each individual package you should create mock implementations of each Interface you define in the domain package.  Keep all your mock implementations in a *mocks* package in the root of your application directory. 

*impl* requires that the package with the interfaces can compile.  Make sure you have a compiling package before using it.
---
name:Exercise One

# Exercise One 


In the *usefulinterfaces/exercises* directory there's a folder called *impl*
```shell
cd impl
cat interface.go
```
Now install the *impl* tool:
```shell	
go get github.com/josharian/impl
```
Try it out:
```shell
impl 'c *CustomerService' github.com/gopheracademy/training/usefulinterfaces/exercises/impl.CustomerGetter
```

---
name:Exercise One

# Exercise One

You should see output like this:

<code src="usefulinterfaces/includes/impl/implout.txt">
---
name:Exercise One

# Impl

That's very useful for both implementing the actual service and implementing the mocks.  But it outputs to *stdout*.  Let's make a postgres implementation of the CustomerGetter service:
```shell
	impl 'c *CustomerService' {full gopath}/impl.CustomerGetter > postgres/customer.go
```
The postgres folder already exists, if it didn't you'd have to create it first.

Now make an implementation for a mock of the CustomerGetter:
```shell
mkdir mocks
impl 'c *CustomerService' {full gopath}impl.CustomerGetter > mocks/customer.go
```
*impl* generates the methods for a type that doesn't exist.  To finish the implementation, you need to add a type called *CustomerService*.
```go
	type CustomerService struct{}
```
Now your code should compile.

---
name:Example Application

# Example Application
```shell
	$ cd $GOPATH/src/github.com/gopheracademy/training/usefulinterfaces/demos/inventory
```
In the root package *inventory* there are three domain objects:

- *Order* in  order.go
- *Product* in product.go
- *Supplier* in supplier.go

---
name:Practical Interfaces

# Practical Interfaces

The *inventory* package also defines five interfaces:

- *OrderStorage* in  order.go
- *ProductStorage* in product.go
- *SupplierStorage* in supplier.go
- *SupplierService* in supplier.go
- *InventoryService* in inventory.go

---
name:Storage Interfaces

# Storage Interfaces

Our application has *storage* interfaces, so there are (incomplete) implementations of those *storage* interfaces in the *postgres* package.

<code src="usefulinterfaces/demos/inventory/postgres/orders.go">

---
name:Supplier Service

# Supplier Service
Similarly there's a *SupplierService* that represents the interactions over an external API to our supplier "ACME Widgets".  The implementation of that service is in the *acme* package.

<code src="usefulinterfaces/demos/inventory/acme/client.go"> 

---
name:Service Interface

# Service Interface

The *Service* interface defines the external API of our application to its consumers.  

<code src="usefulinterfaces/demos/inventory/inventory.go">

---
name:Transport

# Transport

We have customers who demand a REST style interface over HTTP, and we also have internal users in our own company who can use something faster like Go's *net/rpc*

It makes sense to create a new interface to represent the delivery of our service over any type of connection.  In the *transport* package, there's an interface just for this:


<code src="usefulinterfaces/demos/inventory/transport/transporter.go">

---
name:Transport Implementations

# Transport Implementations

Now we can have transport specific implementations in packages under *transport*:

- *http*
- *rpc*

Let's browse through those two implementations quickly.  

Of note, they both implement the *transport.InventoryTransporter* interface, which *embeds* the *inventory.Service*.  

Now our *http* and *rest* implementations of the *transport.InventoryTransporter* interface are required to also implement the *inventory.Service* interface.


* Practical Interfaces

PROTIP:

Prove that your type implements an interface at compile time with this declaration:

<code src="usefulinterfaces/demos/inventory/transport/http/rest.go">

This variable declaration assigns a concrete type *RESTService* to the interface it should implement, then discards the result.

By doing this, your code won't compile if the *RESTService* doesn't implement the *transport.InventoryTransporter* interface. 

- Type Safety -- even when using Interfaces for abstraction


---
name:Tie It Together

# Tie It Together

Now with interfaces all defined and (mostly) implemented, we can tie it all together in *main.go*

`src/github.com/gophertrain/material/usefulinterfaces/demos/inventory/cmd/main.go`

Because the network services take interfaces, we can easily swap out storage drivers when changing databases, or more commonly, for testing.

---
name:Behaviors And Dependencies

# Behaviors And Dependencies

Every interface in this application represents a behavior that is tied to a dependency.

- Storage: postgres
- Transport: http & rpc
- SupplierService: vendor API

With clearly defined interfaces, we can create new implementations of behaviors as business requirements change.

- Migrate to mysql? - implement the Storage interfaces
- New supplier?  Implement the SupplierService interface
- New service delivery method?  Implement the Transport interface

---
name:Benefits of Good Interfaces

# Benefits of Good Interfaces

Your domain types (Order, Supplier, Product, etc) don't know or *care* how they're stored or delivered.

- No circular dependencies
- Easy testing with mocks that implement the interfaces

---
name:stdlib Interface Examples

# Standard Library Interface Examples

Now let's explore some good interfaces in Go's Standard Library.

---
name:encoding/TextMarshaler

# encoding/TextMarshaler

*encoding/TextMarshaler*

[encoding/TextMarshaler](https://sourcegraph.com/github.com/golang/go/-/info/GoPackage/encoding/-/TextMarshaler)

---
name:crypto/Signer

# crypto/Signer

*crypto/Signer*

[crypto/signer](https://sourcegraph.com/github.com/golang/go/-/info/GoPackage/crypto/-/Signer)

---
name:database/sql

# database/sql

- `database/sql` may be the best example of many interfaces used to represent a single aggregate dependency

*database/sql/Driver*

[database/sql/Driver](https://godoc.org/database/sql/driver)

---
name:database/sql

# Interface Power

The *REAL* power of these interfaces is that many of them are so inter-related.

*sql/driver.Rows* is an iterator over a query's results

*rows.Next(dest*[]Value)* populates a slice of *sql/driver.Value*

*sql/driver.Value* represents base types that any *sql/driver* must be able to retrieve and store.

Convert to and from a *sql/driver.Value* by implementing the *sql/driver.Valuer* and *sql/driver.ValueConverter* interfaces.

---
name:Size vs Utility

# Interface Size vs Interface Utility

Rob Pike's quote at the beginning of the module stated clearly that the larger the interface, the weaker the abstraction.

Interface Guidelines:

- Start big 

	Abstract out your dependencies first

- Use interfaces for ONE and only ONE behavior

- Don't get sucked into Interface hell!

	Create interfaces only where multiple types can implement the same behavior
	*AND*
	It is necessary to use that behavior in multiple places without coupling the
	types that implement the behavior.



---
name:Exercise

# Exercise

*usefulinterfaces/exercises/inventory* contains a *mocks* folder.

Create a file in the mocks directory that implements the OrderStorage interface.

- OrderStorage = mocks/orders.go 

Hint: use a type that has a map[int]*DomainModel  (e.g. map[int]*inventory.Order)

Bonus: Implement Supplier and Product

Double Bonus Round: implement a mock for SupplierService in mocks/acme.go
