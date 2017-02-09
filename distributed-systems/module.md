# Getting Started with Distributed Systems

# Why Distributed Systems

- Cost Savings
  - Smaller hardware
- Scalability
  - Horizontal vs. Vertical scaling
  - Speed
- Reliability
  - Failover
  - Redundancy


name: Distributed Concepts
video: 201926132
description: 
level: Intermediate
topic: Distributed Computing
# Distributed Concepts
## Distributed Concepts

---
name: CAP Theorem
video: 
thumb:
github:
# CAP Theorum
## 

In theoretical computer science, the CAP theorem, also named Brewer's theorem after computer scientist Eric Brewer, states that it is impossible for a distributed computer system to simultaneously provide more than two out of three of the following guarantees: Consistency. Availability. Partition tolerance.

- [Consistency](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Availability](https://en.wikipedia.org/wiki/Availability)
- [Partition Tolerance](https://en.wikipedia.org/wiki/Network_partition)

#### References
[Wikipedia Cap Theorem](https://en.wikipedia.org/wiki/CAP_theorem)---
name: Gossip Protocols
video: 
thumb:
github:
# Gossip Protocols
## Gossip Protocols

### [HashiCorp member list](https://github.com/hashicorp/memberlist)

The use cases for such a library are far-reaching: all distributed systems require membership, and memberlist is a re-usable solution to managing cluster membership and node failure detection.

memberlist is eventually consistent but converges quickly on average. The speed at which it converges can be heavily tuned via various knobs on the protocol. Node failures are detected and network partitions are partially tolerated by attempting to communicate to potentially dead nodes through multiple routes.---
name: Consensus
video: 
thumb:
github:
# Consensus
## 

[ Wikipedia Consensus](https://en.wikipedia.org/wiki/Consensus_(computer_science))

- Sharing a common state across systems
- Agreeing on changes across systems
- Distributed consensus
- Raft, Paxos, Zab


name: Data Communication Protocols
video: 201926132
description: 
level: Beginner
topic: Distributed Computing
# Data Communication Protocols
## Data Communication Protocols

---
name: TLV (Type Length Value)
video: 
thumb:
github:
# TLV (Type Length Value)
## TLV (Type Length Value)

Within data communication protocols, optional information may be encoded as a type-length-value or TLV element inside a protocol. TLV is also known as tag-length-value.

The type and length are fixed in size (typically 1-4 bytes), and the value field is of variable size. These fields are used as follows:

**Type**
    A binary code, often simply alphanumeric, which indicates the kind of field that this part of the message represents;
**Length**
    The size of the value field (typically in bytes);
**Value**
    Variable-sized series of bytes which contains data for this part of the message.
Some advantages of using a TLV representation data system solution are:

- TLV sequences are easily searched using generalized parsing functions;
- New message elements which are received at an older node can be safely skipped and the rest of the message can be parsed. This is similar to the way that unknown XML tags can be safely skipped;
- TLV elements can be placed in any order inside the message body;
- TLV elements are typically used in a binary format which makes parsing faster and the data smaller;
It is easier to generate XML from TLV to make human inspection of the data possible.

#### Reference
[Type Length Value](https://en.wikipedia.org/wiki/Type-length-value)---
name: JSON
video: 
thumb:
github:
# JSON
## JSON

JSON is an open-standard format that uses human-readable text to transmit data objects consisting of attribute-value pairs. It is the most common data format used for asynchronous browser/server communication, largely replacing XML which is used by Ajax.

```json
{
  "firstName": "John",
  "lastName": "Smith",
  "isAlive": true,
  "age": 25,
  "address": {
    "streetAddress": "21 2nd Street",
    "city": "New York",
    "state": "NY",
    "postalCode": "10021-3100"
  },
  "phoneNumbers": [
    {
      "type": "home",
      "number": "212 555-1234"
    },
    {
      "type": "office",
      "number": "646 555-4567"
    },
    {
      "type": "mobile",
      "number": "123 456-7890"
    }
  ],
  "children": [],
  "spouse": null
}
```

#### Reference
[JSON](https://en.wikipedia.org/wiki/JSON)---
name: Protocol Buffers (protobuff)
video: 
thumb:
github:
# Protocol Buffers (protobuff)
## Protocol Buffers (protobuff)

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data - think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

#### Reference
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)
- [Protocol Buffers in Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)
# HashiCorp Raft
## HashiCorp Raft

---
name: Building your own consensus using Go libraries
video: 
thumb:
github:
# Building your own consensus using Go libraries
## 

[https://raft.github.io](raft)

We will use the HashiCorp raft library for our solution.

### raft is:
- [Consistent](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Partition Tolerant](https://en.wikipedia.org/wiki/Network_partition)
- [Byzantine fault tolerant](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance)

### raft is NOT:
- [Always Available](https://en.wikipedia.org/wiki/Availability) (but shouldn't be a problem with proper cluster design)

#### References
- [HashiCorp Raft Source](https://github.com/hashicorp/raft)
- [HashiCorp Raft Documentation](https://godoc.org/github.com/hashicorp/raft)---
name: Design Considerations - How many nodes?
video: 
thumb:
github:
# Design Considerations - How many nodes?
## 

There are a lot of good articles on how many nodes to have in a consensus cluster, like the one from [etcd](https://coreos.com/etcd/docs/latest/admin_guide.html#optimal-cluster-size)

The short version is you need at least 3 for fault tolerance, and good practice is to always use odd numbers to avoid a split vote (Byzantine failures).
---
name: Common mistakes
video: 
thumb:
github:
# Common mistakes
## 

- Reading from the finite state machine on a node that isn't the leader (this is a dirty read)
- Not calling `raft.Barrier` before reading to ensure all pending writes have applied.
- Putting business logic in the state machine---
name: Raft - Code Review
video: 
thumb:
github:
# Raft - Code Review
## 

[demo source code](https://github.com/gophertrain/modules/tree/master/src/consensus/demos/raft)

- Command Args
- Joining
- Raft Store
  - State Machine
- API
  - http service---
name: Raft - Exercise
video: 
thumb:
github:https://github.com/gophertrain/modules/tree/master/src/consensus/demos/raft
# Raft - Exercise
## 

Run it as single server, then as three nodes (with different ports)

Start your first node:

```
raft -httpaddr localhost:8180 --raftaddr localhost:8186 /tmp/raft1
```

Start your second node:

```
raft --httpaddr localhost:8280 --raftaddr localhost:8286 --join localhost:8180 /tmp/raft2
```

Start your third node:

```
raft --httpaddr localhost:8380 --raftaddr localhost:8386 --join localhost:8180 /tmp/raft3
```

Write your first value

```
curl -XPOST localhost:8180/key -d '{"foo": "bar"}'
```

Read back the value

```
curl -XGET localhost:8180/key/foo
```---
name: Raft - Reading from a non-leader
video: 
thumb:
github:
# Raft - Reading from a non-leader
## 

This is an error, as you can only get a consensus read from the leader. However, we are nice enough to put the new leader in the header so you could write a client to redirect if needed.

```
$ curl -v -XGET localhost:8380/key/foo
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8380 (#0)
> GET /key/foo HTTP/1.1
> Host: localhost:8380
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 500 Internal Server Error
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< X-Raft-Leader: 127.0.0.1:8186
< Date: Sun, 30 Oct 2016 19:12:59 GMT
< Content-Length: 11
<
not leader
* Connection #0 to host localhost left intact
```---
name: Raft - Things that we didn't do
video: 
thumb:
github:
# Raft - Things that we didn't do
## 

- We did not handle writing or reading to a non-leader. We still need to respond with the leader address, or allow our handlers to forward the request to the leader via redirect.
- Tests... yes, we need a lot of testing for this.
- Add a close method to clean up and shut down our raft and httpd services nicely



name: Etdc Raft
video: 201926132
description: 
level: 
topic: Distributed Computing
# Etdc Raft
## Etdc Raft

---
name: Etcd Raft
video: 
thumb:
github:
# Etcd Raft
## Etcd Raft

- Very fast and scalable
- `Propose` vs. `Set`
- You are responsible to ensure the value was accepted and propagated
- Great for `fire and forget` scenarios if you don't need guaranteed deliver (but unlikely since you are using a consensus protocol...)

- [source code](https://github.com/coreos/etcd/tree/master/raft)
- [documentation](https://godoc.org/github.com/coreos/etcd/raft)


name: Existing Consensus Solutions
video: 201926132
description: 
level: Intermediate
topic: Distributed Computing
# Existing Consensus Solutions
## Existing Consensus Solutions

---
name: Using Etcd and Consul for distributed consensus
video: 
thumb:
github:
# Using Etcd and Consul for distributed consensus
## 

You don't have to write your own consensus system.  You can use Consul or Etcd or Zookeeper as your storage engine for state. This will add an extra dependency to your solution for anyone that needs to consume it.

Advantages of rolling your own consensus layer with raft:
- Two great libraries to choose from
- Removes a layer of complexity from deployment and configuration for your system
