---
title: cloudforce-one-requests_request-constants
page_id: schema-cloudforce-one-requests-request-constants-1d29c5df
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-constants

```yaml
{"type": "object", "properties": {"priority": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_priority"}, "example": ["routine", "high", "urgent"]}, "status": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_request-status"}, "example": ["open", "accepted", "reported", "approved", "completed", "declined"]}, "tlp": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_tlp"}, "example": ["clear", "green", "amber", "amber-strict", "red"]}}, "title": "Request Constants"}
```
