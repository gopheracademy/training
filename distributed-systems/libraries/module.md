name: Building your own consensus using Go libraries
description: 
level: Intermediate
topic: Distributed Computing

---
name: HashiCorp Raft
video: 
thumb:
github:
# HashiCorp Raft

[https://raft.github.io](raft)

We will use the HashiCorp raft library for our solution.

### raft is:
- [Consistent](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Partition Tolerant](https://en.wikipedia.org/wiki/Network_partition)
- [Byzantine fault tolerant](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance)

### raft is NOT:
- [Always Available](https://en.wikipedia.org/wiki/Availability) (but shouldn't be a problem with proper cluster design)

*References*
- [HashiCorp Raft Source](https://github.com/hashicorp/raft)
- [HashiCorp Raft Documentation](https://godoc.org/github.com/hashicorp/raft)

---
name: Design Considerations - How many nodes?
video: 
thumb:
github:
# Design Considerations - How many nodes?

There are a lot of good articles on how many nodes to have in a consensus cluster, like the one from [etcd](https://coreos.com/etcd/docs/latest/admin_guide.html#optimal-cluster-size)

The short version is you need at least 3 for fault tolerance, and good practice is to always use odd numbers to avoid a split vote (Byzantine failures).
---
name: Common mistakes
video: 
thumb:
github:
# Common mistakes

- Reading from the finite state machine on a node that isn't the leader (this is a dirty read)
- Not calling `raft.Barrier` before reading to ensure all pending writes have applied.
- Putting business logic in the state machine

---
name: Code Review
video: 
thumb:
github:
# Code Review
## 

[demo source code](https://github.com/gophertrain/modules/tree/master/src/consensus/demos/raft)

- Command Args
- Joining
- Raft Store
  - State Machine
- API
  - http service

---
name: Exercise
video: 
thumb:
github:https://github.com/gophertrain/modules/tree/master/src/consensus/demos/raft
# Exercise

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
```

---
name: Reading from a non-leader
video: 
thumb:
github:
# Reading from a non-leader

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
```

---
name: Things that we didn't do
video: 
thumb:
github:
# Things that we didn't do

- We did not handle writing or reading to a non-leader. We still need to respond with the leader address, or allow our handlers to forward the request to the leader via redirect.
- Tests... yes, we need a lot of testing for this.
- Add a close method to clean up and shut down our raft and httpd services nicely

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

