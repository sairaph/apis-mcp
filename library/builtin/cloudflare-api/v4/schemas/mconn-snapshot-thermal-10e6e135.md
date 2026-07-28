---
title: mconn_snapshot_thermal
page_id: schema-mconn-snapshot-thermal-10e6e135
path: schemas
description: Snapshot Thermal
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_thermal

Snapshot Thermal

```yaml
{"description": "Snapshot Thermal", "type": "object", "properties": {"critical_celcius": {"description": "Critical failure temperature of the component (degrees Celsius)", "type": "number"}, "current_celcius": {"description": "Current temperature of the component (degrees Celsius)", "type": "number"}, "label": {"description": "Sensor identifier for the component", "type": "string"}, "max_celcius": {"description": "Maximum temperature of the component (degrees Celsius)", "type": "number"}}, "required": ["label"]}
```
