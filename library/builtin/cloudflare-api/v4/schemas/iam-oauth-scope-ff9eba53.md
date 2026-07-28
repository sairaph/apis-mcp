---
title: iam_oauth_scope
page_id: schema-iam-oauth-scope-ff9eba53
path: schemas
description: An available OAuth scope that can be assigned to an OAuth client.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_scope

An available OAuth scope that can be assigned to an OAuth client.

```yaml
{"description": "An available OAuth scope that can be assigned to an OAuth client.", "type": "object", "properties": {"category": {"description": "Category for grouping scopes in the UI.", "type": "string", "example": "account_and_billing"}, "id": {"description": "The scope label to use in the scopes array when creating or updating an OAuth client.", "type": "string", "example": "account.read"}, "name": {"description": "Human-readable name of the OAuth scope.", "type": "string", "example": "Account Read"}, "scopes": {"description": "The underlying resource scopes (Bach scopes) that define which resources this OAuth scope can act upon.", "type": "array", "items": {"type": "string"}, "example": ["com.cloudflare.api.account"]}}, "required": ["name", "id"], "title": "OAuth Scope"}
```
