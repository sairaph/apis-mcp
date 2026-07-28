---
title: r2-data-catalog_namespace-identifier
page_id: schema-r2-data-catalog-namespace-identifier-8e1da8a4
path: schemas
description: |-
    Specifies the hierarchical namespace parts as an array of strings.
    For example, ["bronze", "analytics"] represents the namespace "bronze.analytics".
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_namespace-identifier

Specifies the hierarchical namespace parts as an array of strings.
For example, ["bronze", "analytics"] represents the namespace "bronze.analytics".

```yaml
{"description": "Specifies the hierarchical namespace parts as an array of strings.\nFor example, [\"bronze\", \"analytics\"] represents the namespace \"bronze.analytics\".\n", "type": "array", "items": {"type": "string"}, "example": ["bronze", "analytics"], "maxItems": 16, "minItems": 1}
```
