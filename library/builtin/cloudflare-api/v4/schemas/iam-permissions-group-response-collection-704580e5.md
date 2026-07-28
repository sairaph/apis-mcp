---
title: iam_permissions_group_response_collection
page_id: schema-iam-permissions-group-response-collection-704580e5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_permissions_group_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/iam_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"category": {"description": "Product category that this permission group belongs to.", "type": "string", "enum": ["developer_platform", "ai_and_machine_learning", "dns_and_zones", "app_security", "rules_and_configuration", "cloudflare_one_and_zero_trust", "analytics_and_logs", "network_services", "media", "email_and_messaging", "cache_and_performance", "account_and_billing", "other"], "x-auditable": true}, "id": {"description": "Public ID.", "type": "string", "x-auditable": true}, "name": {"description": "Permission Group Name", "type": "string", "x-auditable": true}, "scopes": {"description": "Resources to which the Permission Group is scoped", "type": "array", "items": {"enum": ["com.cloudflare.api.account", "com.cloudflare.api.account.zone", "com.cloudflare.api.user", "com.cloudflare.edge.r2.bucket"], "type": "string", "x-auditable": true}}}, "type": "object"}, "example": [{"category": "account_and_billing", "id": "7cf72faf220841aabcfdfab81c43c4f6", "name": "Billing Read", "scopes": ["com.cloudflare.api.account"]}, {"category": "network_services", "id": "9d24387c6e8544e2bc4024a03991339f", "name": "Load Balancing: Monitors and Pools Read", "scopes": ["com.cloudflare.api.account"]}, {"category": "network_services", "id": "d2a1802cc9a34e30852f8b33869b2f3c", "name": "Load Balancing: Monitors and Pools Write", "scopes": ["com.cloudflare.api.account"]}, {"category": "developer_platform", "id": "8b47d2786a534c08a1f94ee8f9f599ef", "name": "Workers KV Storage Read", "scopes": ["com.cloudflare.api.account"]}, {"category": "developer_platform", "id": "f7f0eda5697f475c90846e879bab8666", "name": "Workers KV Storage Write", "scopes": ["com.cloudflare.api.account"]}, {"category": "developer_platform", "id": "1a71c399035b4950a1bd1466bbe4f420", "name": "Workers Scripts Read", "scopes": ["com.cloudflare.api.account"]}, {"category": "developer_platform", "id": "e086da7e2179491d91ee5f35b3ca210a", "name": "Workers Scripts Write", "scopes": ["com.cloudflare.api.account"]}]}}, "type": "object"}]}
```
