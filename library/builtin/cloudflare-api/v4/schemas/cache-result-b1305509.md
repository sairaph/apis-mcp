---
title: cache_result
page_id: schema-cache-result-b1305509
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache_result

```yaml
{"type": "object", "properties": {"editable": {"description": "Whether this setting can be updated or not.", "type": "boolean", "readOnly": true}, "id": {"type": "string", "example": "ssl_automatic_mode", "readOnly": true}, "modified_on": {"description": "Last time this setting was modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}, "next_scheduled_scan": {"description": "Next time this zone will be scanned by the Automatic SSL/TLS.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "nullable": true, "readOnly": true}, "value": {"description": "Current setting of the automatic SSL/TLS.", "type": "string", "example": "auto", "enum": ["auto", "custom"], "readOnly": true}}, "required": ["id", "modified_on", "value", "editable"]}
```
