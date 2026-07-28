---
title: cloudforce-one-requests_request-asset-item
page_id: schema-cloudforce-one-requests-request-asset-item-139d4802
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_request-asset-item

```yaml
{"type": "object", "properties": {"created": {"description": "Defines the asset creation time.", "example": "2022-01-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "description": {"description": "Asset description.", "type": "string", "example": "example description", "x-auditable": true}, "file_type": {"description": "Asset file type.", "type": "string", "example": "docx", "x-auditable": true}, "id": {"description": "Asset ID.", "type": "integer", "x-auditable": true}, "name": {"description": "Asset name.", "type": "string", "example": "example.docx", "x-auditable": true}}, "required": ["id", "name"], "title": "Request Asset Item"}
```
