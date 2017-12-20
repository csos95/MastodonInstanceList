function api_call(endpoint, options, success, failure) {
	fetch(endpoint, options)
		.then(response => response.json())
		.then(function (data) {
			console.log(data);
			if (data.status === 'success') {
				success(data);
			} else {
				failure(data);
			}
		});
}

// READ
function readInstances() {
	api_call('/api/instances', {
		method: 'GET',
	}, function (data) {
		globalObject.instances = data.instances;
	}, function (data) {
		console.log('failed to get instances');
	});
}