package storeapi

// To regenerate the API and a Chi server use the following generate command
//go:generate oapi-codegen -config oapi-codegen.cfg.yaml openapi.yaml
//
// The generated Chi server components are not used to construct the server it
// is instead used to verify our web handlers cover the API completely.
