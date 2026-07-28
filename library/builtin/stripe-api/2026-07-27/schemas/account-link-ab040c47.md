---
title: account_link
page_id: schema-account-link-ab040c47
path: schemas
description: |-
    Account Links are the means by which a Connect platform grants a connected account permission to access
    Stripe-hosted applications, such as Connect Onboarding.

    Related guide: [Connect Onboarding](https://docs.stripe.com/connect/custom/hosted-onboarding)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_link

Account Links are the means by which a Connect platform grants a connected account permission to access
Stripe-hosted applications, such as Connect Onboarding.

Related guide: [Connect Onboarding](https://docs.stripe.com/connect/custom/hosted-onboarding)

```yaml
{"title": "AccountLink", "required": ["created", "expires_at", "object", "url"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "expires_at": {"type": "integer", "description": "The timestamp at which this account link will expire.", "format": "unix-time"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["account_link"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL for the account link."}}, "description": "Account Links are the means by which a Connect platform grants a connected account permission to access\nStripe-hosted applications, such as Connect Onboarding.\n\nRelated guide: [Connect Onboarding](https://docs.stripe.com/connect/custom/hosted-onboarding)", "x-expandableFields": [], "x-resourceId": "account_link"}
```
