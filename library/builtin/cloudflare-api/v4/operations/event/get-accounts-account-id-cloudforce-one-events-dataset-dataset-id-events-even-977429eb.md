---
title: Reads an event
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-events-93f2f277
path: operations/event
description: Retrieves a specific event by its UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/events/{event_id}
operation_ids:
    - get_EventRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads an event

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/events/{event_id}`

Operation ID: `get_EventRead`

Retrieves a specific event by its UUID.

## Definition

```yaml
{"operationId": "get_EventRead", "summary": "Reads an event", "description": "Retrieves a specific event by its UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}, {"name": "event_id", "in": "path", "description": "Event UUID.", "required": true, "schema": {"description": "Event UUID.", "type": "string"}}], "responses": {"200": {"description": "Returns the event.", "content": {"application/json": {"schema": {"type": "object", "properties": {"attacker": {"type": "string", "example": "Flying Yeti"}, "attackerCountry": {"type": "string", "example": "CN"}, "attackerCountryAlpha3": {"type": "string", "example": "CHN"}, "category": {"type": "string", "example": "Domain Resolution"}, "datasetId": {"type": "string", "example": "dataset-example-id"}, "date": {"type": "string", "example": "2022-04-01T00:00:00Z"}, "event": {"type": "string", "example": "An attacker registered the domain domain.com"}, "hasChildren": {"type": "boolean"}, "indicator": {"type": "string", "example": "domain.com"}, "indicatorType": {"type": "string", "example": "domain"}, "indicatorTypeId": {"type": "number", "example": 5}, "insight": {"type": "string"}, "killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": " ", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": " ", "type": "string"}}, "numReferenced": {"type": "number"}, "numReferences": {"type": "number"}, "rawId": {"type": "string", "example": "453gw34w3"}, "referenced": {"type": "array", "items": {"example": " ", "type": "string"}}, "referencedIds": {"type": "array", "items": {"type": "number"}}, "references": {"type": "array", "items": {"example": " ", "type": "string"}}, "referencesIds": {"type": "array", "items": {"type": "number"}}, "releasabilityId": {"type": "string"}, "tags": {"type": "array", "items": {"example": "malware", "type": "string"}}, "targetCountry": {"type": "string", "example": "US"}, "targetCountryAlpha3": {"type": "string", "example": "USA"}, "targetIndustry": {"type": "string", "example": "Agriculture"}, "tlp": {"type": "string", "example": "amber"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "date", "targetCountry", "targetCountryAlpha3", "attacker", "attackerCountry", "attackerCountryAlpha3", "targetIndustry", "rawId", "indicatorTypeId", "indicator", "event", "numReferenced", "numReferences", "tlp", "category", "indicatorType", "referenced", "references", "tags", "killChain", "mitreAttack", "mitreCapec", "referencedIds", "referencesIds", "hasChildren", "datasetId"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
