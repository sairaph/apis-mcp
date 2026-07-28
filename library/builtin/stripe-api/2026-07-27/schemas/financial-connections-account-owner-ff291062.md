---
title: financial_connections.account_owner
page_id: schema-financial-connections-account-owner-ff291062
path: schemas
description: Describes an owner of an account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# financial_connections.account_owner

Describes an owner of an account.

```yaml
{"title": "BankConnectionsResourceOwner", "required": ["id", "name", "object", "ownership"], "type": "object", "properties": {"email": {"maxLength": 5000, "type": "string", "description": "The email address of the owner.", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "name": {"maxLength": 5000, "type": "string", "description": "The full name of the owner."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["financial_connections.account_owner"]}, "ownership": {"maxLength": 5000, "type": "string", "description": "The ownership object that this owner belongs to."}, "phone": {"maxLength": 5000, "type": "string", "description": "The raw phone number of the owner.", "nullable": true}, "raw_address": {"maxLength": 5000, "type": "string", "description": "The raw physical address of the owner.", "nullable": true}, "refreshed_at": {"type": "integer", "description": "The timestamp of the refresh that updated this owner.", "format": "unix-time", "nullable": true}}, "description": "Describes an owner of an account.", "x-expandableFields": [], "x-resourceId": "financial_connections.account_owner"}
```
