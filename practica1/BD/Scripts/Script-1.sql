SELECT * FROM juegos;

CREATE table games(
	id int not null primary key auto_increment,
	name varchar(250) not null,
	gender varchar(250) not null,
	multiplayer int not null
);



SELECT * FROM games;


drop table games ;
show databases;


-- Base de Datos 
CREATE database prueba55;
-- DROP database sounoproy2;

CREATE TABLE votes(
	id int not null primary key auto_increment,
	sede int not null,
	municipio varchar(50) not null,
	departamento varchar(50) not null,
	papeleta varchar(50) not null,
	partido varchar(50) not null
);

INSERT INTO votes(sede, departamento, municipio, papeleta, partido ) 
VALUES(1,'Guatemala','San Miguel Petapa','Rosada','VAMOS');




