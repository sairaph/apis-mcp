---
title: capability
page_id: schema-capability-38a5be91
path: schemas
description: |-
    This is an object representing a capability for a Stripe account.

    Related guide: [Account capabilities](https://docs.stripe.com/connect/account-capabilities)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# capability

This is an object representing a capability for a Stripe account.

Related guide: [Account capabilities](https://docs.stripe.com/connect/account-capabilities)

```yaml
{"title": "AccountCapability", "required": ["account", "id", "object", "requested", "status"], "type": "object", "properties": {"account": {"description": "The account for which the capability enables functionality.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "future_requirements": {"$ref": "#/components/schemas/account_capability_future_requirements"}, "id": {"maxLength": 5000, "type": "string", "description": "The identifier for the capability."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["capability"]}, "requested": {"type": "boolean", "description": "Whether the capability has been requested."}, "requested_at": {"type": "integer", "description": "Time at which the capability was requested. Measured in seconds since the Unix epoch.", "format": "unix-time", "nullable": true}, "requirements": {"$ref": "#/components/schemas/account_capability_requirements"}, "status": {"type": "string", "description": "The status of the capability.", "enum": ["active", "inactive", "pending", "unrequested"], "x-stripeBypassValidation": true}}, "description": "This is an object representing a capability for a Stripe account.\n\nRelated guide: [Account capabilities](https://docs.stripe.com/connect/account-capabilities)", "x-expandableFields": ["account", "future_requirements", "requirements"], "x-resourceId": "capability"}
```
