---
title: r2_class_based_metrics
page_id: schema-r2-class-based-metrics-76fa2c59
path: schemas
description: Metrics based on what state they are in(uploaded or published).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_class_based_metrics

Metrics based on what state they are in(uploaded or published).

```yaml
{"description": "Metrics based on what state they are in(uploaded or published).", "type": "object", "properties": {"published": {"$ref": "#/components/schemas/r2_object_size_metrics"}, "uploaded": {"$ref": "#/components/schemas/r2_object_size_metrics"}}}
```
