---
title: cloudflare-pipelines_workers_pipelines_http_source
page_id: schema-cloudflare-pipelines-workers-pipelines-http-source-2947a31c
path: schemas
description: '[DEPRECATED] HTTP source configuration. Use the new streams API instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudflare-pipelines_workers_pipelines_http_source

[DEPRECATED] HTTP source configuration. Use the new streams API instead.

```yaml
{"description": "[DEPRECATED] HTTP source configuration. Use the new streams API instead.", "type": "object", "properties": {"authentication": {"description": "Specifies whether authentication is required to send to this pipeline via HTTP.", "type": "boolean"}, "cors": {"type": "object", "properties": {"origins": {"description": "Specifies allowed origins to allow Cross Origin HTTP Requests.", "type": "array", "items": {"type": "string"}, "example": ["*"], "maxItems": 5}}}, "format": {"description": "Specifies the format of source data.", "type": "string", "enum": ["json"]}, "type": {"type": "string"}}, "deprecated": true, "required": ["type", "format"]}
```
