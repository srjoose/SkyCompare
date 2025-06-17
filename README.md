
✈️ SkyCompare

SkyCompare es una plataforma para comparar rutas y precios de vuelos entre diferentes aeropuertos y compañías aéreas. 
Este repositorio incluye tanto el backend en Go como el frontend en Vue.

✅ Requisitos

Antes de comenzar, asegúrate de tener instalados en tu sistema:

- [Go](https://golang.org/dl/) (versión 1.20 o superior)
- [Node.js y npm](https://nodejs.org/)

🔙 Ejecutar el Backend

1. Abre una terminal y accede al directorio del backend: cd skycompare-backend-main
2. Crea un archivo .env en la raíz de esta carpeta con el siguiente contenido: DATABASE_URL=mysql://root:gPpKUcIGLpkhqgSclwXgHexbppJGlXUM@caboose.proxy.rlwy.net:18660/railway
3. Ejecuta el servidor: go run cmd/server/main.go

🔜 Ejecutar el Frontend

1. Abre una nueva terminal y accede al directorio del frontend: cd skycompare-frontend-main
2. Instala las dependencias del proyecto: npm install
3. Inicia el servidor de desarrollo: npm run serve


🌐 Acceder a la aplicación

Una vez arrancados el backend y el frontend, abre tu navegador y accede a: http://localhost:8080/
Desde ahí podrás utilizar SkyCompare con todas sus funcionalidades.

