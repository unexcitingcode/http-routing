# http-routing

A completely abstract http method and path routing DSL.

This DSL has a focus on simplicity. Simplicity meaning the only thing it does
is express a language for routing http requests. It doesn't make an effort to
complicate itself with the concerns of data validation, request body parsing,
header extraction, response body serialization, etc. All fo these other
concerns can and should be expressed in their own independent DSLs elsewhere.
