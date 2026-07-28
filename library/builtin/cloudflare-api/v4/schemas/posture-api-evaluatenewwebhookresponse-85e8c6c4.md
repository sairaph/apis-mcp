---
title: posture-api_EvaluateNewWebhookResponse
page_id: schema-posture-api-evaluatenewwebhookresponse-85e8c6c4
path: schemas
description: Response body for webhook evaluation test results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_EvaluateNewWebhookResponse

Response body for webhook evaluation test results.

```yaml
{"description": "Response body for webhook evaluation test results.", "type": "object", "properties": {"message": {"description": "Human-readable message describing the test result.", "type": "string", "example": "Webhook test successful"}, "status_code": {"description": "HTTP status code returned by the webhook endpoint. 0 if connection failed.", "type": "integer", "example": 200}, "success": {"description": "Whether the webhook test was successful (received 2xx response).", "type": "boolean", "example": true}}, "required": ["success", "status_code", "message"]}
```
