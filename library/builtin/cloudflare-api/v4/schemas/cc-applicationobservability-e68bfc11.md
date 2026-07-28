---
title: cc_ApplicationObservability
page_id: schema-cc-applicationobservability-e68bfc11
path: schemas
description: Settings for application observability such as logging.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ApplicationObservability

Settings for application observability such as logging.

```yaml
{"description": "Settings for application observability such as logging.", "type": "object", "properties": {"logs": {"$ref": "#/components/schemas/cc_ObservabilityLogs"}, "target_instance_count": {"description": "Fixed number of instances that should receive the application-level observability overlay.\nMutually exclusive with `target_instance_percentage`.\n", "type": "integer", "minimum": 1}, "target_instance_percentage": {"description": "Percentage of instances that should receive the application-level observability overlay.\nThis rounds up so at least this percentage of instances is targeted.\nMutually exclusive with `target_instance_count`.\n", "type": "integer", "maximum": 99, "minimum": 1}}}
```
