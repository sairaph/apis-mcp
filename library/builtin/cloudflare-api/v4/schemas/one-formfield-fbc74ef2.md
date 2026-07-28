---
title: one_FormField
page_id: schema-one-formfield-fbc74ef2
path: schemas
description: A form field within a form_input step.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_FormField

A form field within a form_input step.

```yaml
{"description": "A form field within a form_input step.", "type": "object", "properties": {"label": {"description": "Human-readable field label.", "type": "string"}, "name": {"description": "Field identifier (maps to credentials key).", "type": "string"}, "placeholder": {"description": "Placeholder text.", "type": "string", "nullable": true}, "required": {"description": "Whether field is required.", "type": "boolean"}, "supported_file_types": {"description": "Allowed file extensions for file_upload type.", "type": "array", "items": {"type": "string"}, "nullable": true}, "type": {"description": "Field input type.\n\n* `text` - text\n* `password` - password\n* `email` - email\n* `file_upload` - file_upload", "type": "string", "enum": ["text", "password", "email", "file_upload"]}}, "required": ["label", "name", "placeholder", "required", "supported_file_types", "type"]}
```
