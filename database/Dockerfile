# pull de la imagen, para realizar creacion
FROM postgres:10.3
# esto sera lo primer que se ejecute, con alias 1.sql
COPY up.sql /docker-entrypoint-initdb.d/1.sql
# correr comando postgres de manera permanente, mantendra la base de datos corriendo para poder ser utilizada
CMD ["postgres"]