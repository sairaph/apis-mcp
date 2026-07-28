---
title: api-shield_per_operation_bulk_settings
page_id: schema-api-shield-per-operation-bulk-settings-b53cfb13
path: schemas
description: Operation ID to per operation setting mapping
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_per_operation_bulk_settings

Operation ID to per operation setting mapping

```yaml
{"description": "Operation ID to per operation setting mapping", "type": "object", "example": {"3818d821-5901-4147-a474-f5f5aec1d54e": {"mitigation_action": "log"}, "b17c8043-99a0-4202-b7d9-8f7cdbee02cd": {"mitigation_action": "block"}}, "additionalProperties": {"$ref": "#/components/schemas/api-shield_per_operation_setting"}}
```
