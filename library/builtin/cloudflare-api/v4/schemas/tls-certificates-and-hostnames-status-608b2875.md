---
title: tls-certificates-and-hostnames_status
page_id: schema-tls-certificates-and-hostnames-status-608b2875
path: schemas
description: Status of the zone's custom SSL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status

Status of the zone's custom SSL.

```yaml
{"description": "Status of the zone's custom SSL.", "type": "string", "example": "active", "enum": ["active", "expired", "deleted", "pending", "initializing"], "readOnly": true, "x-auditable": true}
```
