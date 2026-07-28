---
title: posture-api_CreateWebhookJobsResponse
page_id: schema-posture-api-createwebhookjobsresponse-9877cc72
path: schemas
description: Response for webhook job creation requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_CreateWebhookJobsResponse

Response for webhook job creation requests.

```yaml
{"description": "Response for webhook job creation requests.", "type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"created": {"description": "Successfully created webhook jobs.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookJob"}}, "failed": {"description": "Failed webhook job creation attempts.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_WebhookJobFailure"}}}, "required": ["created", "failed"]}}, "required": ["result"], "type": "object"}]}
```
