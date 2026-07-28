---
title: rulesets_TransformResponseHTMLRule
page_id: schema-rulesets-transformresponsehtmlrule-033cd591
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_TransformResponseHTMLRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["transform_response_html"]}, "action_parameters": {"properties": {"link_maze": {"description": "Enables the link maze transformation on the response.", "type": "object", "title": "Link Maze"}}, "required": ["link_maze"]}, "description": {"example": "Apply a HTML transformation to the response."}}, "title": "Transform Response HTML Rule"}]}
```
