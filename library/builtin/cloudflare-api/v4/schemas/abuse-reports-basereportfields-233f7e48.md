---
title: abuse-reports_BaseReportFields
page_id: schema-abuse-reports-basereportfields-233f7e48
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_BaseReportFields

```yaml
{"type": "object", "properties": {"act": {"$ref": "#/components/schemas/abuse-reports_SubmissionReportType"}, "comments": {"description": "Any additional comments about the infringement not exceeding 2000 characters", "type": "string", "maxLength": 2000, "minLength": 1}, "company": {"description": "Text not exceeding 100 characters. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).", "type": "string", "maxLength": 100, "minLength": 1}, "email": {"description": "A valid email of the abuse reporter. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).", "type": "string"}, "email2": {"description": "Should match the value provided in `email`", "type": "string"}, "name": {"description": "Text not exceeding 255 characters. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).", "type": "string", "maxLength": 255, "minLength": 1}, "reported_country": {"description": "Text containing 2 characters", "type": "string", "maxLength": 2, "minLength": 2}, "reported_user_agent": {"description": "Text not exceeding 255 characters", "type": "string", "maxLength": 255, "minLength": 1}, "tele": {"description": "Text not exceeding 20 characters. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).", "type": "string", "maxLength": 20, "minLength": 1}, "title": {"description": "Text not exceeding 255 characters", "type": "string", "maxLength": 255, "minLength": 1}, "urls": {"description": "A list of valid URLs separated by ‘\\n’ (new line character). The list of the URLs should not exceed 250 URLs. All URLs should have the same hostname. Each URL should be unique. This field may be released by Cloudflare to third parties such as the Lumen Database (https://lumendatabase.org/).", "type": "string"}}, "required": ["act", "email", "email2", "name", "urls"]}
```
