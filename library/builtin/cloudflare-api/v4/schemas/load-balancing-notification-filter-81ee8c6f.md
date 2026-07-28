---
title: load-balancing_notification_filter
page_id: schema-load-balancing-notification-filter-81ee8c6f
path: schemas
description: Filter pool and origin health notifications by resource type or health status. Use null to reset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_notification_filter

Filter pool and origin health notifications by resource type or health status. Use null to reset.

```yaml
{"description": "Filter pool and origin health notifications by resource type or health status. Use null to reset.", "type": "object", "properties": {"origin": {"$ref": "#/components/schemas/load-balancing_filter_options"}, "pool": {"$ref": "#/components/schemas/load-balancing_filter_options"}}, "example": {"origin": {"disable": true}, "pool": {"healthy": false}}, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}
```
