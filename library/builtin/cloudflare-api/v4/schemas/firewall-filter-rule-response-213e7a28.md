---
title: firewall_filter-rule-response
page_id: schema-firewall-filter-rule-response-213e7a28
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_filter-rule-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/firewall_filter-rule-base"}, {"properties": {"filter": {"oneOf": [{"$ref": "#/components/schemas/firewall_filter"}, {"$ref": "#/components/schemas/firewall_deleted-filter"}]}}, "type": "object"}]}
```
