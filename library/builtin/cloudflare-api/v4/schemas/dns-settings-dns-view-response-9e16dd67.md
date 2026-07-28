---
title: dns-settings_dns-view-response
page_id: schema-dns-settings-dns-view-response-9e16dd67
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_dns-view-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns-view"}, {"properties": {"id": {"$ref": "#/components/schemas/dns-settings_identifier"}}, "required": ["id"], "type": "object"}], "required": ["id", "name", "zones", "created_time", "modified_time"]}
```
