---
title: rulesets_ManagedTransformsPatch
page_id: schema-rulesets-managedtransformspatch-e07b03ab
path: schemas
description: A Managed Transforms patch object. Both fields are optional; only the sections provided will be updated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ManagedTransformsPatch

A Managed Transforms patch object. Both fields are optional; only the sections provided will be updated.

```yaml
{"description": "A Managed Transforms patch object. Both fields are optional; only the sections provided will be updated.", "type": "object", "properties": {"managed_request_headers": {"description": "The list of Managed Request Transforms.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_ManagedTransform"}, {"properties": {"id": {"example": "add_bot_protection_headers"}}}]}, "title": "Managed Request Transforms", "uniqueItems": true}, "managed_response_headers": {"description": "The list of Managed Response Transforms.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_ManagedTransform"}, {"properties": {"id": {"example": "add_security_headers"}}}]}, "title": "Managed Response Transforms", "uniqueItems": true}}, "title": "Managed Transforms Patch"}
```
