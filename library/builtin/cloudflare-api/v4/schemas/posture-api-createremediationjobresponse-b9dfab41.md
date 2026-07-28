---
title: posture-api_CreateRemediationJobResponse
page_id: schema-posture-api-createremediationjobresponse-b9dfab41
path: schemas
description: Response for remediation job creation requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_CreateRemediationJobResponse

Response for remediation job creation requests.

```yaml
{"description": "Response for remediation job creation requests.", "type": "object", "properties": {"errors": {"description": "Array of error messages.", "type": "array", "items": {"type": "object"}, "example": []}, "messages": {"description": "Array of informational messages.", "type": "array", "items": {"type": "object"}, "example": []}, "result": {"type": "object", "properties": {"created": {"description": "Successfully created remediation jobs.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJob"}}, "failed": {"description": "Failed remediation job creation attempts.", "type": "array", "items": {"$ref": "#/components/schemas/posture-api_RemediationJobFailure"}}}, "required": ["created", "failed"]}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages", "result"]}
```
