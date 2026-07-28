---
title: ConnectionCounts
page_id: schema-connectioncounts-4084db59
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# ConnectionCounts

```yaml
type: object
properties:
    proto:
        type: string
        enum:
            - ah
            - dccp
            - egp
            - esp
            - gre
            - icmp
            - igmp
            - igp
            - ipv4
            - ipv6-icmp
            - sctp
            - tcp
            - udp
        description: IP protocol name (or number if no name used).
        example: ipv4
    src:
        type: string
        description: Source addr:port.
        example: 108.86.185.125:52343
    dst:
        type: string
        description: Destination addr:port.
        example: 108.86.185.126:443
    txPkts:
        type: integer
        description: Number of packets sent.
        example: 10
    txBytes:
        type: integer
        description: Number of bytes sent.
        example: 1000
    rxPkts:
        type: integer
        description: Number of packets received.
        example: 10
    rxBytes:
        type: integer
        description: Number of bytes received.
        example: 1000
```
