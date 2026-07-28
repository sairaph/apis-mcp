---
title: terminal.connection_token
page_id: schema-terminal-connection-token-b3d261a4
path: schemas
description: |-
    A Connection Token is used by the Stripe Terminal SDK to connect to a reader.

    Related guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal.connection_token

A Connection Token is used by the Stripe Terminal SDK to connect to a reader.

Related guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)

```yaml
{"title": "TerminalConnectionToken", "required": ["object", "secret"], "type": "object", "properties": {"location": {"maxLength": 5000, "type": "string", "description": "The id of the location that this connection token is scoped to. Note that location scoping only applies to internet-connected readers. For more details, see [the docs on scoping connection tokens](https://docs.stripe.com/terminal/fleet/locations-and-zones?dashboard-or-api=api#connection-tokens)."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["terminal.connection_token"]}, "secret": {"maxLength": 5000, "type": "string", "description": "Your application should pass this token to the Stripe Terminal SDK."}}, "description": "A Connection Token is used by the Stripe Terminal SDK to connect to a reader.\n\nRelated guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)", "x-expandableFields": [], "x-resourceId": "terminal.connection_token"}
```
