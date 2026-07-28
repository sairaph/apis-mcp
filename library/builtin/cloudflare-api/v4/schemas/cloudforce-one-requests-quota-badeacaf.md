---
title: cloudforce-one-requests_quota
page_id: schema-cloudforce-one-requests-quota-badeacaf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_quota

```yaml
{"type": "object", "properties": {"anniversary_date": {"description": "Anniversary date is when annual quota limit is refreshed.", "example": "2022-04-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "quarter_anniversary_date": {"description": "Quarter anniversary date is when quota limit is refreshed each quarter.", "example": "2022-04-01T00:00:00Z", "allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_time"}]}, "quota": {"description": "Tokens for the quarter.", "type": "integer", "example": 120, "x-auditable": true}, "remaining": {"description": "Tokens remaining for the quarter.", "type": "integer", "example": 64, "x-auditable": true}}, "title": "Quota"}
```
