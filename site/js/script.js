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
		}
	},
	created: function () {
		readInstances();
	}
})