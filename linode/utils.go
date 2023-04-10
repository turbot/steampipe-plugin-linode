package linode

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/linode/linodego"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (linodego.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "linode"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(linodego.Client), nil
	}

	// Default to the env var settings
	token := os.Getenv("LINODE_TOKEN")

	// Prefer config settings
	linodeConfig := GetConfig(d.Connection)
	if linodeConfig.Token != nil {
		token = *linodeConfig.Token
	}

	// Error if the minimum config is not set
	if token == "" {
		return linodego.Client{}, errors.New("token must be configured")
	}

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	oauth2Client := &http.Client{
		Transport: &oauth2.Transport{
			Source: tokenSource,
		},
	}

	conn := linodego.NewClient(oauth2Client)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func base64DecodedData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data, err := base64.StdEncoding.DecodeString(types.SafeString(d.Value))
	if err != nil {
		return nil, nil
	}
	return data, nil
}

func isNotFoundError(err error) bool {
	return err.Error() == "[404] Not found"
}
