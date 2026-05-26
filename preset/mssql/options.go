package mssql

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithAdminPassword sets administrator password that can be used to connect
// (default: Gn0m!ck~).
func WithAdminPassword(password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDatabase creates a database with the provided name in the container. If
// not provided, "mydb" is used by default.  WithQueries, if provided, runs
// against the new database.
func WithDatabase(db string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithQueries executes the provided queries against the database created with
// WithDatabase, or against default "mydb" database.
func WithQueries(queries ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLicense sets EULA acceptance state. To accept the license, use true. See
// https://hub.docker.com/_/microsoft-mssql-server?tab=description for more
// information.
func WithLicense(accept bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithQueriesFile sets a file name to read initial queries from. Queries from
// this file are executed before any other queries provided in WithQueries.
func WithQueriesFile(file string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }
