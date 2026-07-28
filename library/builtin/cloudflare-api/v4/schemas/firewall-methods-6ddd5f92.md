---
title: firewall_methods
page_id: schema-firewall-methods-6ddd5f92
path: schemas
description: The HTTP methods to match. You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_methods

The HTTP methods to match. You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.

```yaml
{"description": "The HTTP methods to match. You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.", "type": "array", "items": {"description": "An HTTP method or `_ALL_` to indicate all methods.", "enum": ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "_ALL_"], "example": "GET", "type": "string", "x-auditable": true}, "example": ["GET", "POST"]}
```
