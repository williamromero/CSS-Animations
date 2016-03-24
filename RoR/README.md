[ RUBY ON RAILS ]

	Es un framework de Ruby que provee módulos de producción.
	Es una serie de herramientas creada por Heinemeier Hansson.

RoR - 1.0 	/ 2005
RoR - 4.2.6 / 2016

Aplicaciones que usan Rails
	- Basecamp
	- Github
	- Shopify
	- Airbnb
	- Twitch
	- Soundcloud
	- Hulu

Porqué Rails?

	- Optimizado para la felicidad del programador.
	- Autogeneración de funciones para reducir la configuración del usuario.
	- Orientado a convenciones.
	- Sintaxis Ruby + Uso de Rails = So much win

Arquitectura MVC
(Modelo - Vista - Controlador)
	- MODELO: Es la representación de la información con la cual se opera.
	 		  Conexión con la DB.

	- CONTROLADORES: Es el encargado de manejar las peticiones HTTP y crear una respuesta adecuada. Son los que realizan las acciones luego del envío de un formulario, la interacción de rutas y peticiones de servidor.

		/usuario /perfil ---> Controlador ---> Modelos: ConDB ---> info ---> Vista

	- VISTA: Es la parte en la que se presentan los recursos de la aplicación.

FIRST TIME RUNNING
	* rails new musicapp
	* cd musicapp
	* ls -a
	* rails server
