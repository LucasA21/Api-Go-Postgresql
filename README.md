# README

#### Init

- `Inicializar mod: ` go mod init ....
- `Obtener paquete golira mux :` go get -u github.com/gorilla/mux
- `Instalar ORM - GORM`: go get -u gorm.io/gorm
    - `Instalar driver de postgres:` go get -u gorm.io/driver/postgres

- `**(OPCIONAL)** Instalar paquete Go - Air`:  go install github.com/air-verse/air@latest 

    Este paquete sirve para no tener que parar la consola y volverla a ejecutar cada vez que se quiere visualizar un cambio. Para inicializarlo se usa: `air init` y luego se corre una sola ves el programa con el comando: `air`


#### Database

- Crear un contenedor de docker con una imagen de postgres: `docker run --name some-postgres -e POSTGRES_USER=lucas -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`

- Crear base de datos : 
    - ejecutar la imagen en bash: `docker exec -it some-postgres bash`
    - conectar a bdd : `psql -U lucas --password`
    - crear bdd: `CREATE DATABASE gorm;`