#Configurar un servidor AWS EC2 para Rails

1. Seleccionar la microinstancia que ofrece AWS. 
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/Create_instanse1.png "Create Instance")

2. Configurar las reglas de seguridad de tu instancia.
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/SecurityGroup.png "Firewall Instance")
![alt text](https://github.com/williamromero/Code-Snippets/blob/master/RoR/Screen%20Shot%202016-04-25%20at%2011.02.59%20AM.png?raw=true "Firewall Instance")

3. Crear o descargar la llave de seguridad privada para poder acceder a tu servidor.
![alt text](https://platzi.com/blog/wp-content/uploads/2015/07/PEM-Key.png "Firewall Instance")

4. Una vez descargada, debemos de instalar la llave en la carpeta ~/.ssh/ de nuestro ordenador MacOSX, por lo tanto vamos a dirigirnos a la carpeta con el siguiente comando:

```
  cd /                                                            # Nos dirige al root del ordenador.
  cd ~/.ssh/                                                      # Nos dirige a la carpeta SSH en donde se guardan las llaves de seguridad de los servicios.
  mkdir webapp                                                    # Creamos una carpeta, con el nombre de la llave de nuestro servicio.
  cd webapp                                                       # Ingresamos a la carpeta que hemos creado.
  cp /Users/tu_nombre_de_usuario/Downloads/webapp.pem webapp.pem  # Copiamos la llave de seguridad de la carpeta en la que se descargó, hasta la carpeta "webapp" que se encuentra en el folder ~/.ssh/.
```
  4.1 Otra forma de iniciar sería colocarlo en el escritorio de la siguiente forma:
  ```
    cd /Users/usuario/Desktop/
    mkdir AWS
    cp /Users/tu_nombre_de_usuario/Downloads/webapp.pem /Users/tu_nombre_de_usuario/Desktop/AWS/webapp.pem
  ```
  
  
* Antes de conectarnos a la instancia tendremos que configurar los permisos de la llave que ya descargamos para poder usarla en modo lectura únicamente. Para alterar los permisos de la llave usamos el comando:

```
  chmod 0400 webapp.pem
```

  *  Debemos de agregar la llave al folder SSH:
  
  ```
    ssh-add webapp.pem
  ```

* Generar una public/private rsa key pair.

```
  ssh-keygen
```
* Luego de ello, aparecerán los siguientes textos a los cuales deberemos solo de dar **enter**:

```
  Generating public/private rsa key pair.
  Enter file in which to save the key (/home/ubuntu/.ssh/id_rsa): 
  /home/ubuntu/.ssh/id_rsa already exists.
  Overwrite (y/n)? y
  Enter passphrase (empty for no passphrase): 
  Enter same passphrase again: 
  Your identification has been saved in /home/ubuntu/.ssh/id_rsa.
  Your public key has been saved in /home/ubuntu/.ssh/id_rsa.pub.
  The key fingerprint is:
  68:7c:20:6v:fw:9x:5y:1z:0a:5b:8c:7d:ae:bf:3g:4h ubuntu@ip-000-00-00-000
  The key's randomart image is:
  +--[ RSA 2048]----+
  |  . E.=+ .o+     |
  |   + = ++.o .    |
  |    @ + .o       |
  |   . B o         |
  |      + S        |
  |     . .         |
  |                 |
  |                 |
  |                 |
  +-----------------+
```
* Luego, ir a recoger la llave que quedó en los registros **~/.ssh/** para que Github pueda conectarse al servidor sin ninguna clave.
```
  ubuntu@ip-000-00-00-000:~/.ssh$ cd ~/.ssh/
  ubuntu@ip-000-00-00-000:~/.ssh$ ls
  authorized_keys  id_rsa  id_rsa.pub
  ubuntu@ip-000-00-00-000:~/.ssh$ nano id_rsa.pub
```

* Copiar el registro **ssh-rsa** que corresponda al servidor al con el que vaya a conectarse.
``` 
  ssh-rsa aaASEADLCAGIOERGALIWQTÑPWERWQOIOIWRT...ASNWGWR ubuntu@ip-000-00-00-000
```
* Ir a Github: Settings/SSH and GPG keys y agregar la clave SSH del servidor a las claves de Github. Colocar el nombre y la clave de SSH.

* Ahora, para acceder por medio de SSH, nos dirigimos a la ventana de comandos y colocamos:

```
  $ ssh -i "webapp.pem" ubuntu@ec2-00-00-000-00.us-west-2.compute-1.amazonaws.com
           *Key Name*                     *Public DNS*
```

* Ya dentro del servidor, debemos de actualizar todas las dependencias:

```
  sudo apt-get update && sudo apt-get -y upgrade
```

* Luego, instalar CURL & git:

```
  sudo apt-get install git CURL
```

* Luego de instalar estas dependencias, instalamos el RVM & Ruby:

```
  \curl -L https://get.rvm.io | bash -s stable --ruby
```

* Si en algún caso, cometieramos un error en seguir los pasos o cualquiera que sea ingresando la llave privada y el registro de esta quedase grabado
dentro de la memoria de las rutas ssh, podemos borrarlo con el siguiente comando:

```
  ssh-keygen -R ec2-00-00-000-00.us-west-2.compute-1.amazonaws.com
```

* Crear un usuario para realizar el deploy: 

```
  sudo adduser william
```

* Ingresar el usuario al grupo sudo:

```
  sudo adduser william sudo
```

* Ir al **"folder /etc/"**, luego insertar el comando:

```
  sudo visudo
```

  - Y agregar debajo del usuario root, el siguiente comando: 

  ```
    root    ALL=(ALL:ALL) ALL
    william ALL=(ALL:ALL) ALL
  ```
  
* Instalar Nginx con el siguiente comando:

  ```
    sudo apt-get install nginx
    
    sudo apt-get install nginx
    nginx -h
    cat /etc/init.d/nginx         // Activar por primera vez todos los servicios.
    /etc/init.d/nginx -h
    sudo service nginx start      // Arrancar el servidor nginx.
    cd /etc/nginx
  ```

* Instalar la gema de Rails & 

```
  gem install rails
```

* Instalar mysql y configurar su servicio:

```
  sudo apt-get install mysql-server mysql-client libmysqlclient-dev
```

* Ingresar la contraseña del usuario root:
```
  New password for the MySQL "root" user:
  ********15
```

* Instalar postgres y configurar su servicio:
```
  sudo apt-get install postgresql postgresql-contrib libpq-dev
```

* Instalar CAPISTRANO :
```
  gem install capistrano
```

* Instalar PUMA:
```
  gem install puma
```

* Instalar PHPMYADMIN
```
  sudo apt-get install phpmyadmin

```
  - Configuring phpmyadmin | Web server to reconfigure automatically: |*| Apache

* Instalar Java Virtual Machine (JVM):
```
  sudo apt-get install openjdk-7-jdk
```



###Artículo en producción

***
####Referencias:
* http://www.sitepoint.com/deploy-your-rails-app-to-aws/
* https://platzi.com/blog/instalar-ghost-amazon-ec2/
* http://docs.aws.amazon.com/gettingstarted/latest/wah-linux/getting-started-deploy-app.html
* https://manuais.iessanclemente.net/index.php/Servidor_Virtual_VPS_con_Amazon_EC2_-_Instalaci%C3%B3n_y_configuraci%C3%B3n
* http://superuser.com/questions/30087/remove-key-from-known-hosts
* https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#lists
* http://railsapps.github.io/installrubyonrails-ubuntu.html


####Deploy
***
* http://www.sitepoint.com/deploy-your-rails-app-to-aws/
* https://github.com/bvmake/WhosGotWhat/wiki/Installing-Rails-on-free-Amazon-EC2-Micro
* https://coderwall.com/p/yz8cha/deploying-rails-app-using-nginx-unicorn-postgres-and-capistrano-to-digital-ocean
* https://gist.github.com/jtadeulopes/7271019
* http://railsapps.github.io/installrubyonrails-ubuntu.html
* http://climber2002.github.io/blog/2015/02/09/deploy-rails-application-on-ubuntu-14-dot-04-part-2/
* https://gorails.com/deploy/ubuntu/16.04
* https://semaphoreci.com/community/tutorials/how-to-use-capistrano-to-deploy-a-rails-application-to-a-puma-server
* https://gist.github.com/ChuckJHardy/f44dda5f94c6bbdba9a4
* http://www.liquidweb.com/kb/how-to-install-and-configure-phpmyadmin-on-ubuntu-14-04/
* https://praveencastelino.wordpress.com/2013/04/24/connecting-to-amazon-ec2-server-on-mac-using-terminal-app/
* http://unix.stackexchange.com/questions/3568/how-to-switch-between-users-on-one-terminal
* https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/
* https://medium.com/@katsuya0515/deploying-rails-app-to-ec2-using-nginx-unicorn-by-capistrano-31fc3556b428#.2g14chwd3


* https://gist.github.com/JamesDullaghan/5941259
* https://www.ralfebert.de/tutorials/rails-deployment/
* http://vladigleba.com/blog/topics/deployment-series/
* https://coursecraft.net/c/rails-on-aws/splash
* https://www.digitalocean.com/community/tutorials/how-to-deploy-rails-apps-using-unicorn-and-nginx-on-centos-6-5


* https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-git-hooks-on-ubuntu-14-04
