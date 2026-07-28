---
title: access_service_token_rule
page_id: schema-access-service-token-rule-64b93dc9
path: schemas
description: Matches a specific Access Service Token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_service_token_rule

Matches a specific Access Service Token

```yaml
{"description": "Matches a specific Access Service Token", "type": "object", "properties": {"service_token": {"type": "object", "properties": {"token_id": {"description": "The ID of a Service Token.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}}, "required": ["token_id"]}}, "required": ["service_token"], "title": "Service Token"}
```
