name: Etdc Raft
video: 
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
