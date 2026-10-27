package infra

import "github.com/go-chi/chi/v5/codegov_acceptance/simple/domain"

func Load(value int) int {
    return domain.NormalizeAmount(value)
}
