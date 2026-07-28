---
title: load-balancing_filter_options
page_id: schema-load-balancing-filter-options-a16d5677
path: schemas
description: Filter options for a particular resource type (pool or origin). Use null to reset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_filter_options

Filter options for a particular resource type (pool or origin). Use null to reset.

```yaml
{"description": "Filter options for a particular resource type (pool or origin). Use null to reset.", "type": "object", "properties": {"disable": {"description": "If set true, disable notifications for this type of resource (pool or origin).", "type": "boolean", "default": false, "nullable": true, "x-auditable": true, "x-stainless-terraform-configurability": "optional"}, "healthy": {"description": "If present, send notifications only for this health status (e.g. false for only DOWN events). Use null to reset (all events).", "type": "boolean", "nullable": true, "x-auditable": true}}, "nullable": true}
```
