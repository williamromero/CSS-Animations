## RAILS DEPLOY 

**1.** Crear aplicación de Rails [Terminal]:
<pre>
  rails new railsapp
</pre>
	
**2.** Instalar las siguientes gemas [Terminal]:
   	<pre>
		gem list  
		gem <b>install figaro</b>
		gem <b>install puma</b>
   	</pre>

**3.** Abrir archivo railsapp/Gemfile de aplicación:
    <pre>
    	gem <b>'figaro'</b>
	    gem <b>'puma'</b>
    	group :development do
    	  # Access an IRB console on exception pages or by using &lt;%= console %&gt; in views
    	  gem 'web-console', '~&gt; 2.0'
    		gem 'capistrano'
    		gem 'capistrano3-puma'
    		gem 'capistrano-rails', require: false
    		gem 'capistrano-bundler', require: false
    		gem 'capistrano-rvm'
    	  # Spring speeds up development by keeping your application running in the background. Read more: https://github.com/rails/spring
    	  gem 'spring'
    	end
	</pre>

**4.** Actualizar gemas en [Terminal]:
    <pre>
	    bundle install
    </pre>

**5.** Crear un STAGE de Producción [Terminal]:
	<pre>
	    cap install <b>STAGES=production</b>
	</pre>
	&nbsp;
	<pre>
		mkdir -p config/deploy
		create config/deploy.rb
		create config/deploy/production.rb
		mkdir -p lib/capistrano/tasks
		create Capfile
		Capified
	&nbsp;
		 _    _                              _
		| |  | |                            (_)
		| |__| | __ _ _ __ _ __ _____      ___  ___
		|  __  |/ _  | '__| '__/ _ \ \ /\ / / |/ _ \
		| |  | | (_| | |  | | | (_) \ V  V /| | (_) |
		|_|  |_|\__,_|_|  |_|  \___/ \_/\_(_)_|\___/
	&nbsp;
		- Continuous Integration & Deployment
		      Built by the team behind Capistrano
	&nbsp;
		- Learn more at http://hrw.io/cap-for-teams
	&nbsp;
		- Free for small projects!
	&nbsp;
		- Test, deploy and collaborate online easily
		   using tools you already know and love!
	&nbsp;
		- Trigger tasks automatically based on Git changes
		  and webhooks. Get notified by email, slack, etc.
	&nbsp;
		- Works seamlessly for PHP, Node.js, Ansible, Python, Go,
		  Capistrano and more!
	&nbsp;
		Try it now?  (Yes/no): Aborting: timeout...
	&nbsp;
	</pre>

**6.** Abrir archivo **"railsapp/Capfile"** de aplicación, remover el **signo '#'**:
	<pre>
		# Load DSL and set up stages
		require "capistrano/setup"
		&nbsp;
		# Include default deployment tasks
		require "capistrano/deploy"
		&nbsp;
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
		&nbsp;
		# Load custom tasks from 'lib/capistrano/tasks' if you have any defined
		Dir.glob("lib/capistrano/tasks/*.rake").each { |r| import r }
	</pre>

**7.** Abrir archivo **"config/deploy.rb"** de aplicación:
    <pre>
    	set :application, <b>'railsapp'</b>
    	set :repo_url, <b>'git@github.com:williamromero/railsapp.git'</b> # Edit this to match your repository
    	set :branch, <b>:master</b>
    	set :deploy_to, <b>'/home/deploy/railsapp'</b>
    	set :pty, true
    	set :linked_files, <b>%w{config/database.yml config/application.yml}</b>
    	set :linked_dirs, %w{bin log tmp/pids tmp/cache tmp/sockets vendor/bundle public/system public/uploads}
    	set :keep_releases, 5
    	set :rvm_type, :user
    	set :rvm_ruby_version, 'ruby-2.3.0' # Edit this if you are using MRI Ruby
    &nbsp;
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
    </pre>
    
**8.** Abrir archivo **"railsapp/.gitignore"** de aplicación:
    <pre>
	    config/database.yml
    </pre>
    
**9.** Crear el repositorio en Github y luego crear el repositorio en la linea de comandos [Terminal]:
    <pre>
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
    </pre>

**10.** Crear AWS EC2 Instance:

* Elegir instancia: 
	+ Ubuntu Server 14.04 LTS (HVM), SSD Volume Type - ami-9abea4fb
	
* Seleccionar el tipo de instancia:	
	+ t2.micro

* Configurar detalles de instancia:
	+ Next: Add Storage

* Elegir almacenamiento:
	+ Next: Tag Instance

**11.** Elegir el valor de llave y de tag:
    <pre>
    	**Key:** 	Name
    	**Value:** 	railsapp
    </pre>
**12.** Configurar un Security Group:
	<pre>
	    Elegir [x] un grupo de seguridad.
		[x] <b>launch-wizard-3</b>
    </pre>
**13.** Presionar el botón Launch:    
    <pre>
	Elegir una "key-pair" nueva o existente:
	&nbsp;
		# Choose an existing key pair
		Select a key pair:
			<b>&gt; webapp</b> 
    </pre>
    
## SERVER INSTALATION

**1.** Abrir la nueva instancia en Terminal:
	<pre>
		ssh -i "webapp.pem" <b>ubuntu@ec2-52-36-26-81.us-west-2.compute.amazonaws.com</b>
	</pre>

**2.** Actualizar servidor:
	<pre>
		sudo apt-get update && sudo apt-get -y upgrade
	</pre>

**3. Crear el usuario DEPLOY:**
	<pre>
		sudo useradd -d /home/deploy -m deploy
	</pre>

**4.** Cambiar el password del usuario DEPLOY:
	<pre>
		sudo passwd deploy
		#PASS: production
	</pre>

**5.** Otorgar todos los accesos del usuario DEPLOY:
	<pre>
	sudo visudo
	```
		root    ALL=(ALL:ALL) ALL
		deploy  ALL=(ALL:ALL) ALL
	```
	</pre>
**6.** Ingresar con el usuario DEPLOY:
	<pre>
		su - deploy
	</pre>

**7.** Crear la clave de seguridad del servidor:
	<pre>
		ssh-keygen
	</pre>

**8.** Generar la clave de seguridad:
	<pre>
		cat ~/.ssh/id_rsa.pub
	</pre>

**9.** Guardar SSH-KEY - Github/SSH and GPG keys/ AWS Ubuntu Server - Deploy User:
	<pre>
		ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDoT2TD5R09yBwP1jGWeXoTk4zCgVB5jPt/c5kWXvKAQwGZk+fgWaxaX8dFDmOvF/pEQGHsJuLw6R0PQWNGccf4DWa99gG7MpG5Hxn9C88ODekCLQHszMIwKrfeEuekvsikuY/XxqizAB72y0/TXY3Ju3+5s8wOsMURKskQw0aeBn1T7duxioADCrLRt1MLRJi2OKWG4uxapoGPlWjK42aO1g5VNzzcrjCUKEYVZdDCLXKvKlqGV/cErlwohYLvouSWP8PY7ypdC9bdqJV/yBr1AbuW8TnicbEjD3u7/zBUJF8HZLPJzfG8YLs+mJdZXyIqarvXBWfuQY0FfQ0t0swt deploy@ip-172-31-23-70
	</pre>

**10.** Hacer la prueba de autenticidad de conexión del host con Github:
	<pre>	
		ssh -T git@github.com
	</pre>

**11.** Ir al **ordenador** de nuevo para tomar la llave SSH:
	<pre>
		cat ~/.ssh/id_rsa.pub
	</pre>
	<pre>	
		ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoBePMIcyWVVhbDsfbzI7NC9ZEDZmt/wWGsUiYBUpQnvrtGceUXjwOfr408bV06HcRS/taFRgZ7bZmzPc05k9VG0tVRmnIdKgtX8I6lyQHAvIzrXrxda1Xiga6E9wElMg65XTwfyX+qj+cK3JhwvspCK1WFH1+ZcGBBAV90BljQXbs06010P24FBtJZDjmqewBB0aVuht8U3APAlKr63a+XzyjQHbpT8CusC5q7Noe6PjCWF9zLekeb4bnM5mK36tUv0RDv063sMHb2tbnjp396tlF/1TBFsLINPwOZX4c/yh9qbHP6i8SazhA0YGMN9XL59nWk9F3NUA6zHTcdtp3 william@162.0.10.in-addr.arpa
	</pre>

**12.** Ir al **servidor** y crear un archivo:
&nbsp;
+ Guardar llave SSH del ordenador en servidor:
	<pre>
		nano ~/.ssh/authorized_keys
	</pre>
	&nbsp;

**13.** Copiar el registro SSH en este archivo, para automatizar la conexión.
	<pre>
		ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoBePMIcyWVVhbDsfbzI7NC9ZEDZmt/wWGsUiYBUpQnvrtGceUXjwOfr408bV06HcRS/taFRgZ7bZmzPc05k9VG0tVRmnIdKgtX8I6lyQHAvIzrXrxda1Xiga6E9wElMg65XTwfyX+qj+cK3JhwvspCK1WFH1+ZcGBBAV90BljQXbs06010P24FBtJZDjmqewBB0aVuht8U3APAlKr63a+XzyjQHbpT8CusC5q7Noe6PjCWF9zLekeb4bnM5mK36tUv0RDv063sMHb2tbnjp396tlF/1TBFsLINPwOZX4c/yh9qbHP6i8SazhA0YGMN9XL59nWk9F3NUA6zHTcdtp3 william@162.0.10.in-addr.arpa
	</pre>

**14.** Ahora podemos entrar sin clave desde cualquier **ordenador** solo poniendo lo siguiente:
	&nbsp;
	<pre>
		ssh deploy@52.36.26.81
		Ahora podemos ingresar al servidor sin clave.
		Ver referencia: https://www.digitalocean.com/community/tutorials/how-to-configure-ssh-key-based-authentication-on-a-linux-server
	</pre>

## PACKAGES INSTALATION:	

**1.*** Instalar GIT en el **Servidor**:
	<pre>
		sudo apt-get install git
	</pre>

**2.** Instalar **NGINX** en el **servidor**::
	<pre>
		sudo apt-get install nginx
	</pre>
	&nbsp;
**3.** Editar el registro del Virtual Host con NGINX
	<pre>
		sudo nano /etc/nginx/sites-available/default
	</pre>

**4.** Agregar el símbolo **"#"** a todo lo que contiene el archivo para anular todos los campos y copiar el siguiente código:
	&nbsp;
	<pre>
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
	</pre>

**5.** Instalar el servidor de POSTGRESQL
<pre>
	sudo apt-get install postgresql postgresql-contrib libpq-dev
</pre>
	&nbsp;
	+ CREAR USUARIO DE POSTGRESQL COMO SUPERUSUARIO
	<pre>
		sudo -u postgres createuser -s deploy
	</pre>
	&nbsp;
	+ INSTALAR PSQL PARA PODER ACCEDER A LA BASE DE DATOS MEDIANTE POSTGRESQL CONSOLE
	<pre>
		sudo -u postgres psql
	</pre>
	&nbsp;	
	+ REGISTRAR EL PASSWORD
	<pre>
		postgres=# \password deploy
			Enter new password: production
			Enter it again: production
		postgres=#
	</pre>
	&nbsp;		
	+ SALIR DE POSTGRESQL
	<pre>
		\q
	</pre>
	&nbsp;
	+ Crear **base de datos** en POSTGRESQL:
	<pre>
		sudo -u postgres createdb -O deploy deploy_production
									 *		*	
									 USER   DB NAME
	</pre>

**7.** Crear la **base de datos** en MYSQL
	<pre>
		sudo apt-get install mysql-server mysql-client libmysqlclient-dev
	</pre>
** New password for the MySQL "root" user: 
	<pre>
		Insert your user password
	</pre>
	
**8.** INSTALAR LA VERSION FIRMADA DE RVM - El RVM a partir de la versión 1.26 introduce versiones firmadas y comprobación automática de las mismas. Para ello necesitamos instalar dicha firma. En caso contrario podría darnos problemas al instalar RVM. http://www.rubyonrails.org.es/instala.html
	<pre>
		gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 
	</pre>

**9.** LUEGO INSTALAR LA ÚLTIMA VERSIÓN ESTABLE DE RVM
	<pre>
		\curl -sSL https://get.rvm.io | bash -s stable --rails --ruby
	</pre>
	&nbsp;
	* To start using RVM you need to run `source /home/deploy/.rvm/scripts/rvm`
    in all your open shell windows, in rare cases you need to reopen all shell windows.
	&nbsp;
**10.** Salir del **"usuario DEPLOY"** para probar que se haya instalado y volver a entrar:
	<pre>
		deploy@ip-172-31-23-70:~$ exit
		logout
		ubuntu@ip-172-31-23-70:~$ su - deploy
		Password: 
		deploy@ip-172-31-23-70:~$ rvm -v
		rvm 1.27.0 (latest) by Wayne E. Seguin <wayneeseguin@gmail.com>, Michal Papis <mpapis@gmail.com> [https://rvm.io/]
	</pre>
	&nbsp;
**11.** Ahora instalar **BUNDLER** para manejar las gemas:
	<pre>
		gem install bundler --no-ri --no-rdoc
	</pre>
	&nbsp;
**12.** Ir al GIT en el ordenador del proyecto y visualizar los siguientes datos en el archivo:
	<pre>
		config/deploy.rb
	</pre>
	<pre>
		set :application, <b>'railsapp'</b>
		set :repo_url, <b>'git@github.com:williamromero/railsapp.git'</b> # Edit this to match your repository
		set :branch, <b>:master</b>
		set :deploy_to, <b>'/home/deploy/railsapp'</b>
	</pre>
	&nbsp;
**13.** Luego volver al servidor y buscar la carpeta en la que nos situamos y luego debemos ir al folder:
	<pre>
		pwd
		/home/deploy/
	</pre>
	&nbsp;
**14.** A continuación, deberemos de crear un folder en el que se alojará el proyecto de GITHUB:
	<pre>
		mkdir railsapp
	</pre>
	&nbsp;
**15.** Lueg, dentro de esa carpeta, deberemos de crear otros dos folders al que hará referencia la configuraición del servidor de aplicaciones PUMA mediante el archivo **"config/deploy.rb"**:
	<pre>
		mkdir -p railsapp/shared/config
	</pre>
	&nbsp;
**16.** Ahora dentro de la carpeta **railsapp/shared/config**, crearemos los archivos que reemplazará del proyecto en el ordenador para que pueda conectarse con el **SERVIDOR DE APLICACIONES PUMA.**
	<pre>
		nano urlshortner/shared/config/database.yml
	</pre>
**16.** En este archivo, introduciremos la configuración de la Base de Datos:
	<pre>
		<b>production:</b>
		  adapter: postgresql
		  encoding: unicode
		  database: deploy_production
		  username: deploy
		  password: production
		  host: localhost
		  port: 5432
	</pre>

**17.** Ahora dentro de la misma carpeta, crear el archivo **"application.yml"**:
	<pre>
		nano application.yml
	</pre>
	&nbsp;
**18.** Desde la aplicación copiada en el ordenador, crear una [**SECRET KEY**](http://www.jamesbadger.ca/2012/12/18/generate-new-secret-token/) para el ambiente PRODUCTION:
	<pre>
		rake secret
	</pre>

**19.** Ahora deberemos de colocarla en el servidor, en el archivo **application.yml** y copial la SECRET_KEY_BASE:
	<pre>
		SECRET_KEY_BASE: "cafe70dc504ce903398ed6ba175620b405a1f1f5773f4d72dcce99045ba15e0986918069d2599535ab54544bf33ddc2d2a9dcbd54112e86d4362e926a2133920"	
	</pre>
	&nbsp;		
**20.** Ahora en el **servidor**, ir al archivo **"config/deploy/production.rb"** y copiar la siguiente linea:
	<pre>
		server <b>'52.36.26.81'</b>, user: <b>'deploy'</b>, roles: %w{web app db}
				IP DNS        USER           
	</pre>
**21.** Ahora ir a la aplicación en el ordenador e ingresar el siguiente comando y esperar la clonación de toda la aplicación al servidor:
	<pre>
		cap production deploy
	</pre>

**22.** Esto hará saltar un error porque el archivo **UGLIFIER** no funciona como JS, entonces hay que instalar **NODEJS** en el servidor:
	<pre>
		sudo apt-get install <b>nodejs</b>
	</pre>

**23.** Luego de instlar **NODEJS** en el servidor, vamos a la computadora remota y volvemos a ejecutar el comando de **capistrano** para hacer el deploy:
	<pre>
		cap production deploy
	</pre>

**24.** Si aparece un problema por la gema **PG** en la computadora, puede instalarse de dos maneras:
	<pre>
		gem install pg / gem install pg -- --with-pg-config=/usr/local/bin/pg_config
	</pre>
	ó
 	<pre>
		brew install libpqxx
	</pre>

**25.** Luego de eso, agregarla al **GEMFILE** del proyecto en computadora:
	<pre>
		# gem 'capistrano-rails', group: :development
		gem 'pg', '~> 0.18.4'
	</pre>

**26.** Luego de ello, realizar un registro de **gemas**:
	<pre>
		bundle install
	</pre>

**27.** Luego hacer un **commit** para guardarlo en el repositorio.

**28.** Luego, volver a ejectuar el comando de **capistrano** para hacer el **deploy**:
	<pre>
		cap production deploy
	</pre>

**29.** Luego ir al servidor y reiniciar NGINX para correr por primera vez la aplicación:
	<pre>
		sudo service nginx restart
	</pre>





### Deploy Rails project in AWS EC2 using capistrano:
&nbsp;
* How to deploy RubyonRails project to AWS EC2 using Capistrano: 
<br/>[YouTube Video](https://youtu.be/imdrYD4ooIk?t=25m34s)
<br/>[Github Repo](https://github.com/rkmmanivannan/rails-ec2-configuration/blob/master/method1/serverconfig)
[Amberbit Link](https://www.amberbit.com/blog/2015/11/28/how-to-deploy-rails-on-a-vps/)
&nbsp;
* Check misconfigurations on NGINX:
<br/>[DigitalOcean Link](https://www.digitalocean.com/community/questions/service-nginx-restart-fail)
<br/>[Postgress Instalation on OS ElCapitan](https://teamtreehouse.com/community/ruby-on-rails-bundle-fails-cannot-gem-install-pg-not-yet-resolved-need-help)
<br/>[Postgress Manual](http://www.postgresql.org/files/documentation/pdf/9.3/postgresql-9.3-A4.pdf)
<br/>[Manual Rails Instalation](https://gorails.com/setup/ubuntu/14.10)
<br/>[Rails Web App instaled on Unicorn]() https://gist.github.com/JamesDullaghan/5941259)
&nbsp;
* Create Secret Key:
<br/>
[Rake Secret](http://www.jamesbadger.ca/2012/12/18/generate-new-secret-token/)

### POSTGRESQL:
&nbsp;
**MacOS:** 
* http://postgresapp.com/documentation/
* http://postgresapp.com/documentation/cli-tools.html<br/>
* https://launchschool.com/blog/how-to-install-postgresql-on-a-mac<br/>
* http://www.postgresql.org/files/documentation/pdf/9.5/postgresql-9.5-A4.pdf<br/>	      
* http://www.linuxscrew.com/2009/07/03/postgresql-show-tables-show-databases-show-columns/<br/>
* https://chartio.com/resources/tutorials/how-to-start-postgresql-server-on-mac-os-x<br/>
* https://eggerapps.at/postico/<br/>
* http://postgresapp.com/<br/>
* https://www.mysql.com/products/workbench/<br/>

### MYSQL:
&nbsp;
* https://gist.github.com/hofmannsven/9164408	
