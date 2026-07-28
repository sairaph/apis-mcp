---
title: financial_connections.account_ownership
page_id: schema-financial-connections-account-ownership-52828b77
path: schemas
description: Describes a snapshot of the owners of an account at a particular point in time.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# financial_connections.account_ownership

Describes a snapshot of the owners of an account at a particular point in time.

```yaml
{"title": "BankConnectionsResourceOwnership", "required": ["created", "id", "object", "owners"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["financial_connections.account_ownership"]}, "owners": {"title": "BankConnectionsResourceOwnerList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/financial_connections.account_owner"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "A paginated list of owners for this account.", "x-expandableFields": ["data"]}}, "description": "Describes a snapshot of the owners of an account at a particular point in time.", "x-expandableFields": ["owners"]}
```
