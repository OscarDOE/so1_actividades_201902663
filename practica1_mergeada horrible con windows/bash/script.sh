#!/bin/bash

# Cantidad de filas sin contar el encabezado
total_filas=$(cat ../back/home/oscar/Desktop/reportes.csv | wc -l)
cantidad_filas=$((total_filas - 1))

echo "Cantidad de filas: $cantidad_filas"

# Cantidad de operaciones totales con error
cantidad_error=$(awk -F, '{if ($4 == "-123457000.00") {print $0}}' ../back/home/oscar/Desktop/reportes.csv | wc -l)

echo "Cantidad de operaciones con error: $cantidad_error"

# Cantidad de operaciones por separado
cantidad_suma=$(awk -F, '{if ($2 == "+" && $4 != "-12345700.00") {print $0}}' ../back/home/oscar/Desktop/reportes.csv | wc -l)
cantidad_resta=$(awk -F, '{if ($2 == "-" && $4 != "-12345700.00") {print $0}}' ../back/home/oscar/Desktop/reportes.csv | wc -l)
cantidad_multiplicacion=$(awk -F, '{if ($2 == "*" && $4 != "-12345700.00") {print $0}}' ../back/home/oscar/Desktop/reportes.csv | wc -l)
cantidad_division=$(awk -F, '{if ($2 == "/" && $4 != "-12345700.00") {print $0}}' ../back/home/oscar/Desktop/reportes.csv | wc -l)

echo "Cantidad de operaciones de suma: $cantidad_suma"
echo "Cantidad de operaciones de resta: $cantidad_resta"
echo "Cantidad de operaciones de multiplicacion: $cantidad_multiplicacion"
echo "Cantidad de operaciones de division: $cantidad_division"

# Mostrar los datos del dÃ­a de hoy
fecha=$(date +%Y-%m-%d)

awk -v fecha="$fecha" -F, '{if ($5 ~ fecha) {print $0}}' ../back/home/oscar/Desktop/reportes.csv

echo  "SSSS: $fecha"

pwd

ls