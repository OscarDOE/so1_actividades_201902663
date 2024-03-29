SELECT * FROM ESTADO ;

DROP TABLE victimas;
DROP TABLE reg_victima_hospital;
DROP TABLE ubicacion_victima;
DROP TABLE tratamiento_victima;
DROP TABLE registro_asociado;
DROP TABLE registro_contacto;


INSERT INTO victimas(nombre,apellido,direccion,fecha_primera_sospecha,fecha_confirmacion,fecha_muerte,idestado)
SELECT DISTINCT DB_EXCEL.NOMBRE_VICTIMA,DB_EXCEL.APELLIDO_VICTIMA,DB_EXCEL.DIRECCION_VICTIMA,
TO_DATE(DB_EXCEL.FECHA_PRIMERA_SOSPECHA,'DD/MM/YYYY HH24:MI'),
TO_DATE(DB_EXCEL.FECHA_CONFIRMACION,'DD/MM/YYYY HH24:MI'),
TO_DATE(DB_EXCEL.FECHA_MUERTE,'DD/MM/YYYY HH24:MI'),
ESTADO.ID
FROM DB_EXCEL 
INNER JOIN ESTADO
	ON DB_EXCEL.ESTADO_VICTIMA = estado.DESCRIPCION 
WHERE NOMBRE_VICTIMA IS NOT NULL AND APELLIDO_VICTIMA IS NOT NULL
AND DIRECCION_VICTIMA IS NOT NULL AND FECHA_PRIMERA_SOSPECHA IS NOT NULL 
AND FECHA_CONFIRMACION IS NOT NULL;
SELECT * FROM VICTIMAS ORDER BY id;

SELECT DISTINCT victimas.id,victimas.nombre, DB_EXCEL.NOMBRE_VICTIMA ,DB_EXCEL.NOMBRE_HOSPITAL, hospitales.ID ,hospitales.NOMBRE 
FROM DB_EXCEL
	INNER JOIN hospitales
	ON DB_EXCEL.NOMBRE_HOSPITAL = hospitales.NOMBRE ;
	
	
	
	
	