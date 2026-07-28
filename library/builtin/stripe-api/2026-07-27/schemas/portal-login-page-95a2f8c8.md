---
title: portal_login_page
page_id: schema-portal-login-page-95a2f8c8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_login_page

```yaml
{"title": "PortalLoginPage", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "If `true`, a shareable `url` will be generated that will take your customers to a hosted login page for the customer portal.\n\nIf `false`, the previously generated `url`, if any, will be deactivated."}, "url": {"maxLength": 5000, "type": "string", "description": "A shareable URL to the hosted portal login page. Your customers will be able to log in with their [email](https://docs.stripe.com/api/customers/object#customer_object-email) and receive a link to their customer portal.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
