---
title: dns-firewall_negative_cache_ttl
page_id: schema-dns-firewall-negative-cache-ttl-6f18b3ce
path: schemas
description: |-
    This setting controls how long DNS Firewall should cache negative
    responses (e.g., NXDOMAIN) from the upstream servers.

    This setting does not affect the TTL value in the DNS response
    Cloudflare returns to clients. Cloudflare will always forward the TTL
    value received from upstream nameservers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_negative_cache_ttl

This setting controls how long DNS Firewall should cache negative
responses (e.g., NXDOMAIN) from the upstream servers.

This setting does not affect the TTL value in the DNS response
Cloudflare returns to clients. Cloudflare will always forward the TTL
value received from upstream nameservers.

```yaml
{"description": "This setting controls how long DNS Firewall should cache negative\nresponses (e.g., NXDOMAIN) from the upstream servers.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.\n", "type": "number", "example": 900, "maximum": 36000, "minimum": 30, "nullable": true, "x-auditable": true}
```
