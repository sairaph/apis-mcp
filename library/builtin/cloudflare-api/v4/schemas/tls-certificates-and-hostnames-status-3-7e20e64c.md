---
title: tls-certificates-and-hostnames_status-3
page_id: schema-tls-certificates-and-hostnames-status-3-7e20e64c
path: schemas
description: Status of the hostname's activation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status-3

Status of the hostname's activation.

```yaml
{"description": "Status of the hostname's activation.", "type": "string", "example": "pending", "enum": ["active", "pending", "active_redeploying", "moved", "pending_deletion", "deleted", "pending_blocked", "pending_migration", "pending_provisioned", "test_pending", "test_active", "test_active_apex", "test_blocked", "test_failed", "provisioned", "blocked"], "x-auditable": true}
```
