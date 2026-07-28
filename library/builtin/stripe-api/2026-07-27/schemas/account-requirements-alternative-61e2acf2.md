---
title: account_requirements_alternative
page_id: schema-account-requirements-alternative-61e2acf2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_requirements_alternative

```yaml
{"title": "AccountRequirementsAlternative", "required": ["alternative_fields_due", "original_fields_due"], "type": "object", "properties": {"alternative_fields_due": {"type": "array", "description": "Fields that can be provided to resolve all fields in `original_fields_due`.", "items": {"maxLength": 5000, "type": "string"}}, "original_fields_due": {"type": "array", "description": "Fields that are due and can be resolved by providing all fields in `alternative_fields_due`.", "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
