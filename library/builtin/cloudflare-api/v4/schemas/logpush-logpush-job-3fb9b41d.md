---
title: logpush_logpush_job
page_id: schema-logpush-logpush-job-3fb9b41d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_logpush_job

```yaml
{"type": "object", "properties": {"dataset": {"$ref": "#/components/schemas/logpush_dataset"}, "destination_conf": {"$ref": "#/components/schemas/logpush_destination_conf"}, "enabled": {"$ref": "#/components/schemas/logpush_enabled"}, "error_message": {"$ref": "#/components/schemas/logpush_error_message"}, "frequency": {"$ref": "#/components/schemas/logpush_frequency"}, "id": {"$ref": "#/components/schemas/logpush_id"}, "kind": {"$ref": "#/components/schemas/logpush_kind"}, "last_complete": {"$ref": "#/components/schemas/logpush_last_complete"}, "last_error": {"$ref": "#/components/schemas/logpush_last_error"}, "logpull_options": {"$ref": "#/components/schemas/logpush_logpull_options"}, "max_upload_bytes": {"$ref": "#/components/schemas/logpush_max_upload_bytes"}, "max_upload_interval_seconds": {"$ref": "#/components/schemas/logpush_max_upload_interval_seconds"}, "max_upload_records": {"$ref": "#/components/schemas/logpush_max_upload_records"}, "name": {"$ref": "#/components/schemas/logpush_name"}, "output_options": {"$ref": "#/components/schemas/logpush_output_options"}}, "nullable": true}
```
