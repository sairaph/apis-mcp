---
title: Get policy file
page_id: operation-get-tailnet-tailnet-acl-d8ad01f9
path: operations/policyfile
description: "Retrieves the current policy file for the given tailnet; \nthis includes the ACL along with the rules and tests that have been defined.\n\nThis method can return the policy file as JSON or HuJSON, depending on the Accept header.\nThe response also includes an `ETag` header, which can be optionally included when [setting the policy file](#tag/policyfile/post/tailnet/{tailnet}/acl) to avoid missed updates.\n\nLearn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).\n\nOAuth Scope: `policy_file:read`."
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/acl
operation_ids:
    - getPolicyFile
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get policy file

`GET /tailnet/{tailnet}/acl`

Operation ID: `getPolicyFile`

Retrieves the current policy file for the given tailnet;
this includes the ACL along with the rules and tests that have been defined.

This method can return the policy file as JSON or HuJSON, depending on the Accept header.
The response also includes an `ETag` header, which can be optionally included when [setting the policy file](#tag/policyfile/post/tailnet/{tailnet}/acl) to avoid missed updates.

Learn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).

OAuth Scope: `policy_file:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/AcceptHeaderParam'
```

## Definition

```yaml
summary: Get policy file
description: "Retrieves the current policy file for the given tailnet; \nthis includes the ACL along with the rules and tests that have been defined.\n\nThis method can return the policy file as JSON or HuJSON, depending on the Accept header.\nThe response also includes an `ETag` header, which can be optionally included when [setting the policy file](#tag/policyfile/post/tailnet/{tailnet}/acl) to avoid missed updates.\n\nLearn more about [policy file ACL syntax](https://tailscale.com/kb/1337/acl-syntax).\n\nOAuth Scope: `policy_file:read`.\n"
operationId: getPolicyFile
tags:
    - PolicyFile
parameters:
    - name: details
      in: query
      required: false
      schema:
        type: boolean
      description: |
        Request a detailed description of the tailnet policy file by providing `details=true` in the URL query string.
        Supplying any other value for `details`, or not sending the param, is treated as sending `details=false`.
        If using this, do not supply an `Accept` parameter in the header.

        The response will contain a JSON object with the fields:
        - `acl`: a base64-encoded string representation of the huJSON format.
        - `warnings`: array of strings for syntactically valid but nonsensical entries.
        - `errors`: an array of strings for parsing failures.
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                examples:
                    exampleJSONPolicyFile:
                        summary: JSON response with details=false
                        $ref: '#/components/examples/ExampleJSONPolicyFile'
                    exampleWithDetails:
                        summary: JSON response with details=true
                        value:
                            acl: Ly8gUG9raW5nIGFyb3VuZCBpbiB0aGUgQVBJIGRvY3MsIGhvcGluZyB5b3UnZCBmaW5kIHNvbWV0aGluZyBnb29kLCBlaD8KLy8gV2UgbGlrZSB5b3VyIHN0eWxlISAgR28gZ3JhYiB5b3Vyc2VsZiBhIFRhaWxzY2FsZSB0LXNoaXJ0IGlmIHRoZXJlIGFyZQovLyBzdGlsbCBzb21lIGF2YWlsYWJsZS4gQnV0IHNoaGguLi4gZG9uJ3QgdGVsbCBhbnlvbmUhCi8vCi8vICAgICAgICAgICAgIGh0dHBzOi8vc3dhZy5jb20vZ2lmdC82a29mNGs1Z3B1ZW95ZDB2NXd6MHJkYmMKewoJLy8gRGVjbGFyZSBzdGF0aWMgZ3JvdXBzIG9mIHVzZXJzIGJleW9uZCB0aG9zZSBpbiB0aGUgaWRlbnRpdHkgc2VydmljZS4KCSJncm91cHMiOiB7CgkJImdyb3VwOmV4YW1wbGUiOiBbInVzZXIxQGV4YW1wbGUuY29tIiwgInVzZXIyQGV4YW1wbGUuY29tIl0sCgl9LAoKCS8vIERlY2xhcmUgY29udmVuaWVudCBob3N0bmFtZSBhbGlhc2VzIHRvIHVzZSBpbiBwbGFjZSBvZiBJUCBhZGRyZXNzZXMuCgkiaG9zdHMiOiB7CgkJImV4YW1wbGUtaG9zdC0xIjogIjEwMC4xMDAuMTAwLjEwMCIsCgl9LAoKCS8vIEFjY2VzcyBjb250cm9sIGxpc3RzLgoJImFjbHMiOiBbCgkJLy8gTWF0Y2ggYWJzb2x1dGVseSBldmVyeXRoaW5nLgoJCS8vIENvbW1lbnQgdGhpcyBzZWN0aW9uIG91dCBpZiB5b3Ugd2FudCB0byBkZWZpbmUgc3BlY2lmaWMgcmVzdHJpY3Rpb25zLgoJCXsiYWN0aW9uIjogImFjY2VwdCIsICJ1c2VycyI6IFsiKiJdLCAicG9ydHMiOiBbIio6KiJdfSwKCV0sCn0K
                            warnings:
                                - '"group:example": user not found: "user1@example.com"'
                                - '"group:example": user not found: "user2@example.com"'
                            errors: null
            application/hujson:
                schema:
                    type: string
                examples:
                    exampleHuJSONPolicyFile:
                        $ref: '#/components/examples/ExampleHuJSONPolicyFile'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
