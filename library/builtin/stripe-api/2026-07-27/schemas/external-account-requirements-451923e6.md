---
title: external_account_requirements
page_id: schema-external-account-requirements-451923e6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# external_account_requirements

```yaml
{"title": "ExternalAccountRequirements", "type": "object", "properties": {"currently_due": {"type": "array", "description": "Fields that need to be resolved to keep the external account enabled. If not resolved by `current_deadline`, these fields will appear in `past_due` as well, and the account is disabled.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}, "errors": {"type": "array", "description": "Details about validation and verification failures for `due` requirements that must be resolved.", "nullable": true, "items": {"$ref": "#/components/schemas/account_requirements_error"}}, "past_due": {"type": "array", "description": "Fields that haven't been resolved by `current_deadline`. These fields need to be resolved to enable the external account.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}, "pending_verification": {"type": "array", "description": "Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": ["errors"]}
```
