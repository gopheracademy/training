name: Distributed Concepts
description: Cover concepts associated with distributed computing.
level: Intermediate
topic: Distributed Computing

---
name: CAP Theorem
video: 
thumb:
github:
# CAP Theorum

In theoretical computer science, the CAP theorem, also named Brewer's theorem after computer scientist Eric Brewer, states that it is impossible for a distributed computer system to simultaneously provide more than two out of three of the following guarantees: Consistency. Availability. Partition tolerance.

- [Consistency](https://en.wikipedia.org/wiki/Consistency_(database_systems))
- [Availability](https://en.wikipedia.org/wiki/Availability)
- [Partition Tolerance](https://en.wikipedia.org/wiki/Network_partition)

*References*
[Wikipedia Cap Theorem](https://en.wikipedia.org/wiki/CAP_theorem)

---
name: Gossip Protocols
video: 
thumb:
github:
# Gossip Protocols

## [HashiCorp member list](https://github.com/hashicorp/memberlist)

The use cases for such a library are far-reaching: all distributed systems require membership, and memberlist is a re-usable solution to managing cluster membership and node failure detection.

memberlist is eventually consistent but converges quickly on average. The speed at which it converges can be heavily tuned via various knobs on the protocol. Node failures are detected and network partitions are partially tolerated by attempting to communicate to potentially dead nodes through multiple routes.u

---
name: Consensus
video: 
thumb:
github:
# Consensus

[ Wikipedia Consensus](https://en.wikipedia.org/wiki/Consensus_(computer_science))

- Sharing a common state across systems
- Agreeing on changes across systems
- Distributed consensus
- Raft, Paxos, Zab

