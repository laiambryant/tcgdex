package tcgdex

import (
"github.com/laiambryant/tcgdex/client"
"github.com/laiambryant/tcgdex/endpoint"
"github.com/laiambryant/tcgdex/models"
)

type TCGDex struct {
	Client *client.Client

	Card  *endpoint.Endpoint[models.Card, models.CardResume]
	Set   *endpoint.Endpoint[models.Set, models.SetResume]
	Serie *endpoint.Endpoint[models.Serie, models.SerieResume]
}

func New(opts ...client.Option) *TCGDex {
	c := client.NewHTTPClient(nil, opts...)
	sdk := &TCGDex{
		Client: c,
	}
	sdk.Card = endpoint.New[models.Card, models.CardResume](c, "cards")
	sdk.Set = endpoint.New[models.Set, models.SetResume](c, "sets")
	sdk.Serie = endpoint.New[models.Serie, models.SerieResume](c, "series")
	return sdk
}
