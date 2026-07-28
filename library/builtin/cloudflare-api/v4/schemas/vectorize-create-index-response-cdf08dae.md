---
title: vectorize_create-index-response
page_id: schema-vectorize-create-index-response-cdf08dae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vectorize_create-index-response

```yaml
{"type": "object", "properties": {"config": {"$ref": "#/components/schemas/vectorize_index-dimension-configuration"}, "created_on": {"description": "Specifies the timestamp the resource was created as an ISO8601 string.", "type": "string", "format": "date-time", "example": "2022-11-15T18:25:44.442097Z", "readOnly": true, "x-auditable": true}, "description": {"$ref": "#/components/schemas/vectorize_index-description"}, "modified_on": {"description": "Specifies the timestamp the resource was modified as an ISO8601 string.", "type": "string", "format": "date-time", "example": "2022-11-15T18:25:44.442097Z", "readOnly": true, "x-auditable": true}, "name": {"$ref": "#/components/schemas/vectorize_index-name"}}}
```
