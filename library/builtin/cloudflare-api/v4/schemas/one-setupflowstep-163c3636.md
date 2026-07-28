---
title: one_SetupFlowStep
page_id: schema-one-setupflowstep-163c3636
path: schemas
description: A single step in the setup flow. Polymorphic based on type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_SetupFlowStep

A single step in the setup flow. Polymorphic based on type.

```yaml
{"description": "A single step in the setup flow. Polymorphic based on type.", "type": "object", "properties": {"component_id": {"description": "Component identifier (for component type).", "type": "string", "nullable": true}, "description": {"description": "Step description with markdown support.", "type": "string", "nullable": true}, "dynamic_content": {"description": "Dynamic content blocks (for instruction/form_input).", "type": "array", "items": {"$ref": "#/components/schemas/one_DynamicContent"}, "nullable": true}, "form_fields": {"description": "Form fields (for form_input).", "type": "array", "items": {"$ref": "#/components/schemas/one_FormField"}}, "is_required": {"description": "Whether step is required (for form_input).", "type": "boolean"}, "parameters": {"description": "Component parameters (for component type).", "additionalProperties": {"type": "string"}, "nullable": true, "type": "object"}, "title": {"description": "Step title (for instruction/form_input/oauth_redirect).", "type": "string"}, "type": {"description": "Step type.\n\n* `component` - component\n* `instruction` - instruction\n* `form_input` - form_input\n* `oauth_redirect` - oauth_redirect", "type": "string", "enum": ["component", "instruction", "form_input", "oauth_redirect"]}}, "required": ["type"]}
```
