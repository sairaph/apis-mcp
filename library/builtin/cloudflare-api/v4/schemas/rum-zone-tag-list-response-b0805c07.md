---
title: rum_zone-tag-list-response
page_id: schema-rum-zone-tag-list-response-b0805c07
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_zone-tag-list-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/rum_api-response-common"}, {"properties": {"result": {"description": "List of zone tag identifiers associated with Web Analytics sites.", "type": "array", "items": {"$ref": "#/components/schemas/rum_zone_tag"}}}, "type": "object"}]}
```
