---
title: mcn_provider_status
page_id: schema-mcn-provider-status-c2a71bd1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_provider_status

```yaml
{"type": "object", "properties": {"credentials_good_since": {"type": "string", "x-auditable": true}, "credentials_missing_since": {"type": "string", "x-auditable": true}, "credentials_rejected_since": {"type": "string", "x-auditable": true}, "discovery_message": {"type": "string", "x-auditable": true}, "discovery_message_v2": {"type": "string", "x-auditable": true}, "discovery_progress": {"$ref": "#/components/schemas/mcn_provider_discovery_progress"}, "discovery_progress_v2": {"$ref": "#/components/schemas/mcn_provider_discovery_progress"}, "in_use_by": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_cloud_platform_client"}}, "last_discovery_completed_at": {"type": "string", "x-auditable": true}, "last_discovery_completed_at_v2": {"type": "string", "x-auditable": true}, "last_discovery_started_at": {"type": "string", "x-auditable": true}, "last_discovery_started_at_v2": {"type": "string", "x-auditable": true}, "last_discovery_status": {"$ref": "#/components/schemas/mcn_provider_discovery_status"}, "last_discovery_status_v2": {"$ref": "#/components/schemas/mcn_provider_discovery_status"}, "last_updated": {"type": "string", "x-auditable": true}, "regions": {"type": "array", "items": {"type": "string", "x-auditable": true}}}, "required": ["last_discovery_status", "last_discovery_status_v2", "regions", "discovery_status", "discovery_status_v2", "discovery_progress", "discovery_progress_v2"]}
```
