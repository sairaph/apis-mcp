---
title: api-shield_old_schema_upload_details_errors_critical
page_id: schema-api-shield-old-schema-upload-details-errors-critical-57b4b9d1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_old_schema_upload_details_errors_critical

```yaml
{"type": "object", "properties": {"critical": {"description": "Diagnostic critical error events that occurred during processing.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_old_schema_upload_log_event"}}, "errors": {"description": "Diagnostic error events that occurred during processing.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_old_schema_upload_log_event"}}}}
```
