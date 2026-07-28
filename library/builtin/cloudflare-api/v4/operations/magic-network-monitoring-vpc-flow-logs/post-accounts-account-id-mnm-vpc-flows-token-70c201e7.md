---
title: Generate authentication token for VPC flow logs export.
page_id: operation-post-accounts-account-id-mnm-vpc-flows-token-e49826cd
path: operations/magic-network-monitoring-vpc-flow-logs
description: Generate authentication token for VPC flow logs export.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/mnm/vpc-flows/token
operation_ids:
    - magic-network-monitoring-vpc-flows-generate-authentication-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate authentication token for VPC flow logs export.

`POST /accounts/{account_id}/mnm/vpc-flows/token`

Operation ID: `magic-network-monitoring-vpc-flows-generate-authentication-token`

Generate authentication token for VPC flow logs export.

## Definition

```yaml
{"operationId": "magic-network-monitoring-vpc-flows-generate-authentication-token", "summary": "Generate authentication token for VPC flow logs export.", "description": "Generate authentication token for VPC flow logs export.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic-visibility-mnm_account_identifier"}}], "responses": {"200": {"description": "Generate authentication token for VPC flow logs export response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_vpc_flows_single_response"}}}}, "4XX": {"description": "Generate authentication token for VPC flow logs export failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic-visibility-mnm_mnm_vpc_flows_single_response"}, {"$ref": "#/components/schemas/magic-visibility-mnm_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Magic Network Monitoring VPC Flow logs"], "x-api-token-group": ["Magic Network Monitoring Admin"]}
```
