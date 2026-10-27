package infra

import "github.com/go-chi/chi/v5/codegov_acceptance/multi/domain"

func Load(value int) int {
    return domain.NormalizeAmount(value)
}
