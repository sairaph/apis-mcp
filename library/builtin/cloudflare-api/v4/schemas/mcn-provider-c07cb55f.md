---
title: mcn_provider
page_id: schema-mcn-provider-c07cb55f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_provider

```yaml
{"type": "object", "properties": {"aws_arn": {"type": "string", "x-auditable": true}, "azure_subscription_id": {"type": "string", "x-auditable": true}, "azure_tenant_id": {"type": "string", "x-auditable": true}, "cloud_type": {"$ref": "#/components/schemas/mcn_cloud_type"}, "description": {"type": "string"}, "friendly_name": {"type": "string"}, "gcp_project_id": {"type": "string", "x-auditable": true}, "gcp_service_account_email": {"type": "string", "x-auditable": true}, "id": {"$ref": "#/components/schemas/mcn_provider_id"}, "last_updated": {"type": "string", "x-auditable": true}, "lifecycle_state": {"$ref": "#/components/schemas/mcn_provider_lifecycle_state"}, "state": {"$ref": "#/components/schemas/mcn_provider_discovery_status"}, "state_v2": {"$ref": "#/components/schemas/mcn_provider_discovery_status"}, "status": {"$ref": "#/components/schemas/mcn_provider_status"}}, "required": ["id", "friendly_name", "cloud_type", "last_updated", "state", "state_v2", "lifecycle_state"]}
```
