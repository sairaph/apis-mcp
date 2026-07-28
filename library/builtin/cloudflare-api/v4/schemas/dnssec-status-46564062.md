---
title: dnssec_status
page_id: schema-dnssec-status-46564062
path: schemas
description: Status of DNSSEC, based on user-desired state and presence of necessary records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_status

Status of DNSSEC, based on user-desired state and presence of necessary records.

```yaml
{"description": "Status of DNSSEC, based on user-desired state and presence of necessary records.", "example": "active", "enum": ["active", "pending", "disabled", "pending-disabled", "error"], "x-auditable": true}
```
