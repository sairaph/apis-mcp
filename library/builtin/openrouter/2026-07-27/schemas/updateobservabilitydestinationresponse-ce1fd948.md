---
title: UpdateObservabilityDestinationResponse
page_id: schema-updateobservabilitydestinationresponse-ce1fd948
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpdateObservabilityDestinationResponse

```yaml
{"example": {"data": {"api_key_hashes": null, "config": {"baseUrl": "https://us.cloud.langfuse.com", "publicKey": "pk-l...EfGh", "secretKey": "sk-l...AbCd"}, "created_at": "2025-08-24T10:30:00Z", "enabled": true, "filter_rules": null, "id": "99999999-aaaa-bbbb-cccc-dddddddddddd", "name": "Production Langfuse", "privacy_mode": false, "sampling_rate": 1, "type": "langfuse", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "properties": {"data": {"allOf": [{"$ref": "#/components/schemas/ObservabilityDestination"}, {"description": "The updated observability destination."}]}}, "required": ["data"], "type": "object"}
```
