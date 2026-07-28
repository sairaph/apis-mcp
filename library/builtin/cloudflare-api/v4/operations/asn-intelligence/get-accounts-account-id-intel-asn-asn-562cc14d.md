---
title: Get ASN Overview.
page_id: operation-get-accounts-account-id-intel-asn-asn-ecda9759
path: operations/asn-intelligence
description: Gets an overview of the Autonomous System Number (ASN) and a list of subnets for it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/asn/{asn}
operation_ids:
    - asn-intelligence-get-asn-overview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ASN Overview.

`GET /accounts/{account_id}/intel/asn/{asn}`

Operation ID: `asn-intelligence-get-asn-overview`

Gets an overview of the Autonomous System Number (ASN) and a list of subnets for it.

## Definition

```yaml
{"operationId": "asn-intelligence-get-asn-overview", "summary": "Get ASN Overview.", "description": "Gets an overview of the Autonomous System Number (ASN) and a list of subnets for it.", "parameters": [{"name": "asn", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_asn"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}], "responses": {"200": {"description": "Get ASN Overview response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_asn_components-schemas-response"}}}}, "4XX": {"description": "Get ASN Overview response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_asn_components-schemas-response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["ASN Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
