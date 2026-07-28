---
title: login_link
page_id: schema-login-link-1056dbad
path: schemas
description: |-
    Login Links are single-use URLs that takes an Express account to the login page for their Stripe dashboard.
    A Login Link differs from an [Account Link](https://docs.stripe.com/api/account_links) in that it takes the user directly to their [Express dashboard for the specified account](https://docs.stripe.com/connect/integrate-express-dashboard#create-login-link)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# login_link

Login Links are single-use URLs that takes an Express account to the login page for their Stripe dashboard.
A Login Link differs from an [Account Link](https://docs.stripe.com/api/account_links) in that it takes the user directly to their [Express dashboard for the specified account](https://docs.stripe.com/connect/integrate-express-dashboard#create-login-link)

```yaml
{"title": "LoginLink", "required": ["created", "object", "url"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["login_link"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL for the login link."}}, "description": "Login Links are single-use URLs that takes an Express account to the login page for their Stripe dashboard.\nA Login Link differs from an [Account Link](https://docs.stripe.com/api/account_links) in that it takes the user directly to their [Express dashboard for the specified account](https://docs.stripe.com/connect/integrate-express-dashboard#create-login-link)", "x-expandableFields": [], "x-resourceId": "login_link"}
```
