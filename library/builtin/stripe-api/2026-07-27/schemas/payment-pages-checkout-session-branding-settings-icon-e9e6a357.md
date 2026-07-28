---
title: payment_pages_checkout_session_branding_settings_icon
page_id: schema-payment-pages-checkout-session-branding-settings-icon-e9e6a357
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_branding_settings_icon

```yaml
{"title": "PaymentPagesCheckoutSessionBrandingSettingsIcon", "required": ["type"], "type": "object", "properties": {"file": {"maxLength": 5000, "type": "string", "description": "The ID of a [File upload](https://stripe.com/docs/api/files) representing the icon. Purpose must be `business_icon`. Required if `type` is `file` and disallowed otherwise."}, "type": {"type": "string", "description": "The type of image for the icon. Must be one of `file` or `url`.", "enum": ["file", "url"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL of the image. Present when `type` is `url`."}}, "description": "", "x-expandableFields": []}
```
