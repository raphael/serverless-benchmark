const US_PER_SEC = 1e6;
const NS_PER_US = 1e3;
const https = require('https');

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

exports.handler = function(event, context, callback) {
	var n = event
	if (n) {
		var start = process.hrtime()
		eratosthenes(n);
		var elapsed = process.hrtime(start);
		var dur = elapsed[0] * US_PER_SEC + elapsed[1] / NS_PER_US;

		var options = {
			host: "optima-tve.appspot.com",
			port: 443,
			path: '/data',
			method: 'POST'
		};
		var req = https.request(options, function (res) {
			var body = "";
			res.on('data', function (b) {
				body += b;
			})
			res.on('end', function() {
				status = res.statusCode;
				if (status != 200) {
					callback(body);
					context.done();
					return;
				}

				callback(null, dur);
				context.done();
			});
		});
		req.write("{" +
			'"service": "lambda",' +
			'"name": "sieve-' + n + '",' +
			'"value":' + dur +
			"}")
		req.end()
	}
	else {
		callback("Please pass a value for N in the request body")
		context.done();
	}
};
