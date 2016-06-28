## RAILS DEPLOY 

1. Crear aplicación de Rails [Terminal]:

    <pre>	
      rails new railsapp
    </pre>
	
** Instalar las siguientes gemas [Terminal]:**
	<pre>
		gem list  
		gem <b>install figaro</b>
		gem <b>install puma</b>
	</pre>
# Abrir archivo railsapp/Gemfile de aplicación:

	gem 'figaro'
	gem 'puma'

	group :development do
	  # Access an IRB console on exception pages or by using <%= console %> in views
	  gem 'web-console', '~> 2.0'
		gem 'capistrano'
		gem 'capistrano3-puma'
		gem 'capistrano-rails', require: false
		gem 'capistrano-bundler', require: false
		gem 'capistrano-rvm'

	  # Spring speeds up development by keeping your application running in the background. Read more: https://github.com/rails/spring
	  gem 'spring'
	end

# Actualizar gemas en [Terminal]:

	bundle install

# Crear un STAGE de Producción [Terminal]:

	cap install STAGES=production
	
	```
		mkdir -p config/deploy
		create config/deploy.rb
		create config/deploy/production.rb
		mkdir -p lib/capistrano/tasks
		create Capfile
		Capified
	
		 _    _                              _
		| |  | |                            (_)
		| |__| | __ _ _ __ _ __ _____      ___  ___
		|  __  |/ _  | '__| '__/ _ \ \ /\ / / |/ _ \
		| |  | | (_| | |  | | | (_) \ V  V /| | (_) |
		|_|  |_|\__,_|_|  |_|  \___/ \_/\_(_)_|\___/

		- Continuous Integration & Deployment
		      Built by the team behind Capistrano

		- Learn more at http://hrw.io/cap-for-teams


		- Free for small projects!

		- Test, deploy and collaborate online easily
		   using tools you already know and love!

		- Trigger tasks automatically based on Git changes
		  and webhooks. Get notified by email, slack, etc.

		- Works seamlessly for PHP, Node.js, Ansible, Python, Go,
		  Capistrano and more!

		Try it now?  (Yes/no): Aborting: timeout...
	
	```

# Abrir archivo railsapp/Capfile de aplicación, remover el signo '#':

	```
		# Load DSL and set up stages
		require "capistrano/setup"

		# Include default deployment tasks
		require "capistrano/deploy"

		# Include tasks from other gems included in your Gemfile
		#
		# For documentation on these, see for example:
		#
		#   https://github.com/capistrano/rvm
		#   https://github.com/capistrano/rbenv
		#   https://github.com/capistrano/chruby
		#   https://github.com/capistrano/bundler
		#   https://github.com/capistrano/rails
		#   https://github.com/capistrano/passenger
		#
		 require 'capistrano/rvm'
		# require 'capistrano/rbenv'
		# require 'capistrano/chruby'
		 require 'capistrano/bundler'
		 require 'capistrano/rails/assets'
		 require 'capistrano/rails/migrations'
		 require 'capistrano/puma'
		# require 'capistrano/passenger'

		# Load custom tasks from 'lib/capistrano/tasks' if you have any defined
		Dir.glob("lib/capistrano/tasks/*.rake").each { |r| import r }
	```

# Abrir archivo "config/deploy.rb" de aplicación:

	set :application, 'railsapp'
	set :repo_url, 'git@github.com:williamromero/railsapp.git' # Edit this to match your repository
	set :branch, :master
	set :deploy_to, '/home/deploy/railsapp'
	set :pty, true
	set :linked_files, %w{config/database.yml config/application.yml}
	set :linked_dirs, %w{bin log tmp/pids tmp/cache tmp/sockets vendor/bundle public/system public/uploads}
	set :keep_releases, 5
	set :rvm_type, :user
	set :rvm_ruby_version, 'ruby-2.3.0' # Edit this if you are using MRI Ruby

	set :puma_rackup, -> { File.join(current_path, 'config.ru') }
	set :puma_state, "#{shared_path}/tmp/pids/puma.state"
	set :puma_pid, "#{shared_path}/tmp/pids/puma.pid"
	set :puma_bind, "unix://#{shared_path}/tmp/sockets/puma.sock"    #accept array for multi-bind
	set :puma_conf, "#{shared_path}/puma.rb"
	set :puma_access_log, "#{shared_path}/log/puma_error.log"
	set :puma_error_log, "#{shared_path}/log/puma_access.log"
	set :puma_role, :app
	set :puma_env, fetch(:rack_env, fetch(:rails_env, 'production'))
	set :puma_threads, [0, 8]
	set :puma_workers, 0
	set :puma_worker_timeout, nil
	set :puma_init_active_record, true
	set :puma_preload_app, false

# Abrir archivo "railsapp/.gitignore" de aplicación:

	config/database.yml

# Crear el repositorio en Github y luego crear el repositorio en la linea de comandos [Terminal]:

	# Initialized empty Git repository in /Users/william/Desktop/AWS/railsapp/.git/
	git init
	# Agregar todos los archivos al repositorio remoto.
	git add .
	# Hacer el commit inicial.
	git commit -m 'Initialized commit'
	# Agregar el registro del git y agregarlo al de Github.
	git remote add origin git@github.com:williamromero/railsapp.git
	# Cargar los registros de Github.
	git push -u origin master

# Crear AWS EC2 Instance:

	Elegir instancia: 
		Ubuntu Server 14.04 LTS (HVM), SSD Volume Type - ami-9abea4fb
	
	Seleccionar el tipo de instancia:	
		t2.micro

	Configurar detalles de instancia:
		Next: Add Storage

	Elegir almacenamiento:
		Next: Tag Instance

# Elegir el valor de llave y de tag:

	Key: 	Name
	Value: 	railsapp

# Configurar un Security Group:

	Elegir [x] un grupo de seguridad.

		[x] launch-wizard-3

# Presionar el botón Launch:

	Elegir una "key-pair" nueva o existente:

		# Choose an existing key pair
		Select a key pair:
			> webapp 

## SERVER INSTALATION

# Abrir la nueva instancia en Terminal:

	ssh -i "webapp.pem" ubuntu@ec2-52-36-26-81.us-west-2.compute.amazonaws.com

# Actualizar servidor:

	sudo apt-get update && sudo apt-get -y upgrade

# Crear el usuario DEPLOY:

	sudo useradd -d /home/deploy -m deploy

# Cambiar el password del usuario DEPLOY:

	sudo passwd deploy
	#PASS: production

# Otorgar todos los accesos del usuario DEPLOY:

	sudo visudo
	```
		root    ALL=(ALL:ALL) ALL
		deploy  ALL=(ALL:ALL) ALL
	```

# Ingresar con el usuario DEPLOY:

	su - deploy

# Crear la clave de seguridad del servidor:

	ssh-keygen

# Generar la clave de seguridad:

	cat ~/.ssh/id_rsa.pub

# Guardar SSH-KEY - Github/SSH and GPG keys/ AWS Ubuntu Server - Deploy User:

	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDoT2TD5R09yBwP1jGWeXoTk4zCgVB5jPt/c5kWXvKAQwGZk+fgWaxaX8dFDmOvF/pEQGHsJuLw6R0PQWNGccf4DWa99gG7MpG5Hxn9C88ODekCLQHszMIwKrfeEuekvsikuY/XxqizAB72y0/TXY3Ju3+5s8wOsMURKskQw0aeBn1T7duxioADCrLRt1MLRJi2OKWG4uxapoGPlWjK42aO1g5VNzzcrjCUKEYVZdDCLXKvKlqGV/cErlwohYLvouSWP8PY7ypdC9bdqJV/yBr1AbuW8TnicbEjD3u7/zBUJF8HZLPJzfG8YLs+mJdZXyIqarvXBWfuQY0FfQ0t0swt deploy@ip-172-31-23-70

# Hacer la prueba de autenticidad de conexión del host con Github:

	ssh -T git@github.com



## IR AL ORDENADOR DE NUEVO PARA TOMAR LA LLAVE SSH:

	cat ~/.ssh/id_rsa.pub

	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoBePMIcyWVVhbDsfbzI7NC9ZEDZmt/wWGsUiYBUpQnvrtGceUXjwOfr408bV06HcRS/taFRgZ7bZmzPc05k9VG0tVRmnIdKgtX8I6lyQHAvIzrXrxda1Xiga6E9wElMg65XTwfyX+qj+cK3JhwvspCK1WFH1+ZcGBBAV90BljQXbs06010P24FBtJZDjmqewBB0aVuht8U3APAlKr63a+XzyjQHbpT8CusC5q7Noe6PjCWF9zLekeb4bnM5mK36tUv0RDv063sMHb2tbnjp396tlF/1TBFsLINPwOZX4c/yh9qbHP6i8SazhA0YGMN9XL59nWk9F3NUA6zHTcdtp3 william@162.0.10.in-addr.arpa

## IR AL SERVIDOR Y CREAR UN ARCHIVO:

	# Guardar llave SSH del ordenador en servidor:

		nano ~/.ssh/authorized_keys

	# Copiar el registro SSH en este archivo, para automatizar la conexión.

		ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoBePMIcyWVVhbDsfbzI7NC9ZEDZmt/wWGsUiYBUpQnvrtGceUXjwOfr408bV06HcRS/taFRgZ7bZmzPc05k9VG0tVRmnIdKgtX8I6lyQHAvIzrXrxda1Xiga6E9wElMg65XTwfyX+qj+cK3JhwvspCK1WFH1+ZcGBBAV90BljQXbs06010P24FBtJZDjmqewBB0aVuht8U3APAlKr63a+XzyjQHbpT8CusC5q7Noe6PjCWF9zLekeb4bnM5mK36tUv0RDv063sMHb2tbnjp396tlF/1TBFsLINPwOZX4c/yh9qbHP6i8SazhA0YGMN9XL59nWk9F3NUA6zHTcdtp3 william@162.0.10.in-addr.arpa

## AHORA PODEMOS ENTRAR SIN CLAVE DESDE CUALQUIER ORDENADOR SOLO PONIENDO LO SIGUIENTE:

	ssh deploy@52.36.26.81
	Ahora podemos ingresar al servidor sin clave.
	Ver referencia: https://www.digitalocean.com/community/tutorials/how-to-configure-ssh-key-based-authentication-on-a-linux-server

## INSTALAR GIT EN EL SERVIDOR:

	sudo apt-get install git

# INSTALAR NGINX EN EL SERVIDOR:

	sudo apt-get install nginx

# EDITAR EL REGISTRO DEL VIRTUAL HOST CON NGINX

	sudo nano /etc/nginx/sites-available/default

# AGREGAR SÍMBOLO "#" EN EL ÁREA DE SERVER A TODOS LAS LINEAS QUE ESTÁN SIN NADA
# PARA ANULAR TODOS LOS CAMPOS Y COPIAR EL SIGUIENTE CÓDIGO

	upstream app {
	  # Path to Puma SOCK file, as defined previously
	  server unix:/home/deploy/urlshortner/shared/tmp/sockets/puma.sock fail_timeout=0;
	}
	server {
	  listen 80;
	  server_name localhost;
	  root /home/deploy/urlshortner/current/public;
	  try_files $uri/index.html $uri @app;
	  location / {
	    proxy_set_header X-Forwarded-Proto $scheme;
	    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	    proxy_set_header X-Real-IP $remote_addr;
	    proxy_set_header Host $host;
	    proxy_redirect off;
	    proxy_http_version 1.1;
	    proxy_set_header Connection '';
	    proxy_pass http://app;
	  }
	  location ~ ^/(assets|fonts|system)/|favicon.ico|robots.txt {
	    gzip_static on;
	    expires max;
	    add_header Cache-Control public;
	  }
	  error_page 500 502 503 504 /500.html;
	  client_max_body_size 4G;
	  keepalive_timeout 10;
	}  
	Status API Training Shop Blog About

## INSTALAR EN EL SERVIDOR POSTGRESQL

	sudo apt-get install postgresql postgresql-contrib libpq-dev

## CREAR USUARIO DE POSTGRESQL COMO SUPERUSUARIO

	sudo -u postgres createuser -s deploy

## INSTALAR PSQL PARA PODER ACCEDER A LA BASE DE DATOS MEDIANTE POSTGRESQL CONSOLE

	sudo -u postgres psql

## REGISTRAR EL PASSWORD

	postgres=# \password deploy
		Enter new password: production
		Enter it again: production
	postgres=#

## SALIR DE POSTGRESQL

	\q

## CREAR BASE DE DATOS EN POSTGRESQL

	sudo -u postgres createdb -O deploy deploy_production
								 *		*	
								 USER   DB NAME
								 
## CREAR LA BASE DE DATOS EN MYSQL
	<pre>
		sudo apt-get install mysql-server mysql-client libmysqlclient-dev
	</pre>
** New password for the MySQL "root" user: 
	<pre>
		Insert your user password
	</pre>
	
## INSTALAR LA VERSION FIRMADA DE RVM - El RVM a partir de la versión 1.26 introduce versiones firmadas y comprobación automática de las mismas. Para ello necesitamos instalar dicha firma. En caso contrario podría darnos problemas al instalar RVM. http://www.rubyonrails.org.es/instala.html

	gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 

## LUEGO INSTALAR LA ÚLTIMA VERSIÓN ESTABLE DE RVM

	\curl -sSL https://get.rvm.io | bash -s stable --rails --ruby

	* To start using RVM you need to run `source /home/deploy/.rvm/scripts/rvm`
    in all your open shell windows, in rare cases you need to reopen all shell windows.

## SALIR DEL USUARIO DEPLOY PARA PROBAR QUE SE HAYA INSTALADO Y VOLVER A ENTRAR

	deploy@ip-172-31-23-70:~$ exit
	logout
	ubuntu@ip-172-31-23-70:~$ su - deploy
	Password: 
	deploy@ip-172-31-23-70:~$ rvm -v
	rvm 1.27.0 (latest) by Wayne E. Seguin <wayneeseguin@gmail.com>, Michal Papis <mpapis@gmail.com> [https://rvm.io/]

## AHORA INSTALAR BUNDLER PARA MANEJAR LAS GEMAS

	gem install bundler --no-ri --no-rdoc

## IR AL GIT REMOTO DEL PROYECTO Y VISUALIZAR LOS SIGUIENTES DATOS EN EL ARCHIVO
## CONFIG/DEPLOY.RB

	set :application, 'railsapp'
	set :repo_url, 'git@github.com:williamromero/railsapp.git' # Edit this to match your repository
	set :branch, :master
	set :deploy_to, '/home/deploy/railsapp'

## LUEGO VOLVER AL SERVIDOR Y VER EN QUE CARPETA NOS ENCONTRAMOS, DEBEMOS DE 
## IR A LA CARPETA /home/deploy/

	pwd
	/home/deploy/

## AHORA VAMOS A CREAR EL FOLDER EN EL QUE SE DEPOSITARÁ EL PROYECTO DE GITHUB

	mkdir railsapp

## AHORA DENTRO DE ESA CARPETA, CREAR OTROS DOS FOLDERS AL QUE HARÁ REFERENCIA LA CONFIGURACIÓN DEL
## SERVIDOR DE APLICACIONES PUMA MEDIANTE EL ARCHIVO "config/deploy.rb"

	mkdir -p railsapp/shared/config

## AHORA DENTRO DE LA CARPETA railsapp/shared/config CREAREMOS LOS ARCHIVOS QUE REEMPLAZARÁ DEL PROYECTO REMOTO
## PARA QUE PUEDA CONECTARSE CON EL SERVIDOR DE APLICACIONES PUMA

	nano urlshortner/shared/config/database.yml

## EN ESTE ARCHIVO INTRODUCIR LA CONFIGURACIÓN DE LA BASE DE DATOS:

	production:
	  adapter: postgresql
	  encoding: unicode
	  database: deploy_production
	  username: deploy
	  password: production
	  host: localhost
	  port: 5432

## AHORA DENTRO DE LA MISMA CARPETA, CREAR EL ARCHIVO application.yml:

	nano application.yml

## DESDE LA APLICACIÓN COPIADA EN REMOTO, CREAR UNA SECRET KEY PARA PRODUCTION
## http://www.jamesbadger.ca/2012/12/18/generate-new-secret-token/

	rake secret

### AHORA COLOCARLA EN EL SERVIDOR, EN EL ARCHIVO application.yml Y COPIAR LA SECRET_KEY_BASE

	SECRET_KEY_BASE: "cafe70dc504ce903398ed6ba175620b405a1f1f5773f4d72dcce99045ba15e0986918069d2599535ab54544bf33ddc2d2a9dcbd54112e86d4362e926a2133920"	

## AHORA EN EL SERVIDOR, IR AL ARCHIVO config/deploy/production.rb Y COPIAR LA SIGUIENTE LINEA

	server '52.36.26.81', user: 'deploy', roles: %w{web app db}
			IP DNS        USER           

## AHORA IR A LA APLICACIÓN REMOTA E INGRESAR EL SIGUIENTE COMANDO Y ESPERAR LA CLONACIÓN DE TODA LA APLICACIÓN AL SERVIDOR:

	cap production deploy

## ESTO HARÁ SALTAR UN ERROR PORQUE EL ARCHIVO UGLIFIER NO FUNCIONA COMO JS, ENTONCES HAY QUE INSTALAR NODEJS EN EL SERVIDOR

	sudo apt-get install nodejs

## LUEGO DE INSTALAR NODEJS EN SERVIDOR, VAMOS A LA COMPUTADORA REMOTA Y VOLVEMOS A EJECUTAR EL COMANDO DE CAPISTRANO PARA HACER EL DEPLOY

	cap production deploy

## SI APARECE UN PROBLEMA POR LA GEMA PG EN LA COMPUTADORA, PUEDE INSTALARSE DE DOS MANERAS:

	gem install pg / gem install pg -- --with-pg-config=/usr/local/bin/pg_config

	ó
 
	brew install libpqxx

## LUEGO DE ESO, AGREGARLA AL GEMFILE DEL PROYECTO EN COMPUTADORA

	# gem 'capistrano-rails', group: :development
	gem 'pg', '~> 0.18.4'


## LUEGO DE ELLO, REALIZAR UN REGISTRO DE GEMAS
	
	bundle install

## LUEGO HACER UN COMMIT PARA GUARDARLO EN EL REPOSITORIO

## LUEGO, VOLVER A EJECUTAR EL COMANDO DE CAPISTRANO PARA HACER EL DEPLOY

	cap production deploy

## LUEGO IR AL SERVIDOR Y REINICIAR NGINX PARA CORRER POR PRIMERA VEZ LA APLICACIÓN

	sudo service nginx restart








#How to deploy RubyonRails project to AWS EC2 using capistrano
=> Video de Youtube: https://youtu.be/imdrYD4ooIk?t=25m34s
=> Ver este link: https://www.amberbit.com/blog/2015/11/28/how-to-deploy-rails-on-a-vps/

##PARA REVISAR PROBLEMAS DE NGINX
	https://www.digitalocean.com/community/questions/service-nginx-restart-fail

##PARA INSTALAR POSTGRES EN EL ORDENADOR CON OS EL CAPITAN
	https://teamtreehouse.com/community/ruby-on-rails-bundle-fails-cannot-gem-install-pg-not-yet-resolved-need-help
	http://www.postgresql.org/files/documentation/pdf/9.3/postgresql-9.3-A4.pdf
	https://gorails.com/setup/ubuntu/14.10

	https://gist.github.com/JamesDullaghan/5941259

##RAKE SECRET

	http://www.jamesbadger.ca/2012/12/18/generate-new-secret-token/

##VIDEOTUTORIAL DEPLOY:

	https://www.youtube.com/watch?v=imdrYD4ooIk
	https://github.com/rkmmanivannan/rails-ec2-configuration/blob/master/method1/serverconfig

## POSTGRES:

	MacOS => http://postgresapp.com/documentation/
	      => http://postgresapp.com/documentation/cli-tools.html
	      => https://launchschool.com/blog/how-to-install-postgresql-on-a-mac
	      => http://www.postgresql.org/files/documentation/pdf/9.5/postgresql-9.5-A4.pdf
	      
	      => http://www.linuxscrew.com/2009/07/03/postgresql-show-tables-show-databases-show-columns/
	      
	      => https://chartio.com/resources/tutorials/how-to-start-postgresql-server-on-mac-os-x

## MANEJADOR DE POSTGRES:

	https://eggerapps.at/postico/
	http://postgresapp.com/
	https://www.mysql.com/products/workbench/


