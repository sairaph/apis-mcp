---
title: smartshield_status
page_id: schema-smartshield-status-fe19a5aa
path: schemas
description: The current status of the origin server according to the health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_status

The current status of the origin server according to the health check.

```yaml
{"description": "The current status of the origin server according to the health check.", "type": "string", "example": "healthy", "enum": ["unknown", "healthy", "unhealthy", "suspended"], "readOnly": true, "x-auditable": true}
```
