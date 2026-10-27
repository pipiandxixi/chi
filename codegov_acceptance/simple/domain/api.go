package domain

import "github.com/go-chi/chi/v5/codegov_acceptance/simple/infra"

func NormalizeAmount(value int) int {
    if value < 0 {
        return 0
    }
    return value
}

func LoadAmount(value int) int { return infra.Load(value) }
