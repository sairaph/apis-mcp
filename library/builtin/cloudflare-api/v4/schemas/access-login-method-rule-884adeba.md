---
title: access_login_method_rule
page_id: schema-access-login-method-rule-884adeba
path: schemas
description: Matches a specific identity provider id.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_login_method_rule

Matches a specific identity provider id.

```yaml
{"description": "Matches a specific identity provider id.", "type": "object", "properties": {"login_method": {"type": "object", "properties": {"id": {"description": "The ID of an identity provider.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}}, "required": ["id"]}}, "required": ["login_method"], "title": "Login Method"}
```
