//go:generate go run ../../testdata/gqlgen.go -config testdata/usefunctionsyntaxforexecutioncontext/gqlgen.yml
package federation

import (
	"testing"

	"github.com/dmitry-drozdov/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"

	"github.com/dmitry-drozdov/gqlgen/client"
	"github.com/dmitry-drozdov/gqlgen/graphql/handler"
	"github.com/dmitry-drozdov/gqlgen/plugin/federation/testdata/usefunctionsyntaxforexecutioncontext"
	"github.com/dmitry-drozdov/gqlgen/plugin/federation/testdata/usefunctionsyntaxforexecutioncontext/generated"
)

func TestFederationWithUseFunctionSyntaxForExecutionContext(t *testing.T) {
	srv := handler.New(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &usefunctionsyntaxforexecutioncontext.Resolver{},
		}),
	)
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	t.Run("Hello entities", func(t *testing.T) {
		representations := []map[string]any{
			{
				"__typename": "Hello",
				"name":       "first name - 1",
			}, {
				"__typename": "Hello",
				"name":       "first name - 2",
			},
		}

		var resp struct {
			Entities []struct {
				Name string `json:"name"`
			} `json:"_entities"`
		}

		err := c.Post(
			entityQuery([]string{
				"Hello {name}",
			}),
			&resp,
			client.Var("representations", representations),
		)

		require.NoError(t, err)
		require.Equal(t, "first name - 1", resp.Entities[0].Name)
		require.Equal(t, "first name - 2", resp.Entities[1].Name)
	})
}
