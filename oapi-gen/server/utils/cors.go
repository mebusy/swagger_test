package utils

import (
	"github.com/rs/cors"
)

var CorsObj = cors.New(cors.Options{
	AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "token"},
})
