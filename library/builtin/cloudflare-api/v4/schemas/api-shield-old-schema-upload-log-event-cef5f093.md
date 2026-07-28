---
title: api-shield_old_schema_upload_log_event
page_id: schema-api-shield-old-schema-upload-log-event-cef5f093
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_schema_upload_log_event

```yaml
{"type": "object", "properties": {"code": {"description": "Code that identifies the event that occurred.", "type": "integer", "example": 28, "x-auditable": true}, "locations": {"description": "JSONPath location(s) in the schema where these events were encountered.  See [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/) for JSONPath specification.", "type": "array", "items": {"description": "JSONPath location in the schema where these events were encountered.  See [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/) for JSONPath specification.", "example": ".paths[\"/user/{username}\"].put", "type": "string", "x-auditable": true}}, "message": {"description": "Diagnostic message that describes the event.", "type": "string", "example": "unsupported media type: application/octet-stream", "x-auditable": true}}, "required": ["code"]}
```
