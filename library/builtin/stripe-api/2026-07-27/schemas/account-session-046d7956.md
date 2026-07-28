---
title: account_session
page_id: schema-account-session-046d7956
path: schemas
description: |-
    An AccountSession allows a Connect platform to grant access to a connected account in Connect embedded components.

    We recommend that you create an AccountSession each time you need to display an embedded component
    to your user. Do not save AccountSessions to your database as they expire relatively
    quickly, and cannot be used more than once.

    Related guide: [Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_session

An AccountSession allows a Connect platform to grant access to a connected account in Connect embedded components.

We recommend that you create an AccountSession each time you need to display an embedded component
to your user. Do not save AccountSessions to your database as they expire relatively
quickly, and cannot be used more than once.

Related guide: [Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components)

```yaml
{"title": "ConnectEmbeddedMethodAccountSessionCreateMethodAccountSession", "required": ["account", "client_secret", "components", "expires_at", "livemode", "object"], "type": "object", "properties": {"account": {"maxLength": 5000, "type": "string", "description": "The ID of the account the AccountSession was created for"}, "client_secret": {"maxLength": 5000, "type": "string", "description": "The client secret of this AccountSession. Used on the client to set up secure access to the given `account`.\n\nThe client secret can be used to provide access to `account` from your frontend. It should not be stored, logged, or exposed to anyone other than the connected account. Make sure that you have TLS enabled on any page that includes the client secret.\n\nRefer to our docs to [setup Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components) and learn about how `client_secret` should be handled."}, "components": {"$ref": "#/components/schemas/connect_embedded_account_session_create_components"}, "expires_at": {"type": "integer", "description": "The timestamp at which this AccountSession will expire.", "format": "unix-time"}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["account_session"]}}, "description": "An AccountSession allows a Connect platform to grant access to a connected account in Connect embedded components.\n\nWe recommend that you create an AccountSession each time you need to display an embedded component\nto your user. Do not save AccountSessions to your database as they expire relatively\nquickly, and cannot be used more than once.\n\nRelated guide: [Connect embedded components](https://docs.stripe.com/connect/get-started-connect-embedded-components)", "x-expandableFields": ["components"], "x-resourceId": "account_session"}
```
