---
title: intel_start_end_params
page_id: schema-intel-start-end-params-d0195f01
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_start_end_params

```yaml
{"type": "object", "properties": {"end": {"description": "Defaults to the current date.", "type": "string", "format": "date", "example": "2021-04-30", "x-auditable": true}, "start": {"description": "Defaults to 30 days before the end parameter value.", "type": "string", "format": "date", "example": "2021-04-01", "x-auditable": true}}}
```
