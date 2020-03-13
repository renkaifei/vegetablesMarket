use vegetablesmarket

create table if not exists vegetablecategory(
	categeoryid int not null auto_increment,
	categoryname varchar(256) not null,
	parentid int not null,
	primary key(categoryid)
)

create table if not exists vegetable(
	vegetableid int not null auto_increment,
	vegetablename varchar(256) not null,
	categoryid int not null
)

create table if not exists merchant(
	merchantid int not null auto_increment,
	merchantname varchar(256) not null,
	merchantaddress varchar(512) not null
	primary key(merchantid)
)
