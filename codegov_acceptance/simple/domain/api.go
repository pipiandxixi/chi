package domain

func NormalizeAmount(value int, tenant string) int {
    if value < 0 {
        return 0
    }
    return value
}
