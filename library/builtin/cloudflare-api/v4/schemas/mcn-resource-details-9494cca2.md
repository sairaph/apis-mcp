---
title: mcn_resource_details
page_id: schema-mcn-resource-details-9494cca2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_resource_details

```yaml
{"type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/mcn_account_id"}, "cloud_type": {"$ref": "#/components/schemas/mcn_cloud_type"}, "config": {"type": "object", "additionalProperties": true}, "deployment_provider": {"$ref": "#/components/schemas/mcn_provider_id"}, "id": {"$ref": "#/components/schemas/mcn_resource_id"}, "managed": {"type": "boolean", "x-auditable": true}, "managed_by": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_cloud_platform_client"}}, "monthly_cost_estimate": {"$ref": "#/components/schemas/mcn_cost"}, "name": {"type": "string"}, "native_id": {"type": "string"}, "observations": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/mcn_observation"}}, "provider_ids": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_provider_id"}}, "provider_names_by_id": {"type": "object", "additionalProperties": {"type": "string"}}, "region": {"type": "string", "x-auditable": true}, "resource_group": {"type": "string", "x-auditable": true}, "resource_type": {"$ref": "#/components/schemas/mcn_resource_type"}, "sections": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_details_section"}}, "state": {"type": "object", "additionalProperties": true}, "tags": {"type": "object", "additionalProperties": {"type": "string"}}, "updated_at": {"type": "string"}, "url": {"type": "string"}}, "required": ["id", "native_id", "name", "account_id", "cloud_type", "resource_type", "managed", "provider_ids", "provider_names_by_id", "region", "resource_group", "tags", "updated_at", "url", "config", "state", "observations", "deployment_provider", "sections", "monthly_cost_estimate"]}
```
