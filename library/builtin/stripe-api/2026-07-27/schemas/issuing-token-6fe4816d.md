---
title: issuing.token
page_id: schema-issuing-token-6fe4816d
path: schemas
description: An issuing token object is created when an issued card is added to a digital wallet. As a [card issuer](https://docs.stripe.com/issuing), you can [view and manage these tokens](https://docs.stripe.com/issuing/controls/token-management) through Stripe.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing.token

An issuing token object is created when an issued card is added to a digital wallet. As a [card issuer](https://docs.stripe.com/issuing), you can [view and manage these tokens](https://docs.stripe.com/issuing/controls/token-management) through Stripe.

```yaml
{"title": "IssuingNetworkToken", "required": ["card", "created", "id", "livemode", "network", "network_updated_at", "object", "status"], "type": "object", "properties": {"card": {"description": "Card associated with this token.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/issuing.card"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/issuing.card"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "device_fingerprint": {"maxLength": 5000, "type": "string", "description": "The hashed ID derived from the device ID from the card network associated with the token.", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "last4": {"maxLength": 5000, "type": "string", "description": "The last four digits of the token."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "network": {"type": "string", "description": "The token service provider / card network associated with the token.", "enum": ["mastercard", "visa"]}, "network_data": {"$ref": "#/components/schemas/issuing_network_token_network_data"}, "network_updated_at": {"type": "integer", "description": "Time at which the token was last updated by the card network. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["issuing.token"]}, "status": {"type": "string", "description": "The usage state of the token.", "enum": ["active", "deleted", "requested", "suspended"]}, "wallet_provider": {"type": "string", "description": "The digital wallet for this token, if one was used.", "enum": ["apple_pay", "google_pay", "samsung_pay"]}}, "description": "An issuing token object is created when an issued card is added to a digital wallet. As a [card issuer](https://docs.stripe.com/issuing), you can [view and manage these tokens](https://docs.stripe.com/issuing/controls/token-management) through Stripe.", "x-expandableFields": ["card", "network_data"], "x-resourceId": "issuing.token"}
```
