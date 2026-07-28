---
title: posture-api_FindingRemediation
page_id: schema-posture-api-findingremediation-1f379127
path: schemas
description: Remediation guide information for a finding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_FindingRemediation

Remediation guide information for a finding.

```yaml
{"description": "Remediation guide information for a finding.", "type": "object", "properties": {"frameworks": {"description": "Relevant Compliance Frameworks.", "type": "array", "items": {"type": "string"}, "example": ["SOC2", "ISO27001"], "readOnly": true}, "guide": {"description": "Remediation guide text.", "type": "string", "example": "To fix this issue, update the file permissions to remove public access."}, "id": {"description": "Remediation Id.", "type": "string", "format": "uuid", "example": "a20895dd-9c3b-43bd-a608-71c98c6c2d94", "readOnly": true}, "impact": {"description": "Description of the potential impact.", "type": "string", "example": "Publicly accessible files may expose sensitive information."}, "locale": {"description": "I18N Locale.", "type": "string", "example": "en-US", "readOnly": true}, "threat": {"description": "Description of the threat.", "type": "string", "example": "Data exposure and potential compliance violations."}}, "required": ["frameworks", "guide", "id", "impact", "locale", "threat"]}
```
