---
title: fraud_username_expressions
page_id: schema-fraud-username-expressions-17ba5ac4
path: schemas
description: |-
    List of expressions to detect usernames in write HTTP requests.

    - Maximum of 10 expressions.
    - Omit or set to null to leave unchanged on update.
    - Provide an empty array `[]` to clear all expressions on update.
    - Invalid expressions will result in a 10400 Bad Request with details in the `messages` array.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# fraud_username_expressions

List of expressions to detect usernames in write HTTP requests.

- Maximum of 10 expressions.
- Omit or set to null to leave unchanged on update.
- Provide an empty array `[]` to clear all expressions on update.
- Invalid expressions will result in a 10400 Bad Request with details in the `messages` array.

```yaml
{"description": "List of expressions to detect usernames in write HTTP requests.\n\n- Maximum of 10 expressions.\n- Omit or set to null to leave unchanged on update.\n- Provide an empty array `[]` to clear all expressions on update.\n- Invalid expressions will result in a 10400 Bad Request with details in the `messages` array.\n", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["http.request.body.form[\"username\"][0]", "lookup_json_string(http.request.body.raw, \"username\")"], "maxItems": 10}
```
