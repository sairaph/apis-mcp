---
title: account_branding_settings
page_id: schema-account-branding-settings-c7ee2dbf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_branding_settings

```yaml
{"title": "AccountBrandingSettings", "type": "object", "properties": {"icon": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) An icon for the account. Must be square and at least 128px x 128px.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "logo": {"description": "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) A logo for the account that will be used in Checkout instead of the icon and without the account's name next to it if provided. Must be at least 128px x 128px.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "primary_color": {"maxLength": 5000, "type": "string", "description": "A CSS hex color value representing the primary branding color for this account", "nullable": true}, "secondary_color": {"maxLength": 5000, "type": "string", "description": "A CSS hex color value representing the secondary branding color for this account", "nullable": true}}, "description": "", "x-expandableFields": ["icon", "logo"]}
```
