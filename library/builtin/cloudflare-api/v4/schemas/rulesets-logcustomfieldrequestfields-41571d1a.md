---
title: rulesets_LogCustomFieldRequestFields
page_id: schema-rulesets-logcustomfieldrequestfields-41571d1a
path: schemas
description: The raw request fields to log.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_LogCustomFieldRequestFields

The raw request fields to log.

```yaml
{"description": "The raw request fields to log.", "type": "array", "items": {"description": "The raw request field to log.", "properties": {"name": {"description": "The name of the header.", "type": "string", "example": "my-request-header", "minLength": 1, "title": "Header Name"}}, "required": ["name"], "title": "Raw Request Field", "type": "object"}, "minItems": 1, "title": "Raw Request Fields"}
```
