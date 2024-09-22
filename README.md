# Gestor de Lista de Reproducción de Canciones

## Visión General
Este proyecto implementa un Gestor de Lista de Reproducción de Canciones utilizando estructuras de datos eficientes. Utiliza Golang para el backend y React para el frontend, integrados a través del framework Wails, resultando en una aplicación de escritorio robusta y dinámica.

## Tecnologías Utilizadas
- **Backend**: Golang
- **Frontend**: React
- **Framework de Aplicación de Escritorio**: Wails

## Características Principales
- Gestión y recuperación eficiente de canciones
- Creación y manipulación dinámica de listas de reproducción
- Capacidades avanzadas de búsqueda y ordenamiento

## Estructuras de Datos Implementadas
1. **ArrayList**: Utilizado para el almacenamiento dinámico de canciones.
2. **Lista Enlazada Simple**: Implementada para operaciones eficientes de inserción y eliminación.
3. **HashMap**: Utilizado para búsquedas rápidas de pares clave-valor.
4. **Índice Invertido**: Permite búsquedas de canciones rápidas y flexibles.
5. **Árbol B+**: Facilita el ordenamiento eficiente y consultas de rango sobre atributos de canciones.
6. **Árbol Trie**: Implementado para búsquedas de canciones basadas en prefijos.

## Algoritmos
- **QuickSort**: Utilizado para el ordenamiento eficiente de canciones basado en varios atributos.

## Diseño del Sistema
El sistema se divide en dos componentes principales:
1. **Lógica Interna**: Maneja las estructuras de datos y las funcionalidades centrales.
2. **Interfaz Gráfica**: Proporciona interacción del usuario a través de componentes React.

### Archivos Clave
- `main.go`: Punto de entrada de la aplicación.
- `app.go`: Archivo de implementación principal que contiene la estructura `App` y los métodos asociados.

## Detalles de Implementación

### Carga de Datos
- Las canciones se cargan desde un archivo CSV utilizando la función `ReadCSV`.
- Cada canción se procesa y se agrega simultáneamente a varias estructuras de datos.

### Búsqueda
- Utiliza tanto el Índice Invertido como el Árbol Trie para búsquedas eficientes.
- El Índice Invertido permite búsquedas flexibles basadas en palabras.
- El Árbol Trie está optimizado para búsquedas basadas en prefijos.

### Ordenamiento
- Se utilizan Árboles B+ para ordenar canciones según año, popularidad y duración.
- El algoritmo QuickSort se emplea para el ordenamiento de la lista de reproducción.

### Gestión de Lista de Reproducción
- Funciones para agregar, eliminar y limpiar canciones de la lista de reproducción.
- Soporta generación aleatoria de listas de reproducción.

## Comparaciones de Rendimiento
- Árbol Trie vs. Índice Invertido:
  - Árbol Trie: Eficiencia de búsqueda O(m), donde m es la longitud de la palabra de búsqueda.
  - Índice Invertido: Tiempo de búsqueda promedio O(1) para palabras exactas.

## Interfaz de Usuario
- Dos vistas principales: Lista general de canciones y Lista de reproducción.
- Las características incluyen ordenamiento, búsqueda y gestión de lista de reproducción.
- Implementa paginación para el manejo y visualización eficiente de datos.

## Conclusiones
- La combinación de Golang y React con Wails demuestra ser efectiva para el desarrollo de aplicaciones de escritorio.
- Las estructuras de datos implementadas optimizan significativamente la organización, búsqueda y ordenamiento de canciones.
- El Árbol B+ sobresale en el manejo de grandes volúmenes de datos.
- El Índice Invertido y el Árbol Trie ofrecen búsquedas rápidas y precisas, con algunas limitaciones para el Árbol Trie.

## Trabajo Futuro
- Mayor optimización de los algoritmos de búsqueda.
- Mejora de la interfaz de usuario para una mejor experiencia del usuario.
- Integración con servicios de streaming de música.

## Referencias
[Lista de referencias académicas como se proporciona en el documento original]

## Repositorio GitHub
Para una implementación detallada del código y documentación adicional, visite nuestro [repositorio GitHub](https://github.com/eluqm/EDA2024_2).



## Diagrama de estructuras 
<img src="https://res.cloudinary.com/dazt6g3o1/image/upload/v1720733134/t05iw7g2gdw2y50k2rep.png">

## Como descargar los datos
install python and pip

```
sudo apt install unzip

pip install gdown

gdown --id 1oJxvXaAtfniambjnVV-i26_okcNTeGWs

unzip archive.zip

mkdir data
mv spotify_data.csv ./data/data.csv

rm archive.zip
```
