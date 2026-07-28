---
title: aaa_audit-logs-v2-history-not-found
page_id: schema-aaa-audit-logs-v2-history-not-found-fc450e07
path: schemas
description: Returned when the source audit log entry referenced by `id` cannot be found within the `action_time` window. The `result_info.history_status` is always set to `unavailable` for this response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-history-not-found

Returned when the source audit log entry referenced by `id` cannot be found within the `action_time` window. The `result_info.history_status` is always set to `unavailable` for this response.

```yaml
{"description": "Returned when the source audit log entry referenced by `id` cannot be found within the `action_time` window. The `result_info.history_status` is always set to `unavailable` for this response.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/aaa_messages-2"}, "result": {"type": "object", "nullable": true}, "result_info": {"$ref": "#/components/schemas/aaa_audit-logs-v2-history-result-info"}, "success": {"description": "Indicates whether the API call was successful", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "result_info"]}
```
