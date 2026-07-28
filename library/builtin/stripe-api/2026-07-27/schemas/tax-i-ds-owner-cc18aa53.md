---
title: tax_i_ds_owner
page_id: schema-tax-i-ds-owner-cc18aa53
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_i_ds_owner

```yaml
{"title": "TaxIDsOwner", "required": ["type"], "type": "object", "properties": {"account": {"description": "The account being referenced when `type` is `account`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "application": {"description": "The Connect Application being referenced when `type` is `application`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/application"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/application"}]}}, "customer": {"description": "The customer being referenced when `type` is `customer`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}]}}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The Account representing the customer being referenced when `type` is `customer`.", "nullable": true}, "type": {"type": "string", "description": "Type of owner referenced.", "enum": ["account", "application", "customer", "self"]}}, "description": "", "x-expandableFields": ["account", "application", "customer"]}
```
