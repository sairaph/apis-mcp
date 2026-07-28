---
title: firewall_threshold
page_id: schema-firewall-threshold-02a6b1b2
path: schemas
description: The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_threshold

The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.

```yaml
{"description": "The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.", "type": "number", "example": 60, "minimum": 1, "x-auditable": true}
```
