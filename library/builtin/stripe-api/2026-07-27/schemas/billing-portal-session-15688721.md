---
title: billing_portal.session
page_id: schema-billing-portal-session-15688721
path: schemas
description: |-
    The Billing customer portal is a Stripe-hosted UI for subscription and
    billing management.

    A portal configuration describes the functionality and features that you
    want to provide to your customers through the portal.

    A portal session describes the instantiation of the customer portal for
    a particular customer. By visiting the session's URL, the customer
    can manage their subscriptions and billing details. For security reasons,
    sessions are short-lived and will expire if the customer does not visit the URL.
    Create sessions on-demand when customers intend to manage their subscriptions
    and billing details.

    Related guide: [Customer management](/customer-management)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_portal.session

The Billing customer portal is a Stripe-hosted UI for subscription and
billing management.

A portal configuration describes the functionality and features that you
want to provide to your customers through the portal.

A portal session describes the instantiation of the customer portal for
a particular customer. By visiting the session's URL, the customer
can manage their subscriptions and billing details. For security reasons,
sessions are short-lived and will expire if the customer does not visit the URL.
Create sessions on-demand when customers intend to manage their subscriptions
and billing details.

Related guide: [Customer management](/customer-management)

```yaml
{"title": "PortalSession", "required": ["configuration", "created", "customer", "id", "livemode", "object", "url"], "type": "object", "properties": {"configuration": {"description": "The configuration used by this session, describing the features available.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/billing_portal.configuration"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/billing_portal.configuration"}]}}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "customer": {"maxLength": 5000, "type": "string", "description": "The ID of the customer for this session."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of the account for this session.", "nullable": true}, "flow": {"description": "Information about a specific flow for the customer to go through. See the [docs](https://docs.stripe.com/customer-management/portal-deep-links) to learn more about using customer portal deep links and flows.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/portal_flows_flow"}]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "locale": {"type": "string", "description": "The IETF language tag of the locale Customer Portal is displayed in. If blank or auto, the customer’s `preferred_locales` or browser’s locale is used.", "nullable": true, "enum": ["auto", "bg", "cs", "da", "de", "el", "en", "en-AU", "en-CA", "en-GB", "en-IE", "en-IN", "en-NZ", "en-SG", "es", "es-419", "et", "fi", "fil", "fr", "fr-CA", "hr", "hu", "id", "it", "ja", "ko", "lt", "lv", "ms", "mt", "nb", "nl", "pl", "pt", "pt-BR", "ro", "ru", "sk", "sl", "sv", "th", "tr", "vi", "zh", "zh-HK", "zh-TW"]}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing_portal.session"]}, "on_behalf_of": {"maxLength": 5000, "type": "string", "description": "The account for which the session was created on behalf of. When specified, only subscriptions and invoices with this `on_behalf_of` account appear in the portal. For more information, see the [docs](https://docs.stripe.com/connect/separate-charges-and-transfers#settlement-merchant). Use the [Accounts API](https://docs.stripe.com/api/accounts/object#account_object-settings-branding) to modify the `on_behalf_of` account's branding settings, which the portal displays.", "nullable": true}, "return_url": {"maxLength": 5000, "type": "string", "description": "The URL to redirect customers to when they click on the portal's link to return to your website.", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "The short-lived URL of the session that gives customers access to the customer portal."}}, "description": "The Billing customer portal is a Stripe-hosted UI for subscription and\nbilling management.\n\nA portal configuration describes the functionality and features that you\nwant to provide to your customers through the portal.\n\nA portal session describes the instantiation of the customer portal for\na particular customer. By visiting the session's URL, the customer\ncan manage their subscriptions and billing details. For security reasons,\nsessions are short-lived and will expire if the customer does not visit the URL.\nCreate sessions on-demand when customers intend to manage their subscriptions\nand billing details.\n\nRelated guide: [Customer management](/customer-management)", "x-expandableFields": ["configuration", "flow"], "x-resourceId": "billing_portal.session"}
```
