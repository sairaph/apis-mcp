---
title: Begin SSO connector verification
page_id: operation-post-accounts-account-id-sso-connectors-sso-connector-id-begin-verificat-75fbd08c
path: operations/sso
description: Validates the user has added the DNS TXT record required for validating ownership of the domain they are trying to set up a connector for.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/sso_connectors/{sso_connector_id}/begin_verification
operation_ids:
    - begin-sso-connector-verification
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Begin SSO connector verification

`POST /accounts/{account_id}/sso_connectors/{sso_connector_id}/begin_verification`

Operation ID: `begin-sso-connector-verification`

Validates the user has added the DNS TXT record required for validating ownership of the domain they are trying to set up a connector for.

## Definition

```yaml
{"operationId": "begin-sso-connector-verification", "summary": "Begin SSO connector verification", "description": "Validates the user has added the DNS TXT record required for validating ownership of the domain they are trying to set up a connector for.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "sso_connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_sso_connector_identifier"}}], "responses": {"200": {"description": "Begin SSO connector verification process response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single"}}}}, "4XX": {"description": "Begin SSO connector verification process response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["SSO"], "x-api-token-group": ["SSO Connector Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
