---
title: rulesets_RuleExposedCredentialCheck
page_id: schema-rulesets-ruleexposedcredentialcheck-f80442e0
path: schemas
description: Configuration for exposed credential checking.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RuleExposedCredentialCheck

Configuration for exposed credential checking.

```yaml
{"description": "Configuration for exposed credential checking.", "type": "object", "properties": {"password_expression": {"description": "An expression that selects the password used in the credentials check.", "type": "string", "example": "url_decode(http.request.body.form[\\\"password\\\"][0])", "minLength": 1, "title": "Password Expression"}, "username_expression": {"description": "An expression that selects the user ID used in the credentials check.", "type": "string", "example": "url_decode(http.request.body.form[\\\"username\\\"][0])", "minLength": 1, "title": "Username Expression"}}, "required": ["username_expression", "password_expression"], "title": "Exposed Credential Check"}
```
