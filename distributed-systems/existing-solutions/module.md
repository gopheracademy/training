name: Existing Solutions
description: 
level: Intermediate
topic: Distributed Computing

---
name: Using Etcd and Consul for distributed consensus
video: 
thumb:
github:
# Using Etcd and Consul for distributed consensus

You don't have to write your own consensus system.  You can use Consul or Etcd or Zookeeper as your storage engine for state. This will add an extra dependency to your solution for anyone that needs to consume it.

Advantages of rolling your own consensus layer with raft:
- Two great libraries to choose from
- Removes a layer of complexity from deployment and configuration for your system
