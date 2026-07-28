---
title: dns-firewall_minimum_cache_ttl
page_id: schema-dns-firewall-minimum-cache-ttl-27fa93b3
path: schemas
description: |-
    By default, Cloudflare attempts to cache responses for as long as
    indicated by the TTL received from upstream nameservers. This setting
    sets a lower bound on this duration. For caching purposes, lower TTLs
    will be increased to the minimum value defined by this setting.

    This setting does not affect the TTL value in the DNS response
    Cloudflare returns to clients. Cloudflare will always forward the TTL
    value received from upstream nameservers.

    Note that, even with this setting, there is no guarantee that a
    response will be cached for at least the specified duration. Cached
    responses may be removed earlier for capacity or other operational
    reasons.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_minimum_cache_ttl

By default, Cloudflare attempts to cache responses for as long as
indicated by the TTL received from upstream nameservers. This setting
sets a lower bound on this duration. For caching purposes, lower TTLs
will be increased to the minimum value defined by this setting.

This setting does not affect the TTL value in the DNS response
Cloudflare returns to clients. Cloudflare will always forward the TTL
value received from upstream nameservers.

Note that, even with this setting, there is no guarantee that a
response will be cached for at least the specified duration. Cached
responses may be removed earlier for capacity or other operational
reasons.

```yaml
{"description": "By default, Cloudflare attempts to cache responses for as long as\nindicated by the TTL received from upstream nameservers. This setting\nsets a lower bound on this duration. For caching purposes, lower TTLs\nwill be increased to the minimum value defined by this setting.\n\nThis setting does not affect the TTL value in the DNS response\nCloudflare returns to clients. Cloudflare will always forward the TTL\nvalue received from upstream nameservers.\n\nNote that, even with this setting, there is no guarantee that a\nresponse will be cached for at least the specified duration. Cached\nresponses may be removed earlier for capacity or other operational\nreasons.\n", "type": "number", "example": 60, "default": 60, "maximum": 36000, "minimum": 30, "x-auditable": true}
```
