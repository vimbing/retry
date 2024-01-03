package retry

import "time"

var PremadeRetryLow = Retrier{Max: 3, Delay: time.Second * 1}
var PremadeRetryMedium = Retrier{Max: 10, Delay: time.Second * 1}
var PremadeRetryHard = Retrier{Max: 20, Delay: time.Second * 1}
