package model

type BuilderOption func(*App) error

type NullableBoolean struct{ Value bool }