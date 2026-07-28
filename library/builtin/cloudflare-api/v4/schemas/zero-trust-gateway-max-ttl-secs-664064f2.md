---
title: zero-trust-gateway_max-ttl-secs
page_id: schema-zero-trust-gateway-max-ttl-secs-664064f2
path: schemas
description: Account-level cap on DNS response TTLs, in seconds. Gateway rewrites DNS responses so returned record TTLs do not exceed this value. Null means no cap. Each DNS location can inherit, override, or disable it through the location `max_ttl` setting.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_max-ttl-secs

Account-level cap on DNS response TTLs, in seconds. Gateway rewrites DNS responses so returned record TTLs do not exceed this value. Null means no cap. Each DNS location can inherit, override, or disable it through the location `max_ttl` setting.

```yaml
{"description": "Account-level cap on DNS response TTLs, in seconds. Gateway rewrites DNS responses so returned record TTLs do not exceed this value. Null means no cap. Each DNS location can inherit, override, or disable it through the location `max_ttl` setting.", "type": "integer", "example": 3600, "maximum": 36000, "minimum": 60, "nullable": true, "x-auditable": true, "x-stainless-terraform-configurability": "optional"}
```
