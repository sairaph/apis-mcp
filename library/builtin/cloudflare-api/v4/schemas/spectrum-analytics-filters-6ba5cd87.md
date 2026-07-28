---
title: spectrum-analytics_filters
page_id: schema-spectrum-analytics-filters-6ba5cd87
path: schemas
description: |-
    Used to filter rows by one or more dimensions. Filters can be combined using OR and AND boolean logic. AND takes precedence over OR in all the expressions. The OR operator is defined using a comma (,) or OR keyword surrounded by whitespace. The AND operator is defined using a semicolon (;) or AND keyword surrounded by whitespace. Note that the semicolon is a reserved character in URLs (rfc1738) and needs to be percent-encoded as %3B. Comparison options are:

    Operator                  | Name                            | URL Encoded
    --------------------------|---------------------------------|--------------------------
    ==                        | Equals                          | %3D%3D
    !=                        | Does not equals                 | !%3D
    \>                        | Greater Than                    | %3E
    \<                        | Less Than                       | %3C
    \>=                       | Greater than or equal to        | %3E%3D
    \<=                       | Less than or equal to           | %3C%3D

    Use the above to construct filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_filters

Used to filter rows by one or more dimensions. Filters can be combined using OR and AND boolean logic. AND takes precedence over OR in all the expressions. The OR operator is defined using a comma (,) or OR keyword surrounded by whitespace. The AND operator is defined using a semicolon (;) or AND keyword surrounded by whitespace. Note that the semicolon is a reserved character in URLs (rfc1738) and needs to be percent-encoded as %3B. Comparison options are:

Operator                  | Name                            | URL Encoded
--------------------------|---------------------------------|--------------------------
==                        | Equals                          | %3D%3D
!=                        | Does not equals                 | !%3D
\>                        | Greater Than                    | %3E
\<                        | Less Than                       | %3C
\>=                       | Greater than or equal to        | %3E%3D
\<=                       | Less than or equal to           | %3C%3D

Use the above to construct filters.

```yaml
{"description": "Used to filter rows by one or more dimensions. Filters can be combined using OR and AND boolean logic. AND takes precedence over OR in all the expressions. The OR operator is defined using a comma (,) or OR keyword surrounded by whitespace. The AND operator is defined using a semicolon (;) or AND keyword surrounded by whitespace. Note that the semicolon is a reserved character in URLs (rfc1738) and needs to be percent-encoded as %3B. Comparison options are:\n\nOperator                  | Name                            | URL Encoded\n--------------------------|---------------------------------|--------------------------\n==                        | Equals                          | %3D%3D\n!=                        | Does not equals                 | !%3D\n\\>                        | Greater Than                    | %3E\n\\<                        | Less Than                       | %3C\n\\>=                       | Greater than or equal to        | %3E%3D\n\\<=                       | Less than or equal to           | %3C%3D\n\nUse the above to construct filters.", "type": "string", "example": "event==disconnect%20AND%20coloName!=SFO"}
```
