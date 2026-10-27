package clients

import "github.com/go-chi/chi/v5/codegov_acceptance/simple/domain"

func Quote(value int) int { return domain.NormalizeAmount(value) }
