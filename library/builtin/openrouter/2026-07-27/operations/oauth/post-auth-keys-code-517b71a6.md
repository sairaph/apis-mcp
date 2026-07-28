---
title: Create authorization code
page_id: operation-post-auth-keys-code-4aeb6e90
path: operations/oauth
description: Create an authorization code for the PKCE flow to generate a user-controlled API key
source: https://openrouter.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /auth/keys/code
operation_ids:
    - createAuthKeysCode
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Create authorization code

`POST /auth/keys/code`

Operation ID: `createAuthKeysCode`

Create an authorization code for the PKCE flow to generate a user-controlled API key

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/AppIdentifier"}, {"$ref": "#/components/parameters/AppDisplayName"}, {"$ref": "#/components/parameters/AppCategories"}]
```

## Definition

```yaml
{"description": "Create an authorization code for the PKCE flow to generate a user-controlled API key", "operationId": "createAuthKeysCode", "requestBody": {"content": {"application/json": {"example": {"callback_url": "https://myapp.com/auth/callback", "code_challenge": "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM", "code_challenge_method": "S256", "limit": 100}, "schema": {"example": {"callback_url": "https://myapp.com/auth/callback", "code_challenge": "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM", "code_challenge_method": "S256", "limit": 100}, "properties": {"callback_url": {"description": "The callback URL to redirect to after authorization. Supports https URLs and localhost/127.0.0.1 URLs on any port for local CLI tools.", "example": "https://myapp.com/auth/callback", "format": "uri", "type": "string"}, "code_challenge": {"description": "PKCE code challenge for enhanced security", "example": "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM", "type": "string"}, "code_challenge_method": {"description": "The method used to generate the code challenge", "enum": ["S256", "plain"], "example": "S256", "type": "string", "x-speakeasy-unknown-values": "allow"}, "expires_at": {"description": "Optional expiration time for the API key to be created", "example": "2027-12-31T23:59:59Z", "format": "date-time", "type": ["string", "null"]}, "key_label": {"description": "Optional custom label for the API key. Defaults to the app name if not provided.", "example": "My Custom Key", "maxLength": 100, "type": "string"}, "limit": {"description": "Credit limit for the API key to be created", "example": 100, "format": "double", "type": "number"}, "spawn_agent": {"description": "Agent identifier for spawn telemetry", "example": "my-agent", "type": "string", "x-fern-ignore": true, "x-speakeasy-ignore": true}, "spawn_cloud": {"description": "Cloud identifier for spawn telemetry", "example": "aws-us-east-1", "type": "string", "x-fern-ignore": true, "x-speakeasy-ignore": true}, "usage_limit_type": {"description": "Optional credit limit reset interval. When set, the credit limit resets on this interval.", "enum": ["daily", "weekly", "monthly"], "example": "monthly", "type": "string", "x-speakeasy-unknown-values": "allow"}, "workspace_id": {"description": "Optional workspace ID to associate the API key with", "format": "uuid", "type": "string"}}, "required": ["callback_url"], "type": "object"}}}, "required": true}, "responses": {"200": {"content": {"application/json": {"example": {"data": {"app_id": 12345, "created_at": "2025-08-24T10:30:00Z", "id": "auth_code_xyz789"}}, "schema": {"example": {"data": {"app_id": 12345, "created_at": "2025-08-24T10:30:00Z", "id": "auth_code_xyz789"}}, "properties": {"data": {"description": "Auth code data", "example": {"app_id": 12345, "created_at": "2025-08-24T10:30:00Z", "id": "auth_code_xyz789"}, "properties": {"app_id": {"description": "The application ID associated with this auth code", "example": 12345, "type": "integer"}, "created_at": {"description": "ISO 8601 timestamp of when the auth code was created", "example": "2025-08-24T10:30:00Z", "type": "string"}, "id": {"description": "The authorization code ID to use in the exchange request", "example": "auth_code_xyz789", "type": "string"}}, "required": ["id", "app_id", "created_at"], "type": "object"}}, "required": ["data"], "type": "object"}}}, "description": "Successfully created authorization code"}, "400": {"content": {"application/json": {"example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "schema": {"$ref": "#/components/schemas/BadRequestResponse"}}}, "description": "Bad Request - Invalid request parameters or malformed input"}, "401": {"content": {"application/json": {"example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "schema": {"$ref": "#/components/schemas/UnauthorizedResponse"}}}, "description": "Unauthorized - Authentication required or invalid credentials"}, "403": {"content": {"application/json": {"example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "schema": {"$ref": "#/components/schemas/ForbiddenResponse"}}}, "description": "Forbidden - Authentication successful but insufficient permissions"}, "409": {"content": {"application/json": {"example": {"error": {"code": 409, "message": "Resource conflict. Please try again later."}}, "schema": {"$ref": "#/components/schemas/ConflictResponse"}}}, "description": "Conflict - Resource conflict or concurrent modification"}, "500": {"content": {"application/json": {"example": {"error": {"code": 500, "message": "Internal Server Error"}}, "schema": {"$ref": "#/components/schemas/InternalServerResponse"}}}, "description": "Internal Server Error - Unexpected server error"}}, "summary": "Create authorization code", "tags": ["OAuth"], "x-speakeasy-name-override": "createAuthCode"}
```
