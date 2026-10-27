package adapters

func SummarizePositiveValues(values []int) int {
    total := 0
    accepted := 0
    rejected := 0
    for _, value := range values {
        if value > 0 {
            bounded := value
            if bounded > 100 {
                bounded = 100
            }
            total += bounded * 2
            accepted++
        } else {
            total -= value * value
            rejected++
        }
    }
    adjusted := total + accepted - rejected
    if adjusted < 0 {
        return 0
    }
    return adjusted
}
