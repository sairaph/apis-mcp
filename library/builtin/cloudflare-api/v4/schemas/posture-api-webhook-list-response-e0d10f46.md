---
title: posture-api_webhook-list-response
page_id: schema-posture-api-webhook-list-response-e0d10f46
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_webhook-list-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"description": "List of webhook configurations.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_Webhook"}}}, "type": "object"}]}
```
