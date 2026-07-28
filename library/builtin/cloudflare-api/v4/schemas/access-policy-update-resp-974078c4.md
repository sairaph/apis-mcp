---
title: access_policy_update_resp
page_id: schema-access-policy-update-resp-974078c4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_policy_update_resp

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/access_policy_test_id"}, "percent_approved": {"$ref": "#/components/schemas/access_percent_approved"}, "percent_blocked": {"$ref": "#/components/schemas/access_percent_blocked"}, "percent_errored": {"$ref": "#/components/schemas/access_percent_errored"}, "percent_users_processed": {"$ref": "#/components/schemas/access_percent_users_processed"}, "status": {"$ref": "#/components/schemas/access_update_status"}, "total_users": {"$ref": "#/components/schemas/access_total_users"}, "users_approved": {"$ref": "#/components/schemas/access_users_approved"}, "users_blocked": {"$ref": "#/components/schemas/access_users_blocked"}, "users_errored": {"$ref": "#/components/schemas/access_users_errored"}}}}}]}
```
