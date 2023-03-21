CREATE TABLE history2(
	id int not null primary key auto_increment,
	num1 float not null,
	operation varchar(5) not null,
	num2 float not null,
	resultado float not null,
	fecha datetime not null
);