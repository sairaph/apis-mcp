---
title: NetworkFlowLog
page_id: schema-networkflowlog-c6e1bb88
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# NetworkFlowLog

```yaml
type: object
properties:
    logged:
        type: string
        description: Timestamp of the flow log, in RFC 3339 format.
        example: '2024-06-06T15:27:26.583893Z'
    nodeId:
        type: string
        description: Identifier of the node.
        example: nBLYviWLGB21DEVEL
    start:
        type: string
        description: Time at which flow started, in RFC 3339 format.
        example: '2024-06-06T15:25:26.583893Z'
    end:
        type: string
        description: Time at which flow ended, in RFC 3339 format.
        example: '2024-06-06T15:26:26.583893Z'
    virtualTraffic:
        type: array
        items:
            $ref: '#/components/schemas/ConnectionCounts'
    subnetTraffic:
        type: array
        items:
            $ref: '#/components/schemas/ConnectionCounts'
    exitTraffic:
        type: array
        items:
            $ref: '#/components/schemas/ConnectionCounts'
    physicalTraffic:
        type: array
        items:
            $ref: '#/components/schemas/ConnectionCounts'
example:
    logged: '2024-06-06T15:27:26.583893Z'
    nodeId: nBLYviWLGB21DEVEL
    start: '2024-06-06T15:25:26.583893Z'
    end: '2024-06-06T15:26:26.583893Z'
    virtualTraffic:
        - proto: ipv4
          src: 108.86.185.125:52343
          dst: 108.86.185.126:443
          txPkts: 10
          txBytes: 10000
          rxPkts: 10
          rxBytes: 10000
    subnetTraffic:
        - proto: ipv4
          src: 108.86.185.125:52343
          dst: 108.86.185.126:443
          txPkts: 10
          txBytes: 10000
          rxPkts: 10
          rxBytes: 10000
    exitTraffic:
        - proto: ipv4
          src: 108.86.185.125:52343
          dst: 108.86.185.126:443
          txPkts: 10
          txBytes: 10000
          rxPkts: 10
          rxBytes: 10000
    physicalTraffic:
        - proto: ipv4
          src: 108.86.185.125:52343
          dst: 108.86.185.126:443
          txPkts: 10
          txBytes: 10000
          rxPkts: 10
          rxBytes: 10000
```
