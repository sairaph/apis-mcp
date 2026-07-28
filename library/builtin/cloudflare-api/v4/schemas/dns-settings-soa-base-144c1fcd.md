---
title: dns-settings_soa-base
page_id: schema-dns-settings-soa-base-144c1fcd
path: schemas
description: Components of the zone's SOA record.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_soa-base

Components of the zone's SOA record.

```yaml
{"description": "Components of the zone's SOA record.", "type": "object", "properties": {"expire": {"description": "Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.", "type": "number", "example": 604800, "maximum": 2419200, "minimum": 86400, "x-auditable": true}, "min_ttl": {"description": "The time to live (TTL) for negative caching of records within the zone.", "type": "number", "example": 1800, "maximum": 86400, "minimum": 60, "x-auditable": true}, "mname": {"description": "The primary nameserver, which may be used for outbound zone transfers. If null, a Cloudflare-assigned value will be used.", "type": "string", "example": "kristina.ns.cloudflare.com", "nullable": true, "x-auditable": true}, "refresh": {"description": "Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.", "type": "number", "example": 10000, "maximum": 86400, "minimum": 600, "x-auditable": true}, "retry": {"description": "Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.", "type": "number", "example": 2400, "maximum": 86400, "minimum": 600, "x-auditable": true}, "rname": {"description": "The email address of the zone administrator, with the first label representing the local part of the email address.", "type": "string", "example": "admin.example.com", "x-auditable": true}, "ttl": {"description": "The time to live (TTL) of the SOA record itself.", "type": "number", "example": 3600, "maximum": 86400, "minimum": 300, "x-auditable": true}}}
```
