function eratosthenes(n) {
	var sieve = [];
	var output = [];

	for (var i = 0; i < n; i++) {
		sieve.push(i);
	}
	var limit = Math.sqrt(n);
	for (var i = 2; i <= limit; i++) {
		if (sieve[i]) {
			for (var j = i * i; j < n; j += i) {
				sieve[j] = false;
			}
		}
	}
	var arr = [];
	for (var k = 2; k < sieve.length; k++) {
		if (sieve[k]) {
			arr.push(sieve[k]);
		}
	}
	return arr;
};

module.exports = function (context, req) {
	var n = req.query.n || (req.body && req.body.n);
	if (n) {
		var start = performance.now();
		eratosthenes(n);
		var elapsed = performance.now() - start;
		context.res = {
			// status: 200, /* Defaults to 200 */
			body: elapsed,
		};
	}
	else {
		context.res = {
			status: 400,
			body: "Please pass a value for N in the request body"
		};
	}
	context.done();
};
