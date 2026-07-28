---
title: apps.secret
page_id: schema-apps-secret-4a790709
path: schemas
description: |-
    Secret Store is an API that allows Stripe Apps developers to securely persist secrets for use by UI Extensions and app backends.

    The primary resource in Secret Store is a `secret`. Other apps can't view secrets created by an app. Additionally, secrets are scoped to provide further permission control.

    All Dashboard users and the app backend share `account` scoped secrets. Use the `account` scope for secrets that don't change per-user, like a third-party API key.

    A `user` scoped secret is accessible by the app backend and one specific Dashboard user. Use the `user` scope for per-user secrets like per-user OAuth tokens, where different users might have different permissions.

    Related guide: [Store data between page reloads](https://docs.stripe.com/stripe-apps/store-auth-data-custom-objects)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# apps.secret

Secret Store is an API that allows Stripe Apps developers to securely persist secrets for use by UI Extensions and app backends.

The primary resource in Secret Store is a `secret`. Other apps can't view secrets created by an app. Additionally, secrets are scoped to provide further permission control.

All Dashboard users and the app backend share `account` scoped secrets. Use the `account` scope for secrets that don't change per-user, like a third-party API key.

A `user` scoped secret is accessible by the app backend and one specific Dashboard user. Use the `user` scope for per-user secrets like per-user OAuth tokens, where different users might have different permissions.

Related guide: [Store data between page reloads](https://docs.stripe.com/stripe-apps/store-auth-data-custom-objects)

```yaml
{"title": "SecretServiceResourceSecret", "required": ["created", "id", "livemode", "name", "object", "scope"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "deleted": {"type": "boolean", "description": "If true, indicates that this secret has been deleted"}, "expires_at": {"type": "integer", "description": "The Unix timestamp for the expiry time of the secret, after which the secret deletes.", "format": "unix-time", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "A name for the secret that's unique within the scope."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["apps.secret"]}, "payload": {"maxLength": 5000, "type": "string", "description": "The plaintext secret value to be stored.", "nullable": true}, "scope": {"$ref": "#/components/schemas/secret_service_resource_scope"}}, "description": "Secret Store is an API that allows Stripe Apps developers to securely persist secrets for use by UI Extensions and app backends.\n\nThe primary resource in Secret Store is a `secret`. Other apps can't view secrets created by an app. Additionally, secrets are scoped to provide further permission control.\n\nAll Dashboard users and the app backend share `account` scoped secrets. Use the `account` scope for secrets that don't change per-user, like a third-party API key.\n\nA `user` scoped secret is accessible by the app backend and one specific Dashboard user. Use the `user` scope for per-user secrets like per-user OAuth tokens, where different users might have different permissions.\n\nRelated guide: [Store data between page reloads](https://docs.stripe.com/stripe-apps/store-auth-data-custom-objects)", "x-expandableFields": ["scope"], "x-resourceId": "apps.secret"}
```
