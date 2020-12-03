# gqlgen-sqlboiler

We want developers to be able to build software faster using modern tools like GraphQL, Golang, React Native without depending on commercial providers like Firebase or AWS Amplify.

This program generates type-safe code between your generated gqlgen and sqlboiler models with support for unique id's across your whole database (also if you're not using string id's). We can automatically generate the implementation of queries and mutations like create, update, delete based on your graphql scheme and your sqlboiler models.

To make this program a success coupling (same naming) between your database and graphql scheme is needed at the moment. The advantage of this program is the most when you have a database already designed. However everything is created with support for change so you could write some extra GrapQL resolvers if you'd like without a problem.

## Why gqlgen and sqlboiler

They go back to a schema first approach which we like. The generated code with these tools are the most efficient and fast in the Golang system (and probably outside of it too).

- sqlboiler: https://github.com/volatiletech/sqlboiler#benchmarks
- gqlgen: https://github.com/appleboy/golang-graphql-benchmark#summary

It's really amazing how fast a generated api with these techniques is!

## Usage

1. Generate database structs with: https://github.com/volatiletech/sqlboiler (--no-back-referencing is IMPORTANT!)
   e.g. `sqlboiler mysql --no-back-referencing`
2. Generate gqlgen structs + converts between gqlgen and sqlboiler with this program  
   e.g. `go run convert_plugin.go` for file contents of that program see bottom of this readme you can chose whether to generate the graphql schema itself too

## Features

- [x] Generate graphql.schema based on sqlboiler structs
- [x] Generate converts between sqlboiler structs and graphql (with relations included)
- [x] Generate converts between input models and sqlboiler
- [x] Genarated code understands the difference between empty and null for update inputs so you can set things empty if you explicicitly set them in your mutation!
- [x] Fetch sqlboiler preloads from graphql context
- [x] Support for foreign keys named differently than their corresponding model
- [x] New plugin which generates CRUD resolvers based on mutations in graphql scheme.
- [x] Support one-to-one relationships inside input types.
- [x] Generate code which implements the generated where and search filters
- [x] Batch update/delete generation in resolvers (Not tested yet).
- [x] Enum support.
- [x] public errors in resolvers + logging via zerolog. (feel free for PR for configurable logging!)

## v3.0

- [x] Edges/connections
- [x] Type safe sorting

## Roadmap v3.1

- [ ] Support gqlgen multiple .graphql files
- [ ] Support for opting out of standard generation if functions exist in {resolverName}\_custom.go

## Roadmap v3.2

- [ ] Adding automatic database migrations and integration with https://github.com/web-ridge/dbifier so faster iteration is possible
- [ ] Crud of adding/removing relationships from many-to-many on edges
- [ ] Support more relationships inside input types
- [ ] Generate tests
- [ ] Run automatic tests in Github CI/CD in https://github.com/web-ridge/gqlgen-sqlboiler-examples
- [ ] Batch create generation in resolvers (based on https://github.com/web-ridge/contact-tracing/blob/master/backend/helpers/convert_batch.go)

## Examples

https://github.com/web-ridge/gqlgen-sqlboiler-examples

More examples are welcome. Send a PR ;-)

### Example result of this plugin

https://github.com/web-ridge/gqlgen-sqlboiler-examples/tree/master/social-network/helpers

**Code snippet**

```go
func PostToGraphQL(m *models.Post) *graphql_models.Post {
	if m == nil {
		return nil
	}

	r := &graphql_models.Post{
		ID:      PostIDToGraphQL(m.ID),
		Content: m.Content,
	}

	if boilergql.UintIsFilled(m.UserID) {
		if m.R != nil && m.R.User != nil {
			r.User = UserToGraphQL(m.R.User)
		} else {
			r.User = UserWithUintID(m.UserID)
		}
	}
	if m.R != nil && m.R.Comments != nil {
		r.Comments = CommentsToGraphQL(m.R.Comments)
	}
	if m.R != nil && m.R.Images != nil {
		r.Images = ImagesToGraphQL(m.R.Images)
	}
	if m.R != nil && m.R.Likes != nil {
		r.Likes = LikesToGraphQL(m.R.Likes)
	}

	return r
}
```

sqlboiler.yml

```yaml
mysql:
  dbname: dbname
  host: localhost
  port: 8889
  user: root
  pass: root
  sslmode: "false"
  blacklist:
    - notifications
    - jobs
    - password_resets
    - migrations
mysqldump:
  column-statistics: 0
```

gqlgen.yml

```yaml
schema:
  - schema.graphql
exec:
  filename: graphql_models/generated.go
  package: graphql_models
model:
  filename: graphql_models/genereated_models.go
  package: graphql_models
resolver:
  filename: resolver.go
  type: Resolver
models:
  ConnectionBackwardPagination:
    model: github.com/web-ridge/utils-go/boilergql.ConnectionBackwardPagination
  ConnectionForwardPagination:
    model: github.com/web-ridge/utils-go/boilergql.ConnectionForwardPagination
  ConnectionPagination:
    model: github.com/web-ridge/utils-go/boilergql.ConnectionPagination
  SortDirection:
    model: github.com/web-ridge/utils-go/boilergql.SortDirection
```

Put this in your program convert_plugin.go e.g.

```go
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	gbgen "github.com/web-ridge/gqlgen-sqlboiler/v3"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	output := gbgen.Config{
		Directory:   "helpers", // supports root or sub directories
		PackageName: "helpers",
	}
	backend := gbgen.Config{
		Directory:   "models",
		PackageName: "models",
	}
	frontend := gbgen.Config{
		Directory:   "graphql_models",
		PackageName: "graphql_models",
	}

	if err = gbgen.SchemaWrite(gbgen.SchemaConfig{
		BoilerModelDirectory: backend,
		// Directives:           []string{"IsAuthenticated"},
		// GenerateBatchCreate:  false, // Not implemented yet
		GenerateMutations:    true,
		GenerateBatchDelete:  true,
		GenerateBatchUpdate:  true,
	}, "schema.graphql", gbgen.SchemaGenerateConfig{
		MergeSchema: false, // uses three way merge to keep your customization
	}); err != nil {
		fmt.Println("error while trying to generate schema.graphql")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	if err = api.Generate(cfg,
		api.AddPlugin(gbgen.NewConvertPlugin(
			output,   // directory where convert.go, convert_input.go and preload.go should live
			backend,  // directory where sqlboiler files are put
			frontend, // directory where gqlgen models live
			gbgen.ConvertPluginConfig{
				// UseReflectWorkaroundForSubModelFilteringInPostgresIssue25: true, // see issue #25 on GitHub
			},
		)),
		api.AddPlugin(gbgen.NewResolverPlugin(
			output,
			backend,
			frontend,
			"", // leave empty if you don't have auth
		)),
	); err != nil {
		fmt.Println("error while trying generate resolver and converts")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
```

`go run convert_plugin.go`

## Help us

We're the most happy with your time investments and/or pull request to improve this plugin. Feedback is also highly appreciated.

If you don't have time or knowledge to contribute and we did save you a lot of time, please consider a donation so we can invest more time in this library: [![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=7B9KKQLXTEW9Q&source=url)
