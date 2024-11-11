# Simulador de Estacionamiento en Go

Este proyecto es un simulador de estacionamiento desarrollado en el lenguaje Go. Fue creado con el objetivo de aprender a resolver desafíos relacionados con la concurrencia, aplicando patrones de diseño y utilizando herramientas eficaces para gestionar el flujo de trabajo en sistemas paralelos.

## Mi Experiencia

Desarrollar el proyecto de simulador de estacionamiento en el lenguaje Go fue una experiencia enriquecedora que me permitió profundizar en el manejo de la concurrencia y en la implementación de patrones de diseño para mejorar la estructura y funcionalidad del código. Este proyecto no solo me dio la oportunidad de resolver desafíos de sincronización y paralelismo, sino también de poner en práctica conceptos avanzados que son cruciales en el desarrollo de sistemas concurrentes.

Uno de los puntos clave fue la aplicación del patrón **Observer**, que me permitió organizar el código de manera modular y estructurada, separando las responsabilidades entre los distintos componentes del sistema. Gracias a este patrón, pude gestionar eventos de forma independiente, haciendo que el código fuese más flexible y sencillo de mantener.

Para la interfaz gráfica, opté por la librería **Fyne**. Esta elección me dio la oportunidad de explorar una herramienta sencilla y funcional, que me permitió desarrollar una interfaz amigable para el usuario sin perderme en complejidades excesivas de diseño gráfico. Fyne facilitó la construcción de un entorno visual intuitivo para el simulador, en el cual el usuario puede interactuar y observar el funcionamiento del estacionamiento en tiempo real.

Para abordar los desafíos de sincronización en la simulación, utilicé **Semáforos** y **Canales con buffer** , técnicas que me permitieron controlar el acceso a los recursos compartidos y manejar de manera segura el acceso concurrente. Cada uno de estos mecanismos resultó fundamental para asegurar que las operaciones se realizaran sin conflictos, evitando condiciones de carrera y otros problemas que suelen surgir en entornos de concurrencia. Fue una gran oportunidad para entender a fondo cómo funcionan estos métodos y cómo su aplicación puede influir en el rendimiento y estabilidad de un sistema concurrente.

En conclusión, este proyecto me permitió mejorar mis habilidades en concurrencia y diseño de software, al tiempo que me brindó una experiencia completa en el desarrollo de un simulador interactivo.

## Requisitos

- Go 1.x
- Fyne (para la interfaz gráfica)

## Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/tuusuario/simulador-estacionamiento.git
