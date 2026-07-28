---
title: one_SetupFlow
page_id: schema-one-setupflow-0c6185ba
path: schemas
description: Setup flow for an application auth method.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_SetupFlow

Setup flow for an application auth method.

```yaml
{"description": "Setup flow for an application auth method.", "type": "object", "properties": {"auth_config": {"description": "OAuth configuration (present for OAuth-based flows).", "allOf": [{"$ref": "#/components/schemas/one_AuthConfig"}]}, "default": {"description": "Whether this is the default auth method.", "type": "boolean"}, "description": {"description": "Flow description.", "type": "string"}, "id": {"description": "Setup flow identifier.", "type": "string"}, "name": {"description": "Human-readable flow name.", "type": "string"}, "steps": {"description": "Ordered list of setup steps.", "type": "array", "items": {"$ref": "#/components/schemas/one_SetupFlowStep"}}, "supported_environments": {"description": "Environments this auth method supports (standard, fedramp).", "type": "array", "items": {"type": "string"}}}, "required": ["default", "description", "id", "name", "steps", "supported_environments"]}
```
