---
title: zero-trust-gateway_applications_review_status_response_content
page_id: schema-zero-trust-gateway-applications-review-status-response-content-324b2017
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_applications_review_status_response_content

```yaml
{"type": "object", "properties": {"approved_apps": {"$ref": "#/components/schemas/zero-trust-gateway_approved_apps"}, "created_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}, "in_review_apps": {"$ref": "#/components/schemas/zero-trust-gateway_in_review_apps"}, "unapproved_apps": {"$ref": "#/components/schemas/zero-trust-gateway_unapproved_apps"}, "updated_at": {"$ref": "#/components/schemas/zero-trust-gateway_read_only_timestamp"}}}
```
