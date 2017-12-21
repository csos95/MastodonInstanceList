let globalObject = {
	instances: [],
	selectedTopic: 'all'
}

let instances = new Vue({
	el: '#instances',
	data: globalObject,
	computed: {
		topics: function() {
			let topics = ['all'];
			for (let i = 0; i < this.instances.length; i++) {
				if (topics.indexOf(this.instances[i].topic) === -1) {
					topics.push(this.instances[i].topic);
				}
			}
			return topics;
		},
		currentInstances: function () {
			let instances = [];
			if (this.selectedTopic === 'all') {
				return this.instances;
			}
			for (let i = 0; i < this.instances.length; i++) {
				if (this.instances[i].topic === this.selectedTopic) {
					instances.push(this.instances[i]);
				}
			}
			return instances;
		}
	},
	methods: {
		instanceurl(instance) {
			return "http://" + instance.uri;
		},
		lastActivity(instance) {
			if (instance.stats.datetime === '0001-01-01T00:00:00Z') {
				return 'never';
			}
			let now = new Date();
			let last = new Date(instance.stats.datetime);
			let diff = parseInt(now - last);
			let seconds = diff / 1000;
			let minutes = seconds / 60;
			let hours = minutes / 60;
			let days = hours / 24;
			if (days > 1) {
				return Math.trunc(days) + ' days ago';
			}
			if (hours > 1) {
				return Math.trunc(hours) + ' hours ago';
			}
			if (minutes > 1) {
				return Math.trunc(minutes) + ' minutes ago';
			}
			return Math.trunc(seconds) + ' seconds ago';
		}
	},
	created: function () {
		readInstances();
	}
})