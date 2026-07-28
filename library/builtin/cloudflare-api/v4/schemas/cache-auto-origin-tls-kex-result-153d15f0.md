---
title: cache_auto_origin_tls_kex_result
page_id: schema-cache-auto-origin-tls-kex-result-153d15f0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache_auto_origin_tls_kex_result

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether Auto-Origin TLS KEX selection is enabled for the zone.", "allOf": [{"$ref": "#/components/schemas/cache_auto_origin_tls_kex_value"}], "readOnly": true}, "id": {"type": "string", "example": "auto_origin_tls_kex", "readOnly": true}, "modified_on": {"description": "Last time this setting was modified.", "type": "string", "format": "date-time", "example": "2014-01-01T05:20:00.12345Z", "readOnly": true}}, "required": ["id", "enabled", "modified_on"]}
```
