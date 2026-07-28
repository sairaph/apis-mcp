---
title: customer_session
page_id: schema-customer-session-00e1ba99
path: schemas
description: |-
    A Customer Session allows you to grant Stripe's frontend SDKs (like Stripe.js) client-side access
    control over a Customer.

    Related guides: [Customer Session with the Payment Element](/payments/accept-a-payment-deferred?platform=web&type=payment#save-payment-methods),
    [Customer Session with the Pricing Table](/payments/checkout/pricing-table#customer-session),
    [Customer Session with the Buy Button](/payment-links/buy-button#pass-an-existing-customer).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# customer_session

A Customer Session allows you to grant Stripe's frontend SDKs (like Stripe.js) client-side access
control over a Customer.

Related guides: [Customer Session with the Payment Element](/payments/accept-a-payment-deferred?platform=web&type=payment#save-payment-methods),
[Customer Session with the Pricing Table](/payments/checkout/pricing-table#customer-session),
[Customer Session with the Buy Button](/payment-links/buy-button#pass-an-existing-customer).

```yaml
{"title": "CustomerSessionResourceCustomerSession", "required": ["client_secret", "created", "customer", "expires_at", "livemode", "object"], "type": "object", "properties": {"client_secret": {"maxLength": 5000, "type": "string", "description": "The client secret of this Customer Session. Used on the client to set up secure access to the given `customer`.\n\nThe client secret can be used to provide access to `customer` from your frontend. It should not be stored, logged, or exposed to anyone other than the relevant customer. Make sure that you have TLS enabled on any page that includes the client secret."}, "components": {"$ref": "#/components/schemas/customer_session_resource_components"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "customer": {"description": "The Customer the Customer Session was created for.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}]}}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The Account that the Customer Session was created for.", "nullable": true}, "expires_at": {"type": "integer", "description": "The timestamp at which this Customer Session will expire.", "format": "unix-time"}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["customer_session"]}}, "description": "A Customer Session allows you to grant Stripe's frontend SDKs (like Stripe.js) client-side access\ncontrol over a Customer.\n\nRelated guides: [Customer Session with the Payment Element](/payments/accept-a-payment-deferred?platform=web&type=payment#save-payment-methods),\n[Customer Session with the Pricing Table](/payments/checkout/pricing-table#customer-session),\n[Customer Session with the Buy Button](/payment-links/buy-button#pass-an-existing-customer).", "x-expandableFields": ["components", "customer"], "x-resourceId": "customer_session"}
```
