---
title: magic_cf1_site_update
page_id: schema-magic-cf1-site-update-87745933
path: schemas
description: Partial update payload for a CF1 Site. All properties are optional; only fields supplied in the request body are modified.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_cf1_site_update

Partial update payload for a CF1 Site. All properties are optional; only fields supplied in the request body are modified.

```yaml
{"description": "Partial update payload for a CF1 Site. All properties are optional; only fields supplied in the request body are modified.", "type": "object", "properties": {"description": {"description": "A human-provided description of the CF1 Site.", "type": "string", "example": "Launch Pad 34"}, "location": {"$ref": "#/components/schemas/magic_cf1_site_location"}, "name": {"description": "A human-provided name describing the CF1 Site that should be unique within the account.", "type": "string", "example": "Pad 34"}}}
```
