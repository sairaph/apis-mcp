---
title: spectrum-config_app_config_collection
page_id: schema-spectrum-config-app-config-collection-fba695d9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-config_app_config_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/spectrum-config_api-response-collection"}, {"properties": {"result": {"oneOf": [{"items": {"$ref": "#/components/schemas/spectrum-config_app_config"}, "type": "array"}, {"items": {"$ref": "#/components/schemas/spectrum-config_paygo_app_config"}, "type": "array"}]}}}]}
```
