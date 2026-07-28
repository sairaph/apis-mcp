---
title: iam_single_user_response
page_id: schema-iam-single-user-response-711349a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_single_user_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"betas": {"description": "Lists the betas that the user is participating in.", "type": "array", "items": {"description": "User feature flag", "example": "zone_level_access_beta", "type": "string"}, "readOnly": true}, "country": {"$ref": "#/components/schemas/iam_country"}, "email": {"description": "Current email address of the user.", "type": "string", "format": "email", "example": "alice@example.com", "readOnly": true}, "first_name": {"$ref": "#/components/schemas/iam_first_name"}, "has_business_zones": {"description": "Indicates whether user has any business zones", "type": "boolean", "default": false, "readOnly": true}, "has_enterprise_zones": {"description": "Indicates whether user has any enterprise zones", "type": "boolean", "default": false, "readOnly": true}, "has_pro_zones": {"description": "Indicates whether user has any pro zones", "type": "boolean", "default": false, "readOnly": true}, "id": {"description": "Identifier of the user.", "type": "string", "example": "6d7f2f5f5b1d4a0e9081fdc98d432fd1", "readOnly": true}, "last_name": {"$ref": "#/components/schemas/iam_last_name"}, "organizations": {"type": "array", "items": {"$ref": "#/components/schemas/iam_organization"}}, "suspended": {"description": "Indicates whether user has been suspended", "type": "boolean", "default": false, "readOnly": true}, "telephone": {"$ref": "#/components/schemas/iam_telephone"}, "two_factor_authentication_enabled": {"$ref": "#/components/schemas/iam_two_factor_authentication_enabled"}, "two_factor_authentication_locked": {"$ref": "#/components/schemas/iam_two_factor_authentication_locked"}, "zipcode": {"$ref": "#/components/schemas/iam_zipcode"}}, "required": ["id", "email"]}}, "type": "object"}]}
```
