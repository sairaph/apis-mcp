---
title: Exchange authorization code for API key
page_id: operation-post-auth-keys-82d996dc
path: operations/oauth
description: Exchange an authorization code from the PKCE flow for a user-controlled API key
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /auth/keys
operation_ids:
    - exchangeAuthCodeForAPIKey
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Exchange authorization code for API key

`POST /auth/keys`

Operation ID: `exchangeAuthCodeForAPIKey`

Exchange an authorization code from the PKCE flow for a user-controlled API key

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Exchange an authorization code from the PKCE flow for a user-controlled API key", "operationId": "exchangeAuthCodeForAPIKey", "requestBody": {"content": {"application/json": {"example": {"code": "auth_code_abc123def456", "code_challenge_method": "S256", "code_verifier": "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"}, "schema": {"example": {"code": "auth_code_abc123def456", "code_challenge_method": "S256", "code_verifier": "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"}, "properties": {"code": {"description": "The authorization code received from the OAuth redirect", "example": "auth_code_abc123def456", "type": "string"}, "code_challenge_method": {"description": "The method used to generate the code challenge", "enum": ["S256", "plain", null], "example": "S256", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "code_verifier": {"description": "The code verifier if code_challenge was used in the authorization request", "example": "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk", "type": "string"}}, "required": ["code"], "type": "object"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"key": "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", "user_id": "user_2yOPcMpKoQhcd4bVgSMlELRaIah"}, "schema": {"example": {"key": "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", "user_id": "user_2yOPcMpKoQhcd4bVgSMlELRaIah"}, "properties": {"key": {"description": "The API key to use for OpenRouter requests", "example": "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", "type": "string"}, "user_id": {"description": "User ID associated with the API key", "example": "user_2yOPcMpKoQhcd4bVgSMlELRaIah", "type": ["string", "null"]}}, "required": ["key", "user_id"], "type": "object"}}}, "description": "Successfully exchanged code for an API key"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Exchange authorization code for API key", "tags": ["OAuth"]}
```
