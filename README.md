# Simulador de Estacionamiento en Go

Este proyecto es un simulador de estacionamiento desarrollado en el lenguaje Go, diseñado como una plataforma de aprendizaje para manejar y aplicar conceptos de concurrencia. La visión fue desarrollar un sistema capaz de gestionar múltiples vehículos que entran, estacionan y salen, usando técnicas de concurrencia para reflejar un comportamiento eficiente y realista de un estacionamiento.

## Mi Experiencia

El desarrollo del simulador fue una experiencia significativa que me permitió profundizar en la programación concurrente en Go y en la implementación de patrones de diseño. Desde el inicio, fue un desafío aplicar técnicas avanzadas de concurrencia para gestionar múltiples vehículos en un espacio compartido. Implementé el patrón **Observer** para separar responsabilidades entre componentes, lo que facilitó la comunicación entre el modelo de datos y la interfaz gráfica, manteniendo una actualización en tiempo real del estado de los vehículos.

Para la interfaz gráfica, opté por la librería **Fyne**, una herramienta eficiente y flexible que simplificó el desarrollo visual del simulador. Fyne permitió construir una interfaz amigable, en la cual los usuarios pueden observar el flujo de entrada y salida de vehículos y ver los espacios ocupados y libres en tiempo real. Esto añadió una dimensión interactiva y visual al proyecto, reforzando el aprendizaje práctico de los conceptos teóricos.

Uno de los mayores desafíos fue el control de acceso mediante una puerta compartida, que gestionaba la entrada y salida simultánea de vehículos. Para resolver esto, utilicé **Locks** y **Semáforos**, asegurando que los recursos compartidos no causaran condiciones de carrera. La sincronización adecuada fue clave para que el sistema operara de forma segura y eficiente. Con estas técnicas, pude implementar funciones como el bloqueo de vehículos cuando el estacionamiento está lleno, la asignación de espacios, y la liberación de los mismos al salir un vehículo, todo en un entorno concurrente.

En conclusión, este proyecto no solo cumplió con los objetivos iniciales, sino que también fue una valiosa experiencia educativa, brindándome las bases para enfrentar futuros proyectos en programación concurrente.

## Funcionalidades Principales

- **Generación de Vehículos**: Se pueden crear hasta 100 vehículos, simulando un flujo de entrada y salida constante.
- **Gestión de Capacidad**: El sistema administra un estacionamiento con una capacidad de 20 vehículos, bloqueando el acceso si está lleno.
- **Control de Acceso**: La puerta de acceso funciona como un recurso compartido, gestionando de forma eficiente la entrada y salida.
- **Asignación y Liberación de Espacios**: Los vehículos se asignan a espacios disponibles y se libera el espacio cuando salen.
- **Interfaz Visual**: La interfaz en tiempo real permite observar el comportamiento del estacionamiento y el estado de los vehículos.

## Arquitectura del Proyecto

El proyecto está organizado en una estructura modular para facilitar el mantenimiento y escalabilidad:

- **assets**: Recursos estáticos como imágenes.
- **models**: Definición y funcionamiento de las entidades, incluyendo el vehículo.
- **scene**: Configuración inicial de la escena visual.
- **view**: Representación gráfica de las entidades dinámicas, como los vehículos.

Esta estructura, junto con el uso del patrón **Observer**, permite la comunicación y sincronización entre el modelo y la vista, manteniendo la coherencia del estado en tiempo real.

## Requisitos

- Go 1.x
- Fyne (para la interfaz gráfica)

## Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/diegobejar1011/projectParking.git

