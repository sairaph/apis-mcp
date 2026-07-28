---
title: tls-certificates-and-hostnames_status-9
page_id: schema-tls-certificates-and-hostnames-status-9-009c41e9
path: schemas
description: Status of the certificate or the association.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status-9

Status of the certificate or the association.

```yaml
{"description": "Status of the certificate or the association.", "type": "string", "example": "active", "enum": ["initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"], "readOnly": true, "x-auditable": true}
```
