---
title: bill-subs-api_bulk_account_subscription_response
page_id: schema-bill-subs-api-bulk-account-subscription-response-5030c877
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_bulk_account_subscription_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/bill-subs-api_api-response-collection"}, {"properties": {"result": {"type": "object", "properties": {"client_secrets": {"$ref": "#/components/schemas/bill-subs-api_client_secrets"}, "subscriptions": {"type": "array", "items": {"$ref": "#/components/schemas/bill-subs-api_subscription"}}}}}, "type": "object"}]}
```
