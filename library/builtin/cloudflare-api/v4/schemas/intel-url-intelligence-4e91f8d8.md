---
title: intel_url_intelligence
page_id: schema-intel-url-intelligence-4e91f8d8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_url_intelligence

```yaml
{"type": "object", "properties": {"content_categories": {"description": "Content categories associated with this URL.", "type": "array", "items": {"$ref": "#/components/schemas/intel_url_intelligence_category_with_source"}, "example": [{"id": 155, "name": "Technology", "source_id": 1, "super_category_id": 26}]}, "full_url": {"description": "The full URL that was looked up.", "type": "string", "example": "https://example.com/path", "x-auditable": true}, "hostname": {"description": "The hostname of the URL.", "type": "string", "example": "example.com", "x-auditable": true}, "risk_type": {"description": "Security risk types associated with this URL.", "type": "array", "items": {"$ref": "#/components/schemas/intel_url_intelligence_category_with_source"}, "example": []}, "url_path": {"description": "The path component of the URL.", "type": "string", "example": "/path", "x-auditable": true}}, "required": ["full_url", "hostname", "url_path", "risk_type", "content_categories"]}
```
