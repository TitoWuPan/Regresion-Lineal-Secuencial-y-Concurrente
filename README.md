# Regresion-Lineal-Secuencial-y-Concurrente

Trabajo Final - Programación Concurrente y Distribuida - Junio 2024

## Integrantes :
- Guillén Rojas, Daniel Carlos		- U201920113
- Wu Pan, Tito Peng 			- U201921200
- Santisteban Cerna, José Mauricio	- U201922760

## Profesor :
Jara Garcia, Carlos Alberto


## Introducción
La regresión lineal es una herramienta fundamental en la estadística y el análisis de datos, y se utiliza en una amplia variedad de aplicaciones, como predicción de ventas, análisis de tendencias, modelado de fenómenos naturales, entre otros. En este proyecto, se utilizo datos de la Comisión de Taxis y Limusinas de Nueva York (TLC) del año 2016, se busca desarrollar un modelo de regresión lineal. Este modelo permitirá estimar con mayor precisión los tiempos de viaje, beneficiando tanto a pasajeros como a conductores en la planificación de rutas y horarios. El interfaz fue implementado en Angular y la comunicacion HTTP API esta conectado mediante Docker.

Para inicializar Angular se utiliza en CMD
```bash
ng serve -o
```

Para crear la imagen e inicializar el contenerdor con los puesrto en Docker
```bash
docker build -t go-app
docker run -p 8080:8080 -p 9090:9090 go-app
```

## Aplicación Web
![image](https://github.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/assets/91169600/37094a50-640e-4852-8dc1-bac6bf6f0782)

