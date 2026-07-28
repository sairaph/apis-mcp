---
title: rulesets_LogCustomFieldRawResponseFields
page_id: schema-rulesets-logcustomfieldrawresponsefields-f6aeb82c
path: schemas
description: The raw response fields to log.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_LogCustomFieldRawResponseFields

The raw response fields to log.

```yaml
{"description": "The raw response fields to log.", "type": "array", "items": {"description": "The raw response field to log.", "properties": {"name": {"description": "The name of the response header.", "type": "string", "example": "my-response-header", "minLength": 1, "title": "Header Name"}, "preserve_duplicates": {"description": "Whether to log duplicate values of the same header.", "type": "boolean", "example": true, "default": false, "title": "Preserve Duplicates"}}, "required": ["name"], "title": "Raw Response Field", "type": "object"}, "minItems": 1, "title": "Raw Response Fields"}
```
