package assets

import "embed"

//go:embed dist/*
var fs embed.FS

var root = "dist"

type Storage struct {
	Fs   embed.FS
	Root string
}

func GetStorage() Storage {
	var storage Storage
	storage.Fs = fs
	storage.Root = root
	return storage
}
