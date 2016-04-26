##MODELO VISTA CONTROLADOR
###Ruby on Rails

Cada aplicación contiene en sí, una carpeta llamada "app" en la que se encuentran los folders "controllers", "models" y "views". Esto debido a que responde a la lógica MVC (aplicación de estructura).

* **Modelo:** Está formado por clases. Cada clase es un modelo y cada modelo, representa una tabla en la DB. Por lo tanto, es el encargado de trabajar con la lógica de la base de datos. Por otra parte, los modelos son "objetos" que se crean para la aplicación web.
Los modelos son la forma como se encapsulara la informacion, es decir es el tipo de informacion que finalmente vera el usuario. 

* **Vista:** Es la representación final de una petición. (HTML, ERB, JS, CSS) Es decir, la forma en la que se presentará la información. Asimismo, es como cuando en un archivo PHP utilizamos diferentes plantillas HTML para desplegar lo que se busca en la DB con scripts de PHP y MySQL. Por lo tanto, las vistas son las que despliegan las funciones que se requieren al explorador.

* **Controlador:** Es el encargado de la lógica de la petición. Es un puente entre la Vista y el modelo. Cada controlador, será una clase con métodos y por cada método habrá una vista que representará la versión procesada de ese método.
El controlador, por lo tanto se traduce en los archivos que operan acciones y scripts, haciendo llamadas a diferentes propiedades del objeto (las clases de la DB) y desplegandolas luego en la vista.

	
	Los controladores son como las operaciones que se van a realizar, es ahi donde se incluyen los metodos a los cuales se podran acceder por medio de una URL: Ejemplo de una url donde se llama a un controlador usuario y a su metodo login, con dos parametros: 
	http://misitioweb.com/usuario/login/{par
	 
	Un controlador puede acceder y utilizar una clase Modelo para pasar datos a las vistas. 

###Proceso de Desarrollo:

	**Paso 01:**
		-> Cuando se visita la página: 0.0.0.0:3000/users se emite una solicitud al servidor, esta solicitud llega al archivo 'config/routes.db'.
		Este, es un archivo que se encuentra en la carpeta config y mapea en base a la URL.
		El archivo "routes.rb", llamará a los recursos de acción del controlador Users o Posts en este ejemplo.

			MiblogDemo::Application.routes.draw do
				resources :posts
				resources :users
			end

		Inicialmente, la envía a la acción index del controlador al que corresponde, en este caso 'users_controller' o 'posts_controller'.

	Paso 02:
		-> Cuando nos dirigmos a las acciones de los controladores, podemos evaluar las acciones predefinidas si utilizamos "Scaffold".

		- def index
			@users = User.all
		end

		En este caso, la acción index solicita al "modelo User" que le provea de la lista de todos los usuarios y luego deposita ese valor en la variable @users.

	Paso 03:
		-> Hace la petición al modelo, que con sus clases, maneja los modelos, que a sus vez son tablas.

	Paso 04:
		-> Esta petición es recibida por la clase User, la cual hereda de la biblioteca "ActiveRecord", la cual realiza una acción "Mapeo - Base Relacional" (Mapea una tabla). La clase del modelo debe de declararse en singular debido a que la tabla deberá de reconocerse en plural.

			class User < ActiveRecord::Base
				has_many :posts
			end

	Paso 05:

		El modelo de User devuelve la lista de usuarios al controlador Users_Controller.

	Paso 06:

		El controlador deposita la lista en la variable @Users, la cual será impresa en la vista 'index', del controlador Users.

	Paso 07:

		La vista utiliza código embebido. Es decir que permite la incorporación del lenguaje Ruby, el cual traducirá a HTML.
		Las variables que comienzan con @, son variables de instancia de un método que representa una vista. En ella se contienen los elementos desplegar en la vista.

	Paso 08: 

		El controlador, pasa código HTML a la vista.
		
¿QUÉ ES RoR?

	- Es un framework creado en Ruby para desarrollar aplicaciones web.
	- Convención sobre Configuración: Significa que el lenguaje toma decisiones respecto a ubicación de archivos, funciones y acciones que originalmente se tomarían sin los elementos de dicha librería.

	Arquitectura de Rails - MVC

	Modelo: Es la representación de la información con la cual el sistema opera. Componentes contables y descriptivos que conviven con métodos para su manipulación.

	Vista:
	
	Controlador: Son los que se encargan de recibir las peticiones HTTP y de acuerdo a ello, realiza una acción para responder de forma adecuada.

El usuario, interactua con el controlador mediante las rutas (las páginas), luego los manipuladores se encargan de realizar la conexión con los modelos. Una vez esto hecho y la conexión con los modelos, el controlador se encarga de obtener los datos para la vista.


Al lío:

1. Rails new musicapp
2. Rails server (El contenido se correo en un servidor de aplicaciones llamado WebRick)


Estructura de Aplicación:
	- app => Se contiene todo el código relacionado a la lógica de la aplicación. Por convención, todo el contenido de controllers, models y views, se situan en esta carpeta.
		-> Assets: son todos los archivos que se requieren. (imgs, js, css)
		-> Controllers: Los controladores son los que se encargan de manejar todas las peticiones HTTP y la devolución de los modelos.
		-> Helpers: Se coloca toda la estructura de código que vamos a reutilizar en todas las vistas.
		-> Mailers: Se 
		-> Modelos: Contienen toda la lógica de la app y el contenido de las clases de las bases de datos.
	- bin =>
	- config => 
		-> Environments: Configuraciones que tenemos cuando estamos desarrollando. Estas son diferentes en prueba, en desarrollo y en producción. Por ejemplo, en método de desarrollo, las vistas se cargan en su totalidad, lo que no sucede en producción.
		-> Initializers: Configuraciones que necesitamos que se carguen al iniciar el servidor. 
			* database.yml => Es en donde se guarda la configuración de las bases de datos. Usuario, tipo de DB, contraseña, etc.
			* routes.rb => Son las rutas que se generan por cada controlador. 
			* secrets.yml => Es el archivo en el que se guardan todas las llaves de usuario.
	- db =>	Es donde se guardan las configuraciones de los modelos.
	- lib => Para librerías de terceros.
	- log => Registro de todas las interacciones con nuestra aplicación.
	- public => Se encuentran todos los archivos estáticos (404.html, 500.html) que no necesitan el procesamiento de Ruby.
	- test => 
	- tmp =>
	- vendor =>
			*gemfile = Es donde se especifican las gemas a funcionar en el proyecto. Se describen todas las dependencias de gemas a funcionar en el proyecto.
				- sass: css precompilado
				- coffee script: javascript precompilado al estilo Rails
				- turbolinks: gema para crear aplicaciones que no tengan que recargarse completas sino solo los elementos a modificar. Es una alternativa parecida a Angular o React.

SCAFFOLD 
	
	Es una función que servirá para crear una estructura de vistas, controladores para manejar los modelos.

	rails generate scaffold track title:string album:string artist:string
						 model_name		name_field:data_type

MIGRACIONES

	Las migraciones son las que le dicen a la DB que tablas y campos debe de generar tanto como cambiar tipo de datos de campos o eliminar una tabla o campo.

	Models-> Nombre en singular
	Tabla -> Nombre en plural

	class CreateTracks < ActiveRecord::Migration
	  def change
	    create_table :tracks do |t|
	      t.string :title
	      t.string :album
	      t.string :artist

	      t.timestamps null: false // hora de creación
	    end
	  end
	end

	rake db:migrate

	"LA FORMA DE INTERACTUAR CON LOS MODELOS, SON LAS RUTAS" por medio
	de los controladores.

	cd /config/routes.rb/

	Rails.application.routes.draw do
  	resources :tracks

  	rake routes

 EXPLICACIÓN MVC en RAILS

 	-> En routes, podemos ver las rutas que responden a las acciones del controlador.

 	Prefix	Verb	URI Pattern			Controller#Action
 	tracks  GET 	/tracks(.:format)	tracks#index

 	Lo que quiere decir, es que cuando vayamos a "/tracks", esta ruta corresponde a una acción de nuestro controlador y que tiene un método index.

 	"LOS MODELOS SON LOS OBJETOS QUE CONTIENEN LA LÓGICA DE NUESTRA APLICACIÓN"
 	
METODOS x CADA VISTA
	-> index 		-> update
	-> show			-> destroy
	-> news			-> create
 	-> set


ACTIVE RECORD -> Funciones de manipulación del modelo. Puede crear, editar, modificar tablas y campos.

	-> rake db:migrate
	   (guarda la configuración de la tabla)

RUTAS:

	resources :track -> Creador de rutas para nuestras acciones.
	Para ver nuestras rutas, presionar:

		rake routes

	Los verbos de HTTP son GET, POST, PUT & PATCH.

	Para ingresar a cada una de las vistas de las acciones, seguimos el patrón brindado por el comando rake routes.

	Las acciones o métodos, tomarán el valor de todas las instancias del modelo para luego delegar a la vista la representación del contenido.

MODELOS 

	-> SHOW: Produce la acción show, del controlador 'tracks', que va al filtro "before_action" y pide a la función "set_track", llame a la método find del modelo que le antecede:

		@track= Track.find(params[:id])

	Y eso lo guarda en la variable de instancia @track.

	Track-> El módelo y toda la abstracción del modelo. Esta hereda del ActiveRecord::Base 

	FILTRO before_action :func

	before_action :set_track, only: [:show, :edit, :update, :destroy]

		def set_track
	      @track = Track.find(params[:id])
	     #Instance = Modelo.find(parametro[:id])
	     #	Var
	    end

ACTIVE RECORD: Es un librería que nos permite conectar con la DB y realizar las peticiones. Los modelos heredan de esta clase para para poder responder a las peticiones del controlador.

	Por ejemplo:

		def index
    		@tracks = Track.all
  		  #Instance = Modelo.all (Del objeto Modelo, traer todo los objetos)
  		   	#Var
  		end

RAILS CONSOLE

	La consola de Rails, permite manipular los modelos que se han creado.

	rails console
	Model.connection -> Conecta con la DB.
	Track -> Este comando mostrará los campos que hemos creado

		Track.count
		(0.2ms) SELECT COUNT(*) FROM "tracks"
		=> 2

		Track.first
		Track Load (0.4ms)  SELECT  "tracks".* FROM "tracks"  ORDER BY "tracks"."id" ASC LIMIT 1
		=> #<Track id: 1, title: "Work", album: "The best", artist: "Rihanna", created_at: "2016-04-14 06:24:11", updated_at: "2016-04-14 06:24:11">


		Track.create title: 'Dejarte de amar', artist: 'Kalimba', album: 'Primero'
		   (0.2ms)  begin transaction
		  SQL (0.6ms)  INSERT INTO "tracks" ("title", "artist", "album", "created_at", "updated_at") VALUES (?, ?, ?, ?, ?)  [["title", "Dejarte de amar"], ["artist", "Kalimba"], ["album", "Primero"], ["created_at", "2016-04-16 01:54:41.597944"], ["updated_at", "2016-04-16 01:54:41.597944"]]
		   (1.4ms)  commit transaction
		 => #<Track id: 3, title: "Dejarte de amar", album: "Primero", artist: "Kalimba", created_at: "2016-04-16 01:54:41", updated_at: "2016-04-16 01:54:41"> 

		 Track.where(artist:"Camila").count
   		(0.3ms)  SELECT COUNT(*) FROM "tracks" WHERE "tracks"."artist" = ?  [["artist", "Camila"]]
 		=> 1 

 INSTANCIA DE CLASE MODELO

 		cancion = Track.new
 			=> #<Track id: nil, title: nil, album: nil, artist: nil, created_at: nil, updated_at: nil>

		 ancion.title = 'Five'
		 	=> "Five" 

		 cancion
 			=> #<Track id: nil, title: "Five", album: nil, artist: nil, created_at: nil, updated_at: nil> 

 		cancion.save

		    (0.1ms)  begin transaction
		  SQL (0.5ms)  INSERT INTO "tracks" ("title", "created_at", "updated_at") VALUES (?, ?, ?)  [["title", "Five"], ["created_at", "2016-04-16 02:41:06.876952"], ["updated_at", "2016-04-16 02:41:06.876952"]]
		   (1.1ms)  commit transaction		

		canciones = Track.all  // lista de todos los objetos del modelo Track

		 Track Load (0.3ms)  SELECT "tracks".* FROM "tracks"
 			=> #<ActiveRecord::Relation [#<Track id: 1, title: "Work", album: "The best", artist: "Rihanna", created_at: "2016-04-14 06:24:11", updated_at: "2016-04-14 06:24:11">, #<Track id: 2, title: "Todo cambio", album: "Primer Album", artist: "Camila", created_at: "2016-04-15 23:12:15", updated_at: "2016-04-15 23:12:15">, #<Track id: 3, title: "Dejarte de amar", album: "Primero", artist: "Kalimba", created_at: "2016-04-16 01:54:41", updated_at: "2016-04-16 01:54:41">, #<Track id: 4, title: "Five", album: nil, artist: nil, created_at: "2016-04-16 02:41:06", updated_at: "2016-04-16 02:41:06">]> 

 		canciones[0]  //mediante índice

 			=> #<Track id: 1, title: "Work", album: "The best", artist: "Rihanna", created_at: "2016-04-14 06:24:11", updated_at: "2016-04-14 06:24:11">

 		cancion = canciones.last

 		 => #<Track id: 4, title: "Five", album: nil, artist: nil, created_at: "2016-04-16 02:41:06", updated_at: "2016-04-16 02:41:06"> 

 		cancion.album = 'Five'
 		cancion.artist = 'Maroon Five'
 		cancion.save

 		PARA AGREGAR CONTENIDO DE PRUEBA:

 		2.3.0 :033 >   10.times do
		2.3.0 :034 >     Track.create title: 'Destiny', album: 'Destiny', artist: 'Stratovarius'
		2.3.0 :035?>   end

		PARA BUSCAR UN ELEMENTO:

		Track.find 12

			PARA BUSCAR TODOS LOS ELEMENTOS QUE CUMPLAN CONDICIÓN

				cancion = Track.where title: 'Destiny' // apareceran los 10 registros creados anteriormente.

				cancion.count => 10

			PARA OBTENER LA PRIMERA COINCIDENCIA DE UNA BUSQUEDA CON CONDICION
				Track.find_by title: 'Destiny' // Traerá solo la primera coincidencia


			BORRAR ELEMENTOS:

				canciones = Track.where artist: 'Stratovarius'
				canciones.count
				// canciones.destroy
				canciones.destroy_all

	"LOS ANTERIORES MÉTODOS SON LOS QUE ENCONTRAMOS EN EL CONTROLADOR PARA MANIPULAR LOS OBJETOS DE RUBY"

RUTAS

	rake routes // Toda la funcionalidad que queremos exponer nuestra aplicacion

	Para manipular las rutas, acceder a:
		-> /config/routes.rb

		resources	:tracks  
		// Este método crea automáticamente todos los recursos que podemos utilizar (edit, show, destroy).

LAYOUTS

	Son las vistas que se predefinen con el header y el footer y que tendrán que aparecer en cada una de las vistas de accion del controlador.

HELPERS

///////////////////////// CREACION DE ESTRUCTURAS SIN SCAFFOLD /////////////////////

rails new app -d mysql

mysql -u root -p
CREATE DATABASE landingmailer;

IR A config/database.yml

		default: &default
		  adapter: mysql2
		  encoding: utf8
		  pool: 5
		  username: root
		  password: Cocacola15
		  socket: /tmp/mysql.sock

		development:
		  <<: *default
		  database: landingmailer

		# Warning: The database defined as "test" will be erased and
		# re-generated from your development database when you run "rake".
		# Do not set this db to the same as development or production.
		test:
		  <<: *default
		  database: landingmailer

CREAR EL MODELO
	Ir a app/models/
		nano playlist.rb

	Crear la clase del modelo:

		#Hereda de la clase AR, para poder manipular las DB con los comandos internos
		#Nombre de clase con Mayúscula
		class Playlist < ActiveRecord::Base 

		end


	Ir a consola:

		rails console
			#Para probar si la tabla existe, claramente no porque no hemos migrado las tablas
			Playlist.connection

			Playlist
			# => Playlist(Table doesn't exist)
	
	LAS MIGRACIONES SON VERSIONES DE NUESTRAS DB QUE PODEMOS IR CRECIENDO Y HACIENDO CAMBIOS
	EN EL PRINCIPIO DE LA CREACIÓN.
	
	Salir de consola: Ctrl + D // Ctrl + R

	Realizar migración para generar la tabla con ActiveRecord, Playlist

		rails generate migration CreatePlaylists
								 #CreatePluralModelName

		# Esto genera el modelo ActiveRecord y una migración con su nombre.
	    # invoke  active_record
		# create    db/migrate/20160416073817_create_playlists.rb


	rake db:migrate:status

		database: /Users/williamromero/Desktop/Platzi/Rails/musicapp/db/development.sqlite3

		 Status   Migration ID    Migration Name
		--------------------------------------------------
		   up     20160414060548  Create tracks
		  down    20160416073817  Create playlists


	IR AL MODELO PLAYLIST Y AGREGAR LOS CAMBIOS A UTILIZAR

		1. Ir a config/db/migrate y buscar la migración de playlist
		2. Luego definir los campos a utilizar:

		class CreatePlaylists < ActiveRecord::Migration
			  def change
			    create_table :playlists do |t|
			    	t.string :name
			    	t.integer :number_of_votes
			    end
			  end
		end

	Recibimos el error de rake db:migrate y por ello vamos a rake db:migrate:status para ver que base de datos no ha sido migrada (actualizada y asignada a la db) para usarse como modelo.

	rake db:migrate:status

		database: /Users/williamromero/Desktop/Platzi/Rails/musicapp/db/development.sqlite3

		 Status   Migration ID    Migration Name
		--------------------------------------------------
		   up     20160414060548  Create tracks
		  down    20160416073817  Create playlists   # No está creada en la DB

	Por ello vamos a consola y presionamos:

		rake db:migrate

			== 20160416073817 CreatePlaylists: migrating ==================================
			-- create_table(:playlists)
			   -> 0.0014s
			== 20160416073817 CreatePlaylists: migrated (0.0015s) =========================


	Luego vamos a rails console

		Playlist.connection

		2.3.0 :002 > Playlist
		 => Playlist(id: integer, name: string, number_of_votes: integer) 

		ActiveRecord crea en todos los modelos, el atributo id para poder realizar búsquedas sobre la tabla.


CREAR EL CONTROLADOR

	Vamos a config/routes.rb

		Rails.application.routes.draw do
  		resources :tracks  

  		resources :name_of_table #Crea todos los recursos que queremos exponer vía URL
  		#Esto creará una estructura REST para que nuestra aplicación pueda ser leída.

					Rails.application.routes.draw do
  					resources :playlists  		

  		Para ver las rutas que crea resources de nuestro modelo, presionar: 

  			rake routes

		       Prefix Verb   URI Pattern                   Controller#Action
		       tracks GET    /tracks(.:format)             tracks#index
		              POST   /tracks(.:format)             tracks#create
		    new_track GET    /tracks/new(.:format)         tracks#new
		   edit_track GET    /tracks/:id/edit(.:format)    tracks#edit
		        track GET    /tracks/:id(.:format)         tracks#show
		              PATCH  /tracks/:id(.:format)         tracks#update
		              PUT    /tracks/:id(.:format)         tracks#update
		              DELETE /tracks/:id(.:format)         tracks#destroy
		    playlists GET    /playlists(.:format)          playlists#index
		              POST   /playlists(.:format)          playlists#create
		 new_playlist GET    /playlists/new(.:format)      playlists#new
		edit_playlist GET    /playlists/:id/edit(.:format) playlists#edit
		     playlist GET    /playlists/:id(.:format)      playlists#show
		              PATCH  /playlists/:id(.:format)      playlists#update
		              PUT    /playlists/:id(.:format)      playlists#update
		              DELETE /playlists/:id(.:format)      playlists#destroy  			


		Si ya tenemos la ruta y el modelo, nos falta el controlador, por lo tanto hay que generarlo:

		TAL COMO ACTIVE RECORD AYUDA AL SISTEMA A COMPRENDER QUE EL MODELO DEBE DE MANEJAR LA BASE DE DATOS, EL APPLICATION CONTROLLER AYUDA AL SISTEMA A COMPRENDER EL SISTEMA DE RUTAS PARA PODER SERVIR VISTAS DE ACUERDO A LAS NECESIDADES DEL CONTROLADOR

			Vamos a app/controllers/
			Creamos el archivo: playlists.rb

				ó

			rails generate controller playlists

		#rails generate controller playlists >> Creará el controlador y las vistas y assets necesarios que están ya determinados en las rutas.

			  Running via Spring preloader in process 3210
		      create  app/controllers/playlists_controller.rb
		      invoke  erb
		      create    app/views/playlists
		      invoke  test_unit
		      create    test/controllers/playlists_controller_test.rb
		      invoke  helper
		      create    app/helpers/playlists_helper.rb
		      invoke    test_unit
		      invoke  assets
		      invoke    coffee
		      create      app/assets/javascripts/playlists.coffee
		      invoke    scss
		      create      app/assets/stylesheets/playlists.scss 

		Vamos al controlador playlists
			app/controllers/

		Abrir el archivo playlists_controller.rb

		Si vamos a la ruta /playlists/, veremos que tenemos el error de que la action "index" no fue encontrada en PlaylistsController, debido a que el mismo, no hace referencia a ningun método, y por lo tanto a ninguna vista ya que como vimos antes, en un controlador existen los métodos y por cada método existe una vista.

		Entonces creamos el evento index dentro del controlador Playlists

			class PlaylistsController < ApplicationController
				def index
				end
			end

		Lo cual nos brindará otro error debido a que aunque el generador de controllers creó varios archivos y entre ellos, el folder de "Playlists", no creo los archivos esenciales que van dentro.

		Por lo tanto hay que crear un archivo index, que será el que responderá a la accion index del controlador.

			app/views/playlists   => crear archivo index.html.erb

		La cual estará vacía pero ya no dará ningún error, porque la RUTA INDEX fue creada con:	"resources :playlists" en el archivo "config/routes.rb"  

		Pero ahora, como ingresar un formulario:

			<h1>Listado de listas de Reproducción</h1>
			<table>
				<thead>
					<tr>
						<th>Name</th>
						<th>Number of Votes</th>
					</tr>
				</thead>
				<tbody>

				</tbody>
			</table>

		Pero para desplegar los datos, tenemos que solicitarle al controlador que conecte con el modelo para requerir el despliegue de su contenido.
		Por lo tanto, vamos al controlador playlists

			controllers/playlists_controller.rb


			class PlaylistsController < ApplicationController

				def index
					@playlists = Playlist.all 
				end

			end

		EL UTILIZAR LA ARROBA, NOS DICE QUE ES UNA INSTANCIA QUE ESTÁ VIVIENDO SOLO EN ESE MÉTODO.

		Hecho esto, vamos a nuestra vista y para mostrarlo, realizamos una iteración de lo que el método index solicitó y guardó en la variable @playlists.

			...
				<tbody>
					<% @playlists.each do |playlist| %>
						<tr>
							<td><%= playlist.name %></td>
							<td><%= playlist.number_of_votes %></td>
						</tr>
					<% end %>
				</tbody>
			</table>
			...

		Luego, creamos contenido desde rails console:

			Playlist.connection    #conectamos con el modelo
			Playlist.create name:'Alternative', number_of_votes: '20'

			   (0.1ms)  begin transaction
  				SQL (0.6ms)  INSERT INTO "playlists" ("name", "number_of_votes") VALUES (?, ?)  [["name", "Alternative"], ["number_of_votes", 20]]
   			   (1.6ms)  commit transaction
   			    => #<Playlist id: 1, name: "Alternative", number_of_votes: 20> 

   			Playlist.create name: 'Latin', number_of_votes: '50'

				(0.1ms)  begin transaction
				  SQL (2.1ms)  INSERT INTO "playlists" ("name", "number_of_votes") VALUES (?, ?)  [["name", "Latin"], ["number_of_votes", 50]]
				(1.3ms)  commit transaction
				 => #<Playlist id: 2, name: "Latin", number_of_votes: 50> 

		Finalmente, si nos dirigimos a la vista de la acción index, veremos los datos desplegados ya en pantalla, mediante la variable de instancia que creamos en su método.

		Ahora agregaremos la ruta SHOW, en la cual se muestra un único elemento.

		controllers/playlists_controller.rb

			class PlaylistsController < ApplicationController

				def index
					@playlists = Playlist.all 
				end

				def show
					@playlists = Playlist.find(params[:id])
					#@track = Track.find(params[:id])
				end
			end


		El método show, sabrá de donde tomar el parámetro de selección porque como se especificó en las rutas, la accion show responder a un parámetro ID.

       Prefix Verb   URI Pattern                   Controller#Action
     playlist GET    /playlists/:id(.:format)      playlists#show		

     	Pero, como toda accion, necesita una vista, por lo tanto tendrémos que crear la vista SHOW:

     	app/views/playlists/ => crear archivo show.html.erb

    Y colocar un código así:

     	<p>
			<strong>Name: </strong>
			<%= @playlist.name %>
			<%= @track.title %>			
		</p>

		<p>
			<strong> Number of votes: </strong>
			<%= @playlist.number_of_votes %>
		</p>

	Y finalmente, podemos ver el registro no.1 del modelo Playlist, en nuestra vista SHOW, de la acción show del controlador Playlists.


	SI QUEREMOS DESPLEGAR UN OBJETO DE OTRO CONTROLADOR, LO INSTANCIAMOS EN LA ACCION Y LO REPLICAMOS EN LA VISTA.


PARA CREAR UN CONTROLADOR NEW

	app/controllers/

		def new
			@playlist = Playlist.new
		end


	El método NEW nos dará los elemento para un formulario de validación de modelos.
	Para ello, creamos el archivo:

		app/views/playlists/ => crear archivo new.html.erb

	Luego, haciendo uso de los "helpers", Rails nos permite crear formularios con la siguiente sintáxis.

		<%= form_for(@playlist) do |f| %>

	Esta etiqueta, haciendo uso del helper form_for, nos permite requerir los valores de la variable de instancia @playlist, la cual fue creada en la acción "def new" del controlador. Luego, cada valor de contenido requerido, se deposita en la variable "f", la cual usaremos como comodín para imprimir las casillas en pantalla.

	<%= form_for(@playlist) do |f| %>
		<%= f.label :name %>
		<%= f.text_field :name %>
		<br/><br/>
		<%= f.label :number_of_votes %>
		<%= f.text_field :number_of_votes %>

		<%=	f.submit %>
	<% end %>

	El helper label, siempre sabrá que desplegará la llave del atributo, el input el casillero para la información y el submit, desplegará el nombre de lo que guardamos.

	Ahora, al presionar el botón, obtendremos un error que hace referencia a que la acción "create" no existe.

		## The action 'create' could not be found for PlaylistsController

	Esto sucede, porque no hemos creado en el controlador esta acción que actua en relación a la acción new. Así mismo, si pensamos, no encontrará la vista por lo que habrá que crearla.

	# PERO SI VEMOS EL RAKE ROUTES DE ESTE MODELO, ESTE EVENTO NO TIENE VISTA POR LO QUE SEGUIREMOS RECIBIENDO EL ERROR. POR LO CUAL LA SIGUIENTE LINEA, NO ES VÁLIDA

		######app/views/playlists/ => crear archivo create.html.erb######

	Y vamos al archivo "playlists_controller.rb" y colocamos lo siguiente:

		def create

		end

	Ahora ya encuentra su método y su vista no existirá, debido a que en rake routes, esa vista corresponde a la principal. Es decir, será como la vista siguiente al formulario en la cual se podrá crear el registro, al igual que update y destroy.

	Ahora, cuando hemos recibido la lista de parámetros tenemos que ahondar en como recibirlos en el método create.

	Ahora, en "create" veremos que necesitamos requerir los parámetros, y para ello, desde Rails 4, existe algo llamado "strongparams", que sirve para extraer los parametros que se necesitan guardar. Esto sabe cuales debe de descargar.

	Para ello, crearemos un método privado, con el motivo de recibir allí los parámetros necesarios para que el evento CREATE guarde los datos.

	def create
		
	end

	private
	
		def playlist_params
			params.require(:playlist).permit(:name, :number_of_votes)
		end

	Esto, limita que se reciban parámetros que no necesitamos o que quizás podamos tener escondidos.

	Luego, vamos al método create para validar los parámetros. En este, guardaremos los datos con la función NEW, la cual hemos utilizado antes para desplegar. Es una función I/O por lo que podemos utilizarla para crear un nuevo registro y guardar también. Es decir, cuando lo utilizamos en "def new", es como si creara un nuevo registro pero como no se salva, se pierde. Ya cuando vamos al evento "create" y lo mandamos a llamar de nuevo, allí si se depositan los datos en el registro que creó "new" con los parámetros requeridos del método privado.

	def create
		@playlist = Playlist.new(playlist_params)
		if @playlist.save
			redirect_to @playlist
		else
			render :new
		end
	end 

	Luego, cuando se han guardado ya, evaluamos el valor de la variable, si se pudo o no guardar y si, sí, redireccionaremos al evento SHOW. 

	Sino, redireccionará la vista NEW de nuevo. 

	VALIDACIÓN DE DATOS:

	En este momento, si guardamos datos vacíos, los mismos se guardaran. Por ello, tenemos que ir al modelo para que esto no suceda y cuando se evalue si se guardan los datos, tire una bandera para que no lo haga.

	class Playlist < ActiveRecord::Base
		validates_presence_of :name, :number_of_votes
	end

	Esto evaluará si los campos están y sí no, no guardará y levantará la bandera para que el evento create haga de nuevo el render del formulario.





Referencias: 
1. http://www.sectornueve.com/programacion/ingenieria-de-software/mvc-para-dummies/
2. http://www.webwindowslinux.com/2013/05/explicando-la-arquitectura-aspnet-mvc.html
3. https://gist.github.com/Elicia/8473536 //-> MVC Expliación
