---
title: argo-config_result_object
page_id: schema-argo-config-result-object-cdc8b30c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# argo-config_result_object

```yaml
{"type": "object", "properties": {"editable": {"$ref": "#/components/schemas/argo-config_editable"}, "id": {"$ref": "#/components/schemas/argo-config_setting_id"}, "modified_on": {"$ref": "#/components/schemas/argo-config_modified_on"}, "value": {"$ref": "#/components/schemas/argo-config_setting_value"}}, "required": ["id", "value", "editable"]}
```
