---
title: custom-indicator-feeds_create_provider_result
page_id: schema-custom-indicator-feeds-create-provider-result-8f60c582
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_create_provider_result

```yaml
{"properties": {"account_id": {"description": "The numeric account ID the provider was created for. Distinct from\nthe path `account_id` parameter, which carries the account\nidentifier string used for routing.\n", "type": "integer", "example": 12345, "x-auditable": true}, "name": {"description": "The name of the provider", "type": "string", "example": "my_provider", "x-auditable": true}, "provider_id": {"description": "The unique identifier for the created provider", "type": "integer", "example": 1, "x-auditable": true}}, "required": ["account_id", "name", "provider_id"]}
```
