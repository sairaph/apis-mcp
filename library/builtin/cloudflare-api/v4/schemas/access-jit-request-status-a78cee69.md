---
title: access_jit_request_status
page_id: schema-access-jit-request-status-a78cee69
path: schemas
description: JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_jit_request_status

JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.

```yaml
{"description": "JIT request status. `SPENT` is deprecated and interpreted as `APPROVED`.", "type": "string", "enum": ["PENDING", "APPROVED", "DENIED", "CANCELED", "SPENT"]}
```
