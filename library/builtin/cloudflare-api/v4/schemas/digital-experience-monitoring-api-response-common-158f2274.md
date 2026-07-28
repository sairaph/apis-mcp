---
title: digital-experience-monitoring_api-response-common
page_id: schema-digital-experience-monitoring-api-response-common-158f2274
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/digital-experience-monitoring_messages"}, "messages": {"$ref": "#/components/schemas/digital-experience-monitoring_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
