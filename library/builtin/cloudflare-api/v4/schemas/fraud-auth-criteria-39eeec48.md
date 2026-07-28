---
title: fraud_auth_criteria
page_id: schema-fraud-auth-criteria-39eeec48
path: schemas
description: |-
    A criterion for determining authentication outcome. The `kind` field determines which
    other fields are used for matching.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# fraud_auth_criteria

A criterion for determining authentication outcome. The `kind` field determines which
other fields are used for matching.

```yaml
{"description": "A criterion for determining authentication outcome. The `kind` field determines which\nother fields are used for matching.\n", "type": "object", "properties": {"kind": {"description": "The type of criterion. Currently only `status_code` is supported.", "type": "string", "example": "status_code", "enum": ["status_code"], "x-auditable": true}, "status_codes": {"description": "HTTP status codes to match against the origin response.\n- Maximum of 10 codes per criterion.\n- Each code must be a valid HTTP status code (100-599).\n- Codes are deduplicated and sorted on save.\n- Omit to leave unchanged on update.\n- Provide an empty array `[]` to clear codes on update.\n", "type": "array", "items": {"maximum": 599, "minimum": 100, "type": "integer"}, "example": [200, 201], "maxItems": 10, "x-auditable": true}}, "additionalProperties": false, "required": ["kind"]}
```
