---
title: vuln_scanner_patch-target-environment-request
page_id: schema-vuln-scanner-patch-target-environment-request-7b72d5cc
path: schemas
description: |-
    Applies a partial update. Only the provided fields change; omitted fields remain unchanged.

    The `description` field supports three states:
    - **omitted**: leave unchanged
    - **`null`**: clear the description
    - **`"value"`**: set to the given string
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# vuln_scanner_patch-target-environment-request

Applies a partial update. Only the provided fields change; omitted fields remain unchanged.

The `description` field supports three states:
- **omitted**: leave unchanged
- **`null`**: clear the description
- **`"value"`**: set to the given string

```yaml
{"description": "Applies a partial update. Only the provided fields change; omitted fields remain unchanged.\n\nThe `description` field supports three states:\n- **omitted**: leave unchanged\n- **`null`**: clear the description\n- **`\"value\"`**: set to the given string\n", "type": "object", "properties": {"description": {"description": "Optional description. Omit to leave unchanged, set to `null`\nto clear, or provide a string to update.\n", "type": "string", "example": "Main production environment", "nullable": true}, "name": {"description": "Human-readable name.", "type": "string", "example": "Production Zone"}, "target": {"$ref": "#/components/schemas/vuln_scanner_target-type"}}}
```
