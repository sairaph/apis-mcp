---
title: rulesets_LogCustomFieldTransformedRequestFields
page_id: schema-rulesets-logcustomfieldtransformedrequestfields-c4301cf8
path: schemas
description: The transformed request fields to log.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_LogCustomFieldTransformedRequestFields

The transformed request fields to log.

```yaml
{"description": "The transformed request fields to log.", "type": "array", "items": {"description": "The transformed request field to log.", "properties": {"name": {"description": "The name of the header.", "type": "string", "example": "my-request-header", "minLength": 1, "title": "Header Name"}}, "required": ["name"], "title": "Transformed Request Field", "type": "object"}, "minItems": 1, "title": "Transformed Request Fields"}
```
