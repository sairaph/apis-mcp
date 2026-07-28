---
title: firewall_match
page_id: schema-firewall-match-590e7920
path: schemas
description: Determines which traffic the rate limit counts towards the threshold.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_match

Determines which traffic the rate limit counts towards the threshold.

```yaml
{"description": "Determines which traffic the rate limit counts towards the threshold.", "type": "object", "oneOf": [{"properties": {"headers": {"type": "array", "items": {"properties": {"name": {"$ref": "#/components/schemas/firewall_header_name"}, "op": {"$ref": "#/components/schemas/firewall_header_op"}, "value": {"$ref": "#/components/schemas/firewall_header_value"}}, "type": "object"}}, "request": {"type": "object", "properties": {"methods": {"$ref": "#/components/schemas/firewall_methods"}, "schemes": {"$ref": "#/components/schemas/firewall_schemes"}, "url": {"$ref": "#/components/schemas/firewall_url"}}}, "response": {"type": "object", "properties": {"origin_traffic": {"$ref": "#/components/schemas/firewall_origin_traffic"}}}}, "type": "object"}]}
```
