---
title: zero-trust-gateway_account-log-options
page_id: schema-zero-trust-gateway-account-log-options-b1e02fdf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_account-log-options

```yaml
{"type": "object", "properties": {"log_all": {"description": "Specify whether to log all requests to this service.", "type": "boolean", "example": false, "default": false, "x-auditable": true}, "log_blocks": {"description": "Specify whether to log only blocking requests to this service.", "type": "boolean", "example": true, "default": false, "x-auditable": true}}}
```
