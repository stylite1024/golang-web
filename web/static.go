package web

import "embed"

//go:embed frontend/*
var FrontendFS embed.FS

//go:embed backend/*
var BackendFS embed.FS
