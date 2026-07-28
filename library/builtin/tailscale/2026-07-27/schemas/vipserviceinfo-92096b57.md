---
title: VIPServiceInfo
page_id: schema-vipserviceinfo-92096b57
path: schemas
description: "An information summary for a Service.\n\nEach Service has a unique name within the tailnet, one IPv4 and one IPv6 address, optional comment, list of ports, \nand optional tags."
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# VIPServiceInfo

An information summary for a Service.

Each Service has a unique name within the tailnet, one IPv4 and one IPv6 address, optional comment, list of ports,
and optional tags.

```yaml
type: object
description: "An information summary for a Service.\n\nEach Service has a unique name within the tailnet, one IPv4 and one IPv6 address, optional comment, list of ports, \nand optional tags.\n"
properties:
    name:
        type: string
        description: The unique name of the Service.
        example: svc:example
    addrs:
        type: array
        description: |
            The IP addresses assigned to the Service: the IPv4 followed by the IPv6.
        items:
            type: string
        example:
            - 100.93.49.180
            - fd7a:115c:a1e0::3456:3cb4
    comment:
        type: string
        description: An optional comment for the Service.
        example: Example Service
    ports:
        type: array
        description: |
            A list of protocol:port pairs to be exposed by the Service.

            The only supported protocol is "tcp" at this time. "do-not-validate" can be used to skip validation.
        items:
            type: string
        example:
            - tcp:80
            - tcp:443
    tags:
        type: array
        description: A list of optional tags associated with the Service.
        items:
            type: string
        example:
            - tag:example
```
