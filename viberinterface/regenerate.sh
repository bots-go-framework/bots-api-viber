#!/usr/bin/env bash
rm keyboard_ffjson.go
rm callbacks_ffjson.go
rm messages_ffjson.go
ffjson -force-regenerate keyboard.go
ffjson -force-regenerate callbacks.go
ffjson -force-regenerate messages.go