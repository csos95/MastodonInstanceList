package main

var (
	createInstanceTableSQL = `
	create table if not exists Instance (
		instance_id int not null primary key auto_increment,
		title varchar(128) not null,
		uri varchar(40) not null,
		description text not null,
		email varchar(128) not null,
		version varchar(20) not null,
		thumbnail varchar(200) not null,
		topic varchar(100) not null,
		note text not null,
		registration varchar(20) not null,
		unique(uri)
	)`

	createStatsTableSQL = `
	create table if not exists Stats(
		id int not null primary key auto_increment,
		instance_id int not null,
		datetime datetime not null,
		users int not null,
		statuses int not null,
		connections int not null,
		foreign key (instance_id) references Instance(instance_id)
	)`

	//CREATE
	createInstanceSQL = `insert into Instance values (null, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// READ
	readInstancesSQL     = `select * from Instance;`
	readInstanceStatsSQL = `select datetime, users, statuses, connections from Stats where instance_id = ? order by datetime desc limit 1`

	// UPDATE
	updateInstanceSQL = `update Instance set title = ?, description = ?, email = ?, version = ?, thumbnail = ?, topic = ?, note = ?, registration = ? where instance_id = ?`
	updateStatsSQL    = `insert into Stats values (null, ?, ?, ?, ?, ?)`

	// DELETE
)
