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

exports.sieve = function (req, res) {
	var n = req.query.n || (req.body && req.body.n);
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
		var rec = https.request(options, function (resp) {
			var body = "";
			resp.on('data', function (b) {
				body += b;
			})
			resp.on('end', function() {
				status = resp.statusCode;
				if (status != 200) {
					res.status(500).json({"error": body}).end();
					return
				}

				res.status(200).json(dur).end()
			});
		});
		rec.write("{" +
			'"service": "google",' +
			'"name": "sieve-' + n + '",' +
			'"value":' + dur +
			"}")
		rec.on("error", function() {
			res.status(500).send("cannot make request to recorder");
			return
		});
		rec.end()
	}
	else {
		res.status(400).send("Please pass a value for N in the request body");
	}
};
