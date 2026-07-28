---
title: payment_pages_checkout_session_branding_settings
page_id: schema-payment-pages-checkout-session-branding-settings-a76c40b9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_branding_settings

```yaml
{"title": "PaymentPagesCheckoutSessionBrandingSettings", "required": ["background_color", "border_style", "button_color", "display_name", "font_family"], "type": "object", "properties": {"background_color": {"maxLength": 5000, "type": "string", "description": "A hex color value starting with `#` representing the background color for the Checkout Session."}, "border_style": {"type": "string", "description": "The border style for the Checkout Session. Must be one of `rounded`, `rectangular`, or `pill`.", "enum": ["pill", "rectangular", "rounded"]}, "button_color": {"maxLength": 5000, "type": "string", "description": "A hex color value starting with `#` representing the button color for the Checkout Session."}, "display_name": {"maxLength": 5000, "type": "string", "description": "The display name shown on the Checkout Session."}, "font_family": {"maxLength": 5000, "type": "string", "description": "The font family for the Checkout Session. Must be one of the [supported font families](https://docs.stripe.com/payments/checkout/customization/appearance?payment-ui=stripe-hosted#font-compatibility)."}, "icon": {"description": "The icon for the Checkout Session. You cannot set both `logo` and `icon`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_pages_checkout_session_branding_settings_icon"}]}, "logo": {"description": "The logo for the Checkout Session. You cannot set both `logo` and `icon`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_pages_checkout_session_branding_settings_logo"}]}}, "description": "", "x-expandableFields": ["icon", "logo"]}
```
