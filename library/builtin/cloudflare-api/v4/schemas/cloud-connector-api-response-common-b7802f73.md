---
title: cloud-connector_api-response-common
page_id: schema-cloud-connector-api-response-common-b7802f73
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloud-connector_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cloud-connector_messages"}, "messages": {"$ref": "#/components/schemas/cloud-connector_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
