---
title: rulesets_ServeErrorRule
page_id: schema-rulesets-serveerrorrule-f9049fca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ServeErrorRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["serve_error"]}, "action_parameters": {"allOf": [{"properties": {"content_type": {"$ref": "#/components/schemas/rulesets_ServeErrorContentType"}, "status_code": {"$ref": "#/components/schemas/rulesets_ServeErrorStatusCode"}}, "required": ["content_type"]}, {"oneOf": [{"properties": {"content": {"$ref": "#/components/schemas/rulesets_ServeErrorContent"}}, "required": ["content"], "title": "Action Parameters (Content)"}, {"properties": {"asset_name": {"$ref": "#/components/schemas/rulesets_ServeErrorAssetName"}}, "required": ["asset_name"], "title": "Action Parameters (Asset)"}]}]}, "description": {"example": "Customize the serving of errors."}}, "title": "Serve Error Rule"}]}
```
