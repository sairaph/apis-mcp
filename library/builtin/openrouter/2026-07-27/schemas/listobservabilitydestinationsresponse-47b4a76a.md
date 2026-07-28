---
title: ListObservabilityDestinationsResponse
page_id: schema-listobservabilitydestinationsresponse-47b4a76a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListObservabilityDestinationsResponse

```yaml
{"example": {"data": [{"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "total_count": 1}, "properties": {"data": {"description": "List of observability destinations.", "items": {"$ref": "#/components/schemas/ObservabilityDestination"}, "type": "array"}, "total_count": {"description": "Total number of destinations matching the filters.", "example": 1, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
