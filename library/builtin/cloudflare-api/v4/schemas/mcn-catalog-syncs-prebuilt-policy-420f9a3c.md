---
title: mcn_catalog_syncs_prebuilt_policy
page_id: schema-mcn-catalog-syncs-prebuilt-policy-420f9a3c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_catalog_syncs_prebuilt_policy

```yaml
{"type": "object", "properties": {"applicable_destinations": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_catalog_sync_destination_type"}}, "policy_description": {"type": "string"}, "policy_name": {"type": "string"}, "policy_string": {"type": "string"}}, "required": ["policy_name", "policy_description", "policy_string", "applicable_destinations"]}
```
