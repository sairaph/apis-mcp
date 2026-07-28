---
title: api-shield_schema_issue_notification
page_id: schema-api-shield-schema-issue-notification-f2656eb0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_schema_issue_notification

```yaml
{"type": "object", "properties": {"code": {"description": "A unique error code that describes the kind of issue with the schema", "type": "integer", "minimum": 1000}, "message": {"description": "A short text explaining the issue with the schema", "type": "string"}, "source": {"type": "object", "nullable": true, "properties": {"locations": {"description": "A list of JSON path expression(s) that describe the location(s) of the issue within the provided resource. See [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/) for JSONPath specification.", "type": "array", "items": {"example": ".paths[\"/user/{username}\"].put", "type": "string"}}}}}, "required": ["code", "message"]}
```
