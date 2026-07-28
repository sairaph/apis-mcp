---
title: tls-certificates-and-hostnames_custom_hostname_quota
page_id: schema-tls-certificates-and-hostnames-custom-hostname-quota-84ea38e2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_hostname_quota

```yaml
{"type": "object", "properties": {"allocated": {"description": "The allocated custom hostname quota.", "type": "integer", "format": "int64", "example": 100, "x-auditable": true}, "exceeded": {"description": "Whether the current usage has exceeded the allocated quota.", "type": "boolean", "example": false, "x-auditable": true}, "hard_cap": {"description": "The maximum number of custom hostnames allowed before create requests are rejected.", "type": "integer", "format": "int64", "example": 200, "x-auditable": true}, "used": {"description": "The number of custom hostnames currently in use.", "type": "integer", "format": "int64", "example": 50, "x-auditable": true}}, "required": ["allocated", "used", "exceeded", "hard_cap"]}
```
