---
title: logpush_error_message
page_id: schema-logpush-error-message-94e1e94b
path: schemas
description: 'If not null, the job is currently failing. Failures are usually. repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_error_message

If not null, the job is currently failing. Failures are usually. repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.

```yaml
{"description": "If not null, the job is currently failing. Failures are usually. repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.", "type": "string", "nullable": true, "x-auditable": true}
```
