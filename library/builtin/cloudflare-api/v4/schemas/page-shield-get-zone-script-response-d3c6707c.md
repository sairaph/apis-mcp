---
title: page-shield_get-zone-script-response
page_id: schema-page-shield-get-zone-script-response-d3c6707c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_get-zone-script-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/page-shield_api-get-response-collection"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/page-shield_script"}, {"properties": {"versions": {"type": "array", "items": {"$ref": "#/components/schemas/page-shield_version"}, "example": [{"cryptomining_score": 20, "fetched_at": "2021-08-18T10:51:08Z", "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b423", "js_integrity_score": 2, "magecart_score": 10, "malware_score": 5}], "nullable": true}}}]}}, "required": ["result"]}]}
```
