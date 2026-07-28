---
title: Create a login link
page_id: operation-post-v1-accounts-account-login-links-dac59688
path: operations/untagged
description: |-
    <p>Creates a login link for a connected account to access the Express Dashboard.</p>

    <p><strong>You can only create login links for accounts that use the <a href="/connect/express-dashboard">Express Dashboard</a> and are connected to your platform</strong>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/accounts/{account}/login_links
operation_ids:
    - PostAccountsAccountLoginLinks
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a login link

`POST /v1/accounts/{account}/login_links`

Operation ID: `PostAccountsAccountLoginLinks`

<p>Creates a login link for a connected account to access the Express Dashboard.</p>

<p><strong>You can only create login links for accounts that use the <a href="/connect/express-dashboard">Express Dashboard</a> and are connected to your platform</strong>.</p>

## Definition

```yaml
{"summary": "Create a login link", "description": "<p>Creates a login link for a connected account to access the Express Dashboard.</p>\n\n<p><strong>You can only create login links for accounts that use the <a href=\"/connect/express-dashboard\">Express Dashboard</a> and are connected to your platform</strong>.</p>", "operationId": "PostAccountsAccountLoginLinks", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/login_link"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
