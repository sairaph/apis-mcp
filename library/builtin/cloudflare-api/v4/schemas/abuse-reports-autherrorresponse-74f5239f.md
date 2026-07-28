---
title: abuse-reports_AuthErrorResponse
page_id: schema-abuse-reports-autherrorresponse-74f5239f
path: schemas
description: 'The Abuse Reports API returns this HTTP 401 authorization error after credentials pass authentication but fail the required Abuse Reports permission check or, for report submission, the account entitlement check. Submission entitlement failures return "Not entitled to use feature: Abuse Report API". Enterprise accounts have the entitlement by default; other accounts must request access.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_AuthErrorResponse

The Abuse Reports API returns this HTTP 401 authorization error after credentials pass authentication but fail the required Abuse Reports permission check or, for report submission, the account entitlement check. Submission entitlement failures return "Not entitled to use feature: Abuse Report API". Enterprise accounts have the entitlement by default; other accounts must request access.

```yaml
{"description": "The Abuse Reports API returns this HTTP 401 authorization error after credentials pass authentication but fail the required Abuse Reports permission check or, for report submission, the account entitlement check. Submission entitlement failures return \"Not entitled to use feature: Abuse Report API\". Enterprise accounts have the entitlement by default; other accounts must request access.", "type": "object", "properties": {"err_code": {"description": "Machine-readable error code.", "type": "string", "example": "unknown_err"}, "error_code": {"description": "Legacy alias for err_code.", "type": "string", "example": "unknown_err"}, "msg": {"description": "Human-readable description of the authorization error. Submission entitlement failures return \"Not entitled to use feature: Abuse Report API\".", "type": "string", "example": "Not entitled to use feature: Abuse Report API"}, "result": {"description": "The result is 'error' for an error response.", "type": "string", "example": "error"}}, "example": {"err_code": "unknown_err", "error_code": "unknown_err", "msg": "Not entitled to use feature: Abuse Report API", "result": "error"}, "required": ["result", "msg", "err_code", "error_code"]}
```
