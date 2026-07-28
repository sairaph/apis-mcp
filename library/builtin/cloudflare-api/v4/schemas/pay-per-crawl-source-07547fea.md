---
title: pay-per-crawl_Source
page_id: schema-pay-per-crawl-source-07547fea
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pay-per-crawl_Source

```yaml
{"type": "object", "properties": {"parameter": {"description": "Parameter is a string indicating which URI query parameter caused the error.", "type": "string"}, "parameter_value_index": {"description": "ParameterPosition indicates position of parameter value which caused the error,\nfor cases when there are multiple values for the same parameter.", "type": "integer"}, "pointer": {"description": "Pointer is a JSON Pointer [RFC6901] to the associated entity in the request document\ne.g. \"/data\" for a primary data object, or \"/data/attributes/title\" for a specific attribute.", "type": "array", "items": {"type": "string"}}}}
```
