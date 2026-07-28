---
title: rulesets_RouteRule
page_id: schema-rulesets-routerule-96b32e65
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RouteRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["route"]}, "action_parameters": {"minProperties": 1, "properties": {"host_header": {"$ref": "#/components/schemas/rulesets_RouteHostHeader"}, "origin": {"$ref": "#/components/schemas/rulesets_RouteOrigin"}, "sni": {"$ref": "#/components/schemas/rulesets_RouteSNI"}}}, "description": {"example": "Select an origin server to route the request to."}}, "title": "Route Rule"}]}
```
