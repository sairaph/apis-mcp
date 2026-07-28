---
title: token
page_id: schema-token-3c469e9d
path: schemas
description: |-
    Tokenization is the process Stripe uses to collect sensitive card or bank
    account details, or personally identifiable information (PII), directly from
    your customers in a secure manner. A token representing this information is
    returned to your server to use. Use our
    [recommended payments integrations](https://docs.stripe.com/payments) to perform this process
    on the client-side. This guarantees that no sensitive card data touches your server,
    and allows your integration to operate in a PCI-compliant way.

    If you can't use client-side tokenization, you can also create tokens using
    the API with either your publishable or secret API key. If
    your integration uses this method, you're responsible for any PCI compliance
    that it might require, and you must keep your secret API key safe. Unlike with
    client-side tokenization, your customer's information isn't sent directly to
    Stripe, so we can't determine how it's handled or stored.

    You can't store or use tokens more than once. To store card or bank account
    information for later use, create [Customer](https://docs.stripe.com/api#customers)
    objects or [External accounts](/api#external_accounts).
    [Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,
    performs best with integrations that use client-side tokenization.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# token

Tokenization is the process Stripe uses to collect sensitive card or bank
account details, or personally identifiable information (PII), directly from
your customers in a secure manner. A token representing this information is
returned to your server to use. Use our
[recommended payments integrations](https://docs.stripe.com/payments) to perform this process
on the client-side. This guarantees that no sensitive card data touches your server,
and allows your integration to operate in a PCI-compliant way.

If you can't use client-side tokenization, you can also create tokens using
the API with either your publishable or secret API key. If
your integration uses this method, you're responsible for any PCI compliance
that it might require, and you must keep your secret API key safe. Unlike with
client-side tokenization, your customer's information isn't sent directly to
Stripe, so we can't determine how it's handled or stored.

You can't store or use tokens more than once. To store card or bank account
information for later use, create [Customer](https://docs.stripe.com/api#customers)
objects or [External accounts](/api#external_accounts).
[Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,
performs best with integrations that use client-side tokenization.

```yaml
{"title": "Token", "required": ["created", "id", "livemode", "object", "type", "used"], "type": "object", "properties": {"bank_account": {"$ref": "#/components/schemas/bank_account"}, "card": {"$ref": "#/components/schemas/card"}, "client_ip": {"maxLength": 5000, "type": "string", "description": "IP address of the client that generates the token.", "nullable": true}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["token"]}, "type": {"maxLength": 5000, "type": "string", "description": "Type of the token: `account`, `bank_account`, `card`, or `pii`."}, "used": {"type": "boolean", "description": "Determines if you have already used this token (you can only use tokens once)."}}, "description": "Tokenization is the process Stripe uses to collect sensitive card or bank\naccount details, or personally identifiable information (PII), directly from\nyour customers in a secure manner. A token representing this information is\nreturned to your server to use. Use our\n[recommended payments integrations](https://docs.stripe.com/payments) to perform this process\non the client-side. This guarantees that no sensitive card data touches your server,\nand allows your integration to operate in a PCI-compliant way.\n\nIf you can't use client-side tokenization, you can also create tokens using\nthe API with either your publishable or secret API key. If\nyour integration uses this method, you're responsible for any PCI compliance\nthat it might require, and you must keep your secret API key safe. Unlike with\nclient-side tokenization, your customer's information isn't sent directly to\nStripe, so we can't determine how it's handled or stored.\n\nYou can't store or use tokens more than once. To store card or bank account\ninformation for later use, create [Customer](https://docs.stripe.com/api#customers)\nobjects or [External accounts](/api#external_accounts).\n[Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,\nperforms best with integrations that use client-side tokenization.", "x-expandableFields": ["bank_account", "card"], "x-resourceId": "token"}
```
