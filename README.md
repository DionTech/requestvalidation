# about
This package is a wrapper for the github.com/go-playground/validator package.

You can validate a request (by struct) and get a status ("success" or "error") and when error, a map[string][]string of errors as a response.
This might be useful for a json-formatted api response.