---
title: connect_account_reference
page_id: schema-connect-account-reference-b8f93e9d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_account_reference

```yaml
{"title": "ConnectAccountReference", "required": ["type"], "type": "object", "properties": {"account": {"description": "The connected account being referenced when `type` is `account`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "type": {"type": "string", "description": "Type of the account referenced.", "enum": ["account", "self"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["account"]}
```
