## Kernel
 Se le considera el corazón del sistema operativo, ya que es la que actúa como sinapsis entre el hardware y el software. Sin el hardware el software no tendrá potencia para realizar alguna acción, con solo el cable de poder, que es hardware, no podrá realizar ni ejecutar ninguna aplicación. Por el otro lado, se puede tener mucho hardware, pero sin nada de software programado, no servirá de nada tener tanta potencia sin nada que lo utilice.
 
 No solamente es el núcleo del sistema, sino que también es un programa que controla los accesos al procesador y la memoria, es responsable de los drivers, ejecutarlos y verificarlos, para poder acceder desde el software a los periféricos. Además de lo ya mencionado, es el encargado de gestionar y administrar los recursos de la forma más eficiente posible.
 
 La tarea principal del kernel es la del procesamiento paralelo de diferentes tareas, o lo que se conoce como “multitask”, por lo que debe de cumplir con el tiempo establecido para cada tarea, y siempre tener espacio para poder ejecutar una tarea extra.
 
 El kernel solo funciona como intermediario entre el software del sistema, el de la aplicación y las bibliotecas necesarias, debido a que solo con la carga del sistema operativo, ya es demasiado trabajo, por lo que la interfaz gráfica es totalmente independiente del 
 
 > Tipos de Kernel:
 Existen varios tipos de kernel, pero entre los más importantes se encuentran:
 1. Monolítico:
 Un kernel único y grande para que maneje todas las tareas. Tiene un gran rendimiento para todo, por lo que al añadir un nuevo servicio este tendrá que compilarse de nuevo, por lo que puede ser rápido, pero si llega a ocurrir algún error, este parará toda operación.
 2. Microkernel:
 En vista de los problemas que ocasionaba el kernel monolítico, el microkernel específicamente para solventar ese problema, para que en caso de cualquier fallo, o actualización, no sea necesario parar todo el sistema operativo.
 Al ser "micro", para que todas las funcioes del kernel monolítico fueran adaptadas, todos los otros servicios se ejecutan como otros procesos, y si alguno fallase cumpliera el cometido principal.
 Unas de las desventajas, como siempre para algo que tiende a mejorar, es que es de mayor complejidad de programación y sincronización entre los módulos que componen el microkernel y su acceso a la memoria. Al ser de menor tamaño, tiende a tener un menor rendimiento que uno completo.
 3. Kernel Híbrido
 Como su nombre lo indica es una combinación de 2 o más sistemas, en este caso, del monolítico y el MicroKernel, en donde se usa el kernel grande, con la respuesta más eficaz al basarse en memoria compartida y este mismo es modular, por lo que los procesos pueden ser detenidos sin tener que involucrar a todo el sistema-
