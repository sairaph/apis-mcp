---
title: tls-certificates-and-hostnames_status-4
page_id: schema-tls-certificates-and-hostnames-status-4-db02cf55
path: schemas
description: Status of the fallback origin's activation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status-4

Status of the fallback origin's activation.

```yaml
{"description": "Status of the fallback origin's activation.", "type": "string", "example": "pending_deployment", "enum": ["initializing", "pending_deployment", "pending_deletion", "active", "deployment_timed_out", "deletion_timed_out"], "x-auditable": true}
```
