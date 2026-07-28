---
title: rulesets_ManagedTransforms
page_id: schema-rulesets-managedtransforms-59dfba3f
path: schemas
description: A Managed Transforms object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ManagedTransforms

A Managed Transforms object.

```yaml
{"description": "A Managed Transforms object.", "type": "object", "properties": {"managed_request_headers": {"description": "The list of Managed Request Transforms.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_ManagedTransform"}, {"properties": {"id": {"example": "add_bot_protection_headers"}}}]}, "title": "Managed Request Transforms", "uniqueItems": true}, "managed_response_headers": {"description": "The list of Managed Response Transforms.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_ManagedTransform"}, {"properties": {"id": {"example": "add_security_headers"}}}]}, "title": "Managed Response Transforms", "uniqueItems": true}}, "required": ["managed_request_headers", "managed_response_headers"], "title": "Managed Transforms"}
```
