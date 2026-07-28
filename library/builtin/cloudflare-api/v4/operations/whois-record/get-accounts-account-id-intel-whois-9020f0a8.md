---
title: Get WHOIS Record
page_id: operation-get-accounts-account-id-intel-whois-5df32608
path: operations/whois-record
description: Retrieves WHOIS registration data for a domain, including registrant and nameserver information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/whois
operation_ids:
    - whois-record-get-whois-record
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get WHOIS Record

`GET /accounts/{account_id}/intel/whois`

Operation ID: `whois-record-get-whois-record`

Retrieves WHOIS registration data for a domain, including registrant and nameserver information.

## Definition

```yaml
{"operationId": "whois-record-get-whois-record", "summary": "Get WHOIS Record", "description": "Retrieves WHOIS registration data for a domain, including registrant and nameserver information.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-whois_identifier"}}, {"name": "domain", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "Get WHOIS Record response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-whois_schemas-single_response"}}}}, "4XX": {"description": "Get WHOIS Record response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-whois_schemas-single_response"}, {"$ref": "#/components/schemas/cloudforce-one-whois_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["WHOIS Record"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
