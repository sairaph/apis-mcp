---
title: one_AuthConfig
page_id: schema-one-authconfig-a813de64
path: schemas
description: OAuth configuration for setup flows that use OAuth.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_AuthConfig

OAuth configuration for setup flows that use OAuth.

```yaml
{"description": "OAuth configuration for setup flows that use OAuth.", "type": "object", "properties": {"authorization_url": {"description": "Full OAuth authorization URL with query parameters.", "type": "string", "nullable": true}, "client_id": {"description": "OAuth client ID.", "type": "string", "nullable": true}, "requires_pkce": {"description": "Whether PKCE is required.", "type": "boolean"}, "scopes": {"description": "OAuth scopes to request.", "type": "array", "items": {"type": "string"}}, "url_placeholders": {"description": "Placeholders in authorization URL that frontend must fill.", "type": "array", "items": {"type": "string"}}}, "required": ["authorization_url", "client_id", "requires_pkce", "scopes", "url_placeholders"]}
```
