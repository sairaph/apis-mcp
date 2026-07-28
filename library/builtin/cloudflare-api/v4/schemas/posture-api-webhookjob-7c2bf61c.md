---
title: posture-api_WebhookJob
page_id: schema-posture-api-webhookjob-7c2bf61c
path: schemas
description: Information about a webhook job.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_WebhookJob

Information about a webhook job.

```yaml
{"description": "Information about a webhook job.", "type": "object", "properties": {"asset_data": {"description": "Asset data associated with this webhook job.", "type": "object", "additionalProperties": true}, "created_at": {"description": "When the webhook job was created.", "type": "string", "format": "date-time", "example": "2025-07-07T18:39:13.123456Z"}, "failure_details": {"description": "Additional details about the failure.", "type": "object", "additionalProperties": true}, "failure_reason": {"$ref": "#/components/schemas/posture-api_WebhookJobFailureReasonEnum"}, "id": {"description": "Unique identifier for the webhook job.", "type": "string", "format": "uuid", "example": "c416bc38-75db-425f-ae25-c37b5df5c37f"}, "integration_id": {"description": "ID of the integration.", "type": "string", "format": "uuid", "example": "9f8e7d6c-5b4a-3210-fedc-ba0987654321"}, "last_updated_at": {"description": "When the webhook job was last updated.", "type": "string", "format": "date-time", "example": "2025-07-07T18:39:13.123456Z"}, "parameters": {"$ref": "#/components/schemas/posture-api_WebhookJobParameters"}, "status": {"$ref": "#/components/schemas/posture-api_WebhookJobStatusEnum"}, "triggered_by_actor": {"$ref": "#/components/schemas/posture-api_WebhookJobActorTypeEnum"}, "triggered_by_id": {"description": "ID of the actor that triggered the job.", "type": "string", "example": "user@example.com"}, "webhook_id": {"description": "ID of the webhook configuration.", "type": "string", "format": "uuid", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}, "required": ["id", "webhook_id", "integration_id", "asset_data", "parameters", "status", "triggered_by_actor", "triggered_by_id", "created_at", "last_updated_at"]}
```
