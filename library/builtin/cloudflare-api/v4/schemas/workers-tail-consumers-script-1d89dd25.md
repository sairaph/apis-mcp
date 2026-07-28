---
title: workers_tail_consumers_script
page_id: schema-workers-tail-consumers-script-1d89dd25
path: schemas
description: A reference to a script that will consume logs from the attached Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_tail_consumers_script

A reference to a script that will consume logs from the attached Worker.

```yaml
{"description": "A reference to a script that will consume logs from the attached Worker.", "type": "object", "properties": {"environment": {"description": "Optional environment if the Worker utilizes one.", "type": "string", "example": "production", "x-auditable": true}, "namespace": {"description": "Optional dispatch namespace the script belongs to.", "type": "string", "example": "my-namespace", "x-auditable": true}, "service": {"description": "Name of Worker that is to be the consumer.", "type": "string", "example": "my-log-consumer", "x-auditable": true}}, "required": ["service"]}
```
