---
title: one_AuthMethodDetail
page_id: schema-one-authmethoddetail-7ed07f27
path: schemas
description: Detailed auth method info including credentials schema and instructions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_AuthMethodDetail

Detailed auth method info including credentials schema and instructions.

```yaml
{"description": "Detailed auth method info including credentials schema and instructions.", "type": "object", "properties": {"display_name": {"description": "Human-readable auth method name.", "type": "string"}, "human_interaction_required": {"description": "Whether setup requires human interaction or integration can be created purely using API (e.g., For OAuth can not be created without user interaction).", "type": "boolean"}, "id": {"description": "Auth method identifier.", "type": "string"}, "instructions": {"description": "Step-by-step instructions for obtaining credentials.", "allOf": [{"$ref": "#/components/schemas/one_Instructions"}]}, "payload_example": {"description": "Example credentials payload with placeholder values.", "type": "object", "additionalProperties": {}, "nullable": true}, "payload_schema": {"description": "JSON Schema for the credentials object in POST /v2/integrations request.", "type": "object", "additionalProperties": {}, "nullable": true}, "redirect_url": {"description": "OAuth redirect URL for vendors requiring human interaction.", "type": "string", "nullable": true}}, "required": ["display_name", "human_interaction_required", "id", "instructions", "payload_example", "payload_schema", "redirect_url"]}
```
