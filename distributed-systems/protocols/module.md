name: Data Communication Protocols
description:
level: Beginner
topic: Distributed Computing
class: center, middle
module-time: 45
# Data Communication Protocols

---
name: TLV (Type Length Value)
# TLV (Type Length Value)

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

*Reference*
[Type Length Value](https://en.wikipedia.org/wiki/Type-length-value)

---
name: JSON
# JSON

JSON is an open-standard format that uses human-readable text to transmit data objects consisting of attribute-value pairs. It is the most common data format used for asynchronous browser/server communication, largely replacing XML which is used by Ajax.

<code src="distributed-systems/protocols/protocol.json" />

#### Reference
[JSON](https://en.wikipedia.org/wiki/JSON)

---
name: Protocol Buffers (protobuff)
# Protocol Buffers (protobuff)
## Protocol Buffers (protobuff)

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data - think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

#### Reference
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)
- [Protocol Buffers in Go Tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)

---
name: exercise
# Exercise - 30 Minutes

TODO: Protocol Buffers
