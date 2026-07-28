---
title: person_requirements
page_id: schema-person-requirements-489ce233
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_requirements

```yaml
{"title": "PersonRequirements", "required": ["currently_due", "errors", "eventually_due", "past_due", "pending_verification"], "type": "object", "properties": {"alternatives": {"type": "array", "description": "Fields that are due and can be resolved by providing the corresponding alternative fields instead. Many alternatives can list the same `original_fields_due`, and any of these alternatives can serve as a pathway for attempting to resolve the fields again. Re-providing `original_fields_due` also serves as a pathway for attempting to resolve the fields again.", "nullable": true, "items": {"$ref": "#/components/schemas/account_requirements_alternative"}}, "currently_due": {"type": "array", "description": "Fields that need to be resolved to keep the person's account enabled. If not resolved by the account's `current_deadline`, these fields will appear in `past_due` as well, and the account is disabled.", "items": {"maxLength": 5000, "type": "string"}}, "errors": {"type": "array", "description": "Details about validation and verification failures for `due` requirements that must be resolved.", "items": {"$ref": "#/components/schemas/account_requirements_error"}}, "eventually_due": {"type": "array", "description": "Fields you must collect when all thresholds are reached. As they become required, they appear in `currently_due` as well, and the account's `current_deadline` becomes set.", "items": {"maxLength": 5000, "type": "string"}}, "past_due": {"type": "array", "description": "Fields that haven't been resolved by `current_deadline`. These fields need to be resolved to enable the person's account.", "items": {"maxLength": 5000, "type": "string"}}, "pending_verification": {"type": "array", "description": "Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.", "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": ["alternatives", "errors"]}
```
