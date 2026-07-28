---
title: dns-firewall_maximum_cache_ttl
page_id: schema-dns-firewall-maximum-cache-ttl-0c8937f7
path: schemas
description: |-
    By default, Cloudflare attempts to cache responses for as long as
    indicated by the TTL received from upstream nameservers. This setting
    sets an upper bound on this duration. For caching purposes, higher TTLs
    will be decreased to the maximum value defined by this setting.

    This setting does not affect the TTL value in the DNS response
    Cloudflare returns to clients. Cloudflare will always forward the TTL
    value received from upstream nameservers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_maximum_cache_ttl

By default, Cloudflare attempts to cache responses for as long as
indicated by the TTL received from upstream nameservers. This setting
sets an upper bound on this duration. For caching purposes, higher TTLs
will be decreased to the maximum value defined by this setting.

This setting does not affect the TTL value in the DNS response
Cloudflare returns to clients. Cloudflare will always forward the TTL
value received from upstream nameservers.

```yaml
{"description": "By default, Cloudflare attempts to cache responses for as long as\nindicated by the TTL received from upstream nameservers. This setting\nsets an upper bound on this duration. For caching purposes, higher TTLs\nwill be decreased to the maximum value defined by this setting.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.\n", "type": "number", "example": 900, "default": 900, "maximum": 36000, "minimum": 30, "x-auditable": true}
```
