---
title: digital-experience-monitoring_ram_used_pct_by_app
page_id: schema-digital-experience-monitoring-ram-used-pct-by-app-7012451b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_ram_used_pct_by_app

```yaml
{"type": "object", "properties": {"name": {"description": "Application name.", "type": "string", "nullable": true}, "ram_used_pct": {"description": "RAM usage percentage, on a scale of 0 to 100.", "type": "number", "format": "float", "maximum": 100, "minimum": 0, "nullable": true}}}
```
