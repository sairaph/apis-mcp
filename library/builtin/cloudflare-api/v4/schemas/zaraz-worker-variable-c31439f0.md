---
title: zaraz_worker_variable
page_id: schema-zaraz-worker-variable-c31439f0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_worker_variable

```yaml
{"type": "object", "properties": {"name": {"type": "string", "x-auditable": true}, "type": {"type": "string", "enum": ["worker"], "x-auditable": true}, "value": {"type": "object", "properties": {"escapedWorkerName": {"type": "string", "x-auditable": true}, "workerTag": {"type": "string", "x-auditable": true}}, "required": ["escapedWorkerName", "workerTag"]}}, "required": ["name", "type", "value"]}
```
