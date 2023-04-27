SELECT * FROM juegos;

CREATE table games(
	id int not null primary key auto_increment,
	name varchar(250) not null,
	gender varchar(250) not null,
	multiplayer int not null
);



SELECT * FROM games;


drop table games ;