---
title: dls_regional_hostname_response
page_id: schema-dls-regional-hostname-response-c88db318
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_regional_hostname_response

```yaml
{"type": "object", "properties": {"created_on": {"allOf": [{"$ref": "#/components/schemas/dls_timestamp"}, {"description": "When the regional hostname was created"}, {"example": "2014-01-01T05:20:00.12345Z"}]}, "hostname": {"$ref": "#/components/schemas/dls_hostname"}, "region_key": {"$ref": "#/components/schemas/dls_region_key"}, "routing": {"$ref": "#/components/schemas/dls_routing"}}, "required": ["hostname", "region_key", "routing", "created_on"]}
```
