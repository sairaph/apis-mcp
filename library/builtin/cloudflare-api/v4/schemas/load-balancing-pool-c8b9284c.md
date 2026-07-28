---
title: load-balancing_pool
page_id: schema-load-balancing-pool-c8b9284c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_pool

```yaml
{"type": "object", "properties": {"check_regions": {"$ref": "#/components/schemas/load-balancing_check_regions"}, "created_on": {"$ref": "#/components/schemas/load-balancing_timestamp"}, "description": {"$ref": "#/components/schemas/load-balancing_schemas-description"}, "disabled_at": {"$ref": "#/components/schemas/load-balancing_schemas-disabled_at"}, "enabled": {"$ref": "#/components/schemas/load-balancing_enabled"}, "id": {"$ref": "#/components/schemas/load-balancing_schemas-identifier"}, "latitude": {"$ref": "#/components/schemas/load-balancing_latitude"}, "load_shedding": {"$ref": "#/components/schemas/load-balancing_load_shedding"}, "longitude": {"$ref": "#/components/schemas/load-balancing_longitude"}, "minimum_origins": {"$ref": "#/components/schemas/load-balancing_minimum_origins"}, "modified_on": {"$ref": "#/components/schemas/load-balancing_timestamp"}, "monitor": {"$ref": "#/components/schemas/load-balancing_monitor_id"}, "monitor_group": {"$ref": "#/components/schemas/load-balancing_monitor_group_id"}, "name": {"$ref": "#/components/schemas/load-balancing_name"}, "networks": {"$ref": "#/components/schemas/load-balancing_networks"}, "notification_email": {"$ref": "#/components/schemas/load-balancing_notification_email"}, "notification_filter": {"$ref": "#/components/schemas/load-balancing_notification_filter"}, "origin_steering": {"$ref": "#/components/schemas/load-balancing_origin_steering"}, "origins": {"$ref": "#/components/schemas/load-balancing_origins"}}}
```
