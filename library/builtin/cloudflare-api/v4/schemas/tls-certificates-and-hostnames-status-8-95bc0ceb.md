---
title: tls-certificates-and-hostnames_status-8
page_id: schema-tls-certificates-and-hostnames-status-8-95bc0ceb
path: schemas
description: Status of the certificate activation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status-8

Status of the certificate activation.

```yaml
{"description": "Status of the certificate activation.", "type": "string", "example": "active", "enum": ["initializing", "pending_deployment", "pending_deletion", "active", "deleted", "deployment_timed_out", "deletion_timed_out"], "x-auditable": true}
```
