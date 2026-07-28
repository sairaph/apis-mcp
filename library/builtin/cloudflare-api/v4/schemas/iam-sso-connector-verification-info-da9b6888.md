---
title: iam_sso_connector_verification_info
page_id: schema-iam-sso-connector-verification-info-da9b6888
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_sso_connector_verification_info

```yaml
{"type": "object", "properties": {"code": {"$ref": "#/components/schemas/iam_dns_verification_code"}, "status": {"description": "The status of the verification code from the verification process.", "type": "string", "example": "pending", "enum": ["awaiting", "pending", "failed", "verified"]}}}
```
