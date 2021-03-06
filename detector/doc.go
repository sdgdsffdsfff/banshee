// Copyright 2015 Eleme Inc. All rights reserved.

/*

Package detector is a tcp server to detect anomalies.

Detector Input protocol

Line based text, for example:

	timer.count_ps.get_user 1452674178 3.4

Detection Algorithms

A simple approach to detect anomalies is to set fixed thresholds, but
the problem is how large/small they should be. By this way, alertings
will be noisily and thresholds are also hard to maintain.

So we explore an automated way.

The well-known 3-sigma rule: http://en.wikipedia.org/wiki/68%E2%80%9395%E2%80%9399.7_rule.

States that nearly all values (99.7%) lie within 3 standard deviations of the mean in
a normal distribution.

That's to say: If the metric value deviates too much from average, it should be an anomaly!

	func IsAnomaly(value float64) bool {
		return math.Abs(value - mean) > 3 * stddev
	}

And we name the ratio of the distance to 3 times standard deviation as score:

	score = math.Abs(value - mean) / (3.0 * stddev)

If score > 1, that means the metric is currently anomalously trending up.

If score < -1, that means the metric is currently anomalously trending down.

If score is larger than -1 and less than 1, the metric is normal.

*/
package detector
