---
title: custom-indicator-feeds_create_provider_request
page_id: schema-custom-indicator-feeds-create-provider-request-f99bed57
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_create_provider_request

```yaml
{"properties": {"account_id": {"description": "The numeric account ID to create the provider for. Distinct from the\npath `account_id` parameter, which carries the account identifier\nstring used for routing.\n", "type": "integer", "example": 12345, "x-auditable": true}, "name": {"description": "The name of the provider", "type": "string", "example": "my_provider", "x-auditable": true}}, "required": ["account_id", "name"]}
```
