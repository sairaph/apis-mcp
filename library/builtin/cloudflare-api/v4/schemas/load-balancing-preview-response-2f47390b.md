---
title: load-balancing_preview_response
page_id: schema-load-balancing-preview-response-2f47390b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_preview_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/load-balancing_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"pools": {"description": "Monitored pool IDs mapped to their respective names.", "type": "object", "example": {"abwlnp5jbqn45ecgxd03erbgtxtqai0d": "WNAM Datacenter", "ve8h9lrcip5n5bbga9yqmdws28ay5d0l": "EEU Datacenter"}, "additionalProperties": {"description": "The pool name associated with the pool ID.", "type": "string", "x-auditable": true}}, "preview_id": {"$ref": "#/components/schemas/load-balancing_identifier"}}}}, "type": "object"}]}
```
